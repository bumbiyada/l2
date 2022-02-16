package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/chatzikalymnios/goget/internal/utils"
)

func main() {
	//current work directory, if dir isn`t specified by flag -d , it takes current dir
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	// all flags of this utility
	showHelp := flag.Bool("h", false, "Show help message.")
	inputFile := flag.String("i", "", "Input file to read URLs from.")
	outputDirectory := flag.String("d", cwd,
		"Directory to save downloaded files to. Default value is the current working directory.")
	concurrencyLevel := flag.Int("c", runtime.GOMAXPROCS(0),
		"Number of concurrent downloads allowed. Default value is GOMAXPROCS (nproc, if unset).")
	flag.Parse()

	// help
	if *showHelp {
		fmt.Printf("Usage: %s [OPTION]... [URL]...\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	// some error  handling
	if *concurrencyLevel < 1 {
		fmt.Fprintf(os.Stderr, "Error: concurrency level can't be less than 1\n")
		os.Exit(1)
	}
	// firts block of code
	lines := flag.Args()
	// if flag -i specified and urls are from file and not from args
	if *inputFile != "" {
		//internal/utils/parser
		lines, err = utils.ReadLines(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
	}
	// second block || internal/utils/parser
	urls, err := utils.StringToURL(lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing URLs: %v\n", err)
		os.Exit(1)
	}
	// third block
	err = utils.DownloadURLs(urls, *concurrencyLevel, *outputDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading: %v\n", err)
		os.Exit(1)
	}
}
