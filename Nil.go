// * Nil
/**
Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
Namun di Go-Lang ada data nil, yaitu data kosong
Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
*/

package main

import "fmt"

func NewMap(name string) map[string]string{
	if name == "" {
		return nil
	}else{
		return map[string]string{
			"name": name,
		}
	}
}


func main() {
	var person map[string]string = NewMap("Jerry")

	if person == nil {
		fmt.Println("Person is nil")
	}else{
		fmt.Println("Person is not nil")
	}

}