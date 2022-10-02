package main

import (
	"fmt"
)

var books = []string{}

func main() {
	//新增
	books = append(books, "01")
	fmt.Println(books)

	//刪除
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice[0:2], slice[3:5]...)
	fmt.Println(slice)

	fmt.Println("==========testPointer01=============")

	testPointer01()

	fmt.Println("==========testPointer02=============")

	testPointer02()
}

func testPointer01() {
	var (
		hello = "Hello World!"
		// 宣告pointer時，型別上要加上*符號，後方reference的變數要加上&符號
		greeting *string = &hello
	)

	fmt.Println(*greeting)
	// => "Hello World!"
	fmt.Println(greeting)
	// => 0x1040a120
	fmt.Println(hello)
	// => "Hello World!"
	fmt.Println(&hello)
	// => 0x1040a120

	//兩變數位址相同
}

func testPointer02() {
	var (
		hello = "Hello World!"
		// 宣告pointer時，型別上要加上*符號，後方reference的變數要加上&符號
		// greeting *string = &hello
		hello2 = hello
	)
	fmt.Println(hello)
	//Hello World!
	fmt.Println(&hello)
	//0xc00010e160

	fmt.Println(hello2)
	//Hello World!
	fmt.Println(&hello2)
	//0xc00010e170

	//兩變數位址不同。與testPointer01測試的結果不一樣。
}
