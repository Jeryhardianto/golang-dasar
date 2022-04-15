// * Recover = function yang bisa untuk menangkap data Panic
// * dengan recover proses panic akan berhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func endApp(){
	//Recover wajib didalam func defer
	message := recover();
	if message != nil{
		fmt.Println("Ini error: ", message);
	}
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
	runApp(false);
	fmt.Println("Hallo");


}