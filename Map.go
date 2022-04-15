// * Map

package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Jery Hardianto",
		"address": "Janti",
	}

	person["title"] = "Backend Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	//* Functions in maps
	Books := make(map[string]string)
	Books["title"] = "Go Programming"
	Books["author"] = "Jery Hardianto"
	Books["publisher"] = "Gramedia"
	fmt.Println(Books)
	fmt.Println(len(Books))
	
	delete(Books, "publisher")
	fmt.Println(Books)
	fmt.Println(len(Books))

}