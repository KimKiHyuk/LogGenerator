package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"sync"
	"time"
)

func GetBodyFromUrl(url string) {
	resp, _ := http.Get(url)
	robots, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", robots)
}


func main()  {

	for {
		var wait sync.WaitGroup
		wait.Add(1)

		go func(url string) {
			defer wait.Done()
			GetBodyFromUrl(url)
		}("https://google.com/robots.txt")

		wait.Wait()
		fmt.Println(time.Now())
		time.Sleep(3 * time.Second) // logging by 3 sec
	}
}