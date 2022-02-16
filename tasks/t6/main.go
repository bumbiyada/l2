package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// function
func main() {
	// flags initialization
	separator := flag.Bool("s", false, "only with separator")
	fields := flag.String("f", "", "specify columns")
	delimeter := flag.String("d", "\t", "use another delimeter")
	flag.Parse()
	path := flag.Arg(0)
	// open file
	file, err := os.Open(path)
	defer file.Close()

	// err handling
	if err != nil {
		log.Println("Cannot open the file, make sure you open something correct")
		return
	}
	// parsing -f value to int type
	colls, err := strconv.ParseInt(*fields, 10, 64)
	if err != nil {
		log.Println("Cannot parse fields param, value must be int type")
		return
	}
	// scanner init
	scanner := bufio.NewScanner(file)
	// iterating by lines
	for i := 1; scanner.Scan(); i++ {
		tmp := scanner.Text()
		// strings split action
		columns := strings.Split(tmp, *delimeter)
		// hadling some cases
		if int(colls) > len(columns) || int(colls) < 0 {
			fmt.Println("")
		} else if *separator == true && len(columns) == 1 {
			fmt.Println("")
		} else {
			// printing the result
			fmt.Println(columns[colls-1])
		}
	}
}
