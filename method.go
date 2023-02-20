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
		// 객체란 데이터(state)와 기능(Function)을 묶은 것이다.
	}
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
