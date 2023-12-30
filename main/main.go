package main

import (
	"fmt"
	"go_study/main/lib"
	"go_study/main/lib2"
	"math/rand"
	"os"
	"sort"
	"strings"
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
	//initTest()

	//typeSection()
	//stringTest()
	//arrayTest2()
	//slice2()
	//slice3()
	//slice4()

	mapUseAge2()
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

func typeSection() {
	var b1 bool = true
	//var str1 string = "안녕하세요."
	//var str1 string = `c:\go_study\src\go_study\main\main.go`
	if &b1 != nil {
		fmt.Println("b1: ", b1)
	}

	//fmt.Println("str1 length: ", utf8.RuneCountInString(str1))

	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("str1: ", rune(str1[0]), str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("str2: ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("str3: ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	for i, r := range str1 {
		fmt.Println("index: ", i, "value: ", string(r))

	}

	constStr := []rune(str3)
	for i, r := range constStr {
		fmt.Printf("index: %d, value: %d\n", i, r)
	}
}

func stringTest() {
	var str1 string = "The Go Programming Language"
	var str2 string = "is an open source programming language"
	var str3 string = str1 + " " + str2

	fmt.Println(str3)
	var strSet []string
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println(strings.Join(strSet, " "))
}

func arrayTest() {

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 [5]string = [5]string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1: ", arr1[i])
	}

	for index, value := range arr1 {
		fmt.Print("\nindex: ", index, " value: ", value)
	}

	for _, value := range arr2 {
		fmt.Print("\nvalue: ", value)
	}

	for index := range arr2 {
		fmt.Print("\nvalue: ", index)
	}

}

func arrayTest2() {
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Printf(" arr1: %p, arr2: %p\n", &arr1, &arr2)
	fmt.Printf(" arr1: %p, %v, arr2: %p, %v\n", &arr1, arr1, &arr2, arr2)
}

func slice1() {
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	slice3[4] = 10
	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	var slice5 []int = make([]int, 5, 10)
	slice6 := make([]int, 5)
	slice7 := make([]int, 5, 10)
	slice8 := make([]int, 5)

	slice6[2] = 7
	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

}

func slice2() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	fmt.Printf(" arr1: %p, %v\n", &arr1, arr1)
	fmt.Printf(" arr2: %p, %v\n", &arr2, arr2)

	println()
	slice1 := make([]int, 5, 10)
	slice2 := slice1

	slice1[2] = 7

	fmt.Printf("slice1: %p, %v", slice1, slice1)
	println()
	fmt.Printf("slice2: %p, %v", slice2, slice2)
}

func slice3() {
	arr := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		arr = append(arr, i)
		fmt.Printf("len: %d, cap: %d, %v\n", len(arr), cap(arr), arr)
	}

	//만약 용량에서 벗어나면 용량의 2배를 한 만큼 용량이 늘어난다. -> 용량을 처음 할당할때 주의!
}

func slice4() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice3 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice1[:])
	slice2 := []int{3, 6, 10, 9, 2, 7, 5, 8, 4, 1}
	fmt.Println(sort.IntsAreSorted(slice2))
	sort.Ints(slice2)
	fmt.Println(slice2)

	c := [5]int{1, 2, 3, 4, 5}
	var d []int = make([]int, 5, 10)
	d = c[:]
	d = append(d, 0) // 인덱스 5
	d = append(d, 7) // 이제 인덱스 6에 7을 추가

	c[0] = 100
	fmt.Println(c)
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]
	fmt.Println("e: ", len(e), cap(e), e)
	fmt.Println("f: ", len(f), cap(f), f)

	f = append(f, 11, 12, 13)
	fmt.Println("e: ", len(e), cap(e), e)
	fmt.Println("f: ", len(f), cap(f), f)
}

func mapUseAge() {

	map5 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	for key, value := range map5 {
		fmt.Println("key: ", key, "value: ", value)
	}

}

func mapUseAge2() {
	map1 := map[string]string{
		"naver":  "https://www.naver.com",
		"daum":   "https://www.daum.net",
		"google": "https://www.google.com",
		"yahoo":  "",
	}

	value1, isExist := map1["naver"]
	value2, isExist2 := map1["daum"]
	value3, isExist3 := map1["google"]
	value4, isExist4 := map1["kakao"]
	value5, isExist5 := map1["yahoo"]
	fmt.Println("value1: ", value1, "isExist: ", isExist)
	fmt.Println("value2: ", value2, "isExist: ", isExist2)
	fmt.Println("value3: ", value3, "isExist: ", isExist3)
	fmt.Println("value4: ", value4, "isExist: ", isExist4)
	fmt.Println("value5: ", value5, "isExist: ", isExist5)

	if value1, isExist := map1["naver"]; isExist {
		fmt.Println("value1: ", value1, "isExist: ", isExist)
	}
}
