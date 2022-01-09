package image

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func SaveImg(url string, folder string) error {
	// url := "http://i.imgur.com/m1UIjW1.jpg"
	// don't worry about errors
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fileName, err := buildFileName(url)
	if err != nil {
		return err
	}

	//open a file for writing
	file, err := os.Create(fmt.Sprintf("./%s/%s", folder, fileName))
	if err != nil {
		fmt.Println("tuka neshto")
		return err
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func buildFileName(fullUrlFile string) (string, error) {
	fileUrl, err := url.Parse(fullUrlFile)
	if err != nil {
		return "", err
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	return segments[len(segments)-1], nil
}
