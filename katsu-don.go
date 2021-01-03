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
)

func main() {
	// User inputs the extension number via the -extNum argument.
	var extNum string
	flag.StringVar(&extNum, "extNum", "", "Extension Number.")

	// Same as above but with -fileName
	var fileName string
	flag.StringVar(&fileName, "fileName", "", "File Name.")

	flag.Parse()

	err := DownloadFile(fileName + ".xpi", "https://addons.mozilla.org/firefox/downloads/file/" + extNum)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded " + fileName + ".xpi")
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