package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Begin struct {
	filename  string
	peerId    int
	ipAddress string
}

func main() {

	fmt.Println("||========================================================================================||")
	fmt.Println("||                           PEER-TO-PEER FILE SHARING SYSTEM                             ||")
	fmt.Println("||                       ========================================                         ||")
	fmt.Println("||========================================================================================||")
	fmt.Println("\n <CENTRAL INDEX SERVER IS UP AND RUNNING....>")
	fmt.Println(" ============================================")

	registerRequestThread := make(chan struct{})
	searchRequestThread := make(chan struct{})

	go func() {
		defer close(registerRequestThread)
		RegisterRequestThread()
	}()

	go func() {
		defer close(searchRequestThread)
		SearchRequestThread()
	}()

	<-registerRequestThread
	<-searchRequestThread
}

func RegisterRequestThread() {
	listen, err := net.Listen("tcp", "localhost:2001")
	if err != nil {
		fmt.Println("error listening", err)
		return
	}
	defer listen.Close()

	fmt.Println("Listen for Register")
	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("error accepting", err)
			return
		}
		fmt.Println("Connection Received from " + connection.RemoteAddr().String() + " for registration")
		var strVal string
		decodedConnection := gob.NewDecoder(connection)
		decodedConnection.Decode(&strVal)
		fmt.Println(strVal)
		fmt.Println("<====Registered====>")
	}

}

func SearchRequestThread() {
	listen, err := net.Listen("tcp", "localhost:2002")
	if err != nil {
		fmt.Println("error listening", err)
		return
	}
	defer listen.Close()
	fmt.Println("Listen for Search")

}
