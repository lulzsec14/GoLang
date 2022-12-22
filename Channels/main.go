package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range websites {
		go checkStatus(link, c)
	}

	// fmt.Println("Not paused!")

	// for i := 0; i < len(websites); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		// time.Sleep(time.Second * 5)
		// Not good approach as the main go routine will wait for 5 seconds
		// before recieving next fecth message which will lead piling up of
		// many fetch messages in channel

		// go checkStatus(l, c)

		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	checkStatus(l, c)
		// }()
		// If we do this then the child go routine will have access
		// to copy of varaiable l and not the original variable in
		// main go routine as (two different go routine are trying to access
		// to one common variable) l is outside the functions scope
		// to solve this problem the variable needs to be passed to
		// the function as an arguement


		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)

	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// fmt.Println(<-c)

}

func checkStatus(link string, c chan string) {
	// time.Sleep(time.Second * 5)
	// Not a good approach either as it will delay every fetching by 5 sconds
	// but main go routine will continue

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
