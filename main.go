package main

import (
	"example/test/structTest"
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

	fmt.Println("==========testPointer03=============")
	testPointer03()

	fmt.Println("==========testPointer04=============")
	testPointer04()

	fmt.Println("==========testPointer05=============")
	testPointer05()

	fmt.Println("==========testPointer06=============")
	testPointer06()

	fmt.Println("==========testPointer07=============")
	testPointer07()

	fmt.Println("==========testPointer08=============")
	testPointer08()

	fmt.Println("==========testPointer09=============")
	testPointer09()
}

func testPointer01() {
	var (
		hello = "Hello World!"
		// 宣告pointer時，型別上要加上*符號，後方reference的變數要加上&符號
		// greeting *string = &hello
		greeting = &hello
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

	*greeting = "changed value"
	fmt.Println(hello)
	// => "changed value"
}

func testPointer02() {
	var (
		hello  = "Hello World!"
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

	hello2 = "changed value"
	fmt.Println(hello)
	//Hello World!
	//沒有改變值，因為兩個變數位址不一樣
}

func testPointer03() {

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	//42

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i
	//21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	//73

}

func testPointer04() {
	jim1 := person{
		firstName: "Jim1",
	}

	jim2 := person{
		firstName: "Jim2",
	}

	jim1.updateNameFromPointer("Aaron") // It works as expected
	fmt.Println("01: ", jim1.firstName)

	jim2.updateName("Aaron") // It doesn't work as expected
	fmt.Println("02: ", jim2.firstName)
}

type person struct {
	firstName string
}

func (p *person) updateNameFromPointer(newFirstName string) {
	// *variable 表示把該指摽對應的值取出
	(*p).firstName = newFirstName
	//golang允許直接簡略如下
	// p.firstName = newFirstName
}

// 當沒有使用 *type 時
// 每次傳進來的 p 都會是複製一份新的（by value）
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func testPointer05() {
	structTest.Run()
}

func testPointer06() {
	structTest.Run2()
}

func testPointer07() {
	structTest.Run3()
}

func testPointer08() {
	structTest.Run4()
}

func testPointer09() {
	structTest.Run5()
}
