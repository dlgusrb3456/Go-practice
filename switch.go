package main

import (
	"fmt"
)

func main() {
	{
		fmt.Println("simple example")
		a := 3

		// if / elseif 가 많이 쓰이는 경우 가독성을 위해 switch를 사용하기도 한다
		// 비교가 값을 기준으로 진행되므로, 조건문을 사용해야 한다면 if문을 사용해라
		switch a {
		case 1:
			fmt.Println("a==1")
		case 2:
			fmt.Println("a==2")
		case 3, 4: //이처럼 값을 여러개 비교할 수도 있다.
			fmt.Println("a==3")
		default:
			fmt.Println("a...?")
		}
	}

	{
		temp := 18

		fmt.Println("case에 조건문을 넣는 법")
		//작동 원리
		// switch문에 들어간 값(ex: true)과 case문의 조건문의 결과가 == 이라면 해당 case문을 사용한다.
		switch true {
		case temp < 10, temp > 30:
			fmt.Println("나가기 싫어")
		case temp >= 10 && temp < 20:
			fmt.Println("나가도 되나?")
		case temp >= 15 && temp < 25:
			fmt.Println("나가야지")
		default:
			fmt.Println("뜨시다!")
		}
	}

	{
		fmt.Println("Switch 초기문")

		switch age := 32; age {
		case 10:
			fmt.Println("Teenage")
		case 20:
			fmt.Println("이십대")
		case 32:
			fmt.Println("hehe")
		default:
			fmt.Println("호에에엥")
		}
	}

	{
		fmt.Println("break와 fallthrough")
		//다른 언어와 다르게 go에서는 case마다 break를 안적어도 된다(적어도 상관은 없음). 알아서 빠져나온다.

		//fallthrough를 통해 case에 걸려도 빠져나오지 않고 다음 case를 이어서 실행한다.
		a := 3

		switch a {
		case 1:
			fmt.Println("a==1")
		case 2:
			fmt.Println("a==2")
		case 3, 4:
			fmt.Println("a==3")
			fallthrough
		default:
			fmt.Println("a...?")
		}
		//a==3
		//a...? 이렇게 두개가 모두 출력 된다.
	}

	{
		fmt.Println("my favorite color is", colorToString(getMyFavoriteColor()))
	}

}

type ColorType int // 타입 선언. int타입과 똑같은데 이름을 ColorType이라고 지정
// int라고 해도 되지만, 보다 의미를 분명하게 하기 위해 이렇게 사용한다.
const (
	Red ColorType = iota // 0,1,2,3 ... 으로 저장됨
	Blue
	Green
	Yellow
)

func getMyFavoriteColor() ColorType {
	return Green
}
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "RED"
	case Blue:
		return "BLUE"
	case Green:
		return "GREEN"
	case Yellow:
		return "YELLOW"
	default:
		return "hoho"
	}
}
