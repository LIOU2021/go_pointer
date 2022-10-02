package structTest

func Run8() {
	t1 := Person{
		Name: "t11",
	}

	t2 := Person{
		Name: "t22",
	}

	t1.Introduce() //Hi, I'm t11
	t2.Introduce() //Hi, I'm t22
}
