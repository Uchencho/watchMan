package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"sync"
	"time"

	"github.com/Uchencho/watchMan/config"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	BlogPage = "https://uchencho.pythonanywhere.com/"
	host     = "smtp.gmail.com"
	port     = "587"
	address  = host + ":" + port
)

func sendMail(message []byte) error {

	from := config.GmailEmail
	password := config.GmailPassword
	recipient := []string{"aloziekelechi17@gmail.com"}

	// Authentication
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, recipient, message)
	if err != nil {
		fmt.Println("Could not send email with error: ", err)
		return err
	}
	fmt.Println("Email sent")
	return nil
}

func hitPeak() bool {
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
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK

}

func hitPythonAnywhere() bool {
	resp, err := http.Get(BlogPage)
	if err != nil {
		fmt.Println("Error doing request ", err)
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	startTime := time.Now()

	go func() {
		if hitPeak() {
			fmt.Println("\nPeak API running effectively")
		}
		wg.Done()
	}()

	go func() {
		if hitPythonAnywhere() {
			fmt.Println("\nBlog Page is running smoothly")
		}
		wg.Done()
	}()

	fmt.Println("Started running the concurrent program")
	wg.Wait()
	err := sendMail([]byte("All APIs are up and running, reporting from Golang headquarters"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("It took the concurrent program ", time.Since(startTime), "to finish")
	fmt.Println("\nFinished running the concurrent program")

}
