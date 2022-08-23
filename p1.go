package main

import "fmt"

func main() {
	printstuff()
}

func printstuff() {
	fmt.Println("hello world - this is a test!")

	for i := 1; i <= 10; i++ {
		println("Second line with basic go command ", i, " !")
	}
}
