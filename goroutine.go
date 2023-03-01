package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("고루틴!")

	{
		// 쓰레드는 실행 흐름과 같다. 초기 개발에서는 단 하나의 실행 흐름만 존재했다.
		// 시간이 지남에 따라 멀티 쓰레드가 등장. => 하나의 실행 흐름이 아닌 여러 실행 흐름이 가능해짐. (코어가 빠르게 쓰레드를 교체함) 이게 어떻게 가능한가?
		// 쓰레드마다 본인의 IP포인트가 있음. cpu가 빠르게 쓰레드를 교체하며 각각의 쓰레드가 독립적으로 돌아가게 함.
		// cpu에게 명령을 내리는 os가 이것을 하는 멋진놈임.
	}
	{
		// 그럼 고루틴은 뭘까? : 고에서 만든 경량 쓰레드임.
		// 메인 함수도 고루틴임 - 메인 고루틴
		// 새로운 고루틴은 단순히 go 함수_호출로 만듦.

		// go PrintHangul()
		// go PrintNumber()

		/*
			총 세개의 고루틴이 생성됨
			1. 메인 루틴
			2. PrintHangul()
			3. PrintNumber()
			이 세개의 고루틴이 서로 다른 흐름을 갖고 실행되고 있음
		*/

		time.Sleep(3 * time.Second) //3초간 sleep. 만약 sleep을 하지 않으면 main이 종료된다. main의 종료는 프로그램의 종료를 의미하므로 위의 두 go루틴이 실행중이지만 프로그램이 종료된다.
		// 그럼 go루틴 사용할때마다 매번 초를 계산해서 기다려줘야 하는 번거로운 작업을 해줘야 하나?
		// 이를 도와주는 것이 있음. 서브 고루틴이 종료될때까지 기다리는 sync.WaitGroup

	}
	{
		// wg.Add(10) //10개의 고루틴 작업을 기다릴 것임. (Done받을때마다 하나씩 숫자가 감소하고 이게 0이되면 wait를 빠져나감 )
		// for i := 0; i < 10; i++ {
		// 	go SumAtoB(1, 1000000000)
		// }

		// wg.Wait() // 10개의 wg 작업이 다 끝날때까지 대기. 메인함수가 종료되지 않음.
	}

	{
		// 이런 고루틴의 동작원리.
		// 고루틴은 OS 쓰레드를 이용하는 경량 쓰레드임. 즉 쓰레드 != 고루틴임
		// 고루틴은 쓰레드를 이용함
		// 코어1 - OS 쓰레드 1 - 고루틴 1 ... 이런 구조로 고루틴이 실행되며 1대1 매칭된다. 즉 코어가 6개라면 동시에 돌릴 수 있는 고루틴은 6개뿐이라는 것이다.
		// 그럼 8개의 고루틴을 실행하면 나머지 2개의 고루틴은 어디두느냐? 얘네는 대기상태에 들어가 대기한다.
		// 그리고 실행중인 고루틴들에서 시스템콜이나 I/O 요청 네트워크 요청 등에 의해 대기상태로 들어가면, 기존의 대기하고 있던 고루틴이 실행되며 고립상태를 벗어난다. (교체하는 것임)

		// 이게 좋은 점은 OS 단위에서의 context 스위칭이 일어나지 않기 때문이다. 하나의 코어당 하나의 OS 쓰레드만 사용하므로 각 코어가 쓰레드간의 스위칭을 할 이유가 없다.
		// => 오해하면 안되는 것은 OS 단위에서의 context 스위칭이 없는 것이지, OS쓰레드에 붙은 go루틴들은 위에서의 설명대로 교체 작업이 일어나므로 여기서는 context switching이 일어난다.
		// => 하지만 스위칭의 비용이 OS 단위에서의 스위칭보다 훨씬 적다(경량화 했기 때문에 switching 할 내용이 적음). 그러므로 다른 것에 비해 장점이 있다.

		//Go 쓰레드 안에 Local Run Queue가 존재함. 여기에 대기 큐가 순차적으로 쌓이고 순차적으로 실행된다.
	}

	{
		// 동시성 프로그래밍의 주의점.
		// => 동일한 메모리 자원을 여러 고루틴에서 접근할 때 동시성 문제가 발생함. (race condition 문제) heap메로리의 자원 사용할때 조심할 것.
		// var wg1 sync.WaitGroup
		// account := &Account{10}
		// wg1.Add(10)
		// for i := 0; i < 10; i++ {
		// 	go func() {
		// 		for {
		// 			DepositAndWithdraw(account) //10개의 고루틴이 무한하게 돌아감. 하지만 panic이 일어나며 종료됨. 이유는 balance라고 하는 자원에 동시에 접근하기 때문임. race condition (너가 아는 그거 맞음 )
		// 			// 이를 해결하기 위해선 고루틴에서 각각 하나의 메모리 자원만 접근 해야함. 혹은 Lock 사용(Mutex Lock)
		// 		}
		// 		wg1.Done()
		// 	}()
		// }

		// wg1.Wait()

		/*
			str := "Alice"
			go func(name string) {
				fmt.Println("Your name is", name)
			}(str)

			Is same as:

			str := "Alice"
			f := func(name string) {
				fmt.Println("Your name is", name)
			}
			go f(str)
		*/
	}
	{

		// rand.Seed(time.Now().UnixNano())

		// wg.Add(2)
		// fork := &sync.Mutex{}
		// spoon := &sync.Mutex{}

		// go diningProblem("A", fork, spoon, "fork", "spoon") // 순서를 똑같이 spoon, fork로 하면 deadlock이 발생하지 않음. 잘기다리니까
		// go diningProblem("B", spoon, fork, "spoon", "fork")
		// wg.Wait()

		/*
			A 이 fork 을 획득함.
			B 이 spoon 을 획득함.
			fatal error: all goroutines are asleep - deadlock!
		*/
	}

	{
		// 영역을 나눠서 race condition 문제 피하기
		var jobList [10]Job
		for i := 0; i < 10; i++ {
			jobList[i] = &SquareJob{i}
		}

		wg.Add(10)
		for i := 0; i < 10; i++ {
			job := jobList[i] // 서로 다른 메모리에 접근 함.
			go func() {
				job.Do()
				wg.Done()
			}()
		}
		wg.Wait()
	}

}

