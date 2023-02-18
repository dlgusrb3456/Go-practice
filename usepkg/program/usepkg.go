package main

import (
	"fmt"
	"vscode/for_go/Go-practice/usepkg/custompkg"
	"vscode/for_go/Go-practice/usepkg/exinit"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	//custompkg.printCustom2() => 함수명의 첫글자가 소문자인 경우 외부로 공개되지 않음.
	expkg.PrintSample()

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	exinit.PrintD() 
	/*
		f() d:  4
		f() d:  5  패키지 실행시 먼저 변수를 초기화함. 
		exinit.init function 6 // 이후 init 함수를 실행. (존재한다면)
		D:  6
		단, 여러 패키지에서 이 패키지를 사용해도 매번 호출되는 것이 아닌 최초 1회만 수행된다.
	*/
}
