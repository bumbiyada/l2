package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// cfg struct
type cfg struct {
	after   int
	before  int
	context int
	count   int
	ignore  bool
	invert  bool
	fixed   bool
	line    bool
}

// nef config of flags
func NewCfg(a, b, context, count int, ignore, invert, f, l bool) cfg {
	var c cfg
	c.after = a
	c.before = b
	c.context = context
	c.count = count
	c.ignore = ignore
	c.invert = invert
	c.fixed = f
	c.line = l
	return c
}

// other variables
var root, query string
var exit_code = 1
var wg sync.WaitGroup

// compare ints
func return_bigger_int(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// prints before values
func print_before(arr []string, n int, cfg *cfg) {
	len := len(arr)                                                // 4
	offset := len - return_bigger_int(cfg.before, cfg.context) - 1 // 2
	if offset < 0 {
		offset = 0
	}
	for i := offset; i < len-1; i++ {
		fmt.Print("before\t")
		print_result(arr[i], i+1, cfg.line)
	}
}

//printing result
func print_result(text string, n int, flag bool) {
	if flag == true {
		fmt.Printf("%d: %s\n", n, text)
	} else {
		fmt.Printf("%s\n", text)
	}
}

// finding result
func find_result(txt, query string, cfg *cfg) (result bool) {
	if cfg.ignore == true {
		txt = strings.ToLower(txt)
		query = strings.ToLower(query)
	}
	if cfg.fixed == false {
		return strings.Contains(txt, query)
	} else {
		return txt == query
	}
}

//main function of searching
func readFile(wg *sync.WaitGroup, path string, cfg *cfg) {
	defer wg.Done()
	// open file
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}
	// new scanner
	scanner := bufio.NewScanner(file)
	// READING FILE
	//
	// case we have no counter
	if cfg.count < 0 {
		// case we have to invert all  stuff
		if cfg.invert == true {
			for i := 1; scanner.Scan(); i++ {
				tmp := scanner.Text()
				if !find_result(tmp, query, cfg) {
					exit_code = 0

					print_result(tmp, i, cfg.line)
				}
			}
		} else {
			text_buffer := []string{}
			var n_after int
			for i := 1; scanner.Scan(); i++ {
				tmp := scanner.Text()
				text_buffer = append(text_buffer, tmp)
				if n_after != 0 {
					fmt.Print("after\t")
					print_result(tmp, i, cfg.line)
					n_after--
				}
				if find_result(tmp, query, cfg) {
					exit_code = 0
					print_before(text_buffer, i, cfg)
					fmt.Print("\t")
					print_result(tmp, i, cfg.line)
					n_after = return_bigger_int(cfg.after, cfg.context)
				}

			}
		}
	} else {
		// case we have counter
		cnt := cfg.count
		// case we have to invert
		if cfg.invert == true {
			for i := 1; scanner.Scan(); i++ {
				tmp := scanner.Text()
				if cnt == 0 {
					return
				}
				if !find_result(tmp, query, cfg) {
					exit_code = 0
					cnt--
					print_result(tmp, i, cfg.line)
				}
			}
		} else {
			text_buffer := []string{}
			var n_after int
			for i := 1; scanner.Scan(); i++ {
				tmp := scanner.Text()
				text_buffer = append(text_buffer, tmp)
				if cnt == 0 && n_after == 0 {
					return
				}
				if n_after != 0 {
					fmt.Print("after\t")
					print_result(tmp, i, cfg.line)
					n_after--
				}
				if find_result(tmp, query, cfg) {
					exit_code = 0
					cnt--
					print_before(text_buffer, i, cfg)
					print_result(tmp, i, cfg.line)
					n_after = return_bigger_int(cfg.after, cfg.context)
				}

			}
		}
	}

}

func main() {
	//flags of application
	f_after := flag.Int("A", 0, "After")
	f_before := flag.Int("B", 0, "Before")
	f_context := flag.Int("C", 0, "Context A + B")
	f_count := flag.Int("c", -1, "Count")
	f_ignore_case := flag.Bool("i", false, "Ignore case")
	f_invert := flag.Bool("v", false, "Invert")
	f_fixed := flag.Bool("F", false, "Fixed")
	f_line_num := flag.Bool("n", false, "Line numeration")

	flag.Parse()
	// args of application
	query = flag.Arg(0)
	root = flag.Arg(1)
	// get all flags and put them in struct
	cfg := NewCfg(*f_after, *f_before, *f_context, *f_count, *f_ignore_case, *f_invert, *f_fixed, *f_line_num)
	if query == "" || root == "" {
		fmt.Println("No arguments are passed\nPass query string in first arg and dir. in seccond")
		os.Exit(0)
	}
	// starting searching in files
	filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			wg.Add(1)
			go readFile(&wg, path, &cfg)
		}
		return nil
	})
	wg.Wait()
	defer os.Exit(exit_code)
}
