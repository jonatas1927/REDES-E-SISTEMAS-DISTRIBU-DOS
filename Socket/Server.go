package main

import (
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.255:8000")
	CheckError(err)
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		tamanho, addr, err := ServerConn.ReadFromUDP(buf)
		var msg = "-> " + addr.String() + " diz: " + string(buf[0:tamanho])
		ServerConn.Write(buf[0:tamanho])
		fmt.Println(msg)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
