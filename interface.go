package main

import (
	"fmt"
)

func main() {
	fmt.Println("interface Let's Go!!")
	{
		// 인터페이스 : 구체화된 객체가 아닌 추상화된 상호작용으로 관계를 표현. Method는 관계 + 구현이라면 interface는 구현을 제외하고 오로지 관계만을 표현한 것이다.
		// 인터페이스도 Type임!!
		/*
			인터페이스 규칙!
			1. 메서드는 반드시 메서드 명이 있어야 한다
			2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다.
			3. 인터페이스에서는 메서드 구현을 포함하지 않는다.
		*/
	}

	{
		student1 := Student{"lee", 3}
		student2 := Student{"kim", 4}
		var infoer Infoer
		infoer = student1 // Student 구조체가 infoer interface가 가져야 하는 info()메소드를 갖고 있으므로 대입이 된다. (interface에 정의된 모든 메소드를 구현하지 않으면 안됨)
		fmt.Println(infoer.info())
		// infoer.getAge() => infoer 인터페이스가 getAge()는 메소드로 갖고 있지 않기 때문에 불가능 하다.

		infoer = student2
		fmt.Println(infoer.getName())
	}

	{
		//이런 인터페이스 어디다 써먹냐?
		// 예시 : fedex에서 사용하던 API를 우체국 api로 바꾸려고 하니 각 회사에서 제공하는 api의 타입이 달라 오류가 생김. 이럴때 interface를 통해 해결 가능
		// interfaceUse 폴더에서 설명
	}

	{
		//이런 추상화를 통해 내부 동작을 감춰 서비스 제공자와 사용자 모두에게 자유도를 준다.
		// 위에서의 예시처럼 어느 택배사이든 상관 없이 Send라는 함수 하나만으로 우편을 보낼 수 있다는 것이다. (Fedex이든 koreaPost이든 상관이 없다)
	}

	{
		// 덕타이핑! : 어떤 새를 봤는데 그 새가 오리처럼 걷고, 오리처럼 날고, 오리처럼 소리내면 나는 그 새를 오리라고 부르겠다.
		// Go나 python은 덕타이핑을 제공함.
		// 자바의 interface같은 경우 패키지 내부에서 interface를 정의하고 구조체에 implement로 붙이는
		/*	(패키지 쪽에서 정의함)
			type Sender interface{
				Send(string)
			}
			type koreaSender struct implements Sender{

			}
		*/
		// 이런식으로 인터페이스를 사용함. => 외부 패키지 제공자가 스스로를 밝힘. 난 오리야!!!
		// 하지만 go 나 python은 서비스 제공자는 제공만 하고 소비자가 거기에 이름을 붙임.
		/*
			(소비자 쪽에서 정의함)
			=> 너가 뭔진 모르지만 Send라는 함수만 존재하면 난 Sender라고 할거야.
			type Sender interface{
				Send(string)
			}
		*/

		// 이런 덕타이핑의 장점 : 사용자 중심의 코딩이다.
		// 인터페이스 구현 여부를 타입 선언시 하지 않고, 인터페이스가 사용될 때(사용자가 사용할 때) 결정하기 때문에 서비스 제공자는 구체화된 객체를 제공하고 사용자가 필요에 따라 인터페이스를 정의해 사용할 수 있음.
	}
	{
		//인스턴스 사이즈. [인스턴스 메모리 주소][타입 정보] 로 구성됨. (8byte 주소크기)(8바이트 타입 정보) 도합 16바이트로 interface는 구성됨.
		// u = User{}, stringer(인터페이스) = u 라고 하면 [u변수][User타입]
	}

	{
		// 빈인터페이스
		// interface {} => 메소드가 없는 인터페이스
		// => 모든 타입이 가능
		PrintVal(10)
		PrintVal(3.14)
		PrintVal(Student{"hihi", 10})
	}

	{
		//인터페이스 타입 변환
		/*
			var a Interface
			t,err := a.(concreteType)
			보다 구체적인 타입으로 변환 가능
		*/

	}
}

type ExInterface interface { // interface 예시
	sample()                 // 구현이 빠진 Method를 적는다.
	sample2(example int) int // sample이라고 메소드 명을 칭하면 메소드 명이 겹쳐서 안됨.
	//_(x int)                 // 메서드는 이름이 반드시 존재해야함!
} // 이 두 메소드를 지니고 있으면 ExInterface로 보겠다! 라는 의미임. 즉

type Student struct {
	name string
	age  int
}

type Infoer interface {
	info() string
	getName() string
}

func (s Student) info() string {
	return fmt.Sprintln("I'm", s.name, "and", s.age, "year's old") // Sprintf => 반환 타입이 String이다
}
func (s Student) getAge() string {
	return fmt.Sprintln(s.age, "year's old") // Sprintf => 반환 타입이 String이다
}

func (s Student) getName() string {
	return fmt.Sprintln(s.name) // Sprintf => 반환 타입이 String이다
}

// 빈 인터페이스
func PrintVal(v interface{}) { //어느 타입이든 들어올 수 있음.
	switch t := v.(type) {
	case int:
		fmt.Println("type is int", t)
	case float64:
		fmt.Println("type is float64", t)
	default:
		fmt.Println("no support type", t)
	}
}
