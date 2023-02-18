package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	a := test(&slice)
	fmt.Print(*a)
}

func test(slice *[]string) *[]string {
	(*slice) = append((*slice), "a")
	// fmt.Println(slice)
	(*slice)[0] = "b"
	(*slice)[1] = "b"
	// fmt.Print(*slice)
	return slice
}
