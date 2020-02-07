package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type systemStruct []struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	CompiledAttributes struct {
		Sudo struct {
			Enabled         bool `json:"enabled"`
			WithoutPassword bool `json:"withoutPassword"`
		} `json:"sudo"`
	} `json:"compiledAttributes"`
	Paths [][]struct {
		Attributes struct {
			Sudo struct {
				Enabled         bool `json:"enabled"`
				WithoutPassword bool `json:"withoutPassword"`
			} `json:"sudo"`
		} `json:"attributes"`
		To struct {
			Attributes interface{} `json:"attributes"`
			ID         string      `json:"id"`
			Type       string      `json:"type"`
		} `json:"to"`
	} `json:"paths"`
}

func main() {
	const commandID = "5dfba32fc8965f1cf2563a02"
	userID := ""
	api := ""

	listID := getSystems(userID, api)

	// Add the system/s to the list of computers that the command will run on////////////////////////////
	for _, a := range listID {
		url := "https://console.jumpcloud.com/api/v2/systems/" + a + "/associations"
		method := "POST"

		payload := strings.NewReader(" {\n   \"attributes\": {\n      \"sudo\": {\n         \"enabled\": true,\n         \"withoutPassword\": false\n      }\n   },\n    \"op\": \"add\",\n    \"type\": \"command\",\n    \"id\": \"" + commandID + "\"\n}")

		client := &http.Client{
		}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("x-api-key", api)

		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
	}
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
	req2.Header.Add("x-api-key", api)
	res2, err := client2.Do(req2)
	defer res2.Body.Close()
	_, err = ioutil.ReadAll(res2.Body)
	////////////////////////////////////////////////////////////////////////////////////////////////////

	// Remove the system/s from  the list of computers that the command will run on////////////////////////////
	for _, a := range listID {
		url3 := "https://console.jumpcloud.com/api/v2/systems/" + a + "/associations"
		method3 := "POST"

		payload3 := strings.NewReader(" {\n   \"attributes\": {\n      \"sudo\": {\n         \"enabled\": true,\n         \"withoutPassword\": false\n      }\n   },\n    \"op\": \"remove\",\n    \"type\": \"command\",\n    \"id\": \"" + commandID + "\"\n}")

		client3 := &http.Client{
		}
		req3, err := http.NewRequest(method3, url3, payload3)

		if err != nil {
			fmt.Println(err)
		}
		req3.Header.Add("Accept", "application/json")
		req3.Header.Add("Content-Type", "application/json")
		req3.Header.Add("x-api-key", api)

		res3, err := client3.Do(req3)
		defer res3.Body.Close()
		_, err = ioutil.ReadAll(res3.Body)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////
}

func getSystems(a, x string,) []string{
	// Get all the systems for the user///////////////////////////////////
	url4 := "https://console.jumpcloud.com/api/v2/users/" + a + "/systems"
	method4 := "GET"

	client4 := &http.Client {
	}
	req4, err := http.NewRequest(method4, url4, nil)

	if err != nil {
		fmt.Println(err)
	}
	req4.Header.Add("Accept", "application/json")
	req4.Header.Add("Content-Type", "application/json")
	req4.Header.Add("x-api-key", x)

	res4, err := client4.Do(req4)
	defer res4.Body.Close()
	body4, err := ioutil.ReadAll(res4.Body)

	jumpCloudData := systemStruct{}

	err = json.Unmarshal(body4, &jumpCloudData)

	var listID []string

	for i, _ := range jumpCloudData{
		systemID := jumpCloudData[i].ID
		listID = append(listID, systemID)
	}
	//////////////////////////////////////////////////////////////////////
	return listID
}

