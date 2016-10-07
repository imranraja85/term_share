package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")

	flag.Parse()

	fmt.Println("Is a host?", isHost)

	go func() {
		// begin listening on ip and port
		ipAndPort := "127.0.0.1:3009"
		listener, _ := net.Listen("tcp", ipAndPort)
		fmt.Println("Listening on: ", ipAndPort)

		// accept any connection you can get!
		conn, _ := listener.Accept()
		fmt.Println(conn)
	}()

	// fire up a terminal
	cmd := exec.Command("vim")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	log.Printf("Run = %v", err)

	time.Sleep(5 * time.Second)
}
