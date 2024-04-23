package receivers

type Person struct {
	name string
	age  int
}

// Function with receiver
func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}
