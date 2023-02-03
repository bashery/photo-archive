package main

import (
	"fmt"
	"net"
)

func main() {
	s := newServer()
	go s.run() // run server in bachgound

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("unable to strart server %s\n", err)
	}
	defer listener.Close()

	fmt.Println("server chat ranning at port :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("faild to accept %s\n", err.Error())
			continue
		}
		go s.newClient(conn)
	}

}
