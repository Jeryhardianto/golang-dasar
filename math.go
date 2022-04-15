// * Package Math

package main

import (
	"fmt"
	"math"
)

func main() {
// *  math.Round(float64) = Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
fmt.Println(math.Round(1.7))// 2
fmt.Println(math.Round(1.3))// 1


// *  math.Floor(float64) = Membulatkan float64 kebawah
fmt.Println(math.Floor(1.7))// 1


// *  math.Ceil(float64) = Membulatkan float64 keatas
fmt.Println(math.Ceil(1.3))// 2



// * math.Max(float64, float64) = Mengembalikan nilai float64 paling besar
fmt.Println(math.Max(1.3, 2.7))// 2.7


// * math.Min(float64, float64) = Mengembalikan nilai float64 paling kecil
fmt.Println(math.Min(1.3, 2.7))// 1.3

}