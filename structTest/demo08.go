package structTest

import "fmt"

func Run8() {
	t1 := Person{
		Name: "t11",
	}

	t2 := Person{
		Name: "t22",
	}

	t1.Introduce() //Hi, I'm t11
	t2.Introduce() //Hi, I'm t22
	t1.Introduce() //Hi, I'm t11

	fmt.Println(&t1.Name) //0xc0001043e0
	fmt.Println(&t2.Name) //0xc0001043f0
}
