package main

import (
	"fmt"
	"go_study/main/lib"
	"go_study/main/lib2"
	"math/rand"
	"os"
)

/**
 * init 함수 실행 시 import 된 패키지의 init 함수가 먼저 실행된다.
 * init 함수는 패키지가 로드되면서 가장 먼저 실행되는 함수이다.
 */
func main() {
	//fmt.Println("hello")

	//Variable2()
	//Variable3()
	//Const1()
	//Enumeration1()
	//Enumeration2()

	//if1()
	//if2()
	//random1()

	//for1()

	//packageUseage()
	//packageUseage2()
	initTest()
}

func initTest() {
	fmt.Println("initTest start")
	fmt.Println("initTest end")

}

func Variable2() {
	var (
		name   string  = "jeong"
		age    int     = 1234567891011121314
		height float32 = 175.5
	)

	fmt.Println("name: ", name, "age: ", age, "height: ", height)
}

func Variable3() {
	shortVar1 := 3
	shortVar2 := "test"
	shortVar3 := false

	fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)
	}
}

func Const1() {
	const (
		name  string = "kang seung jun"
		age   int    = 25
		house string = "관악드림타운"
	)

	fmt.Println("name: ", name, "age: ", age, "house: ", house)
}

func Enumeration1() {
	const (
		JAN int = 1
		FEB int = 2
		MAR int = 3
	)

	fmt.Println("Jan: ", JAN, "FEB: ", FEB, "MAR: ", MAR)
}

func Enumeration2() {
	const (
		FIRST = iota + 1
		SECOND
		THIRD
	)

	fmt.Println(FIRST)
	fmt.Println(SECOND)
	fmt.Println(THIRD)
}

func Enumeration3() {
	const (
		_ = iota + 1
		_
		THREE
		FOUR
	)
}

func if1() {
	//if 문은 반드시 true로 검사하기
	//소괄호를 사용하지 않는다

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("a는 15이상이다")
	}

	if b < 15 {
		fmt.Println("b는 15이하이다")

	}

	/*	//에러 발생 1
		if b > 15
		{
			fmt.Println("~~")
		}*/

	/*	if error2 := 1; error2=1{
			//에러 발생 2
			에러입니다!!!
		}
	*/
	// 짧은 선언 후 쓰기
	if c := 1; c == 1 {
		fmt.Println("c는 1 입니다")
	}

}

func if2() {
	var a int = 50
	b := 70

	if a > 65 {
		fmt.Println("65 이상")
	} else if a < 30 && b > 40 {
		fmt.Println("30 이하")
	} else {
		fmt.Println("50 입니다.")
	}

	/*
		else 도 마찬가지로

		else
		{
		}

		이렇게 하면 에러가 생긴다.

		go는 하나의 문장 끝마다 ;
		를 붙이기때문이다.
	*/
}

func switch1() {
	//제어문(조건문)
	a := -7

	switch {
	case a < 0:
		fmt.Println("a는 음수")
	case a > 0:
		fmt.Println("a는 양수")
	}

	switch b := 7; b {
	case 0:
		fmt.Println("b는 음수")
	case 1:
		fmt.Println("b는 1")
	default:
		fmt.Println("b는 7")
	}

}

func random1() {
	var randomInt = rand.Intn(100)

	fmt.Println("This is random int = ", randomInt)

	switch e := "go"; e {
	case "go":
		fmt.Println("this is go! ")

	}

}

func for1() {
	loc := []string{"seoul", "busan", "incheon"}
	for index, name := range loc {
		fmt.Println("for1 case = ", index, name)

	}

	sum := 0

	for i := 0; i <= 100; i++ {
		sum += 1

		fmt.Println("this is index: i", i)
		fmt.Println("this is sum: ", sum)

	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i, "*", j, "=", i*j, "\t")

		}
		fmt.Println()

	}

Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 && i == 4 {
				break Loop1
			}
			fmt.Println("i :", i, "j: ", j)

		}
	}

}

func packageUseage() {

	var name string

	fmt.Println("이름은? : ")

	_, _ = fmt.Scanf("%s", &name)

	fprintf, err := fmt.Fprintf(os.Stdout, "Hi %s\n", name)
	if err != nil {
		return
	}

	fmt.Println("fprintf : ", fprintf)

	//패키지 : 코드 구조화 및 재사용
	//응집도, 결합도
	//Go : 패키지 단위의 독립이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성할 것을 권고
	//패키지 이름 = 디렉토리 이름
	//같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	//네이밍 규칙 : 소문자 private, 대문자 public
	//Go : main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 x, 프로그램의 시작점 start point
	//패키지 : 패키지 레벨 변수, 함수, 메소드 등 식별자 사용을 제한 -> 구조체, 인터페이스 등은 제한 x
	//패키지 별칭 : alias
	//패키지 임베디드 : 패키지를 익명으로 선언 후 사용 가능
	//패키지 접근제어 : 소문자 private, 대문자 public

}

func packageUseage2() {
	fmt.Println("9보다 큰 수", lib.CheckNum1(10))
	fmt.Println("1000보다 큰 수", lib2.CheckNum1(1001))

}

func init() {
	fmt.Println("init start")
	fmt.Println("init end")
}
