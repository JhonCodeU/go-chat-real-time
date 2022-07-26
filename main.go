package main

import "net"

func logFatal(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	openConnection = make(map[net.Conn]bool)
	newConnection  = make(chan net.Conn)
	deadConnection = make(chan net.Conn)
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	logFatal(err)

	defer ln.Close()

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				logFatal(err)
			}
			openConnection[conn] = true
			newConnection <- conn
		}
	}()
}
