package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {

		//	fmt.Println(err)
		s := fmt.Sprintf("%s is Down!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)
		c <- url // sending into the channel
	} else {

		s := fmt.Sprintf("%s -> status Code: %d \n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url

	}

}

func main() {

	urls := []string{"https://golang1.org", "https://google.com", "https://www.medium.com"}

	//1.
	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	/*for {
		go checkUrl(<-c, c)
		fmt.Println(strings.Repeat("#", 30))
		time.Sleep(time.Second * 2)
	}*/

	/*for url := range c {
		time.Sleep(time.Second * 2)
		go checkUrl(url, c)
	}*/

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)

		}(url)

	}
}
