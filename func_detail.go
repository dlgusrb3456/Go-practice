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
		var operator func(int, int) int
		operator = getOperator("+") //add 함수의 주소가 반환됨. * 이었다면 mul이 반환됐을 것임
		var result = operator(3, 4) // add(3,4)와 동일함.
		fmt.Println(result)
	}

	{
		//함수 리터럴 (람다)
		var operator func(int, int) int
		operator = getOperator2("*") //=> 이 함수 보기
		var result = operator(3, 4)
		fmt.Println(result)
		// 일반 함수는 상태를 가질 수 없지만, 함수 리터럴은 내부 상태를 가질 수 있다.
		// 이게 뭔 소리야!
		i := 0
		f := func() { // 같은 지역이 변수를 사용할 수 있다. 인자로 받지 않아도.
			i += 10
		}
		i++
		f()
		fmt.Println(i) // 11
		// 중요한건 캡쳐(외부 인자)는 값복사가 아닌 레퍼런스 복사임. 위에서 i += 10한 i는 *i와 같다. (같은 메모리 공간을 가리킨다)
		//그래서 주의할 점!
		CaptureLoop()  //3 3 3이 출력됨.
		CaptureLoop2() // 0 1 2가 출력됨.
		// Go루틴에서 이 부분을 잘 알아둬야함. 레퍼런스 복사!
	}

	{
		//의존성 주입. 어디다 쓰일지는 잘 모르겠다..
		f, err := os.Create("test2.txt")
		if err != nil {
			fmt.Println("failed to create test2.txt")
			return
		}
		defer f.Close()

		writeHello(func(msg string) { //리터럴 함수
			fmt.Fprintln(f, msg)
		})

		writeHello(func(msg string) { //리터럴 함수
			fmt.Println(msg)
		})

	}
}

type Writer func(string)

func writeHello(write Writer) { //Hello World를 어디다가 적는진 알 수없지만 그저 Hello World를 적기만 한다. 어디에 어떻게 적을지는 사용자가 정한다. 이를 의존성 주입이라 한다.
	write("Hello World")
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

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func getOperator2(op string) func(int, int) int {
	if op == "+" {
		return func(a, b int) int { //리터럴
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int { //리터럴
			return a * b
		}
	} else {
		return nil
	}
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]() //3 3 3 이 출력됨. => 얘네가 가리키는 i는 모두 하나의 i를 가리키고 있으므로 132번째 줄에서의 i가 출력된다. (함수를 만들 당시에 i)
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]() //0,1,2 이 출력됨. => 얘네가 가리키는 v는 {}이 시작할 때마다 새롭게 생성된다. 그러므로 각 v가 모두 다른 주소 값을 갖는다.
	}
}
