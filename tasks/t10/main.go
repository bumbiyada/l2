package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
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
	var (
		network string = "tcp"
		address string
		host    string
		port    string
		wg      sync.WaitGroup
	)
	// flag
	duration := flag.Int("timeout", 10, "set timeout for program")
	flag.Parse()

	host = flag.Arg(0)
	port = flag.Arg(1)

	TimeOut := time.Duration(*duration) * time.Second

	// reading args
	if flag.NArg() < 2 {
		log.Println("not enough args, exiting")
		os.Exit(1)
	}

	address = net.JoinHostPort(host, port)
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("Trying to connect this address = %s | network = %s | timeout = %v\n", address, network, *duration)
	// launching app
	wg.Add(1)
	go func() {
		defer wg.Done()
		f2(ctx, network, address, TimeOut)
	}()

	// listening SigINT
	wg.Add(1)
	go func(cancel context.CancelFunc) {
		defer wg.Done()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		for {
			select {
			case <-sig:
				log.Println("Got sigInt, press enter to exit")
				cancel()
				return
			default:

			}
		}
	}(cancel)
	wg.Wait()
}

//deprecated function
func f1(network, address string, TimeOut time.Duration) {
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

// connect to server
func connect(ctx context.Context, network, address string, t time.Duration) (conn net.Conn, err error) {
	ch := time.After(t)
	n := t.Seconds() / 10
	log.Println(n)
	for {
		select {
		case <-ch:
			err := errors.New("Timeout to connect")
			return conn, err
		case <-ctx.Done():
			err := errors.New("Got SIGINT, exiting")
			return conn, err
		default:
			conn, err := net.Dial(network, address)
			if err != nil {
				log.Println("can`t connect , retrying...", err)
				time.Sleep(time.Second * time.Duration(n))
				continue
			}
			return conn, err
		}
	}
}

//main function of app
func f2(ctx context.Context, network, address string, t time.Duration) {
	conn, err := connect(ctx, network, address, t)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting successfully")
			return
		default:
			// read stdin
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			// send to socket
			fmt.Fprintf(conn, text+"\n")
			// get response
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Print(message)
		}

	}
}
