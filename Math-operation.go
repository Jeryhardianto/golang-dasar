// * Math Operation / operasi matematika

package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	//Argument Assignment
	var d = 10
	d += d // ==> d = d + 10

	fmt.Println(d)

	// Unary Oparator => Operasi ke 1 buah variabel
	d++
	fmt.Println(d)

	var negatif = -10
	var positif = +10 // positif flekasibel, tanda + tidak harus didefinisikan karena default dari golangnya sendiri sudah positif

	fmt.Println(negatif)
	fmt.Println(positif)


}