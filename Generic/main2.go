package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s)) //s와 길이가 같은 [](슬라이스)를 만들어라
	for i, v := range s {
		rst[i] = f(v) //s를 순회하며 rst[i]에 f(s[i])를 넣어라
	}
	return rst
}

func main() {
	fmt.Println("Generic #2, Generic function")

	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	fmt.Println(doubled)

	uppered := Map([]string{"hello", "world", "green"}, func(v string) string {
		return strings.ToUpper(v)
	})
	fmt.Println(uppered)

	toString := Map([]int{1, 2, 3}, func(v int) string {
		return "str" + strconv.Itoa(v)
	})
	fmt.Println(toString)

	// 여러가지 함수를 정의할 수 있음
}
