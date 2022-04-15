// * Operasi Boolean

package main

import "fmt"

func main() {
	var ujian = 80
	var absensi    = 80

	// var lulusUjian = ujian >= 80
	// var lulusAbsensi = absensi >= 80

	// var lulus = lulusAbsensi && lulusUjian

	// fmt.Println("Lulus Ujian? ", lulus)

	// Atau

	fmt.Println("Lulus Ujian? ", ujian >= 80 && absensi >= 80)



}