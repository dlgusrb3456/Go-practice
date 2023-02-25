package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("detail func")
	{
		//가변 인수 함수
		fmt.Println(sum(1, 2, 3, 4, 5))
		// 그럼 1,2,3,4, "Hello" 와 같이 여러 타입 종류를 넣을 순 없나?
	}
	{
		// 빈 인터페이스를 이용해 여러 타입 종류를 넣을 수 있다.
		AllType(1, 2, 3, 4, "asf")
	}
	{
		// defer 지연 실행
		// defer 명령문 => 함수 종료 전에 실행을 보장
		deferTest() // defer가 제일 늦게 실행됨.
		// 왜 사용하나?  주로 OS 자원을 반납해야 할때 사용함. 작업 이후 함수 종료 직전 자원 반납하는 느낌
	}

	{
		// 함수 타입 변수. 함수를 값으로 갖는 변수
		// 어떻게 함수를 값으로 갖을 수 있나..?
	}
}

func sum(nums ...int) int { // ...을 통해 인자의 수를 고정하지 않고 넣을 수 있다. 가변 인수 함수
	// 사실상 []int 타입임. 타입을 []int 라고 해도 됨
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func AllType(args ...interface{}) {
	for _, v := range args {
		switch f := v.(type) {
		case bool:
			fmt.Println("Bool", f)
		case int:
			fmt.Println("int", f)
		default:
			fmt.Println("who r u", f)
		}
	}
}

func deferTest() {
	defer fmt.Println("when you run")
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}
	defer f.Close()
	defer fmt.Println("delete OS")
	defer fmt.Println("1") // defer끼리는 순서가 보장되지 않는듯?

}