type Job interface {
	Do()
}
type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작 \n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료\n", j.index)
}

func diningProblem(name string, m1 *sync.Mutex, m2 *sync.Mutex, mName1 string, mName2 string) {
	for i := 0; i < 100; i++ {
		fmt.Println(name, "이 밥을 먹는다")
		m1.Lock()
		fmt.Println(name, "이", mName1, "을 획득함.")
		m2.Lock()
		fmt.Println(name, "이", mName2, "을 획득함.")

		fmt.Println(name, "이 밥을 먹는중")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		m2.Unlock()
		m1.Unlock()
	}
	wg.Done()
}

var mutex sync.Mutex

/*
	뮤텍스의 문제점
	1. 동시성 프로그래밍으로 인한 성능 향상을 얻을 수 없다. (어차피 임계영역에 접근하지 못하기 때문) 심지어 과도한 lock은 성능이 하락되기도 한다.
	2. Deadlock 문제가 발생할 수 있다. Deadlock 발생시 고루틴은 완전히 멈춘다. lock 2개가 필요한데 각 고루틴이 한개씩 들고 놔주지 않아 영원히 실행되지 않고 죽는 경우와 같음.
	결국 mutex를 아주 조심히 사용해야 함.
	혹은 또 다른 자원관리 기법을 사용한다.
		1. 영역을 나누는 방법
		2. 역할을 나누는 방법
*/

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {

	mutex.Lock()         // mutex lock을 통해 임계구역을 정함
	defer mutex.Unlock() // defer를 통해 함수 종료 직전 임계구역에 대해 unlock함

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value : %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(1 * time.Millisecond)
	account.Balance -= 1000
}

var wg sync.WaitGroup // wait group 생성

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d to %d sum is %d \n", a, b, sum)
	wg.Done() // wait group에게 해당 func이 끝났다고 알림
}

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마'} // rune 타입. rune is an alias for int32 and is equivalent to int32 in all ways. it is used, by convention, to distinguish character values from integer values
	// UTF-8 인코딩 방식을 사용함. 한글은 3byte 사용하고 영어는 1byte 사용함.
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond) //0.3초
		fmt.Printf("%c ", v)
	}
}

func PrintNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
