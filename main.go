package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Uchencho/watchMan/config"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const link = "https://peak-tutors-ub.herokuapp.com/api/accounts/login/"

func main() {

	l := login{
		Username: config.PeakUsername,
		Password: config.PeakPassword,
	}
	reqBody, err := json.Marshal(l)
	if err != nil {
		fmt.Println("Error Marshalling request ", err)
	}

	req, err := http.NewRequest("POST", link, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error doing request ", err)
	}
	fmt.Println("Status code is ", resp.StatusCode)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Reading body ", err)
	}
	fmt.Println(string(body))
}
