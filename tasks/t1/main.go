package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	// current time by this external library
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// print error in stderr
		fmt.Fprintln(os.Stdout, err)
		// print that we have error in stdout and os.exit(1)
		log.Fatalln("!!! ERROR !!! ")
	}
	// printing result
	fmt.Println(time)
}
