// * Struct dalam OOP itu sama seperti CLASS
/**
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field
Sederhananya struct adalah kumpulan dari field

Struct adalah template data atau prototype data
Struct tidak bisa langsung digunakan
Namun kita bisa membuat data/object dari struct yang telah kita buat

*/
package main

import "fmt"

type Customer struct {
	/**
	Name String
	Address String
	Age int
	*/
	// Short hand	
	Name, Address string
	Age int
}

func main() {

	// Recoment
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Jalan Raya"
	eko.Age = 10
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	// Struct literal
	// Recoment
	jery := Customer{
		Name:"Jery", 
		Address:"Jalan Raya",
		Age: 10,
	}
	fmt.Println(jery)

	//Ini tidak remomended lah karena bergantung pada posisi field struct diatas
	budi := Customer{"Budi", "Jalan Raya", 10}
	fmt.Println(budi)


}