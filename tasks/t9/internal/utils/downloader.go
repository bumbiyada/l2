package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

// DownloadURLs downloads all provided URLs concurrently. It limits
// the number of concurrent downloads to the provided concurrencyLevel.
func DownloadURLs(urls []url.URL, concurrencyLevel int, outputDirectory string) error {
	sem := NewSemaphore(concurrencyLevel)
	stop := make(chan struct{})
	done := false
	errors := make(chan error, len(urls))
	var wg sync.WaitGroup

	pp := NewProgressPrinter(len(urls))

	for i, u := range urls {
		wg.Add(1)

		go func(u url.URL, i int) {
			defer wg.Done()

			sem.Down()
			if u.Scheme == "http" || u.Scheme == "https" {
				err := httpFetch(outputDirectory, path.Base(u.Path), u.String(), &pp.progressWriters[i])
				errors <- err
			}
			sem.Up()

		}(u, i)
	}

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		wg.Wait()
		time.Sleep(1 * time.Second)
		ticker.Stop()
		close(stop)
		close(errors)
	}()

	for !done {
		select {
		case <-ticker.C:
			pp.PrintProgress()
		case <-stop:
			done = true
		}
	}

	var allErrors []string

	for err := range errors {
		if err != nil {
			allErrors = append(allErrors, err.Error())
		}
	}

	if len(allErrors) > 0 {
		return fmt.Errorf(strings.Join(allErrors, "\n"))
	}

	return nil
}

func httpFetch(outputDirectory string, fileName string, url string, pw *ProgressWriter) error {
	filePath := path.Join(outputDirectory, fileName)
	partFilePath := filePath + ".part"

	out, err := os.Create(partFilePath)
	if err != nil {
		log.Println("HERE IT IS")
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	pw.filename = fileName
	pw.startTime = time.Now()
	pw.total = uint64(resp.ContentLength)

	_, err = io.Copy(out, io.TeeReader(resp.Body, pw))
	if err != nil {
		log.Println("HERE IT IS 2")
		return err
	}

	pw.endTime = time.Now()
	log.Println(partFilePath)
	log.Println(filePath)
	err = os.Rename(partFilePath, filePath)
	if err != nil {
		log.Println("HERE IT IS 3")
		return err
	}

	return nil
}
