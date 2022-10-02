package structTest

func Run7() {
	var p1 *Person = &Person{
		Name: "p11",
	}

	var p2 *Person = &Person{
		Name: "p22",
	}

	p1.Introduce() //Hi, I'm p11
	p2.Introduce() //Hi, I'm p22
}
