package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("for Let's go~")
	//기본 구조
	// for 초기문; 조건문; 후처리{
	//		코드블록 => 조건문이 true인 경우 수행됨.
	//} ==> 다른 언어의 반복문과 크게 상이하지 않아보임.

	{
		fmt.Println("simple case")

		//반복문 안에서 변수 선언 및 초기화
		for i := 0; i < 10; i++ {
			if i == 9 {
				fmt.Print(i, "\n")
				break
			}
			fmt.Print(i, ",")
		}
	}

	{
		//반복문 바깥의 변수 사용경우
		i := 0
		for ; i < 10; i++ { //초기화 부분을 생략. 단 ;으로 구분은 해줘야 함.
			if i == 9 {
				fmt.Print(i, "\n")
				break
			}
			fmt.Print(i, ",")
		}
	}

	{
		//반복문 후처리 생략
		// for 초기문; 조건문; { => 후처리 생략
		//		코드블록
		//}

		i := 0
		for i < 10 { // 초기화, 후처리 다 없고 조건문만 있는 경우, 이렇게 적어도 됨.
			if i == 9 {
				fmt.Print(i, "\n")
				break
			}
			fmt.Print(i, ",")
			i++
		}
	}

	{
		// 무한 루프
		i := 0
		j := 0

		//방법 1. 조건문에 true 넣기
		for true {
			i++
			fmt.Println(i)
			if i > 10 {
				break
			}
		}

		//방법 2. 아무것도 안쓰기
		for {
			j++
			fmt.Println(j)
			if j > 10 {
				break
			}
		}
	}

	{
		fmt.Println("continue와 break")
		//continue의 경우 만나는 순간 반복문의 후처리로 바로 이동
		//break의 경우 만나는 순간 반복문 바로 종료
		//늘 쓰던대로 하면 될듯.

		stdin := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("input Number")
			var number int
			_, err := fmt.Scanln(&number) //Go에서는 선언한 변수를 무조건 사용해야함. return 받는 값중에 사용하지 않을 값이 있으면 _으로 처리.
			if err != nil {
				fmt.Println("input number~")
				//키보드 버퍼 지우기.
				stdin.ReadString('\n')
				continue
			}

			//fmt.Println("number is", number)
			fmt.Printf("number is %d", &number)
			if number%2 == 0 {
				break
			}
		}

		fmt.Println("for is done")
	}

	{
		fmt.Println("flag 변수, label 활용")
		//특정 조건일때 for문을 종료하고 싶을때 활용할 수 있음.
		a := 1
		b := 1
		found := false // 플래그 변수 역할
		for ; a <= 9; a++ {
			for b = 1; b <= 9; b++ {
				if a*b == 45 { // 이중 포문 내부의 포문 종료
					found = true
					break
				}
			}

			if found { // 플래그 변수를 이용해 바깥 포문도 종료
				break
			}
		}
		// ... 이렇게 하면 중첩되는 반복문의 수만큼 break를 걸어줘야하고 그에 따른 flag 변수도 많아짐
		// label을 활용해 이런 문제를 해결할 수 있음

	OuterFor: //Label 레이블
		for i := 1; i <= 9; i++ {
			for j := 1; j <= 9; j++ {
				if i*j == 45 {
					break OuterFor // 특정 라벨이 달린 포문 자체를 break함.
				}
			}
		}
		// 하지만 Label이나 goto문 같은 경우 IP를 강제로 바꾸기 때문에 스택이 꼬이는? 문제가 발생할 수 있음.
		// 그러므로 좀 복잡하더라도 flag를 쓰는게 낫다.
		fmt.Println()

	}
}
