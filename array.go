package main

import (
	"fmt"
)

func main() {
	fmt.Println("array 레고")
	// 배열은 같은 타입의 데이터들로 이루어진 타입.
	// 자료구조 중 한 종류임. 자료구조란 데이터들을 어떤 혀태로 저장할지를 나타내는 구조를 뜻함

	{
		fmt.Println("simple array 생성")
		// var 변수명 [개수]타입, 정적 배열
		var arr [5]float64 = [5]float64{12.1, 12.2, 12.3, 12.4, 12.5}

		for i := 0; i < 5; i++ {
			fmt.Println(arr[i])
		}

		//인덱스 값 변경
		arr[0] = 13.1
		fmt.Println(arr[0])
	}

	{
		// 다양한 배열 변수 선언
		// var nums [5]int                                     // 초기화 하지 않을시 int의 경우 모두 0으로 채워짐
		// days := [3]string{"monday", "tuesday", "wednesday"} //변수 선언과 마찬가지로 var 생략 가능. 그리고 선언과 동시에 초기화
		// var temps [5]float64 = [5]float64{24.3, 26.7}       // 앞의 두 값은 저 값으로 채워지지만, 나머지 값은 default(0.0)으로 채워진다.
		// var s = [5]int{1: 10, 3: 30}                        //int타입 5사이즈 배열을 변수 s에 넣어라. 단 1번째 인덱스에 10을, 3번째 인덱스에 30을 넣어라.
		// x := [...]int{10, 20, 30}                           // [...]은 길이의 값을 {10,20,30}의 길이로 정하겠다는 것. 3이 되는것 이 경우는...
		// [...] 랑 [] 랑은 완전 다름. 전자는 길이가 고정된 배열이고, 후자는 동적 배열임

	}

	{
		//배열 선언시 배열의 길이는 언제나 상수이다.
		const Y int = 3
		// x:= 5
		// a := [x]int{1,2,3,4,5}   //x는 변수이므로 오류임
		// b := [Y]int{1,2,3}		// 얘는 괜찮음
	}

	{
		//배열의 길이를 아는 내장 함수 len(배열 변수)
		nums := [...]int{10, 20, 30, 40, 50}
		for i := 0; i < len(nums); i++ {
			fmt.Println(nums[i])
		}
	}

	{
		//range를 이용한 배열 순회
		// for문에서 range 사용시 데이터의 요소를 순회하며 두개의 값을 반환한다.
		// 배열의 경우 (인덱스, 값) 을 반환한다.
		var t [5]float64 = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
		for i, v := range t { // (i,v) == (인덱스, 값)
			fmt.Println(i, v)
		}

		//만약 i,v 둘중에 한 값이 필요하지 않다면?
		for _, v := range t { // (_,v) == (인덱스, 값) 늘 그랬듯 안받으면 됨.
			fmt.Println(v)
		}
	}

	{
		//배열 요소 찾아가기

		a := [5]int{1, 2, 3, 4, 5}
		b := [5]int{100, 200, 300, 400, 500}

		for i, v := range a {
			fmt.Printf("a[%d]= %d\n", i, v)
		}
		fmt.Println()
		for i, v := range b {
			fmt.Printf("b[%d]= %d\n", i, v)
		}

		b = a // 복사하는 공간의 크기가 같아야 함.
		fmt.Println()
		for i, v := range b {
			fmt.Printf("b[%d]= %d\n", i, v) // b에 a의 주
		}

	}

	{
		// 다차원 배열
		fmt.Println("이차원 배열")
		b := [2][1]int{{1}, {2}}
		for i, _ := range b {
			for _, y := range b[i] {
				fmt.Println(y)
			}
		}

		for _, arr := range b { //이렇게 쓰도록 하자!
			fmt.Println(arr)
			for _, y := range arr {
				fmt.Println(y)
			}
		}

	}
}
