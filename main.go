package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port :6969")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	con, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer con.Close()

	for {
		buf := make([]byte, 1024)

		_, err := con.Read(buf)

		if err != nil {

			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client", err.Error())
			os.Exit(1)
		}

		con.Write([]byte("+OK\r\n"))
	}
}
