// * Break and Continue

package main

import "fmt"

func main() {
	//* Break = menghentikan perluangan saat ini
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	//* Continue = menghentikan perluangan saat ini dan melanjutkan ke perulangan berikutnya
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}

}