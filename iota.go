package main

import (
	"fmt"
)

const PI = 3.14              // 상수는 타입을 정하지 않아도 됨
const FloatPI float64 = 3.14 // 얘는 타입을 정함

func constTypeCheck() {
	var a int = PI * 100 // 얘는 오류가 발생하지 않음 => 타입을 지정하지 않은 상수는 사용될때 타입을 정함
	// 하지만 PI가 3.145와 같이 100을 곱해도 int가 되지 않으면 오류를 뱉음.
	//var b int = FloatPI * 100 // 얘는 타입을 정했기 때문에 오류남

	fmt.Println(a)
}

// iota 활용. 반복문을 사용하지 않아도 순차적인 값을 상수에 저장할 수 있음.
const (
	MasterRoom uint8 = 1 << iota // 자동으로 밑에 애들도 좌쉬프트 연산을 해주는 값이 저장됨
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room // 비트 or연산
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room // 비트 clear 연산자
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room // 비트 and 연산자
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불키기")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불키기")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불키기")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불키기")
	}

}

// rooms = 00000000 의 비트라 생각
// 각 비트가 각 방의 불 켜짐 여부라고 보면 됨
// 변수 사용의 수가 훨씬 줄어들어 메모리 사용을 적게 한다는 장점이 있음
func main() {
	var rooms uint8 = 0

	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)
	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)

	constTypeCheck()
}
