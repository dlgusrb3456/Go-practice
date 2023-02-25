package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("game start!")

	{
		/* time 패키지!
		Time 객체 : 시각을 나타냄. 			정해진 현재 시각
		Duration 객체 : 시간을 나타냄.  	시각의 흐름. 시각 - 시각을 통한 표현. 100미터를 12초에 뛰었다 느낌
		Location 객체 : 타임존을 나타냄. 	나라별 시차 느낌.
		*/

		loc, _ := time.LoadLocation("Asia/Seoul")
		const longForm = "Jan 2, 2006 at 3:04pm"
		t1, _ := time.ParseInLocation(longForm, "Feb 18, 2023 at 09:51pm", loc) // longform과 같은 형식으로 읽겠다, 실제 시간, 그리고 어느 zone인지
		fmt.Println(t1, t1.Location(), t1.UTC())

		// const YYYYMMDD = "2006-01-02"
		// s := "2022-03-23"
		// t, err := time.Parse(YYYYMMDD, s)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(t)

		const shortForm = "2006-01-02"
		t2, err := time.Parse(shortForm, "2023-02-18") //당장 오늘은 왜 안됨..? 이제 또 되네...
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(t2, t2.Location())
		}

		d := t2.Sub(t1)
		fmt.Println(d)
	}

	{
		// 내가 짠 코드. 이 코드는 int가 아닌 값을 넣는 상황에 대응이 불가능함. input에 대해서 오류시 버퍼를 지워줘야 함.
		/*
			answer := rand.Intn(100) => 이건 매번 값이 같음
			t := time.Now() // random data를 뽑기위한 seed를 설정해줘야 매번 다른 real rand 값이 나옴.
			// 하지만 seed도 매번 바껴야 하므로 time을 seed로 설정하면 seed도 매순간 변함
			rand.Seed(t.UnixNano())
			answer := rand.Intn(100)

			a := 0
			for {
				_, err := fmt.Scanln(&a) //&a => a변수에 할당된 데이터가 담긴 메모리 주소. Go에서는 함수의 인자는 모두 Rvalue로 동작함.
				if err != nil {
					fmt.Println(err)
				} else {
					if a > answer {
						fmt.Println("bigger than answer")
						continue
					}

					if a < answer {
						fmt.Println("smaller than answer")
						continue
					}

					fmt.Println("Correct!")
					break

				}
			}
		*/
	}

	{
		//강의 코드
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(100)
		cnt := 1
		for {
			fmt.Println("숫자 입력 하세요: ")
			n, err := InputIntValue()
			if err != nil {
				fmt.Println("숫자만 입력하세요")
			} else {
				if n > r {
					fmt.Println("입력한 숫자가 더 큼")
				} else if n < r {
					fmt.Println("입력한 숫자가 더 작음")
				} else {
					fmt.Println("정답입니다!")

					break
				}
				cnt++
			}
		}
	}

}

var stdin = bufio.NewReader(os.Stdin) // Input으로부터 값을 읽어올 수 있는, 표준 입력 스트림. 표준입력으로부터 값을 읽어오겠다.

func InputIntValue() (int, error) { // int와 error를 반환받겠다.
	var n int
	_, err := fmt.Scanln(&n) //한줄을 입력받는다. &으로 n의 주소값을 넘겨줌.
	if err != nil {
		stdin.ReadString('\n') //오류시 버퍼에 남아있는 오류문자를 읽어서 오류 부분에 대한 버퍼를 지움.
		// 만약 버퍼를 지우지 않는다면, 잘못 입력된 입력마다 모두 오류를 뱉음. a입력시 a 한번 그리고 \n 한번 해서 총 두번의 오류를 뱉게됨.
	}

	return n, err
}
