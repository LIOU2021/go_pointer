package main

import "fmt"

var books = []string{}

func main() {
	//新增
	books = append(books, "01")
	fmt.Println(books)

	//刪除
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice[0:2], slice[3:5]...)
	fmt.Println(slice)

	fmt.Println("=======================")

	testPointer01()
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
}
