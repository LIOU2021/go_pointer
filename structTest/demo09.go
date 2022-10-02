package structTest

import "fmt"

func Run9() {
	var (
		hello = "Hello World!"
		// 宣告pointer時，型別上要加上*符號，後方reference的變數要加上&符號
		// greeting *string = &hello
		greeting = &hello
		v2       = 22
		t2       = &v2
	)

	fmt.Println(greeting)
	// => 0xc000044460
	fmt.Println(&hello)
	// => 0xc000044460

	fmt.Println("greeting-------")

	fmt.Println(&v2)
	//0xc000016170
	fmt.Println(t2)
	//0xc000016170

	fmt.Println("t2-------")

	var v3 = 33
	var t3 = &v3
	fmt.Println(&v3)
	//0xc000016178
	fmt.Println(t3)
	//0xc000016178

	fmt.Println("t3-------")

	var v4 = 11
	var t4 *int = &v4

	fmt.Println(&v4) //0xc000016180
	fmt.Println(t4)  //0xc000016180

	fmt.Println("t4-------")

	v := 11 //唯獨這種v的宣告方式取出來的位址，後面參考的變數的每個都不一樣位址
	var p1 *int = &v
	var p2 *int = &v

	//底下這種宣告也是一樣結果
	// var p1 = &v
	// var p2 = &v

	fmt.Println(&v)  //0xc000016170
	fmt.Println(&p1) //0xc00000a048
	fmt.Println(&p2) //0xc00000a050
	fmt.Println("p1-------")

	var p3 = &v
	var p4 = &v
	*p4 = 1
	fmt.Println(v)   //1
	fmt.Println(&p3) //0xc00000a058
	fmt.Println(&p4) //0xc00000a060
	fmt.Println("p4-------")
}
