package main

import "fmt"

func main() {
	fmt.Println("Int value: ")
	Print[int]([]int{1, 2, 3})
	fmt.Println()

	fmt.Println("String value: ")
	Print[string]([]string{"a", "b", "c"})
	fmt.Println()

	fmt.Println("Float value: ")
	Print[float32]([]float32{1.0, 2.1, 3.2})
	fmt.Println()

	fmt.Println("Bool value: ")
	Print[bool]([]bool{true, false})
	fmt.Println()
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v) 
		fmt.Print("\t")
	}
}