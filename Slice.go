// * Slice

package main

import "fmt"

func main() {
	var months = [...]string{
		 "Januari",
		 "Februari", 
		 "Maret", 
		 "April", 
	     "Mei", 
		 "Juni", 
		 "Juli",
		 "Agustus", 
		 "September", 
		 "Oktober", 
		 "November", 
		"Desember",
	}

	var Slice1 = months[4:7]
	// var Slice2 = months[6:9]
	fmt.Println(Slice1) 
	fmt.Println(len(Slice1))
	fmt.Println(cap(Slice1))

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//* Membuat slice baru dari awal / Make slice 
	newSlice := make([]string, 2,5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//* Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//* Array x Slice
	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)





}