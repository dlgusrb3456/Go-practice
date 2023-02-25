package main

import (
	"fmt"
)

func main() {
	fmt.Println("method gi")
	// method는 func와 같다. 메소드 == 함수.
	// 함수의 한 종류로 method가 존재함. method는 타입에 속한 함수임.
	// 차이점 : 함수는 독립적이지만 메서드는 타입에 종속적이다. 타입 > 메서드

	/*
		구조
			func ( r Rabbit ) info() int{ => r Rabbit이라는 구조체(모든 패키지 지역내(이 패키지 내부에서만 정의된 타입이란 뜻. int, float, bool 과 같이 전역 타입은 안됨.) 타입 가능. 구조체, 별칭타입 ...)에 속한, info()라는 함수, 반환 타입은 int. 라는 뜻이다. (여기서 r Rabbit은 리시버라고 부른다)
				return r.width * r.height
			}
	*/

	{
		A := account{100}
		withdrawFunc(&A, 10)   // 그냥 함수로 사용
		fmt.Println(A.balance) //90

		A.withdrawMethod1(10)  // 구조체에 속한 메소드로 사용
		fmt.Println(A.balance) //80

		A.withdrawMethod2(10)
		fmt.Println(A.balance) //80 (포인터를 넘겨주지 않으면 call by value로 넘어가서 값이 변하지 않음)
	}
	{
		var A myInt = 15
		A.Add(10)
		fmt.Println(A) // 정상적으로 더해지지 않음.

		A.Add2(10)     // 방법 1
		fmt.Println(A) // 정상적으로 더해짐.

		A = A.Add3(10) // 방법 2 => 이게 보다 직관적인듯.
		fmt.Println(A) // 정상적으로 더해짐.

	}

	{
		//객체로의 진!! 화!!
		// 객체란 데이터(state)와 기능(Function)을 묶은 것이다. low coupling high cohension
		// oop에서 말하는 object란 Data와 Function을 묶은 것 이라고 생각한다.
		// 즉 func (s *Student) SendReport (p *Professor, r *Report) ==> Student와 Professor Student와 Report 사이의 관계를 SendReport를 통해 정의한다. 즉, 객체와 객체간의 관계를 정의한게 method이다.
	}

	{
		A := account1{100, "Lee", "Hyun"}
		A.withdrawPointer(10)
		fmt.Println(A.balance) // 90
		A.withdrawValue(10)
		fmt.Println(A.balance) // 90

		withdrawPointer2(&A, 10)
	}

	{
		// 함수 인자로 포인터를 사용할 때와 value를 사용할 때의 차이...
		// 구조체의 field 값을 변경해도 구조체의 본질 자체가 변하는 것이 아니라면 포인터를 사용한다. (학생 구조체에서 학생의 나이를 하나 증가시킨다고 해서 학생이 다른 학생이 되는것이 아닌 것과 같음)
		// 하지만 Temperature와 같이 값 변경마다 본질적인 의미가 변하는 것이라면 value를 사용하는 것이 좋다.

		// 그리고 Go에서는 생성자, 소멸자 개념이 없다. 상속도 없다...
		vip := VipUser{User{"Lee"}, 2}
		fmt.Println(vip.userName()) // embedded field의 method도 바로 사용 가능하다.
	}
}

type User struct {
	name string
}

type VipUser struct {
	User
	Level int
}

func (user *User) userName() string {
	return user.name
}

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { //#1
	a.balance -= amount
}

func (a *account) withdrawMethod1(amount int) { //#2
	a.balance -= amount
}

// #1과 #2는 동일하게 작동한다. 단순히 의미상의 차이만 존재할 뿐....
func (a account) withdrawMethod2(amount int) {
	a.balance -= amount
}

type myInt int //별칭 타입에도 메소드를 붙일 수 있다.
func (m myInt) Add(amount myInt) {
	m += amount
}

func (m *myInt) Add2(amount myInt) {
	*m += amount
}

func (m myInt) Add3(amount myInt) myInt {
	m += amount
	return m
}

type account1 struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *account1) withdrawPointer(amount int) {
	a1.balance -= amount
}
func withdrawPointer2(a1 *account1, amount int) {
	a1.balance -= amount
}
func (a2 account1) withdrawValue(amount int) {
	a2.balance -= amount
}
