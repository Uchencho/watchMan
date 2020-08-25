package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Uchencho/watchMan/config"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const BlogPage = "https://uchencho.pythonanywhere.com/"

func hitPeak() {
	l := login{
		Username: config.PeakUsername,
		Password: config.PeakPassword,
	}
	reqBody, _ := json.Marshal(l)

	req, err := http.NewRequest("POST", config.Peaklink, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error doing request ", err)
	}
	fmt.Println("The Status code for peak is ", resp.StatusCode)
	defer resp.Body.Close()
}

func hitPythonAnywhere() {
	resp, err := http.Get(BlogPage)
	if err != nil {
		fmt.Println("Error doing request ", err)
	}
	defer resp.Body.Close()

	fmt.Println("The status code for blog page is ", resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	startTime := time.Now()

	go func() {
		hitPeak()
		wg.Done()
	}()

	go func() {
		hitPythonAnywhere()
		wg.Done()
	}()

	fmt.Println("Started running the concurrent program")
	wg.Wait()
	fmt.Println("It took the concurrent program ", time.Since(startTime), "to finish")
	fmt.Println("\nFinished running the concurrent program")

}
