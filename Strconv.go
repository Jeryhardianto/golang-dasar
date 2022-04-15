// * Strconv = String Convertion

package main

import (
	"fmt"
	"strconv"
)

func main() {
// *  strconv.parseBool(string) = Mengubah string ke bool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	}else{
		fmt.Println(err.Error())
	}

// *  strconv.parseFloat(string) = Mengubah string ke float


// *  strconv.parseInt(string) = Mengubah string ke int64
 number, err := strconv.ParseInt("10000", 10, 64)
 if err == nil {
	fmt.Println(number)
 }else{
 	fmt.Println(err.Error()) 	
 }

// *  strconv.FormatBool(bool) = Mengubah bool ke string


// *  strconv.FormatFloat(float, … ) = Mengubah float64 ke string


// *  strconv.FormatInt(int, … ) = Mengubah int64 ke string
 value := strconv.FormatInt(10000, 10)
 fmt.Println(value)

 valueInt, _ := strconv.Atoi("10000") // _ ignore error (REKOMENDASI)
 fmt.Println(valueInt)

}