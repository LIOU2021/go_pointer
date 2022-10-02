package structTest

import "fmt"

func Run10() {
	var p1 = Person{
		Name: "p11",
	}

	var p2 = Person{
		Name: "p22",
	}

	p1.Introduce() //Hi, I'm p11
	p2.Introduce() //Hi, I'm p22
	p1.Introduce() //Hi, I'm p11

	fmt.Println(&p1.Name) //0xc000044470
	fmt.Println(&p2.Name) //0xc000044480
}
