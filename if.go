package main

import (
	"fmt"
)

func main() {
	fmt.Println("if Start!")
	{
		temp := 33
		//간단한 if문
		if temp > 28 {
			fmt.Println("에어컨 켜기")
		} else if temp <= 3 {
			fmt.Println("히터 켜기")
		} else {
			fmt.Println("뽀엥")
		}
	}

	{
		age := 25
		// 조건 여러개 쓰기 And : && , OR : ||
		if age >= 10 && age <= 15 {
			fmt.Println("아가")
		} else if age > 30 && age <= 20 {
			fmt.Println("친구")
		} else {
			fmt.Println("누구세요")
		}
	}

	{
		//쇼트서킷
		// A && B 인 경우 A의 결과가 False라면 우변의 B는 무시되고 False를 뱉는다.
		// A || B 인 경우 A의 결과가 True라면 우변의 B는 무시되고 True를 뱉는다.
	}

	{
		// if 초기문; 조건문. 초기문 : initializer(초기화). 초기문의 코드를 수행하고 조건문을 확인 (두줄로 쓸거 한줄로 쓴거랑 같음)
		// 구조
		// if 초기문; 조건문{
		//	문장
		//}
		a := 10
		if a += 20; a >= 30 {
			fmt.Println("Great!")
		} else {
			fmt.Println("ㅋㅋㅋㅋ")
		}

	}

}
