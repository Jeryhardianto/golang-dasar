// * Regexp

package main

import (
	"fmt"
	"regexp"
)

func main() {
// *  regexp.MustCompile(string) = Membuat Regexp
regexp := regexp.MustCompile("e([a-z])o")


// *  Regexp.MatchString(string) bool =Mengecek apakah Regexp match dengan string
fmt.Println(regexp.MatchString("ejok"))
fmt.Println(regexp.MatchString("etok"))
fmt.Println(regexp.MatchString("eDok"))


// *  Regexp.FindAllString(string, max) = Mencari string yang match dengan maximum jumlah hasil
search := regexp.FindAllString("ejok, eko, eto, aku", 2)
fmt.Println(search)


}