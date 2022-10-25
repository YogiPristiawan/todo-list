package main

import (
	"fmt"
)

func main() {
	test()
}

func test() (s string) {
	defer func() {
		fmt.Println("oke")
	}()

	defer func() {
		fmt.Println("halo")
	}()

	return "oke"

	// defer func() {
	// 	fmt.Println("xixi")
	// }()
}
