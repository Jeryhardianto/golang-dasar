// * Struct Method
/**
Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
Method adalah function

*/
package main

import "fmt"

type Customer struct {
	// Short hand	
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string)  {
	fmt.Println("Hello", name,"My name is",customer.Name)
	
}


func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.sayHello("Jery")

}