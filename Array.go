// * Array pada golang

package main

import "fmt"

func main() {
	// var names [jumlah element array] tipe data
	var names [3]string
	names[0] = "Jery"
	names[1] = "Hardianto"
	names[2] = "Eko"

	fmt.Println(names[0])

	// atau

	var values = [3]int{29,1,99}
	fmt.Println(values)

	// Fanction Array
	// len => Panjang array 
	fmt.Println(len(names))
	fmt.Println(len(values))

	

}