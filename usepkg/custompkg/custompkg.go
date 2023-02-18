package custompkg // main이 아니므로 보조 패키지임
import "fmt"

func PrintCustom() {
	fmt.Println("This is sub package code1")
}

func printCustom2() { // 소문자로 사용했기 때문에 외부에 공개되지 않음.
	fmt.Println("This is sub package code2")
}
