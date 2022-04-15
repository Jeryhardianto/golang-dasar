// * Variable
// * Penamaan variable pada golang hanya boleh menyimpan tipe data yang sama,
// * Ketika ingin menyimpan tipe data yang berbeda kita harus membuat variable baru
// * Unutk membuat variable kita bisa menggunakan kata kunco var, lalu diikuti dengan nama variable dan tipe datanya
// * penulisan var dapat digantikan dengan kata kunci ( := ) di awal
// * Contoh : var name string
package main

import "fmt"

func main() {
	// Menyebut tipe data
   var name string
   name = "Jery"
   fmt.Println(name)

   name = "Hardianto"
   fmt.Println(name)

   //* Tanpa meyebut tipe data
   var Aliasname = "Hardiningrat"
   fmt.Println(Aliasname)

   var age = 23
   fmt.Println(age)

	//* var dapat digantikan dengan kata kunci ( := ) di awal
	  name2 := "Jery Hardianto,S.Kom"
	  fmt.Println(name2)

	  name2 = "Jery Hardianto,S.Kom.M.Eng"
	  fmt.Println(name2)

	  name2 = "Backend Engineer"
	  fmt.Println(name2)

	  var (
		  firstname = "Jery"
		  lastname = "Hardianto"
		  role = "Backend Engineer"
	  )
	  fmt.Println("Hallo nama saya ", firstname + " " + lastname + " Saya seorang " + role)
  
}