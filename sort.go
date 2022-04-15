// * Package Sort

package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type UserSlice []user 

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
 users := []user{
	{"Bob", 31},
	{"John", 42},
	{"Michael", 17},
	{"Jenny", 26},

 }
 fmt.Println(users)

 sort.Sort(UserSlice(users))
 
 fmt.Println(users)
}