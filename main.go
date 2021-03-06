package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/fatih/color"
)

func printStatusCodeWithColor(code int) {
	white := color.New(color.FgWhite)
	black := color.New(color.FgBlack)

	if code >= 200 && code <= 226 {
		white.Add(color.BgGreen).Printf(" %d ", code)
	} else if code >= 300 && code <= 308 {
		black.Add(color.BgYellow).Printf(" %d ", code)
	} else if code >= 400 && code <= 499 {
		white.Add(color.BgRed).Printf(" %d ", code)
	} else {
		black.Add(color.BgWhite).Printf(" %d ", code)
	}

}

func validateURL(input string) bool {
	httpPrefix := strings.HasPrefix(input, "http")

	if !httpPrefix {
		return false
	}

	_, err := url.ParseRequestURI(input)

	if err != nil {
		return false
	}

	return true
}

func getStatusCode(url string) {

	if !validateURL(url) {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		printStatusCodeWithColor(resp.StatusCode)
		fmt.Println(" ", url)
	}
	resp.Body.Close()
}

func getAndParseFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// windows new line problem solved.
	fixedURLList := strings.ReplaceAll(string(data), "\r\n", "\n")
	urlList := strings.Split(fixedURLList, "\n")
	return urlList, nil
}

func main() {
	url := flag.String("u", "", "Give an url")
	path := flag.String("f", "", "Give a file path")

	flag.Parse()

	getStatusCode(*url)

	if len(*path) > 0 {
		if urlList, err := getAndParseFile(*path); err == nil {
			for _, item := range urlList {
				getStatusCode(item)
			}
		}
	} else {
		fmt.Println("You must specify a file path or url.")
	}
}
