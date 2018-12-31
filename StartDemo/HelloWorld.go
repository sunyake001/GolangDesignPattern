package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Run... function
func Run() {
	fmt.Print("Hello World.")
}

func main() {
	go Run()
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Printf("http Error : %v", err)
	}

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Contents Read Error: %v", err)
	}
	log.Printf("contents : %s", string(contents))

	defer res.Body.Close()
}
