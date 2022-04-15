// * Container List

package main

import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")
	data.PushFront("Budi")

	// fmt.Println(data.Front().Value) // Eko
	// fmt.Println(data.Back().Value) // Khannedy
	
	// Dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// Dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	




}