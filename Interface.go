// * Interface

/**
Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
Sebuah interface berisikan definisi-definisi method
Biasanya interface digunakan sebagai kontrak

*/
package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(hasName HasName){
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

//Exp 2
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"
	SayHello(eko)

	cat := Animal{Name: "Kucing"}
	SayHello(cat)

}