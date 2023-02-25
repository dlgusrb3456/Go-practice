package main

import (
	"fmt"
	"math/rand"

	// "text/template" // 만약 이렇게 패키지 명이 겹친다면?
	// "html/template"
	htemplate "html/template"
	ttemplate "text/template" // 별칭을 붙여서 해결!
)

func main() {
	fmt.Println("module & package let's go")
	// 모듈이란 패키지의 모음으로, 한 개의 모듈은 다수의 패키지를 포함할 수 있음.
	// 모듈은 패키지 관리 시스템으로 활용됨. 집합관계 : module > package

	// 그럼 패키지란? 코드를 묶는 단위
	// 모든 코드는 반드시 패키지로 묶어야 한다.

	// 프로그램 == 실행파일. 실행 시작 지점을 포함한 패키지. 즉 main함수를 포함한 main 패키지.
	// main 패키지가 아닌 그외 패키지 : 실행 시작 지점을 포함하지 않은 패키지로 프로그램의 보조 패키지로 동작.

	{
		// math/rand 패키지!
		fmt.Println(rand.Int()) //5829807356788033392 => random data
	}

	{
		// 패키지 명이 겹치는 경우
		ttemplate.New("foo").Parse(`{{define "T"}}Hello`) //별칭 사용
		htemplate.New("foo").Parse(`{{define "T"}}Hello`)
	}

	{
		// 따라하기
		/*
			1. Go-practice/usepkg 폴더 생성
			2. go mod init Go-practice/usepkg
			3. mkdir custompkg
			4. Custompkg.go
			5. mkdir program
			6. usepkg.go
			7. go mod tidy  => 외부 패키지를 다운받음. go.sum이라는 폴더에 명시됨. 다운 받은 패키지는 Go Path에 저장됨
			8. go build

			Go에서는 폴더를 기준으로 패키지를 나눔. 하나의 폴더 아래에서 다른 패키지를 사용할 수 없음
			그래서 같은 폴더 아래에서는 패키지를 공유함.
		*/
	}

	{
		// 패키지 외부 공개
		// 변수명, 구조체명, 함수명, 전역변수 등 패키지 내의 것들의 이름의 첫 글자가 대문자라면 외부 공개가 됨다.
		// 반면 소문자로 시작시 외부공개가 되지 않는다.
		// 또 구조체 필드명도 마찬가지임. (특정 필드를 감출 수 있음)
	}
	{
		//패키지 초기화
		// 패키지가 프로그램에 포함되어 초기화 될때. 패키지 내 init()이 한번만 호출된다.
		// init()을 통해서 패키지 내 전역 변수를 초기화 한다.
		// usepkg/exinit 폴더에서 확인
	}

}
