// * Conversi Tipe Data
package main

import "fmt"
func main() {
	var nilai32 int32 = 100000
	var nili64 int64 = int64(nilai32)
	var nilai8 int8     = int8(nilai32)

	fmt.Println("nilai nili32 = ", nilai32)
	fmt.Println("nilai nili64 = ", nili64)
	fmt.Println("nilai nilai8 = ", nilai8)


	var name = "Jery Hardianto"
	var j byte = name[0] 
	// var k byte = name[0] //tidak mesti var harus j
	var eString string = string(j)

	fmt.Println("name = ", name)
	fmt.Println("J = ", eString)
}