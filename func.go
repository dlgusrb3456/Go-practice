package main

import (
	"fmt"
)

// 함수의 위치는 상관 없음
func Add(a int, b int) int { // Add 라는 함수를 만들고 인자로는 a int타입, b int타입을 받고 반환값으로는 int 타입을 반환할 것이다.
	return a + b
}

// 함수를 왜 쓸까..
// 반복 작업이 싫어잇!! 자주 사용되는걸 템플릿처럼 만들어서 반복 사용하자!
// JMP라는 명령어를 만들게 됨. 50번째 줄에서 함수 호출시 함수가 존재하는 곳으로 IP(Instruction Point)를 옮겨줘야함. 이때 JMP 사용 (순차적으로 실행하지 않으니까)
// 함수가 끝난 후 다시 IP를 51번째 줄로 JMP함

// 근데 함수가 존재하는 곳은 고정이지만 함수가 끝나고 JMP할 곳은 매번 바뀜..
// 이 문제를 해결하기 위해 함수로 JMP 하기 전, 메모리에 돌아올 IP를 미리 적어둠 RP(Return Point)

// 근데 또 이 함수안에서 쓰이는 데이터가 매번 달라.. 그래서 RP처럼 인자도 미리 JMP 뛰기 전에 메모리에 copy해둔다
// 근데!! 또 이 값들이 너무 많아지면? 메모리에 데이터가 분산되있어 접근이 힘듦. 그래서 그냥 RP가 저장된 메모리 뒷주소에 인자를 연속적으로 붙임
// 함수에서는 이 주소를 끝에서부터 읽어서 인자의 수 만큼 값을 읽어들임. 그리고 돌아갈때는 마지막 남은 RP로 돌아감.

// 멀티 반환 함수
func Divide(a int, b int) (int, bool) { //반환값이 여러개인 경우
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// 멀티 반환 함수의 반환 값에 이름 미리 붙이기
func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return //return시 값 지정 필요 없음
	}
	result = a / b
	success = true
	return
}

// 재귀호출 recursive
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	printNo(10)

}

/*
	함수에 인자로 넘기는 모든 값은 call by value로 넘어간다. 즉 Rvalue로 넘어간다.
	그래서 만약 array를 넘긴다면 array를 이루는 구조체를 넘기는 것이고, slice를 넘긴다면 sclice의 구조체를 넘기는 것이다.
*/
