package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	Data string
	Err  error
}

var (
	workers    = 4
	numOfTasks = 20
)

func drain(ch chan result) {
	for k := 0; k < numOfTasks; k++ {
		d := <-ch
		log.Printf("Message is %+v\n", d)
	}
}

func hitGoogle(ch chan (result), counter int) {

	time.Sleep(time.Second * 1)

	log.Println("Invoked website, current count: ", counter)

	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		ch <- result{Err: err}
		return
	}
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		ch <- result{Err: err}
		log.Println("Error in making request")
		return
	}
	defer resp.Body.Close()
	ch <- result{Data: fmt.Sprintf("Success retrieve: %v Made the request at %+v", counter, time.Now())}
}

func esusuWay() {
	ch := make(chan result, workers)

	for i := 0; i < numOfTasks; i++ {
		go hitGoogle(ch, i)
	}

	drain(ch)
}

func main() {
	esusuWay()
}
