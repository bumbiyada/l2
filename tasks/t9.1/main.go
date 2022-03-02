package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// returns filepath local
func GetFilePath(outputDirectory string, fileName string) (filePath string) {
	if fileName == "/" {
		fileName = "index.html"
	}
	filePath = path.Join(outputDirectory, fileName)
	if filePath == "." {
		filePath = "index.html"
	}
	hasSuffix := regexp.MustCompile(`[.][a-z]{2,}$`)
	if !hasSuffix.Match([]byte(filePath)) {
		filePath = filePath + ".html"
	}
	if strings.HasSuffix(filePath, "/") || strings.HasSuffix(filePath, `\`) {
		filePath = filePath + "index.html"
	}
	return filePath
}

//create dir
func createDir(url *url.URL) {
	// MAKE Directory for file
	dirName := path.Join(url.Host, url.Path)
	if url.Path == "/" || url.Path == `\` {
		log.Println("url.Path EMPTY")
		err := os.Mkdir(url.Host, os.FileMode(600))
		if err != nil {
			log.Println("Can`t create Directory")
		}
		log.Println("dir created ", url.Host)
	} else {
		err := os.MkdirAll(path.Dir(dirName), os.FileMode(600))
		if err != nil {
			log.Println("can`t create Directory tree")
		}
		log.Println("dir created", path.Dir(dirName))
	}
}

// outputDirectory string, fileName string, url string
func fetch(depth int, urlname string, fileName string, outputDirectory string) error {
	var wg sync.WaitGroup

	filePath := GetFilePath(outputDirectory, fileName)
	// create file
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatalln(filePath, err)
		return err

	}
	defer out.Close()

	//get data from url
	resp, err := http.Get(urlname)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// if this is leaf, then just get file, else read file and try replace url in []byte
	if depth == 0 {
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Println("can`t copy body to file")
		}
	} else {
		// read data
		buffer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		// parse URLS
		regexUrl := regexp.MustCompile(`"(http[s]?|ftp):\/\/[a-zA-Z\d\/.=?#-]*"`)
		arrayUrls := regexUrl.FindAll(buffer, -1)
		if len(arrayUrls) == 0 {
			fmt.Println("Didn`t find any urls")
		}

		for _, rawurl := range arrayUrls {
			// preparation
			tmp := strings.ReplaceAll(string(rawurl), `"`, ``)
			url, err := url.Parse(tmp)

			if err != nil {
				fmt.Println(tmp, err)
			}

			// MAKE Directory for file
			createDir(url)

			// replace in text links
			workdir, err := os.Getwd()
			if err != nil {
				log.Fatalln(tmp, err)
			}
			replaceUrl := path.Join(workdir, GetFilePath(url.Host, url.Path))
			log.Printf("replacing %s by this %s", tmp, replaceUrl)

			buffer = bytes.Replace(buffer, []byte(tmp), []byte(replaceUrl), 1)
			// recursively fetch other branches
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				err = fetch((depth - 1), url.String(), url.Path, url.Host)
				if err != nil {
					log.Fatalln(tmp, err)
				}
			}(&wg)

		}
		log.Println("copying ")
		_, err = out.Write(buffer)
		if err != nil {
			log.Println("can`t write to file")
		}
		wg.Wait()

	}

	return nil
}

func main() {
	//flags
	rec_depth := *flag.Int("r", 1, "reccursion depth")
	base_path := *flag.String("o", ".", "set output dir, default is current '.'")
	flag.Parse()
	// parsing url
	rawurl := flag.Arg(0)
	url, err := url.Parse(rawurl)
	//starting fetching
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err = fetch(rec_depth, url.String(), path.Base(url.Path), base_path)
		if err != nil {
			log.Println(err)
		}
	}(&wg)
	wg.Wait()
	log.Println("COMPLITED")
}
