package main

import (
	"fmt"
)

func main() {
	fmt.Println("solid!")
	{
		/*
			SOLID란? 객체지향 설계 5가지 원칙의 약자
			1. 단일 책임 원칙 (single responsibility principle, SRP)
			2. 개방-폐쇄 원칙 (open-closed principle, OCP)
			3. 리스코프 치환 원칙 (liskov substitution principle, LCP)
			4. 인터페이스 분리 원칙 (Interface segregation principle, ISP)
			5. 의존 관계 역전 원칙 (dependency inversion principle, DIP)
			=> oop (객체 중심 프로그래밍)가 되기 위한 원칙 solid
			=> oop를 만족 하기 위한 방법인 상속, 추상화 등의 개념을 어떤 방향으로 사용해야 할지에 대한 내용이 solid임
			oop중심 프로그래밍을 하며 solid를 지향해야함.
		*/
	}
	{
		/*
			좋지 않은 설계
			1. 경직성(rigidity) : 모듈 간의 결합도가 너무 높아 코드를 변경하기 매우 어려운 구조. 때론 모듈간 의존 관계가 거미줄처럼 얽혀 있어 어디부터 손대야 할지 모를 정도로 복잡한 구조
			2. 부서지기 쉬움(fragility) : 한 부분을 건드렸더니 다른 부분까지 망가지는 경우. 변경이 힘듦
			3. 부동성 (immobility) : 코드 일부분을 분리해 다른 곳에서도 사용하고 싶지만, 모듈 간 결합도가 높아 옮길 수 없는 경우. 재사용성이 떨어지므로 매번 새로 구현하는 문제가 생김
			=> 셋다 코드간의 결합도(high coupling & low cohension)가 너무 높아서 생기는 문제인듯.
			=> 즉, SOLID는 결합도를 낮추고 응집도를 높이는 바를 지향함. (high cohension low coupling)
		*/
	}
	{
		/* 단일책임 원칙(single responsibility principle) SRP
		정의 : "모든 객체는 책임을 하나만 져야 한다"
		이점 : "코드의 재사용성을 높여줌"

		경제리포트 생성, 경제리포트 전송 / 사회리포트 생성, 사회리포트 전송 ... => 이런식으로 모든 리포트마다 따로 만들면 재사용성이 떨어짐
		리포트 인터페이스 생성, 리포트전송 함수 생성, 리포트 인터페이스 메소드 구현한 각 리포트 타입 생성 => 재사용성이 좋음 & 전송은 리포트를 전송만, 인터페이스는 생성만 담당함. => 각각의 기능이 완전히 분리됨
		*/
		type FinanceReport struct{ //보고서
			report string
		}

		func (r *FinanceReport) SendReport(email string){ //보고서 전송
			fmt.Println("send",email)
		}

		type MarketingReport struct{ //보고서
			report string
		}

		func (r *MarketingReport) SendReport(email string){ //보고서 전송 => 보고서를 전송하는 기능은 동일하지만 따로 만들어줘야 하는 불편함이 존재함.
			fmt.Println("send",email)
		}

		// To solve this problem. 보고서와 보고서 전송을 분리해야함.
		type Report interface{
			Report() string
		}

		type FinanceReport struct{
			report string
		}
		func (r *FinanceReport) Report() string{
			return r.report
		}
		type ReportSender struct{

		}

		func (s *ReportSender) SendReport(report Report){

		}
	}
	{
		/*
			개방-폐쇄 원칙(open-closed principle) OCP
			정의 : 확장(새로운 기능 추가)에는 열려(쉽게 추가 가능) 있고, 변경(기존 코드 변경)에는 닫혀(하기 힘듦) 있다. => 기존 코드 변경 없이 새로운 기능 추가가 용이함.
			이점 : 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 됨.

			위에서 리포트 전송 함수를 다음과 같이 구현한다
			func SendReport(r *Report, method SendType, receiver string){
				switch method{
				case Email:
				case Fax:
					...
				}
			}
			=> 보통 이런식으로 많이 구현함. 이 경우 새로운 타입이 추가될 때 마다 저 코드를 수정해야함.

			이를 OCP에 맞게 구현하기 위해선 각각의 경우를 따로 구현해야함. 여기서도 interface를 통해 구현함

			type ReportSender interface{
				Send(r *Report)
			}

			type EmailSender struct{

			}
			func (e *EmailSender) Send(r *Report){
				// 이메일 전송 구현
			}
			type FaxSender struct{

			}
			func (f *FaxSender) Send(r *Report){
				// 팩스 전송 구현
			}

			구현 이후 switch-case 문에서 methodType 확인후, 이 인터페이스를 활용해 concrete class 사용

			=> 아예 기존 코드를 안건드리는 것은 아님. 조금은 수정해야 하는데 그저 항목 추가 개념정도? 로 조금 수정함
		*/
	}
	{
		/*
			리스코프 치환 원칙(liskov substitution principle) LCP
			정의 : "q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야 한다."
				=> 상위 타입에서 동작하는 함수라면 하위 타입에서도 오류 없이 동작해야 한다.
			이점 : 예상치 못한 작동을 예방할 수 있음

			(다른 언어로 예시)
			class Rectangle{
				width int
				height int
				setWitdh(w int) {width = w}
				setHeight(h int) {height = h}
			}

			class Square extends Rectangle{
				@override
				setWidth(w int) {width = w ; height = w;}
				@override
				setHeight(h int) {height = h; width = h;}
			}

			func FillScreenWidth(screenSize Rectangle, imageSize *Rectangle){
				if imageSize.width < screenSize.width{
					imageSize.setWidth(screenSize.width)
				}
			}
			=> imageSize *Rectangle의 Rectangle이 Square라면 width만 바꾸고 싶은 이 함수가 height까지 바뀌어 버리는 문제가 생김. (func의 기능 자체가 바뀜)
			=> 근데 Go에는 상속 없어서 LSP 위반이 잘 일어나지 않지만, 일어나는 경우가 있음.

			Go에서의 위반 사례. 인터페이스 타입 변환시 문제 발생 가능. interface => concrete object. (지양하자는거임)
			type Report interface{
				Report() string
			}
			type MarketingReport{

			}
			func (m *MarketingReport) Report() string{

			}

			func SendReport(r Report){
				if _,ok := r.(*MarketingReport); ok{ //r이 마케팅 보고서일 경우 패닉
					panic("Can't send MarketingReport")
				}
			}

			var report = &MarketingReport{}
			SendReport(report) // 패닉 발생
		*/
	}
	{
		/*
			인터페이스 분리 원칙 (Interface segregation principle, ISP)
			정의 : "클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 함"
			이점 : 인터페이스를 분리하면 불필요한 메서드들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있음

			위반사례
			type Report interface{
				Report() string
				Pages() int
				Author() string
			}

			func SendReport(r Report){
				send(r.Report())
			}
			=> Report에는 3가지 메소드가 있지만 SendReport는 한가지만 사용함. 즉, 나머지 2개는 불필요한 메소드임
			=> 사용하지 않는 불필요한것에 의존할 필요 없음

			바꾸면
			type Report interface{
				Report() string
			}

			type WrittenInfo interface{
				Pages() int
				Author() string
				WrittenDate() time.Time
			}

			func SendReport(r Report){
				send(r.Report())
			}

			이런식으로 인터페이스를 분리해, 불필요한 의존 관계를 없엔다.

		*/
	}

	{
		/*
			의존 관계 역전 원칙 (dependency inversion principle, DIP)
			정의 : "상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다."
			원칙 :
				1. "상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다."
				2. "추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다."

			[키보드] -> [전송] -> [네트워크] 이런 구체화된 모듈간이 직접적인 관계보다는
			[키보드] -> [<<interface>> 입력] -> [전송] -> [<<interface>> 출력] -> [네트워크] 와 같이 추상화된 모듈과 관계를 맺자.
			=> 입력 모듈이 늘어나도 확장이 편함. 출력 모듈이 늘어나도 확장이 편함.
			=> SRP에서 이미 DIP를 활용함.

			원칙 2 : 구체화된 모듈간의 직접적인 연관을 맺지 말아야 함. [concrete] - [concrete] => [concrete] - [interface] - [interface] - [concrete]
		*/
	}
	{
		//DIP 구현
		var mail = &Mail{}
		var listener EventListener

		listener = &Alarm{}

		mail.Register(listener)
		mail.OnRecv()
	}
}

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("메일이 왔습니다")
}
