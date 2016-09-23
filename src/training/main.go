package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	fmt.Println(swap("a", "b"))
}

func swap(x, y string) (string, string)  {
	return y, x
}
