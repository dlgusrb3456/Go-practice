package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointer...!!")
	// java, python 같은 경우 pointer를 개발자가 다루지 않지만, 내부적으로 사용이 된다.
	// go는 C와 C++과 마찬가지로 pointer를 개발자가 다루게 해주지만, 정도가 낮다.

	{
		// 포인터란~~? 포인터는 메모리 주소를 값으로 갖는 타입
		var a int  // int 타입 변수
		var p *int // int 타입 변수의 메모리 주소를 담는 변구
		p = &a     // &a == a 변수에 할당된 메모리 주소를 포인터 변수 p에 대입
		*p = 20
		fmt.Println("p: ", p, " *p: ", *p) //p:  0xc000016088  *p:  20
	}

	{
		//여러 포인터 변수가 하나의 변수를 가리킬 수 있음
		var a int
		var p1 *int
		var p2 *int
		var p3 *int
		p1 = &a
		p2 = &a
		p3 = &a
		fmt.Println(p1, p2, p3) // 0xc0000160c0 0xc0000160c0 0xc0000160c0
	}

	{
		// 포인터 변수의 기본값은 nil이다
		var p *int
		if p != nil {
			// p가 nil이 아니라는 뜻이 아닌, p가 유효한 메모리를 가리킨다는 뜻
			// 즉, nil이란 것은 p가 가리키는 메모리 공간이 유효하지 않다는 것. 정상적이지 않은 공간을 가리킨다는 것
			fmt.Println("not nil")
		} else {
			fmt.Println("nil, ", p)
		}

	}

	{
		// 포인터 변수도 변수다.
		// 다른 변수처럼 메모리 주소를 갖는다. => C나 C++의 이중포인터 느낌..
	}

	{
		// 이런 포인터 왜 쓰냐?
		// => 같은 공간을 다루기 위해
		// => 주소 값만 전달하기 때문에, 주소 값 (8byte만 사용)

		var data Data // 밑에 선언
		ChageData(data)
		fmt.Printf("value = %d\n", data.value)         // 0
		fmt.Printf("data[100] = %d\n", data.data[100]) // 0
		// => 왜 0, 0이 나올까? ChageData로 값 바꿔줬는데...
		// 인자로 넘긴 data를 copy해서 값만 넘겨줬기 때문이다.
		// 만약 인자로 넘어간 값을 변경하고자 한다면 인자를 포인터 변수 타입으로 선언해야 한다.

		ChageData2(&data)
		fmt.Printf("value = %d\n", data.value)         // 999
		fmt.Printf("data[100] = %d\n", data.data[100]) // 999
		// => 데이터가 저장된 메모리 주소를 넘겼기 때문에 copy된 데이터가 아닌 메모리의 데이터를 조작했다.
	}

	{
		//인스턴스란?
		// 메모리가 가리키는 주소에 존재하는 데이터의 실체

		// 그럼 이 경우는 인스턴스가 몇개 인가?

		var data1 Data = Data{3, [200]int{}}
		var data2 Data = data1
		var data3 Data = data1
		// 인스턴스를 copy해서 사용하므로 세개의 변수에 각각 다른 메모리가 할당된다.
		// 즉, 값만 같다. (주소는 다름!)
		fmt.Println(&data1.value, &data1.data[0], &data2.value, &data2.data, &data3.value, &data3.data)

		// 포인터 변수 배열. 포인터 변수 타입을 인덱스로 갖는 배열.
		a := [11]*int{}
		ChangeArr(a)

		// 위에서 처럼 변수 데이터를 변경하기 위해 배열도 주소 자체를 넘겨주어야 함.
		b := [11]int{}
		ChangeArr2(&b)
		fmt.Println(b)
	}

	{
		// new 내장함수
		p1 := &Data{}      // &를 사용해 초기화
		var p2 = new(Data) // new를 사용해 초기화
		// 값을 넣으며 초기화 할거 아니면 new 쓰는게 편함

		fmt.Println(p1, p2)
	}

	{
		// 인스턴스는 언제 사라지는가? => Go의 메모리 관리 방식?
		// 아무도 찾지 않을때 사라진다.

		u := new(Data)
		u.value = 10
		fmt.Println(u)
		// 코드블럭을 종료하면서 stack에 쌓인 내부 변수 u가 사라진다. 이와 동시에 인스턴스를 참조하는 변수가 존재하지 않아, 인스턴스도 같이 사라진다.
		// 파이썬의 reference count가 0이되면 메모리 뺏는거랑 비슷한듯?
		// 다른점이라면 최근의 언어는, 각 언어의 가상머신을 통해 가비지 컬렉션 기능을 사용하지만 Go의 경우 실행파일 안에 가비지 컬렉터가 내장되어 있다.
		// 그래서 Go는 프로그램이 가벼우면서 GC에 의해 생산성이 높은 특징을 가진다.
		// engineering.linecorp.com/ko/blog/detail/342/ 나중에 요거 보자

	}

	{
		// 스택 메모리와 힙메모리
		nData := NewData(10, [200]int{})
		fmt.Println(nData)

	}
}

type Data struct {
	value int
	data  [200]int
}

func NewData(value int, data [200]int) *Data {
	nData := Data{value, data}
	return &nData

	// 함수가 끝나면 변수는 메모리가 해제된다.
	// 그러면 저기서 return한 주소는 메모리가 해제되어 없는 주소인거 아닌가?
	// Dangling Error! => 유효하지 않는 주소를 사용.
	// 근데 왜 잘 되나? (C나 C++ 같은 경우는 이 코드가 오류남)

	// func 내부의 지역 변수는 stack에 저장된다. func이 종료되면 pop을 하면서 메모리에서 삭제함.
	// 그러나 Go에서는 escapte analyzing을 함. (탈출 분석)
	// 컴파일러가 코드를 분석해 어떤 인스턴스가 func 밖으로 탈출하는지 검사한다.
	// 탈출하는걸 확인하면 stack 공간의 메모리를 할당하지 않고, Heap에 할당한다.
	// Heap에서의 메모리는 쓰임이 당하지 않는 이상 삭제되지 않는다.
}

func ChageData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChageData2(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChangeArr(arr [11]*int) {
	a := 10
	var aPoint *int = &a
	arr[10] = aPoint
	fmt.Println(arr)
}

func ChangeArr2(arr *[11]int) {
	arr[10] = 3
}
