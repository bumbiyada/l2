package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// whether string is numeric or not
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// if []string contains string or not
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//read file
func read_file(path string, unique *bool) []string {
	result := []string{}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println("error while opening file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Text()
		if *unique == false {
			result = append(result, tmp)
		} else {

			if contains(result, tmp) == true {
				continue
			} else {
				result = append(result, tmp)
			}

		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return result
}

// sorting file
func sort_file(array []string, numeric *bool) []string {
	if *numeric == true {
		numbers := []float64{}
		strs := []string{}
		for _, s := range array {
			if isNumeric(s) == true {
				n, _ := strconv.ParseFloat(s, 64)
				numbers = append(numbers, n)
			} else {
				strs = append(strs, s)
			}
		}
		sort.Strings(strs)
		sort.Float64s(numbers)
		for _, f := range numbers {
			s := strconv.FormatFloat(f, 'f', -1, 64)
			strs = append(strs, s)
		}
		return strs
	} else {
		sort.Strings(array)
		return array
	}

}

// printing file
func print_file(array []string, reverse *bool) {
	if *reverse == false {
		fmt.Println("It`s NOT reversed")
		for _, str := range array {
			fmt.Println(str)
		}
	} else {
		fmt.Println("It`s reversed")
		for i := (len(array) - 1); i >= 0; i-- {
			fmt.Println(array[i])
		}
	}
}

//main function
func main() {
	//flags
	showHelp := flag.Bool("h", false, "Show help message.")
	reverse := flag.Bool("r", false, "Sorting Desc")
	numeric := flag.Bool("n", false, "Sorting lines like numbers, not like strings")
	unique := flag.Bool("u", false, "All duplicated lines will be excluded")
	flag.Parse()
	if *showHelp {
		fmt.Printf("Usage: %s [OPTION]... [URL]...\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	//arg
	filepath := flag.Arg(0)
	//read file, sort and print
	obj := read_file(filepath, unique)
	obj = sort_file(obj, numeric)
	print_file(obj, reverse)
}
