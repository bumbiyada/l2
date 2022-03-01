package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/bumbiyada/l2/tasks/t11/internal/based"
	"github.com/bumbiyada/l2/tasks/t11/internal/data"
	"github.com/bumbiyada/l2/tasks/t11/internal/httplistener"
)

func main() {
	log.Println("\t[APP] STARTED")
	var (
		wg         sync.WaitGroup
		http_to_db = make(chan based.Data_to_db, 1)
		db_to_http = make(chan []byte, 1)
	)
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		data.DB_handler(ctx, http_to_db, db_to_http)
	}()
	//based.Based()
	wg.Add(1)
	go func() {
		defer wg.Done()
		httplistener.HttpListener(ctx2, http_to_db, db_to_http)
	}()
	// exit
	func(cancel context.CancelFunc) {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		cancel()
		cancel2()
	}(cancel)
	wg.Wait()
}
