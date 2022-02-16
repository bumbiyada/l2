package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/chatzikalymnios/goget/internal/constants"
)

// ReadLines attempts to open the file indicated by filename and read it.
// It returns a slice containing all lines in the file and any error encountered.
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// StringToURL accepts a slice of raw URL strings and attempts to convert them
// to URL structures. It returns an error if the URL scheme is not supported.
// It returns a slice of URL structures and any error encountered.
func StringToURL(rawurls []string) ([]url.URL, error) {
	var urls []url.URL
	for _, rawurl := range rawurls {
		url, err := url.Parse(rawurl)
		if err != nil {
			return nil, err
		}

		if url.Scheme == "" {
			url.Scheme = constants.DefaultScheme
		}

		_, ok := constants.SupportedUrlSchemes[url.Scheme]
		if !ok {
			return nil, fmt.Errorf("url scheme not supported: %s in %s", url.Scheme, rawurl)
		}

		urls = append(urls, *url)
	}

	return urls, nil
}
