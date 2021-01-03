/*
Katsu-Don: A Firefox Extensions Downloader

This is a tool that allows you to download extensions from addons.mozilla.org 
and saves them as a .xpi file.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var start string
	flag.StringVar(&start, "start", "", "Starting number.")

	var end string
	flag.StringVar(&end, "end", "", "Ending number.")

	flag.Parse()

	s, err := strconv.Atoi(start)
	e, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}

	for i := s; i <= e; i++ {
		num := strconv.Itoa(i)
		err := DownloadFile(num + ".xpi", "https://addons.mozilla.org/firefox/downloads/file/" + num)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded " + num + ".xpi")
	}
}

// DownloadFile will download a url to a local file
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}