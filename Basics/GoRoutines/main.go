package main

import (
	"fmt"
	"net/http"
	
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.in",
		"http://thebigk.ml",
	}

	c:=make(chan string)

	for _, link := range links {
		go checkLink(link,c)

	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
		
	}

}
func checkLink(link string ,c chan string) {
	_, err := http.Get(link) //blocking call
	if err != nil {
		fmt.Println(link + "Site is down")
		c<-"Might be down"
		

		return

	}
	fmt.Println(link, "is up")
	c<- "Yep its up"

}
