package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type PeerServer struct {
	cisIp          string `default:"localhost"`
	clientId       string `default:"1001"`
	regMessage     string
	searchFileName string
}

func (ps *PeerServer) Init() {
	//Open the file for reading
	file, err := os.Open("./Peers/peer_1/main/indexServerIP.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("File closed successfully")
		}
	}(file)

	//Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		ps.cisIp = scanner.Text()
		fmt.Printf("IndexServer IP is: %v\n", ps.cisIp)
	} else {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("||========================================================================================||")
	fmt.Println("||                           PEER-TO-PEER FILE SHARING SYSTEM                             ||")
	fmt.Println("||                       ========================================                         ||")
	fmt.Println("||                                       MENU:                                            ||")
	fmt.Println("||========================================================================================||")

	for {
		fmt.Println("============================================================================================")
		fmt.Println("Enter The Option:")
		fmt.Println("==================")
		fmt.Println("1. Registering the File")
		fmt.Println("2. Searching on CentralIndexServer")
		fmt.Println("3. Downloading from peer server")
		fmt.Println("4. Exit")

		var input string
		_, err = fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input from user:", err)
		}
		ps.regMessage = input
		fmt.Printf("Input from user is: %v\n", ps.regMessage)

		if ps.regMessage == "1" {
			fmt.Println("Enter the String in format: 4_Digit id and File Names seperated by Space")
			_, err = fmt.Scanln(&ps.regMessage)
			if err != nil {
				fmt.Printf("Error while reading input:%v\n", err)
			}
			val := strings.Split(ps.regMessage, " ")
			pearPort, err := strconv.ParseInt(val[0], 10, 0)
			if err != nil {
				fmt.Println("Error while String to Int conversion: ", err)
			}
			RegisterWithCentralServer(ps)
			AttendFileDownloadRequest(pearPort)
		} else if ps.regMessage == "2" {
			SearchWithIServer()
		} else if ps.regMessage == "3" {
			DownloadFromPeerServer()
		} else if ps.regMessage == "4" {
			fmt.Println("Exiting")
			os.Exit(0)
		} else {
			fmt.Println("Please Enter a valid option")
		}

	}

}

func RegisterWithCentralServer(ps *PeerServer) {

	conn, err := net.Dial("tcp", "localhost:2001")
	if err != nil {
		fmt.Println("Error Connecting:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to Register on Central Index Server on port 2001")
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err1 := enc.Encode(ps.regMessage); err1 != nil {
		fmt.Println("Error while encoding", err1)
	}
	fmt.Println("Registered Successfully!!")

}

func AttendFileDownloadRequest(port int64) {

}

func SearchWithIServer() {

}

func DownloadFromPeerServer() {

}

func main() {
	ps := new(PeerServer)
	ps.Init()

}
