package main

import (
	"bufio"
	"net"
	"fmt"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	for {
		message,_ := bufio.NewReader(conn).ReadString('\n')
		//output.SetText(message)
		fmt.Println("``````````````````")
		fmt.Println(message)
		fmt.Println("``````````````````")
	}
}
