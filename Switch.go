// *Switch Expressions

package main

import "fmt"

func main() {
	name := "jerome112345"
	switch name {
	case "jery":
		fmt.Println("Hello jery")
	case "jerome":
		fmt.Println("Hello jerome")
	default:
		fmt.Println("Hi, boleh kenalan ?")
	}

	//* Switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kamu panjang")
	case false:
		fmt.Println("Nama kamu pendek")
	}

	//* Switch dengan tanda kondisi
	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama kamu terlalu panjang")
	case length > 5:
		fmt.Println("Nama kamu panjang")
	default:
		fmt.Println("Nama kamu pendek")
		
	}




}