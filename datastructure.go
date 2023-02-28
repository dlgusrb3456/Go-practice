package main

import (
	"container/list" //리스트 패키지
	"container/ring" //링 패키지
	"fmt"
)

func main() {
	fmt.Println("Data Structure!")
	{
		//리스트. 링크드리스트임.
		/*
			type Element struct{
				Value interface{}  //빈 인터페이스로 무슨 값이든 올 수 있음.
				Next *Element
				Prev *Element
			}
		*/
		v := list.New()
		e4 := v.PushBack(4)
		e1 := v.PushBack(1)
		v.InsertBefore(3, e4) // e4 전에 3을 넣어라
		v.InsertAfter(2, e1)  // e1 뒤에 2를 넣어라
		v.InsertAfter("asdf", e1)
		for e := v.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
		}
		fmt.Println()
		for e := v.Back(); e != nil; e = e.Prev() {
			fmt.Print(e.Value, " ")
		}
		fmt.Println()

	}
	{
		//데이터 지역성.
		// 데이터가 인접해 있을수록 캐시 성공률이 올라가고 성능도 증가한다.
		//보통 데이터를 가져갈때 근처의 값을 한번에 가져감.
		//연속된 메모리 주소를 사용하는 배열의 경우 캐시 성공률이 높을 수 있다. (요소 수가 적은 경우 리스트보다 배열이 빠름. 1000개까진 무난하게 배열이 빠르고 10000개부터는 고민해보는게..)
		// 리스트는 메모리가 떨어져있어서 캐시 fault확률이 높은거임.
	}

	{
		//큐 구현
		que := NewQueue()
		for i := 0; i < 5; i++ {
			que.Push(i)
		}
		value := que.Pop()
		for value != nil {
			fmt.Printf("value: %v ", value)
			value = que.Pop()
		}
		fmt.Println()
	}
	{
		//스택 구현
		stk := NewStack()
		for i := 0; i < 5; i++ {
			stk.PushStack(i)
		}
		value := stk.PopStack()
		for value != nil {
			fmt.Printf("value: %v ", value)
			value = stk.PopStack()
		}
		fmt.Println()
	}
	{
		//링구조
		r := ring.New(5)
		n := r.Len()
		for i := 0; i < n; i++ {
			r.Value = 'A' + i
			r = r.Next()
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%c ", r.Value)
			r = r.Next()
		}
		fmt.Println()
		for i := 0; i < n; i++ {
			fmt.Printf("%c ", r.Value)
			r = r.Prev()
		}
		fmt.Println()
		// 이런 링은 언제 사용하나?
		// 실행 취소 기능: 문서 편집기 등에서 일정한 개수의 명령을 저장하고 실행 취소 하는 경우
		// 고정 크기 버퍼: 데이터에 따라 버퍼가 증가되지 않고 고정된 길이로 쓸 때
		// 리플레이 기능 : 고정된 길이의 리플레이 기능 제공
	}

	{
		// map , 딕셔너리, 해시맵 ... 등으로 부르는!! (키-값 구조로 데이터를 저장하는 자료구조)
		// map[key]value 형식으로 사용함. map[string]int : 키타입은 string이고 value타입은 int이다.
		m := make(map[string]string)
		m["Lee"] = "HyunGyu"
		fmt.Println(m["Lee"])                   //HyunGyu
		m1 := make(map[interface{}]interface{}) //빈 인터페이스를 사용하면 어느 타입이든 사용 가능함.
		m1[123] = "asdf"
		m1["asdf"] = 123
		fmt.Println(m1[123], m1["asdf"])

		//맵 순회
		for key, value := range m1 {
			fmt.Println("key: ", key, "value: ", value)
		}

		person := make(map[int]People)
		person[1] = People{"Lee", 24}
		person[2] = People{"H", 22}
		person[3] = People{"L", 23}
		for _, v := range person {
			fmt.Println("people: ", v) //순회시 순서 보장이 되진 않음.
		}

		//특정 키 삭제. delete(map,key)
		delete(person, 1)
		//특정 키 존재 여부
		v, is := person[2] //is == true or false로 반환됨
		if is == true {
			fmt.Println("value:!!", v)
		} else {
			fmt.Println("no data")
		}
	}
	{
		//해쉬 함수 동작의 이해
		// 해쉬란 잘게 부순다는 뜻. 해쉬 브라운 == 잘게 부순 감자
		// 해쉬함수는 1. 같은 입력이 들어오면 같은 결과가 나옴. 2. 다른 입력이 들어오면 되도록 다른 결과가 나옴. 3. 입력값의 범위는 무한대이고 결과는 특정 범위를 갖음.
	}
	{
		//map의 두 종류
		// 1. hashmap, 2. sortedmap
		// hashmap의 경우 언제나 O(1)을 보장함.
	}
}

type People struct {
	Name string
	Age  int
}

type Queue struct {
	v *list.List
}

type Stack struct {
	v *list.List
}

func (stk *Stack) PushStack(val interface{}) {
	stk.v.PushBack(val)
}

func (stk *Stack) PopStack() interface{} {
	back := stk.v.Back()
	if back != nil {
		return stk.v.Remove(back)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	que := list.New()
	return &Queue{que} //구조체 필드 값 대입
	//return &Queue{ list.New() } 위와 같음.
}
