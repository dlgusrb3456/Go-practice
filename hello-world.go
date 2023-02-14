// 언제나 코드는 package로 시작해야 한다.
// 패키지 == 코드를 묶는 단위임. 이 코드가 어떤 패키지에 속하는지 알 수 있음.
// 아무렇게나 패키지 명을 쓸 수 있지만, main은 사전에 정의된 의미가 있다.
// main은 프로그램 시작점을 의미한다. cpu가 읽어올 명령어의 시작점 (중간에 시작도 ㄱㄴ)
// 결국 go는 하나의 main 패키지와 다른 package들로 이루어져 있다.
package main

// import == 가져오다
// ex) fmt == fmt 패키지를
// "import fmt" == "fmt 패키지의 코드를 이 package에 가져와 사용하겠다"
import "fmt"

// 함수명도 아무거나 상관 없지만 main은 위와 마찬가지로 사전에 정의된 의미가 있다.
// "이 함수부터 시작해라" 라는 의미를 지녔고, package main은 즉 main() 함수를 갖고 있는 패키지를 뜻한다.
// main()이 시작하면 프로그램이 시작되고 끝나면 프로그램이 끝난다.
func main() {
	fmt.Println("Hello world!")
}
