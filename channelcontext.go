package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("channel and context let's go!")
	{
		//채널은 고루틴끼리 메세지를 전달할 수 있는 메세지 큐 (FIFO)
		// 이는 Thread-safe queue라고도 불리며, 멀티 쓰레드 환경에서 lock없이도 사용 가능한 큐이다.
		// 쓰레드간의 통신?

		// make()로 채널 인스턴스 생성
		// var message chan string = make(chan string)
		// message : 채널 인스턴스 변수
		// chan string : 채널 타입
		// chan :  채널 키워드
		// string : 메세지 타입

		// 채널에 데이터 넣기
		// message <- "This is a message"
		// 빼기
		// var msg string = <- message (FIFO)

		// 결론적으로 하나의 데이터에 대해 여러 쓰레드에서 접근해도 racecondition 문제가 나지 않도록 알아서 Thread-safe하게 해준다.
	}
	{
		// var wg sync.WaitGroup
		// ch := make(chan int)
		// wg.Add(1)
		// go square(&wg, ch) // channel은 그 자체가 레퍼런스 타입이라 &를 이용해 주소 값으로 바꾸지 않아도 된다.
		// ch <- 9
		// wg.Wait()
	}

	{
		//채널의 크기. default == 0

		// 예상 시나리오 : main 함수가 끝나면서 square2가 실행되지 않고 "is it print?" 출력후 종료된다.
		// 실제 시나리오 : square2의 무한 루프가 계속 실행되고 "is it print?"는 실행되지 않음.
		// 이유는 ch 에 9를 넣었지만, 가져가는 애가 없어서 무한히 대기하기 때문임
		//ch := make(chan int) // 문제
		// ch := make(chan int, 2) // 하지만 이렇게 큐에 사이즈를 줘서 빈공간을 만들어주면 대기하지 않고 종료됨. 즉, 빈공간이 없이 재고가 쌓이면 대기함.
		// go square2()
		// ch <- 9
		// fmt.Println("is it print?")
	}
	{
		//채널에서 데이터 대기
		// ch := make(chan int)
		// var wg sync.WaitGroup
		// wg.Add(1)
		// go square3(&wg, ch) //대기 함수
		// for i := 0; i < 10; i++ {
		// 	ch <- i * 2
		// }
		// // => 무한 대기 함수에서 wg.done으로 넘어가지 않아 메인 고루틴과, square3 고루틴이 deadlock에 빠져버림
		// // 이를 해결하기 위해 close()로 채널을 닫아준다.
		// close(ch) // => 채널에 데이터를 보내줬으니 이제 닫아줌. 좀비 고루틴 방지.
		// wg.Wait()
	}
	{
		// close()로 채널 닫기
		// 위에서의 문제와 같은 go루틴을 좀비 고루틴이라함.
		// 좀비 고루틴 : 채널을 닫지 않아 무한대기 하는 고루틴. 고루틴 릭(leak) 이라고도 함.
	}

	{
		//select문. 여러 채널에서 동시에 데이터를 기다릴때 사용
		// 일정 간격으로 실행. time 패키지의 Tick()은 일정 간격으로 신호를 주는 채널을 반환. After()는 일정 시간 대기후 한번만 신호를 주는 채널 반환.
		// ch := make(chan int)
		// var wg sync.WaitGroup
		// wg.Add(1)
		// go square4(&wg, ch)
		// for i := 0; i < 10; i++ {
		// 	ch <- i * 2
		// }
		// //채널을 닫아주지 않아도 무한 루트를 After를 통해 빠져나옴.
		// wg.Wait()

	}
	{
		// 역할 나누기 방법. deadlock 회피
		// producer, consumer 패턴 구현
		// 하나의 고루틴에서 생산(생성자), 채널에 값 넣기 , 다른 고루틴에서 채널의 데이터 받아서 사용 (소비자)

		// tireCh := make(chan *Car) //타이어 만드는 고루틴에서 다 만들어서 이 채널에 넣고 다른 고루틴에서 이를 받아서 수행하고, 아래의 채널로 넘김. 그럼 아래의 채널은 다른 고루틴이 받아서 처리함.
		// paintCh := make(chan *Car)

		// fmt.Printf("Start Factory\n")

		// wgCar.Add(3)
		// go MakeBody(tireCh)             // 차체 만들어서 tireCh에 car 넘기기
		// go InstallTire(tireCh, paintCh) // tirech에서 car 받아서 tire 만들고 paintch에 넘기기
		// go PaintCar(paintCh)            // paintch에서 car 받아서 pain칠하기

		// wgCar.Wait()
		// fmt.Println("Close the factory")
		// 컨베이어벨트 시스템과 같음.
	}
	{
		//컨텍스트. 작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할
		// 고루틴에다가 명세서를 넘겨 실행하게 하는 듯.
		// wgContext.Add(1)
		// ctx, cancel := context.WithCancel(context.Background()) // ❶ 컨텍스트 생성. 기본 컨텍스트 위에 덮어쓰는 구조. Cancel 할 수 있는 기능을 추가해줌
		// go PrintEverySecond(ctx)
		// time.Sleep(5 * time.Second)
		// cancel() // ❷ 취소

		// wgContext.Wait()
	}
	{
		//컨텍스트로 작업 시간 설정
		// ctx, cancel := context.WithTimeout(context.Background(),3*time.Second) => 3초 뒤에 Done 시그널 알아서 발생

		// 컨텍스트로 특정 값 지정 가능, key-value로 값 가져오는 듯
		wgSquare.Add(1)
		ctx := context.WithValue(context.Background(), "number1", 9) // ❶ 컨텍스트에 값을 추가한다
		// 컨텍스트 랩핑. 여러 데이터를 넣을 수 있음
		ctx = context.WithValue(ctx, "number2", 10)
		ctx = context.WithValue(ctx, "keyword", "Lilly")
		go square5(ctx)
		wgSquare.Wait()

	}
	{
		// pub-sub 패턴. 옵저버 패턴과 굉장히 유사함.
		// 옵저버 : 특정 subject가 각각의 옵저버에게 이벤트 발생시마다 notify하며 알려줌
		// pub-sub도 마찬가지임.
	}

}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch
	time.Sleep(time.Second)
	fmt.Println("square: ", n*n)
	wg.Done()
}

