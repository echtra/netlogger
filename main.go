package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"time"
)

func startLogging(conn *net.TCPConn, dir string) {
	now := time.Now().UnixNano()
	fileName := fmt.Sprintf("netlog.%d", now)
	file := path.Join(dir, fileName)
	log.Printf("Starting log %s", fileName)

	defer conn.Close()

	out, err := os.Create(file)
	if err != nil {
		log.Printf("Error opening %s: %v", file, err)
		return
	}

	defer out.Close()

	n, err := io.Copy(out, conn)
	if err != nil {
		log.Printf("Error in %s: %v", fileName, err)
	} else {
		log.Printf("Wrote %d bytes to %s", n, fileName)
	}
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	port := flag.Int("p", 30000, "Port to listen on")
	dir := flag.String("d", cwd, "Directory to write log files to")
	flag.Parse()

	listenAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go startLogging(conn, *dir)
	}
}
