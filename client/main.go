package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8000")
	logFatal(err)

	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	logFatal(err)
	strings.Trim(username, "\r\n")

	fmt.Sprintf("welcome %s", username)
}

func logFatal(err error) {
	if err != nil {
		panic(err)
	}
}
