package exinit

import "fmt"

var (
	a = c + b //a 를 정의하기 위해 c,b를 정의함
	b = f() 
	c = f() // c정의를 위해 f()가 실행됨.
	d = 3
)

func init() {
	d++
	fmt.Println("exinit.init function", d)
}

func f() int {
	d++
	fmt.Println("f() d: ", d)
	return d
}

func PrintD() {
	fmt.Println("D: ", d)
}
