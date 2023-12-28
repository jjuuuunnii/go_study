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
	Enumeration2()

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
