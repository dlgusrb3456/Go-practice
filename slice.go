package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slice let's go!")
	// 슬라이스 : Go에서 제공하는 동적 배열 타입.

	{
		/*
			var slice []int
			if len(slice) == 0 {
				fmt.Println("slice is empty", slice)
			}

			slice[1] = 10 // 여기서 오류가 발생함. 현재 slice는 비어있기 때문
			fmt.Println(slice)
		*/
	}

	{
		var slice []int = []int{1, 2, 3} // 위와 달리 이렇게 몇개의 값을 초기화 한 후에 해당 인덱스에 접근하면 ㄱㅊ다.
		if len(slice) == 0 {
			fmt.Println("slice is empty", slice)
		}

		slice[1] = 10 // 여기서 오류가 발생허자 얺움
		fmt.Println(slice)
	}

	{
		// 슬라이스 초기화
		var slice1 = []int{1, 2, 3}
		var slice2 = []int{1, 5: 2, 10: 3}
		fmt.Println(slice1, slice2)

		// Make를 이용한 초기화
		var slice3 = make([]int, 3) // => 요소 개수 3개를 갖는 슬라이스. 3개는 기본 값인 0으로 초기화 됨
		fmt.Println(slice3)
	}
	{
		// 슬라이스 순회
		var slice = []int{1, 2, 3}
		for i := 0; i < len(slice); i++ {
			slice[i] += 10
		}
		fmt.Println(slice)
		for i, v := range slice {
			slice[i] = v * 2
		}

		fmt.Println(slice)
	}

	{
		// 요소 추가 - append()
		var slice = []int{1, 2, 3}
		slice2 := append(slice, 4) // 기존의 slice에서 4를 추가해 반환해준다.

		fmt.Println(slice2)
		//여러개추가도 가능함! 4,5,6,7,8 이런 식으로!
	}

	{
		// 이런 슬라이스의 동작 원리.
		/*

			type SliceHeader struct{
				Data uintptr 	// 실제 배열을 가리키는 포인터
				Len int			// 요소 개수
				Cap int			// 실제 배열 길이 (이놈이 String과 다름) 최대로 몇개까지 쓸 수 있느냐를 알려줌.
			}
		*/

		// 즉, 슬라이스는 Go에서 제공하는 배열을 가리키는 포인터 타입임. 단 배열의 길이가 바뀌고 어디까지 사용할 수 있는지 표시함.
		// => 함수 인자로 arr 넘길시 포인터를 넘기지 않으면 값이 변하지 않지만, slice의 경우는 값이 변경됨. (애초에 포인터 타입이니까)
	}

	{
		// append 동작 원리
		// append는 슬라이스에 요소를 추가한 새로운 슬라이스를 반환...
		// 기존 슬라이스가 바뀔 수도 있고, 아닐수도 있다.
		// append는 새로운 배열을 만들고 append할 배열의 값을 copy 후 뒤에 값을 추가한다. 이후 이 새로 만든 배열의 struct를 만들어서 반환한다.
		// 위의 방식대로라면 언제나 새롭운 슬라이스가 생기는거 아닌가?

		// 동작원리를 살펴보면 그렇지 않다.
		// 우선 요소를 추가할 빈 공간이 충분한지 확인한다.
		// 충분하다면 빈 공간에 요소를 추가해 반환한다.
		// 충분하지 않다면 새로운 배열을 할당하고, 복사, 요소 추가의 단계를 거친다.

		// 그럼 빈공간은 어케 아는가?
		// 슬라이스 struct의 cap - len을 한 값이 빈 공간이다. 이 빈공간에 append하는 요소가 들어가는지 확인하는 것.

		// 그래서 새로 할당받은 slice라면 값을 변경해도 새로운 slice이기 때문에 이전의 slice에 서로 독립적으로 작동한다.
		// 하지만 기존의 slice를 이어서 사용한다면 서로 같은 주소를 가리켜 상호간에 영향을 미친다.
	}

	{
		//흔히 하는 실수
		slice := []int{1, 2, 3}
		addNum(slice)
		fmt.Println(slice) // 기존 값이 그대로 나옴... 왜? 슬라이스는 포인터 타입 아닌가?
		// 저 함수 안에서 append를 하면서 새로운 슬라이스를 할당하고 그 값에 append를 했기 때문에 기존의 값은 변동이 없음..
		// 이를 해
		fmt.Printf("slice memory: %p \n", slice)
		addNum2(&slice)

		fmt.Println("addNum2:", slice)
		fmt.Printf("slice memory: %p \n", slice)
		// 혹은 새로운 slice를 반환한다.
		slice = addNum3(slice) // struct가 보내짐. data, len, cap 24 바이트. => 주소만 넘기는(8바이트만 넘기는) 것보다 효율이 안좋아진다고 느껴지겠지만 사실상 별 차이 없음
		// 그래서 보통 이렇게 쓰는걸 선호함. 이렇게 쓰는게 slice를 slice답게 쓰는 것임. slice는 값타입이기 때문에 걍 값으로 써라...
		fmt.Println(slice) // 적용 잘 됨.
	}

	{
		// 슬라이싱. 배열의 일부를 가리키는 기능. 슬라이싱의 결과가 슬라이스임
		// array[stasrtidx:endIdx] stasrt ~ end-1까지
		// 이 결과인 slice는 arr의 start를 시작 주소로 가리키고 len을 start~end-1까지의 개수로, cap은 전체 길이 - startIdx로 정해진다.
		arr := [5]int{1, 2, 3, 4, 5}
		slice := arr[1:2]
		fmt.Println("arr: ", arr)
		fmt.Println("slice: ", slice, len(slice), cap(slice)) //[2] 1 4

		arr[1] = 100
		fmt.Println("arr: ", arr)
		fmt.Println("slice: ", slice, len(slice), cap(slice)) //[100] 1 4 => 새로운 배열을 할당받는 것이 아니라, 기존의 배열의 인덱스 주소를 가리키는 struct이므로 값이 영향을 미침

		slice = append(slice, 500)
		fmt.Println("arr: ", arr)                             //[1 100 500 4 5] => 얘도 바뀜... 아직 cap이 충분해서 append를 해도 새로운 배열을 할당받지 않기 때문에 얘도 영향을 받았음.
		fmt.Println("slice: ", slice, len(slice), cap(slice)) //[100 500] 2 4
	}

	{
		// 슬라이스도 슬라이스 할 수 있음
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := slice1[1:2]
		fmt.Println(slice1)
		fmt.Println(slice2) // [2], len = 1 , cap = 4
	}

	{
		array := [100]int{1: 1, 2: 2, 99: 100}
		slice1 := array[1:10]
		slice2 := slice1[2:99] //  slice1의 len은 작지만 slice가 가리키는게 array이기 때문에 가능함. (끝을 정확하게 지정해주는 것이 좋은듯..)
		fmt.Println(slice1)
		fmt.Println(slice2)

		// 혹은! slice에 새로운 인자를 추가함. slice[시작 인덱스: 끝인덱스: 최대인덱스] => 위에서 말한 오류를 줄일 수 잇음. 즉, 최대인덱스 == cap이 됨
	}

	{
		//유용한 슬라이싱 기능 - 복사
		// 슬라이스를 그대로 복사해 새로운 슬라이스를 만들어 두 슬라이스가 서로 독립적이게 함
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := slice1[:] // => 이건 복사가 아님... 서로 영향을 미침
		slice1[0] = 0
		fmt.Println(slice1, "  ", slice2)

		// 그럼 복사를 어케 하냐?

		// 1. 반복문을 통해 함.
		for i, v := range slice1 {
			slice2[i] = v
		} // => 하나씩 값 복사
		slice1[0] = 1
		fmt.Println(slice1, "  ", slice2)

		// 2. append 사용
		slice3 := append([]int{}, slice1...) // 비어있는 slice에 append를 진행함. slice1...은 슬라이스 전체를 의미함. 빈슬라이스이기 때문에 cap이 0이라 새로운 배열을 할당해줘 서로 독립적인 슬라이스가 됨.
		slice3[0] = 9
		fmt.Println(slice1, "  ", slice3)

		// 3. make & copy 사용
		slice4 := make([]int, len(slice1))
		copy(slice4, slice1)
		slice4[0] = -1
		fmt.Println(slice1, "  ", slice4)

	}

	{
		// 요소 삭제!!
		// 증긴 요소 날리면, 그 뒤에 있는 값들을 앞으로 한칸씩 당기고 마지막 값은 삭제해버림.
		slice := []int{1, 2, 3, 4, 5}
		idx := 2

		// 방법 1.
		// for i := idx + 1; i < len(slice); i++ {
		// 	slice[i-1] = slice[i]
		// }
		// slice = slice[:len(slice)-1]

		// 방법2.
		slice = append(slice[:idx], slice[idx+1:]...) // 딱 인덱스만 빼고 됨

		// 방법3.
		copy(slice[idx:], slice[idx+1:])
		fmt.Println("slice", slice)
	}

	{
		//요소 삽입. 1. 맨뒤에 요소 추가 2. 하나씩 뒤로 복사 (넣길 원하는 idx까지) 3. 값 수정
		slice := []int{1, 2, 3, 4, 5}
		slice = append(slice, 0) // 1. 맨 뒤에 요소 추가 (아무 값이나)
		idx := 2

		for i := len(slice) - 2; i >= idx; i-- { //2. 하나씩 뒤로 복사
			slice[i+1] = slice[i]
		}

		slice[idx] = 100 // 3. 하나씩 뒤로 당겨서 빈 공간을 만들어둔 곳에 값 수정.
		fmt.Println(slice)

		// 한줄로 하는 방벙
		slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...) // 한줄로 가능하지만 임시 버퍼를 할당해야 해서 위의 for문에 비해 효율적이지 않다.

		// copy를 이용해 불필요한 메모리 할당 없에기
		slice = append(slice, 0)         // 1. 뒤에 요소 추가
		copy(slice[idx+1:], slice[idx:]) // 2. 값 복사
		slice[idx] = 100                 // 3. 값 변경
	}

	{
		// 슬라이스 정렬. sort라는 패키지가 존재함.
		slice := []int{3, 1, 2, 6, 3, 67, 8, 9, 100}
		sort.Ints(slice)
		fmt.Println(slice) //[1 2 3 3 6 8 9 67 100]
		// 이외에도 다양한 타입에 대한 sort를 지원해줌.
	}
	{
		//구조체 슬라이스 정렬. 구조체의 특정 필드 값을 기준으로 정렬 하고 싶다면?

		s := []Student{{"lee", 2}, {"kim", 3}, {"lim", 4}, {"hwang", 1}}
		// sort.Sort(s) => Sort()는 interface를 인자로 받기 때문에 오류가 난다. method와 인터페이스 부분에서 확인.
		// interface는 len, less, swap을 갖고있는 타입만 됨. 밑에 func를 통해 정의함.
		sort.Sort(Students(s))
		fmt.Println(s)

	}

}

type Student struct {
	Name string
	Age  int
}
type Students []Student // 구조체 슬라이스

func (s Students) Len() int {
	return len(s)
}
func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func addNum(slice []int) {
	slice[0] = 10            // 이건 적용됨
	slice = append(slice, 4) // 이건 안됨
}

func addNum2(slice *[]int) {
	// 이건 적용됨
	*slice = append(*slice, 4, 5, 6) // 이건 됨
}

func addNum3(slice []int) []int { // 값을 리턴해 버리는 방법
	slice[0] = 10
	slice = append(slice, 4)
	return slice
}
