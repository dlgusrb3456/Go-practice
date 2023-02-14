package main

import "fmt"

// 변수 : 값을 저장하는 메모리 공간을 가리키는 이름
// 왜 중요한가? 프로그램이란 결국 데이터를 연산/조작 하는 일. 이 데이터는 메모리에 저장됨
// 이 메모리에 있는 데이터에 접근하기 위한 방법이 변수

var g int = 10 // 전역 변수로 패키지가 끝날때까지 계속 메모리에 저장되어 있다.

func main() {
	fmt.Println("var start!")

	{
		// 변수 할당 및 변경 구조

		// var => Variable (변수를 선언하겠다!)
		// a => 변수명, int => 타입, = => 대입연산자 (우측의 값을 좌측에 대입해라)
		// 20 = 대입 연산자를 통해 좌변에 대입될 Data

		// 1. 컴퓨터는 메모리에 int 를 담을 빈 공간을 찾는다
		// 2. 메모리에 Data를 담는다 (Copy)
		// 3. 변수명이 Data가 담긴 메모리 주소를 가리키게 한다
		var a int = 20

		// 변수 값 변경
		// 이미 a는 값이 선언되어 메모리가 할당되었으므로 var를 써서 재선언 할 필요 없다
		// a변수가 가리키는 메모리 공간의 Data를 20으로 copy한다
		a = 20

		// 위의 두 변수가 가리키는 주소의 Data를 가져와서 출력
		fmt.Println(a) // 20 , Good Morning
	}

	{
		// 변수 타입, 다른 언어와 같음. 단 Go는 초 강타입언어라 다른 타입의 변수 연산시 타입 변환 잘 해줘야 함.

		var a int16 = 3456   // 2바이트 정수
		var b int8 = int8(a) // 2바이트 정수인 a를 1바이트 정수로 타입 변환
		fmt.Println(a, b)    //3456, -128  .... 왜 b는 -128인가?
		// => 적은 타입에 많은 값을 넣어서 나머지 값이 버려진 것
		// 3456 : 00001101 10000000 => 앞의 1바이트가 날아가고 10000000(-128)만 남음

		//숫자타입
		//uint8,uint16,uint32,uint64 => 부호 없이 양수형 정수 표현
		//int8,int16,int32,int64 => 부호 있이 정수 표현
		//float32,float64 => 부호 있이 실수 표현

		//그외
		//bool, string, 배열, 슬라이스(동적 배열), 구조체, 포인터, 함수타입, 맵, 인터페이스, 채널(난중에)
	}

	{
		fmt.Println()
		//타입 변환
		//int + int16 이것도 오류를 내는 초초초강타입 언어임.
		a := 3
		var b float64 = 3.5

		// var c int = b => 오류! 타입이 안맞음!
		var c int = int(b)

		// d := a * b => 당연히 안됨!
		d := a * int(b)

		var e int64 = 7
		// f := a * e int와 int64도 다른 타입으로 인식해서 오류! 임
		f := a * int(e)

		fmt.Println(a, b, c, d, e, f) //3 3.5 3 9 7 21
	}

	{
		//변수 선언법
		var a0 int = 10             // 선언과 동시에 초기화
		var a1 int                  // 선언만
		var a2 = 10                 // 타입 생략. 단, 이경우는 무조건 값을 초기화 해줘야함. 초기화된 값의 타입을 갖게됨.
		a3 := 10                    // var 와 type 생략. := 기호를 사용 (: 이 var와 같음). 이경우도 위와 같은 이유로 초기화를 꼭 해줘야 함.
		fmt.Println(a0, a1, a2, a3) // 10 0(default 값) 10 10

		//타입별 default 값
		// 모든 정수 : 0 , 모든 실수 : 0.0, boolean : false, 문자열 : ""(빈 문자열), 그외 : nil(정의되지 않은 메모리 주소를 나타내는 Go의 키워드)
	}

	{
		// 2진수 정수 표현
		// 음수, 양수를 표현해야 하므로 맨 앞 비트는 부호비트로 표현됨. 0이면 양수 1이면 음수 (보수법 확인)
		// 2진수 실수 표현
		//152.345 => 0.152345*10^3 => 소수부 : 152345 지수부 : 3, 이렇게 저장함
		// 마찬가지로 부호비트가 존재함. 그리고 지수부와 소수부로 나뉨. float32 기준 지수부 8비트, 소수부 23(7자리)비트 할당함
		// 소수부가 잘리는 문제가 발생해 계산에 오차가 낼 수 있음

		var a float32 = 1234.523
		var b float32 = 3456.123
		var c float32 = a * b
		var d float32 = c * 3
		fmt.Println(a) // 1234.523
		fmt.Println(b) // 3456.123
		fmt.Println(c) // 4.266663e+06 ...? 4266663.334329 값이지만 뒤에 34329가 잘림
		fmt.Println(d) // 1.2799989e+07 여기도 마찬가지로 잘림 (심지어 오차난 수를 배수 하니 오차가 더 커짐)
	}
}