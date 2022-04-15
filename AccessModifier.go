// * Access Modifier
/**
Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
Jika nama nya diawali dengan hurup BESAR, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup KECIL, artinya tidak bisa diakses dari package lain
*/
//* FUNCTION atau VARIABEL yang diawali dengan hurup BESAR, artinya bisa diakses dari package lain, jika dimulai dengan hurup KECIL, artinya tidak bisa diakses dari package lain

package main

import "golang-dasar/helper"

func main() {
	helper.Hallo("Jery") // ini bisa diakses dari package lain/luar karena diawali dengan hurup besar
	helper.sayGoodBye("Jery")  //error => sayGoodBye not exported by package helper karena function nya diawali dengan hurup kecil maka tidak bisa diakses dari luar

	helper.version = 2 // error => variabel ini tidak bisa diakses dari package lain karena diawali dengan hurup kecil 
	helper.Application = "Belajar Golang" //variabel ini bisa diakses dari package lain/luar karena diawali dengan hurup besar

}