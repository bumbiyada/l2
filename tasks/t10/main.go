package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case <-c:
			log.Println("exit")
		}
	}
}
func copyTo(dst io.Writer, src io.Reader) {
	log.Println("action")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal("can`t read/write", err)
	}
}
func main() {
	//parse flags
	var (
		TimeOut time.Duration = 10 * time.Second
		network string        = "tcp"
		address string
		host    string
		port    string
	)
	duration := *flag.Int("timeout", 10, "set timeout for program")

	flag.Parse()
	TimeOut = time.Duration(duration) * time.Second

	// reading args
	if flag.NArg() < 2 {
		log.Println("not enough args, exiting")
		os.Exit(1)
	}
	host = flag.Arg(0)
	port = flag.Arg(1)

	address = net.JoinHostPort(host, port)
	fmt.Printf("Trying to connect this address = %s | network = %s | timeout = %v\n", address, network, duration)
	conn, err := net.DialTimeout(network, address, TimeOut)
	if err != nil {
		log.Println("Can`t connect to dial, or timeout \n", err)
		os.Exit(1)
	}
	defer conn.Close()

	log.Println("Connected...")
	go copyTo(os.Stdout, conn)
	copyTo(conn, os.Stdin)
	time.Sleep(2 * time.Second)
	log.Println("exiting")
}
