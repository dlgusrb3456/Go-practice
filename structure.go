package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Structure Let's Go~")
	// 구조체는 여러 필드를 묶어서 사용하는 타입.
	// 특이점은, 구조체 안에 필드만 존재, 매소드는 외부에 따로 정의
	// (go에서의 class인듯 하다)

	// low coupling, high cohension (낮은 결합도, 높은 응집도)
	// => 여러 타입을 하나의 타입으로 묶어, 응집도를 높임
	// 함수도 관련 코드블록을 묶어 응집도와 재사용성을 높임
	// 배열은 같은 타입의 데이터를 묶어 응집도를 높임
	// 구조체는 관련된 데이터(타입이 달라도 됨)를 묶어 응집도와 재사용성을 높임
	{
		// 구조체를 선언하는 간단한 방법
		type Student struct {
			Name  string
			Class int
			No    int
			Score float64
		}

		//구조체 변수 선언 및 초기화 방법1
		var std1 Student
		std1.Name = "Lee"
		std1.Class = 1
		std1.No = 12
		std1.Score = 3.1

		fmt.Println(std1) //{Lee 1 12 3.1}
		fmt.Printf("이름 : %s 반: %d 번호: %d 점수: %0.1f \n", std1.Name, std1.Class, std1.No, std1.Score)
		//이름 : Lee 반: 1 번호: 12 점수: 3.1

		//구조체 변수 선언 및 초기화 방법2
		var std2 Student = Student{"kim", 1, 1, 3.2}
		fmt.Println(std2) //{kim 1 1 3.2}

		//구조체 변수 선언 및 초기화 방법3
		//특정 필드만 초기화. 빈 필드는 default로 채워짐
		var std3 Student = Student{Name: "kim", Class: 1}
		fmt.Println(std3) //{kim 1 0 0}
	}

	{
		//구조체를 포함하는 구조체
		type User struct {
			Name string
			ID   string
			Age  int
		}

		type VIPUser struct {
			UserInfo User //다른 구조체 타입을 포함
			VIPLevel int
			Price    int
		}

		user1 := User{"hg", "lackm", 24}
		vipUser1 := VIPUser{user1, 1, 10}
		fmt.Println(user1)                                 //{hg lackm 24}
		fmt.Println(vipUser1)                              //{{hg lackm 24} 1 10}
		fmt.Println("user Name: ", vipUser1.UserInfo.Name) //user Name:  hg

	}

	{
		//포함된 필드 방식..? (embedded field)
		type User struct {
			Name  string
			ID    string
			Age   int
			Level int
		}

		type VIPUser struct {
			User     //타입만 적어놓음 => 이걸 embedded field 라고 함.
			VIPLevel int
			Price    int
		}

		//왜 하는가?
		user1 := User{"hg", "lackm", 24, 1}
		vipUser1 := VIPUser{user1, 1, 10}
		fmt.Println("user Name: ", vipUser1.Name) //user Name:  hg
		// => 위와 다른점은 vipUser1의 이름에 접근하는데 User를 거치지 않고 감. "."을 여러번 찍지 않는 장점이 있다.
		// 종속의 개념이 아닌 대입의 개념인듯

		//그럼 embedded field의 field명이 구조체의 field 명과 같으면 어떻게 되는가?
		// => 구조체의 field값이 우선됨.
		type User1 struct {
			Name  string
			ID    string
			Age   int
			Level int
		}

		type VIPUser1 struct {
			User1    //타입만 적어놓음 => 이걸 embedded field 라고 함.
			VIPLevel int
			Price    int
			Level    int
		}

		user2 := User1{"hg", "lackm", 24, 1}
		vipUser2 := VIPUser1{user2, 1, 10, 2}

		fmt.Println("Level is ~ ", vipUser2.Level) //Level is ~  2
		// embedded field 내부의 값에 접근하려면 .을 통해 드가야함.
		fmt.Println("Level is ~ ", vipUser2.User1.Level) //Level is ~  1

	}

	{
		// 구조체 크기 : 구조체 내의 모든 필드의 사이즈를 더한 값
		type User struct {
			Age   int     //8바이트
			Score float64 //8바이트
		} // => 16바이트 구조체

	}

	{
		// 구조체 복사 : 모든 필드 값이 복사됨
		// 대입 연산자만 사용해도 필드값이 복사됨
		type User struct {
			Age   int
			Score float64
		}

		user1 := User{1, 1.1}
		user2 := User{2, 2.2}
		user1 = user2
		fmt.Println(user1) // {2 2.2}
	}

	{
		type User struct {
			Age   int32   //4
			Score float64 //8
		} //12

		user := User{21, 1.1}
		fmt.Println(unsafe.Sizeof(user)) // 16...? 위 구조체는 12바이트 크기인데 왜 사이즈가 16바이트로 나올까?
		// => 메모리 정렬. (Memory Alignment)
		// 레지스터에서 값을 가져올때 8바이트 단위로 끊어오기 때문. => 추가로 찾기
		// 즉 4바이트여도 8바이트 크기에 할당하기 때문에 4바이트(int) + 4바이트(빈공간, padding)을 함.
	}

}
