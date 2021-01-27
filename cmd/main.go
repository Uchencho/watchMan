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
	workers    = 3
	numOfTasks = 10
)

func drain(ch chan result) {
	for k := 0; k < numOfTasks; k++ {
		d := <-ch
		log.Printf("Message is %+v\n", d)
	}
}

func testConcurrency() {

	log.Println("Started test")

	jobChan := make(chan int, numOfTasks)

	ch := make(chan result)

	for i := 0; i < workers; i++ {

		go invokeJobs(jobChan, ch) // Concurrently run 4 workers
	}

	for w := 0; w < numOfTasks; w++ {
		jobChan <- w
	}
	close(jobChan)

	drain(ch)

	log.Println("Done")
}

func invokeJobs(joblist <-chan int, results chan result) {
	for j := range joblist {
		hitGoogle(results, j)
	}
}

func hitGoogle(ch chan (result), counter int) {

	time.Sleep(time.Second * 1)

	log.Println("Invoked current count: ", counter)

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

func main() {
	testConcurrency()
}
