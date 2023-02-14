package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("IO Start!")

	{
		//표준 출력 함수
		// Print() : 함수 입력값들을 출력
		// Println() : 함수 입력값들을 출력하고 개행(\n)
		// Printf() : 서식(format)에 맞도록 입력값들을 출력

		a := 10
		b := 20
		f := 3279943873.234

		fmt.Print("a: ", a, " b: ", b, " f: ", f)   // a: 10 b: 20 f: 3.279943873234e+09 (출력값 사이에 빈칸을 넣어주지 않음)
		fmt.Println("a: ", a, " b: ", b, " f: ", f) // a:  10  b:  20  f:  3.279943873234e+09 (출력값 사이에 빈칸을 넣어줌, 그래서 두칸 띄어져있음)
		fmt.Printf("a: %d b: %d f: %f", a, b, f)    // a: 10 b: 20 f: 3279943873.234000 (%라는 formatter로 값을 넣음 d:decimal f:float(6자리) s:string v:데이터 타입 맞춰서 알아서 해준다 / 여기선 float가 위와 달리 지수표현이 아닌 실수 표현으로 출력됨 지수형태로 할거면 e사용 g는 자동으로 지수,실수 해줌)
	}

	{
		fmt.Println()

		a := 123
		b := 456
		c := 123456789
		fmt.Printf("%5d,%5d\n", a, b)    //  123,  456, 5자리로 출력하는데 숫자가 모자르면 앞에 빈칸으로 출력
		fmt.Printf("%05d,%05d\n", a, b)  //00123,00456, 위와 동일하지만 빈칸이 아닌 0으로 채워서 출력
		fmt.Printf("%-5d,%-05d\n", a, b) //123  ,456, 똑같이 5자리로 표현하지만 뒤가 아닌 앞으로 밀어서 출력

		fmt.Printf("%5d,%5d\n", c, c)    //123456789,123456789 넘는건 걍 출력
		fmt.Printf("%05d,%05d\n", c, c)  //123456789,123456789
		fmt.Printf("%-5d,%-05d\n", c, c) //123456789,123456789

	}

	{
		fmt.Println("표준 입력!")
		//Scan() 표준 입력에서 값을 입력받는다.
		//Scanf() 표준 입력에서 서식 형태로 값을 입력받는다.
		//Scanln() 표준 입력에서 한 줄을 읽어서 값을 입력받는다.

		//Scanln() 예재
		var a, b int
		n, err := fmt.Scanln(&a, &b) //&a => a변수에 할당된 데이터가 담긴 메모리 주소. Go에서는 함수의 인자는 모두 Rvalue로 동작함.
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n, a, b)
		}

	}

	{
		stdin := bufio.NewReader(os.Stdin) //stdin은 표준 입력을 나타냄. 표준입력을 통해 입력된 표준 입력 버퍼 접근 하기 위함

		var a int
		var b int

		n, err := fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println(err)
			stdin.ReadString('\n') // 표준 입력 버퍼에서 \n이 나올때까지 읽어라. (에러시 다음 입력을 읽어오기 위해 \n까지 읽어오라는 뜻. 버퍼 비우기)
		} else {
			fmt.Println(n, a, b)
		}

	}
}
