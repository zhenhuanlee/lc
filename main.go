package main

import "fmt"

func main() {
	a := make(map[int]int)

	a[1] = 1
	a = nil
	a[2] = 1
	fmt.Println(nil)
}
