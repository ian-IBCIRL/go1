package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	printstuff()

	readstuff()

}

func printstuff() {
	fmt.Println("hello world - this is a test!")

	for i := 1; i <= 10; i++ {
		println("Second line with basic go command ", i, " !")
	}
}

func readstuff() {
	status := ""
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err = bufio.NewReader(conn).ReadString('\n')
	println(status)
}
