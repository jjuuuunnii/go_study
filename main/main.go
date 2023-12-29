package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hello")

	//Variable2()
	//Variable3()
	//Const1()
	//Enumeration1()
	//Enumeration2()

	//if1()
	if2()
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
