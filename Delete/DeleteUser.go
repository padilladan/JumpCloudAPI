package main

import (
	"bufio"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	apiKey := os.Getenv("api")
	if apiKey == "" {
		fmt.Println("No API key provided, exiting.")
		os.Exit(1)
	} else {
		fmt.Println("The current API key used: ", apiKey)
	}

	//ID Number input
	fmt.Println("Enter ID: ")
	ID := bufio.NewReader(os.Stdin)
	const inputdelimiter = '\n'
	selectedID, err := ID.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	selectedID = strings.Replace(selectedID, "\n", "", -1)

	url := "https://console.jumpcloud.com/api/systemusers/" + selectedID
	method := "DELETE"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("x-api-key", apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	apiKey = ""
}