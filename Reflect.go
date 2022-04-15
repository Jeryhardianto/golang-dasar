// * Reflect

package main

import (
	"fmt"	
	"reflect"
)


// type Sample struct {
// 	Name string
// 	Age  int
// }

type Sample struct {
	Name string `required:"true", max:"10"`
	Age  int `required:"true", max:"1"`
}

//* Validation Library
func IsValid(data interface{}) bool{
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).String() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Jeryy", 10}

	var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(1).Name)


	// * Struct Tag
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	
	//* Validation Library
	// sample.Name = ""
	fmt.Println(IsValid(sample))

}