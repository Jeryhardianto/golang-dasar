// * Constant untuk variable yang tidak pernah berubah

package main

import "fmt"

func main() {
	const name = "Jery Hardianto,S.Kom.M.Eng"
	const age = 23
	const role = "Backend Engineer"
	fmt.Println("Hallo nama saya ", name + " " + "Saya seorang " + role)
	fmt.Println("Umur saya ", age)

	const (
		firstname = "Jery"
		lastname = "Hardianto"
	)
	fmt.Println("Hallo nama saya ", firstname + " " + lastname + ", Saya seorang " + role)
}