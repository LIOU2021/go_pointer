package structTest

type Mary struct {
	*Person
	Power int
}

func Run4() {
	mary := &Mary{
		Person: &Person{"mary"},
	}

	mary.Introduce()

}
