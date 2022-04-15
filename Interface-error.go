// * Interface error
/**
Untuk membuat error, kita tidak perlu manual.
Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan kita bahas secara detail di materi tersendiri)

*/
package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error ){
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else{
		result := nilai / pembagi
		return result, nil
	}
}


func main() {
	result, err := pembagi(100, 0)
	if err == nil {
		fmt.Println("hasil :", result)
		}else{
		fmt.Println("Error :", err.Error())
	}

}