package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	time.Sleep(30 * time.Second)
    ln, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Fatal("Error starting server:", err)
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("Error accepting connection:", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buffer := make([]byte, 1024)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            log.Println("Error reading:", err)
            return
        }
        fmt.Println("Received:", string(buffer[:n]))
    }
}
