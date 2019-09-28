package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 0}
	fmt.Println(a[:2])
}

func aa(b [5]int) {
	b[1] = 10
}
