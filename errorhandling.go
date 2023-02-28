package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) //Scanner는 일정(한줄, 한단어)하게 데이터를 가져오기 편함. 인자로는 I/O Reader타입을 받음. str을 읽어오는 Reader객체를 받아서 사용함.
	scanner.Split(bufio.ScanWords)                      //한 단어씩 읽어옴

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil

}

// 다음 단어를 읽어 숫자로 변환해 반환.
// 변환된 숫자, 읽은 글자 수, 에러 메세지를 반환함.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to Scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) //  "24" -> 24 , "abc" -> error
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { //err 인자에서 &numError로 감싸진게 있는지 찾고, 있으면 걔로 변환해서 출력. 없으면 return이 nil이 됨
			fmt.Println("NumberError", numError) //numError.func 하면 어떤 함수 하다가 오류 났는지 알 수 있음.
		}
	}
}
func main() {
	fmt.Println("Error Handling!!")
	{
		//에러 핸들링의 방식 두가지
		/*
			1. 빠르게 프로그램 죽이기
			2. 빠르게 에러 처리해서 프로그램 지속시키기
			=> 개발 단계마다, 프로그램 성격마다 처리방법이 다르다.
		*/
	}

	{
		line, err := ReadFile(filename)
		if err != nil {
			err = WriteFile(filename, "start file")
			if err != nil {
				fmt.Println("파일 생성 실패.", err)
				return
			}
			line, err = ReadFile(filename)
			if err != nil {
				fmt.Println("파일 읽기 실패", err)
				return
			}
		}
		fmt.Println("파일 내용: ", line)
	}
	{
		//사용자 에러 반환
		sqr, err := sqrt(-2)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			print(sqr)
		}

		pwdError := passwordError{1, 2}
		fmt.Println(pwdError.Error())

		error := RegisterAccount("이현규", "ㅁㄴ")
		if error != nil {
			if errInfo, ok := error.(passwordError); ok { //인터페이스 타입변환
				fmt.Printf("%v Len:%d RequireLe: %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
			} else {
				fmt.Println("회원가입 완료")
			}
		}
	}
	{
		//에러 감싸기
		readEq("123 3")
		readEq("123 abc")
		/*
						Failed to readNextInt(), pos:4 err:Failed to convert word to int, word:abc, err:strconv.Atoi: parsing "abc": invalid syntax
			NumberError strconv.Atoi: parsing "abc": invalid syntax
						감싼 에러 안의 에러 안의 에러까지 출력함
						즉, 정확히 에러가 어디서 났는지 알 수 있다.
		*/
		fmt.Println("handling done")
	}

	{
		// 두번째 방법. 패닉. => 처리하기 힘든 에러를 만났을 때, 프로그램을 조기 종료해버림. 빠르게 종료시켜서 오류를 해걸 하기 위해 사용
		//divide(3, 0) //panic: b는 0일 수 없습니다

		// 패닉 생성. func panic(interface{}) => 어느 타입이든 상관 없음
		// 패닉 전파 그리고 복구
		/*
			프로그램 개발시에는 어디서 문제인지 빠르게 알아야 하기 때문에 panic을 자주 사용하지만
			배포 후에는 최대한 안죽는게 중요하기 때문에 panic으로 죽이면 안된다.
			panic을 error로 다 바꾸기 힘들기 때문에 이를 도와주는 복구가 있음/
			main() => f() => g() => h() 인 상황에서 h()에서 패닉이 발생한 경우 역순으로 돌아가며 복구를 확인한다.
			즉 g()에서 복구확인 f()에서 복구확인 main()에서 복구 확인을 함. 이때 main()에서도 복구가 안됐으면 종료된다. (그 전에 기회를 많이 줌)

			복구에는 recover()가 사용된다.
		*/

	}
	{
		// 패닉 복구. func recover() interface{} (패닉에서 사용된 인자(패닉 객체)가 나옴. 같은 빈인터페이스임),
		// defer와 함께 사용됨
		f()
		fmt.Println("program in still running")

		// 복구는 최대한 사용할 일이 없게 해야함. panic이 발생했다는 것이니 이를 꼭 db에 적던가 email로 알려주던가 해야함.
	}

	{
		//Go는 SEG(Structured Error Handling)를 지원하지 않음. try-catch 이런거 말하는거임.
		/*
			왜 안하냐?
			1. 성능 문제. => 오류가 발생하지 않는 상황에도 계속 seh를 지원하기 위해 성능을 사용함.
			2. 에러를 먹어버림. (오히려 에러 처리를 등한시 함) 걍 최상위 Exception으로 다 받아버리고 로그하나 찍고 끝내는 경우가 빈번히 발생함

			!!=> 에러처리는 매우 중요함!
			에러코드에서 반환하는 함수에서 반환되는 에러를 제대로 처리해야 한다! _ 과 같은 빈칸 지시자로 무시하면 안됨. 웬만하면 err != nil로 다 확인하고 문제를 해결해야함.
			에러는 최대한 드러내고 조기에 발견해 더 큰 문제를 미연에 방지해야함
		*/
	}
}

func f() {
	fmt.Println("f() 시작")
	defer func() {
		if r := recover(); r != nil { //함수 맨 마지막에 이 함수가 실행됨. 이 함수가 실행되며 recover()가 실행됨. 그리고 이 실행결과인 r이 존재하면 panic이 있다는 의미이니 함수를 통해 복구 작업을 실행함.
			fmt.Println("panic 복구 - ", r) //main까지 panic이 전파되지 않게 해줌. (main까지 전파되면 프로그램이 죽기 때문)
		}
	}()

	g()
	fmt.Println("f() 끝!")
}
func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))

}

func h(a, b int) int {
	{
		if b == 0 {
			panic("b는 0일 수 없습니다")
		}
		fmt.Printf("%d / %d = %d\n", a, b, a/b)
	}
	return a / b
}

func divide(a, b int) {
	{
		if b == 0 {
			panic("b는 0일 수 없습니다")
		}
		fmt.Printf("%d / %d = %d\n", a, b, a/b)
	}
}

type error interface {
	Error() string //Error() 메소드가 있고, string만 반환하면 무슨 타입이든 에러 타입이 될 수 있음.
}

type passwordError struct {
	Len        int
	RequireLen int
}

func (err passwordError) Error() string { //Error() string만 있으면 error로 쓰일수있다!
	return fmt.Sprintln("암호 길이가 짧습니다. 암호 길이는", err.RequireLen, "이상이어야 합니다.")
}
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// return fmt.Errorf("암호 길이가 짧습니다. 필요한 길이: %d",8) 해도 같은 결과를 뱉지만 매번 작성해주어야 한다.
		return passwordError{len(password), 8} // 그러므로 error 객체를 만들어서 사용하면 편하고, 보다 자세한 작업이 가능하다.
	}
	//아니면 등록~
	return nil
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수이어야 함 f: %g", f)
	}
	return math.Sqrt(f), nil
}

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n') // \n까지 읽어라 => 한줄 읽기
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "error.txt"
