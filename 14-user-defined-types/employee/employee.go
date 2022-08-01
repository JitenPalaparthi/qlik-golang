package employee

import "fmt"

type Employee struct {
	ID      int
	Name    string
	Email   string
	Address // promoted type. Can access directly through Employee object.. e.City etc
}

type Address struct {
	Line1, Street, City, PinCode, State, Country string
}

func (e *Employee) ShowDetails() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Email:", e.Email)
	fmt.Println("Address Line1:", e.Line1)
	fmt.Println("Address Street:", e.Street)
	fmt.Println("Address City:", e.City)
	fmt.Println("Address State:", e.State)
	fmt.Println("Address PinCode:", e.PinCode)
	fmt.Println("Address Country:", e.Country)
}

type Person struct {
	int
	string
	bool
}

func NewPerson(id int, name string, isMarried bool) *Person {
	return &Person{int: id, string: name, bool: isMarried}
}

func (p *Person) ShowDetails() {
	fmt.Println("ID:", p.int)
	fmt.Println("Name:", p.string)
	fmt.Println("IsMarried:", p.bool)
}

type Empty struct {
}

func (e Empty) Greet() {
	fmt.Println("Hello World!")
}
