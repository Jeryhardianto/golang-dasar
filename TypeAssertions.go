// * Type Assertions
/**
Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

*/
package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	//Switch with type assertion (REKOMEN)
	result = random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}

}