package main

import (
	//"github.com/stretchr/testify/assert"
	"testing"
)

/*
	실행파일에서 이 테스트 코드는 실행되지 않음.
	터미널에서 실행 : go test
	일부만 실행 : go test -run TestSquare2

	테스트 도움을 주는 패키지 : "github.com/stretchr/testify/assert"
	패키지 다운 : go get github.com/stretchr/testify/assert, or go mod tidy로 필요한 외부 패키지 다운.
*/

/*
	TDD (Test-Driven development)
	코드구현보다 먼저 테스트 코드를 작성.
	장점
	1. 테스트 케이스가 자연스럽게 늘어남
	2. 테스트가 촘촘해짐
	3. 자연스러운 회기 테스트가 가능
	4. 리팩토링이 쉬워짐
	5. 개발이 즐거워짐
	6. 코드 커버리지가 자연히 증가함

	단점
	1. 모듈간 의존성이 높은 경우 테스트 케이스를 만들기 힘듦
	2. 동시성 테스트에 취약함
	3. 진정한 TDD가 아닌 형식적인 테스트로 전락할 수 있음
	4. 지속적인 모니터링과 관리가 필요함
*/

/*
	벤치마크: 성능 검사시 사용
	1. 파일명이 _test.go로 끝나야 함
	2. testing 패키지를 임포트해야 함
	3. 벤치마크 코드는 func BenchmarkXxxx(b *testing.B) 형태이어야 함
*/

// func TestSquare1(t *testing.T) {
// 	assert.Equal(t, 81, square(9), "square(9) shoud be 81") // 아래 코드를 간단하게 만들 수 있음, (t, 결과, 입력, 오류메세지)
// 	// rst := square(9)
// 	// if rst != 81 {
// 	// 	t.Errorf("square(9) shoud be 81 but returns %d", rst)
// 	// }
// }

// func TestSquare2(t *testing.T) {
// 	rst := square(3)
// 	if rst != 9 {
// 		t.Errorf("square(9) shoud be 9 but returns %d", rst)
// 	}
// }

func BenchmarkFibo1(b *testing.B) { // BenchmarkFibo1-8           39200(동일 시간동안 이만큼 코드를 돌렸다)            31535 ns/op (한 operation당 이만큼의 시간이 걸렸다.)
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}
func BenchmarkFibo2(b *testing.B) { // BenchmarkFibo2-8        91419516                11.01 ns/op
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
