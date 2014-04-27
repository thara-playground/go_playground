package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

type Named interface {
	Name() string
}

type Persan struct {
	FirstName string
	LastName string
}

func printName(named Named) {
	fmt.Println(named.Name())
}

func main() {
	person := &Person{"Tomochika", "Hara"}
	named, ok := interface{}(person).(Named)
	if ok {
		printName(named)
	} else {
		fmt.Println("Person is not Named interface")
	}

	persan := &Persan{"Tomochika", "Hara"}
	named, ok = interface{}(persan).(Named)
	if ok {
		printName(named)
	} else {
		fmt.Println("persan is not Named interface")
	}
}
