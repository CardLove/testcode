package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("unix", "/root/workspace/my/testcode/socket/sockct.sockct")
	if err != nil {
		fmt.Println("error starting socket")
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection")
			continue
		}
		go hadnleRequest(conn)

	}

}

func hadnleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading data.")
		}

		if n == 0 {
			fmt.Println("connection closed")
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
