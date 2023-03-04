package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func print[T any](a T) { // [T any] => 이 함수에서 사용할 Generic Type T를 any(타입 제한자)로(모든 타입이 가능함) 정의함
	fmt.Println(a)
}

// func min2(a, b interface{}) interface{} { //=> 빈 인터페이스는 연산자를 지원해주지 않음
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func min3[T any](a, b T) T { // 이 또한 오류가 남. T에서 받을 수 있는 타입들이 모두다 연산자 계산이 가능하지 않기 때문임. 그래서 any가 아닌 다른 타입 제한자를 줘야함.
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func min4[T int | int8 | int16 | int32 | float32 | float64 | int64](a, b T) T { // 타입 제한자로 준 타입들이 모두 대소 연산이 가능하므로 오류가 발생하지 않음
	if a < b {
		return a
	}
	return b
}

// 근데 매번 저 타입을 다 쓰기 귀찮으니 저걸 하나의 타입 제한자로 선언해서 사용하자
type Integers interface {
	int | int8 | int16 | int32 | int64
}

type Floats interface {
	float32 | float64
}

func min5[T Integers | Floats](a, b T) T { // 커스텀 타입 제한자 사용
	if a < b {
		return a
	}
	return b
}

// 이런 타입을 사용하기 편하게 Go에서 만들어준 패키지가 있음
// "golang.org/x/exp/constraints" , constraints 패키지로, 여러개 많음 https://pkg.go.dev/golang.org/x/exp/constraints 여기서 확인
// 보면 하나 특이한점이 있음. ~int 이렇게 타입을 씀. 이건 int를 타입으로 하는 별칭타입까지 모두 포함한다는 의미임.
func min6[T constraints.Ordered](a, b T) T { // 커스텀 타입 제한자 사용
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println("Generic In Go")
	{
		a := 10
		b := 20
		fmt.Println(min(a, b))

		// var c int16 = 10
		// var d int16 = 20
		// fmt.Println(min(c,d)) // 자동으로 타입 변환을 해주지 않음.

		// 즉 같은 기능을 하더라도 타입별로 함수를 만들어줘야 하는 번거로움이 있음
		// Generic을 사용해 이런 번거로움을 없에자!
	}
	{
		var a int = 10
		print(a)
		var b float32 = 3.14
		print(b)
		var c string = "Hello"
		print(c)
		// => 모두 정상적으로 동작함.
		// 빈 인터페이스를 사용해도 가능한데 왜 Generic 을 사용해야 함..?
		// 위의 min2() 함수처럼 빈인터페이스에서 지원해주지 않는 여러 기능이 존재함 (연산자 비교)
		// 하지만 min3()처럼 T 타입을 any로 주는 경우도 안됨.
		// min4()처럼 해당 함수 안에서 사용하는 타입 연산이 타입 제한자의 모든 타입이 사용 가능해야 오류가 발생하지 않음.
		// 즉, 타입 제한자를 잘 줘야함.
	}

}
