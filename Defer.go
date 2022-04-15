// * Defer

package main

import "fmt"

func loggin(){
	fmt.Println("Selesai memanggil function");
}

func runnApplication(value int){
	defer loggin(); //executed pertama
	result := 10 / value
	fmt.Println(result);
	fmt.Println("Run function");
}


func main() {

	runnApplication(0);
}