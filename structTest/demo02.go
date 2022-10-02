package structTest

type mark struct {
	Person
	Power int
}

func Run2() {
	mark1 := mark{
		Person: Person{"Mark"},
	}

	mark2 := mark{
		Person: Person{"Mark2"},
	}

	mark1.Introduce()
	mark2.Introduce()
}
