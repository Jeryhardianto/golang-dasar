// * For Loops

package main

import "fmt"

func main() {
	counter := 0

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	
	//* For dengan statement
	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	//* For range
	//* without for range
	names := []string{"jerome", "jery", "jerry"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Nama ke", i, "adalah", names[i])
	}

	//* with for range
	for i, values := range names {
		// Print index and value
		fmt.Println("Nama ke", i, "adalah", values)
		// Print value 
		fmt.Println(values)
	}

	person := make(map[string]string)
	person["name"] = "Jery"
	person["age"] = "23"

	for key, value := range person {
		fmt.Println( key, "=", value)
	}



}