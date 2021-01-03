/*
Katsu-Don: A Firefox Extensions Downloader

This is a tool that allows you to download extensions from addons.mozilla.org 
and saves them as a .xpi file.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Asks user for extension number and saves it as `extNum`
	// TODO: Make it so the user can search up the name and not look for the number needed.
	fmt.Println("Please input the extension number:")
	var extNum string
	fmt.Scanln(&extNum)

	// Asks user what to name said file to fileName
	// TODO: Automatically figure out the file name.
	fmt.Println("Input file name (.xpi is applied automatically)")
	var fileName string
	fmt.Scanln(&fileName)

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