func square2() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}

func square3(wg *sync.WaitGroup, ch chan int) {

	for n := range ch { //무한반복을 돌면서 ch에 데이터가 올때까지 대기함. 데이터가 들어오면 출력하고 다시 대기.
		fmt.Println("Square : ", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func square4(wg *sync.WaitGroup, ch chan int) {

	tick := time.Tick(time.Second)            // 1초에 한번씩 신호를 주는 channel tick
	terminate := time.After(10 * time.Second) // 10초 후에 신호를 주는 channel terminate

	for {
		select { //select에서 여러 case가 동시에 걸리는 경우에는 random하게 case를 실행함. 그래서 출력 결과가 다를 수 있음.(Tick이 매 초마다 나오지 않음)
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square: ", n*n)
			time.Sleep(time.Second)
		}
	}
}

var wgSquare sync.WaitGroup

func square5(ctx context.Context) {
	if v := ctx.Value("number1"); v != nil { // ❷ 컨텍스트에서 값을 읽어온다.
		n := v.(int)
		fmt.Printf("Square:%d \n", n*n)
	}
	if v := ctx.Value("number2"); v != nil { // ❷ 컨텍스트에서 값을 읽어온다.
		n := v.(int)
		fmt.Printf("Square:%d \n", n*n)
	}
	if v := ctx.Value("keyword"); v != nil { // ❷ 컨텍스트에서 값을 읽어온다.
		n := v
		fmt.Println(n)
	}
	wgSquare.Done()
}

type Car struct {
	Body  string
	Tire  string
	Color string
}

var startTime = time.Now()
var wgCar sync.WaitGroup

func MakeBody(tireCh chan *Car) { // ❷ 차체 생산
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			// Make a body
			car := &Car{}
			car.Body = "Sports car" //car struct의 body field만 담당.
			tireCh <- car           // 이후 다 만든 차를 tirech에 넣기
		case <-after: // ❸ 10초 뒤 종료
			close(tireCh)
			wgCar.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) { // ❹ 바퀴 설치
	for car := range tireCh {
		// Make a body
		time.Sleep(time.Second)
		car.Tire = "Winter tire" //위와 마찬가지
		paintCh <- car
	}
	wgCar.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) { // ➎ 도색
	for car := range paintCh {
		// Make a body
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime) // ➏ 경과 시간 출력
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wgCar.Done()
}

var wgContext sync.WaitGroup

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // ❸ 취소 확인. cancel() 확인
			wgContext.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(time.Duration(w.z) * time.Second)
		out <- w
	}
}
