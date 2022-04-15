// * Panic function = untuk menghetikan program
// * Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// * Saat panic function dipanggil, program akan langsung berhenti dan namun defer function tetap akan dieksekusi

package main

import "fmt"

func endApp(){
	fmt.Println("Ending Application");
}

func runApp(error bool){
	defer endApp();
	if(error){
		panic("APLIKASI ERROR BRO");
	}
	fmt.Println("Run Application");

}

func main() {
	runApp(true);

}