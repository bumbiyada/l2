package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// main listener of all things
func event_listener() {
	//starting shell
	log.Println("You are now in shell")
	var tmp string
	var arr []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>>")
		scanner.Scan()
		tmp = scanner.Text()
		arr = strings.Split(tmp, " ")
		// if tree
		if arr[0] == "exit" || arr[0] == "exit()" || arr[0] == "q!" {
			fmt.Println("Exiting Succesfully")
			return
		}
		if arr[0] == "help" {
			help()
			continue
		}
		if arr[0] == "cd" && len(arr) == 2 {
			os.Chdir(arr[1])
		} else {
			execute(arr)
		}

	}
}

// execute applications
func execute(arr []string) {
	name := arr[0]
	args := arr[1:]
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error while executing program")
	}

}

// help function
func help() {
	fmt.Println(
		`------------------------------------Go SHELL------------------------------------
-USAGE:
	- type command to execute it
	- type exit | exit() | q! to exit`)
}

// main function
func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	//starting listener
	go func() {
		defer wg.Done()
		event_listener()
	}()
	wg.Wait()
}
