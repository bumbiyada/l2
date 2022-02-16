package main

import (
	"fmt"
	"log"
	"sync"
)

// main listener of all things
func event_listener() {
	var tmp string
	var (
		command string
		arg     string
	)
	for {
		n, err := fmt.Scanln(&tmp)
		//r := strings.NewReader("5 true gophers")
		//n, err := fmt.Fscanf(r, "%s %s", &command, &arg)
		if err != nil {
			log.Println("error while getting input")
			continue
		}
		tmp = command + arg
		fmt.Println(n, "\n", tmp)
		if tmp == "exit" {
			fmt.Println("EXITED ")
			return
		}
	}
}

// main function
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		event_listener()
	}()
	wg.Wait()
}
