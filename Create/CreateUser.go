package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	// Input new API key
	apiKey := os.Getenv("api")

	//USERNAME input
	fmt.Println("Enter username: ")
	username := bufio.NewReader(os.Stdin)
	const inputdelimiter = '\n'
	newUsername, err := username.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	newUsername = strings.Replace(newUsername, "\n", "", -1)

	//EMAIL input
	fmt.Println("Enter email: ")
	email := bufio.NewReader(os.Stdin)
	const inputdelimiter2 = '\n'
	newEmail, err := email.ReadString(inputdelimiter2)
	if err != nil {
		fmt.Println(err)
		return
	}
	newEmail = strings.Replace(newEmail, "\n", "", -1)

	//FIRSTNAME input
	fmt.Println("Enter first name: ")
	firstname := bufio.NewReader(os.Stdin)
	const inputdelimiter3 = '\n'
	newFirstname, err := firstname.ReadString(inputdelimiter3)
	if err != nil {
		fmt.Println(err)
		return
	}
	newFirstname = strings.Replace(newFirstname, "\n", "", -1)

	//LASTNAME input
	fmt.Println("Enter last name: ")
	lastname := bufio.NewReader(os.Stdin)
	const inputdelimiter4 = '\n'
	newLastname, err := lastname.ReadString(inputdelimiter4)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLastname = strings.Replace(newLastname, "\n", "", -1)

	url := "https://console.jumpcloud.com/api/systemusers"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("{\"username\":\"%v\",\"email\":\"%v\",\"firstname\":\"%v\",\"lastname\":\"%v\"}", newUsername, newEmail, newFirstname, newLastname))

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}