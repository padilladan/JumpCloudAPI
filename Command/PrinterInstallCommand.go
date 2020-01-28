package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"net/http"
	//"io/ioutil"
	"os"
)

func main() {

	// Setup API/commandID/systemID///////////////////////////////////////////////////////////////////////////////
	apiKey := os.Getenv("api")
	if apiKey == "" {
		fmt.Println("No API key provided, exiting.")
		os.Exit(1)
	} else {
		fmt.Println("The current API key used: ", apiKey)
	}

	systemID := os.Getenv("system")
	if systemID == "" {
		fmt.Println("No systemID key provided, exiting.")
		os.Exit(1)
	} else {
		fmt.Println("The current systemID key used: ", systemID)
	}

	commandID := os.Getenv("command")
	if commandID == "" {
		fmt.Println("No commandID key provided, exiting.")
		os.Exit(1)
	} else {
		fmt.Println("The current commandID key used: ", commandID)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////





	// Add the system to the list of computers that the command will run on////////////////////////////
	url := "https://console.jumpcloud.com/api/v2/systems/" + systemID + "/associations"
	method := "POST"

	payload := strings.NewReader(" {\n   \"attributes\": {\n      \"sudo\": {\n         \"enabled\": true,\n         \"withoutPassword\": false\n      }\n   },\n    \"op\": \"add\",\n    \"type\": \"command\",\n    \"id\": \"" + commandID + "\"\n}")

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
	//////////////////////////////////////////////////////////////////////////////////////////////////



	// Run the command using the webhook////////////////////////////////////////////////////////////////
	url2 := "https://console.jumpcloud.com/api/command/trigger/InstallPrinter"
	method2 := "POST"
	client2 := &http.Client {
	}
	req2, err := http.NewRequest(method2, url2, nil)
	if err != nil {
		fmt.Println(err)
	}
	req2.Header.Add("x-api-key", apiKey)
	res2, err := client2.Do(req2)
	defer res2.Body.Close()
	_, err = ioutil.ReadAll(res2.Body)
	////////////////////////////////////////////////////////////////////////////////////////////////////





	// Remove the system to the list of computers that the command will run on////////////////////////////
	url3 := "https://console.jumpcloud.com/api/v2/systems/" + systemID + "/associations"
	method3 := "POST"

	payload3 := strings.NewReader(" {\n   \"attributes\": {\n      \"sudo\": {\n         \"enabled\": true,\n         \"withoutPassword\": false\n      }\n   },\n    \"op\": \"remove\",\n    \"type\": \"command\",\n    \"id\": \"" + commandID + "\"\n}")

	client3 := &http.Client {
	}
	req3, err := http.NewRequest(method3, url3, payload3)

	if err != nil {
		fmt.Println(err)
	}
	req3.Header.Add("Accept", "application/json")
	req3.Header.Add("Content-Type", "application/json")
	req3.Header.Add("x-api-key", apiKey)

	res3, err := client3.Do(req3)
	defer res3.Body.Close()
	_, err = ioutil.ReadAll(res3.Body)
	////////////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Println("Printers Installed")
}