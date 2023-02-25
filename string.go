package main

import (
	"fmt"
)

func main() {
	fmt.Println("string let's Go!")
	//ASCII라는 문자코드를 사용했었음. 얜 1byte로 문자를 표현함. 0 제외하고 총 255개를 표현할 수 있음
	//중국어나 일본어 한국어 등 사용해야 할 문자가 많아지면서 255개로는 택도 없어짐
	// 그래서 UniCode가 나옴. Go에서는 UTF-8을 사용함. 더 큰 바이트를 사용해버림.

	{
		//문자열 표현하는 두 방법. "" 와 ``
		poet1 := "동해물과 백두산이 마르고 닳도록.\n 하느님이 보우하사 우리 나라 만세"
		poet2 := `무궁화 삼천리 화려강산 \n대한사람 \n대한으로 
		길이 보전하세.`
		fmt.Println(poet1) //동해물과 백두산이 마르고 닳도록.
		//하느님이 보우하사 우리 나라 만세

		fmt.Println(poet2) //무궁화 삼천리 화려강산 \n대한사람 \n대한으로
		//길이 보전하세.
	}

	{
		//문자열 순회. len()을 이용. => 바이트 길이 반환
		str := "Hello world 망할"

		for i := 0; i < len(str); i++ {
			fmt.Printf("타입: %T 값:%d 문자값:%c\n", str[i], str[i], str[i]) // 영어는 1byte로 표현되지만, 한글은 3byte로 하나의 글을 표현함.
		}
		// 이 방식의 순회는 제대로된 순회가 이루어지지 않음. str[7]이 의미하는게 7번째 문자가 아닌 7번째 byte의 값을 의미하기 때문.
	}

	{
		//문자열 순회. []rune 타입으로 변환 후 한 글짜씩 순회
		str := "Hello world 망할"
		arr := []rune(str) //rune 배열. rune -> int32의 별칭타입. 배열 한칸이 4byte가 되는것!
		for i := 0; i < len(arr); i++ {
			fmt.Printf("타입: %T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i]) // 영어는 1byte로 표현되지만, 한글은 3byte로 하나의 글을 표현함.
		}
		//4byte를 한칸의 크기로 하는 배열로 문자열을 타입 변환 했기 때문에 값이 문자 단위로 나온다.
	}

	{
		//문자열 순회. range 사용
		str := "Hello world 망할"
		for _, v := range str {
			fmt.Printf("타입: %T 값:%d 문자값:%c\n", v, v, v)
		}

	}

	{
		//문자열 합산
		str1 := "hello"
		str2 := "world"

		str3 := str1 + " " + str2
		fmt.Println(str3)

		str1 += " " + str2
		fmt.Println(str1)
		// 오직 문자열 더하기만 Go에서는 지원해줌. 빼기, 곱하기 안됨!
	}

	{
		//문자열 비교!
		// == , != 으로 equal 비교
		// < , > , <= , >= 으로 대소 비교
		// 사전식으로 크기를 비교함. 대문자거 더 작음. A-Z : 65-90, a-z : 97-122
		str1 := "hi?"
		if str1 == "hi?" {
			fmt.Println("hello~")
		} else {
			fmt.Println("fuck you!")
		}

	}

	{
		//문자열 구조.
	}

}
