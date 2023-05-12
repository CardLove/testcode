package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("unix", "/root/workspace/my/testcode/socket/sockct.sockct")
	if err != nil {
		fmt.Println("Error connecting to server")
		return
	}
	defer conn.Close()

	msg := " hello, server"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("error sending data.")
		return
	}
}
