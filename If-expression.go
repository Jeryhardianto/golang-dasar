// * If Expression

package main

import "fmt"

func main() {

	//* if else
	name := "jery hardiningraat"

	if name == "jery" {
		fmt.Println("Hello jery")
	}else{
		fmt.Println("Hi, boleh kenalan ?")
	}

	//* else if
	if name == "jery" {
		fmt.Println("Hello jery")
	}else if name == "jerome" {
		fmt.Println("Hello jerome")
	}else{
		fmt.Println("Hi, boleh kenalan ?")
	}

	//* if dengan short statement
	
	if length := len(name); length > 5 {
		fmt.Println("Nama kamu panjang")
	}else{
		fmt.Println("Nama kamu pendek")
	}

}