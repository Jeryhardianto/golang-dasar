// * Package Time
  
package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())


	// Membuat waktu manual
	utc := time.Date(2020, time.April, 30, 20, 30, 0, 0, time.UTC)
	fmt.Println(utc)

	// Parse waktu
	// layout := time.RFC3339
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1999-01-29")
	fmt.Println(parse)
}