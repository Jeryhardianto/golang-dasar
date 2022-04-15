// * Function

package main

import "fmt"

//* Functions
func sayHallo(){
	fmt.Println("Hallo Golang")
}

//* Functions Parameters
func sayHalloWithName(name string, age int){
	fmt.Println("Hallo", name, "umur", age)
}

//* Function with Return Value
func getHello(name string) string {
	if name == "" {
		return "Hello World"
	} else {
		return "Hello " + name
	}
}

//* Function with Multiple Return Value
func getFullName() (string, string, int) {
	return "Jery", "Hardianto", 23
}

//* Function with Named Return Values
func getComplateName() (firstName, lastName, middleName string) {
	 firstName = "Jery"
	 lastName = "Hardianto"
	 middleName = "Saputra"
	 return 
}

//* Function variadic
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

//* Function As Value
func getGoodBye(name string) string {
	return "Good bye " + name
}

//* Function as Parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFilter := filter(name)
	// return filter(nameFilter)
	fmt.Println("Hallo ",nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing"{
		return "...."
	}else{
		return name
	}
}

//* Function Type Declaration
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hallo ",nameFilter)
}

//* Anonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "is not allowed")
	} else {
		fmt.Println("User", name, "is allowed")
	}
}

//* Recursive Function / Function memanggil dirinya sendiri
// Case Study : Factorial
// Factorial for Loops
func factorialForLoop(value int) int {
	total := 1
	for i := value; i >0 ; i--{
		total *= i
	}
	return total
}

// Factorial for Recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	for i := 0; i <= 1; i++ {
		sayHallo()
	}

	//* Function with Parameters
	sayHalloWithName("Jery", 23)

	//* Function with Return Value
	result := getHello("Jery Hardianto")
	fmt.Println(result)

	//* Function with Multiple Return Value
	// firstName, lastName, age := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Println(age)
	//* untuk mengambil nilai firstNamenya saja / mengignore nilai lastName, ad=ge
	firstName, _, _ := getFullName()
	fmt.Println(firstName)

	//* Function with Named Return Values
	firstName, lastName, middleName := getComplateName()
	fmt.Println(firstName + " " + middleName + " " + lastName)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(middleName)

	//* Function variadic
	total := sumAll(10, 90, 30, 50, 40)
	fmt.Println(total)

	//* Function variadic with slice
	numbers := []int{10, 90, 30, 50, 40}
	total = sumAll(numbers...)
	fmt.Println(total)

	//* Function as Value
	goodBye := getGoodBye

	result = goodBye("Jery")
	fmt.Println(result)
	fmt.Println(getGoodBye("Hardianto"))

	//* Function as Parameter
	sayHelloWithFilter("Jery", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//* Function Type Declaration
	sayHelloWithFilter2("Jery", spamFilter)
	sayHelloWithFilter2("Anjing", spamFilter)

	//* Anonymous Function
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Jery", blacklist)

	registerUser("eko", func(name string) bool {
		return name == "eko"
	})

	//* Recursive Function / Function memanggil dirinya sendiri
	// Case Study : Factorial
	// Factorial for Loops
	loop := factorialForLoop(5)
	fmt.Println(loop)
	fmt.Println(5*4*3*2*1)

	// Factorial for Recursive
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
	fmt.Println(5*4*3*2*1)

}