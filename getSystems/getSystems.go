package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
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

	userID := ""
	api := ""

	url := "https://console.jumpcloud.com/api/v2/users/" + userID + "/systems"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", api)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	jumpCloudData := systemStruct{}

	err = json.Unmarshal(body, &jumpCloudData)

	var listID []string

	for i, _ := range jumpCloudData{
		systemID := jumpCloudData[i].ID
		listID = append(listID, systemID)
	}

	fmt.Println(listID)
}