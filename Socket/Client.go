package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.255:8000")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.255:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		i++
		buf := []byte(text)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(text, err)
		}
	}
}
