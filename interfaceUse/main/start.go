package main

import (
	"Go-practice/interfaceUse/fedex" // 하기 전에 mod를 만들어 줘야함. go mod init Go-practice/interfaceUse
	"Go-practice/interfaceUse/koreaPost"
)

func SendBook1(name string, sender *fedex.FedexSender) { // fedex 패키지 내의 FedexSender 타입 구조체와 FedexSender의 메소드를 사용함.
	sender.Send(name) // 메소드 사용
}

func SendBook2(name string, sender *koreaPost.PostSender) { // koreaPost 패키지 내의 PostSender 타입 구조체와 PostSender 메소드를 사용함.
	sender.Send(name) // 메소드 사용
}

// 패키지가 바뀔때마다 SendBook을 정의 해야함...
// ------------------------------------------------------------

type Sender interface { // interface를 통해 번거로움 해결!
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

// -----------------------------------------------------------

func main() {
	{
		// 외부 패키지인 fedex를 import하면 해당 패키지를 통해 Send 기능을 사용할 수 있음을 보여줌.
		sender := &fedex.FedexSender{}
		SendBook1("어린 왕자", sender)
	}

	{
		// 새로운 외부 패키지를 사용하게 되는 경우 기존의 작업을 모두 다시 한번 진행해야한다는 번거로움이 있음. (기존 코드의 수정이 필요)
		// 만약 외부 패키지 사용을 빈번히 바꿔야 하는 상황이 온다면? => 매번 코드를 싸그리 갈아 엎어? goto 16 line
		sender := &koreaPost.PostSender{}
		SendBook2("어린 왕자", sender)
	}

	{
		sender1 := &fedex.FedexSender{}    // 외부 패키지에서 이미 Send() 함수를 정의했음
		sender2 := &koreaPost.PostSender{} // 이렇게 변수와 패키지만 바꿔주면 됨.
		SendBook("cuty!", sender1)
		SendBook("sexy!", sender2)

		// 패키지의 반환 타입이 다르더라도 함수가 같으면 이렇게 가능. (근데 외부 패키지들이 서로 함수 명을 같게 해주려나... (우리가 만든 외부 패키지면 ㄱㄴ인데 우리꺼는 우리가 수정할테고..))
	}

}
