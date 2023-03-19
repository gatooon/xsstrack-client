package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
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

	fmt.Println("Server Listening")

	defer connection.Close()
	for {
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		currentTime := time.Now()
		formatedTime := strconv.Itoa(currentTime.Hour()) + ":" + strconv.Itoa(currentTime.Minute()) + ":" + strconv.Itoa(currentTime.Second())
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println("Time : " + formatedTime)
		fmt.Println("Received: ", string(buffer[:mLen]))
		fmt.Println("----------------------------------------------------------------------------")
	}
}
