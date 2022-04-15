// * Type Declarations = membuat type data dengan nama yang kita inginkan

package main

import "fmt"

func main() {
	type NoKTP string
	type Marride bool

	var NIK NoKTP = "123456789012"
	var status Marride = false

	fmt.Println(NIK)
	fmt.Println(status)

}