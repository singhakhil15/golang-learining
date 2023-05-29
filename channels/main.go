package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	// creating channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// receiving data from channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// Repreating Routine
	for _, link := range links {
		go checkLinkRepeate(link, c)
	}
	for l := range c {
		//go checkLinkRepeate(l, c)

		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLinkRepeate(li, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")

		// sending data from channel
		c <- "Might be down"
		return
	}

	fmt.Println(link, " is up")

	// sending data from channel
	c <- "Its up"
}

func checkLinkRepeate(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")

		// sending data from channel
		c <- link
		return
	}

	fmt.Println(link, " is up")

	// sending data from channel
	c <- link
}
