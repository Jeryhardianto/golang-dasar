// * Package String

package main

import (
	"fmt"
	"strings"
)

func main() {

// * strings.Trim(string, cutset) = Memotong cutset di awal dan akhir string
fmt.Println(strings.Trim("    Hello, World   ", " ")) //Hello, World
fmt.Println(strings.Trim("a    Hello, World   a", " ")) //a    Hello, World   a

// * strings.ToLower(string) = Membuat semua karakter string menjadi lower case
fmt.Println(strings.ToLower("Jery Hardianto")) // jery hardianto

// * strings.ToUpper(string) = Membuat semua karakter string menjadi upper case
fmt.Println(strings.ToUpper("Jery Hardianto")) // JERY HARDIANTO

// *  strings.Split(string, separator) = Memotong string berdasarkan separator
fmt.Println(strings.Split("Jery hardianto", " ")) //[Jery hardianto]

// *  strings.Contains(string, search) = Mengecek apakah string mengandung string lain
fmt.Println(strings.Contains("sekolah koding", "koding")) // true
fmt.Println(strings.Contains("sekolah koding", "dasar")) // false

// *  strings.ReplaceAll(string, from, to) = Mengubah semua string dari from ke to
fmt.Println(strings.ReplaceAll("Jery Hardianto", "Hardianto", "Hardiningrat")) // Jery Hardiningrat

}