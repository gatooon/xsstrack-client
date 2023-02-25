package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClient() {
	fmt.Println("Enter server IP Address :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	connection, err := net.Dial("tcp", scanner.Text()+":8800")
	if err != nil {
		fmt.Println("Error : " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected !")
	fmt.Println("Enter Url to listen on: ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()

	connection.Write([]byte(scanner.Text()))

	fmt.Println("Server Listening for " + scanner.Text())

	defer connection.Close()
	for {
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println("Received: ", string(buffer[:mLen]))
		fmt.Println("----------------------------------------------------------------------------")
	}
}
