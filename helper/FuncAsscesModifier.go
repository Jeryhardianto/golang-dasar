package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func Hallo(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("GoodBye", name)
}