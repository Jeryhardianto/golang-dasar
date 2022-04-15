// * Pointer
/**
=> Pass by Value
Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

=> Pointer
Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
  => Operator & (AND)
	Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya
  => Operator * (STAR)
    Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
	Semua variable yang mengacu ke data yang sama tidak akan berubah
	Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *


*/
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

 //*Pointer di function
  func changeAddress(address *Address) {
    address.City = "Malinau"
  }

  //*Pointer method
  type Man struct {
    Name string
  }

  func (man *Man) Marride(){
    man.Name = "Mr. " + man.Name
    fmt.Println("Dimethod ", man.Name)
  }

func main() {
   //* Pass by Value
   address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
   address2 := address1

   address2.City = "Bantul"

   fmt.Println(address1) //{Sleman Yogyakarta Indonesia}
   fmt.Println(address2) //{Bantul Yogyakarta Indonesia}

   //* Pass by Reference || pointer 
   //* Operator & (AND)
   address3 := Address{"Sleman", "Yogyakarta", "Indonesia"}
   address4 := &address3

   address4.City = "Gunung Kidul"

   fmt.Println(address3) //{Gunung Kidul Yogyakarta Indonesia} 
   fmt.Println(address4) //&{Gunung Kidul Yogyakarta Indonesia}

   //* Operator * (STAR)
   address5 := Address{"Sleman", "Yogyakarta", "Indonesia"}
   address6 := &address5
   address7 := &address5

   address6.City = "Bantul"

   *address6 = Address{"Nunukan", "Kalimantan Utara", "Indonesia"}
    
   fmt.Println(address5) //{Nunukan Kalimantan Utara Indonesia}
   fmt.Println(address6) //&{Nunukan Kalimantan Utara Indonesia}
   fmt.Println(address7) //&{Nunukan Kalimantan Utara Indonesia}

   //* Function new
   address8 := new(Address)
//    address8.City = "Sleman"
   fmt.Println(address8) //&{  }

   //* Pointer di function
    address9 := Address{"Sleman", "Yogyakarta", "Indonesia"}
    changeAddress(&address9)
    fmt.Println(address9) //{Malinau Yogyakarta Indonesia} 

  //* Pointer method
  eko := Man{"eko"}
  eko.Marride()
  fmt.Println(eko.Name) //Eko





}