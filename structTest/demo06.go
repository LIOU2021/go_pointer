package structTest

import "fmt"

type Vertex struct {
	X, Y float64
}

func Run6() {
	v := new(Vertex) // new 回傳的是指稱到該變數的 pointer
	v.X = 10
	fmt.Println(v)  //&{10 0}
	fmt.Println(*v) //{10 0}
	fmt.Println(&v) // 0xc00000a040
}
