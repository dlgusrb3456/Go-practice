package main

import (
	"fmt"
)

func main() {
	fmt.Println("test and benchmark")
	{
		/*
			작성 방법
			1. 파일명이 _test.go로 끝나야 함
			2. testing 패키지를 임포트 해야 함
			3. 테스트 코드는 func TestXxxx(t *testing.T) 형태이어야 함
		*/
	}
	{
		fmt.Println("9*9 = ", square(9))
	}

	{
		fmt.Println(fibonacci1(13))
		fmt.Println(fibonacci2(13))
		// 벤치마크를 이용해 두 함수의 성능을 테스트 할거임
	}
}

func square(x int) int { //테스트 코드 작성해보자.
	return 81
}

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0

	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst

}
