package main

import "fmt"

func main() {
	array := [10]int{}

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	fmt.Println(array)
	fmt.Println("Slice example")

	slice := []int{}
	for i := 0; i <= 25; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice)
}
