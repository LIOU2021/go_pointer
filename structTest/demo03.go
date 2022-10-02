package structTest

type Jay struct {
	Person
	Power int
}

func Run3() {
	jay := Jay{
		Person: Person{"Jay"},
	}

	//不宣告struct，直接建立並且賦值
	jayTest := struct {
		Person
		Power int
	}{
		Person: Person{"JayTest"},
		Power:  100,
	}

	jay.Introduce()     //Hi, I'm Jay
	jayTest.Introduce() //Hi, I'm JayTest
}
