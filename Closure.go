// * Closure

package main

import "fmt"

func main() {
	name := "JeEkory"
	counter := 0

	increment := func() {
		name = "Jery"
		fmt.Println("Incrementing")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)


}