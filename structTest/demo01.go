package structTest

import (
	"fmt"
)

// STEP 1：建立 Person struct 與其 Method
type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// STEP 2：建立 Saiyan struct，並將 Person embed 在內
// 意思是 Saiyan 是 Person，而不是 Saiyan「有一個」Person
type Saiyan struct {
	*Person
	Power int
}

func Run() {
	// STEP 3：建立 goku
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}

	// STEP 4：可以直接使用 goku.Name，也可以使用 goku.Person.Name
	fmt.Println(goku.Name)        // Goku
	fmt.Println(goku.Person.Name) // Goku

	// STEP 5：方法在使用時也一樣
	goku.Introduce()        // Hi, I'm Goku
	goku.Person.Introduce() // Hi, I'm Goku

}
