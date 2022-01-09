package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"

	"github.com/AlexTsIvanov/WebScrapper/pkg/image"
	"github.com/AlexTsIvanov/WebScrapper/pkg/links"
)

func makeFolder(name string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s/%s", path, name), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var folder string
	fmt.Print("Enter folder name: ")
	fmt.Scanln(&folder)
	var url string
	fmt.Print("Enter URL: ")
	fmt.Scanln(&url)
	var maxRoutines int
	fmt.Print("Enter number of routines: ")
	fmt.Scanln(&maxRoutines)
	var maxResults int32
	fmt.Print("Enter max results number: ")
	fmt.Scanln(&maxResults)

	err := makeFolder(folder)
	if err != nil {
		log.Fatal(err)
	}

	uniqueHref := make(map[string]struct{})
	uniqueImg := make(map[string]struct{})

	var wg sync.WaitGroup
	var counter int32
	ch := make(chan string, maxResults)
	ch <- url
	for i := 1; i <= maxRoutines; i++ {
		wg.Add(1)
		go func(ch chan string) {
			defer wg.Done()
			for url := range ch {
				if counter >= maxResults {
					close(ch)
					break
				}
				atomic.AddInt32(&counter, 1)

				siteslice, err := links.GetHrefLinks(url, uniqueHref)
				if err != nil {
					fmt.Println("Wrong Link")
				}

				imgslice, err := links.GetImgLinks(url, uniqueImg)
				if err != nil {
					fmt.Println("Wrong Link")
				}

				for _, imgUrl := range imgslice {
					err = image.SaveImg(imgUrl, folder)
					if err != nil {
						log.Fatal(err)
					}
				}
				for _, v := range siteslice {
					ch <- v
				}
			}
		}(ch)

	}
	wg.Wait()
	http.NewServeMux()
}
