package main

import (
	"bufio"
	"context"
	"strings"
	"fmt"
	"os"

	jcapiv1 "github.com/TheJumpCloud/jcapi-go/v1"
	"github.com/labstack/gommon/log"
)

type JumpCloudConfig struct {
	Client      *jcapiv1.APIClient
	Auth        context.Context
	ContentType string
	Accept      string
}

func main() {
	// Input new API key
	apiKey := os.Getenv("api")


	// Searched for the staff to terminate
	fmt.Print("\nEnter name of staff: ")
	reader := bufio.NewReader(os.Stdin)
	const inputdelimiter = '\n'
	searchedUser, err := reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert CRLF to LF
	searchedUser = strings.Replace(searchedUser, "\n", "", -1)


	fmt.Println("The current API key used: ", apiKey)
	fmt.Println("Searching for: ", searchedUser, "\n")
	if apiKey == "" {
		log.Error("No API key provided, exiting.")
		os.Exit(1)
	}
	if searchedUser == "" {
		log.Error("No searched user existing.")
		os.Exit(1)
	}

	config := JumpCloudInit(apiKey)
	users := config.getLdapUsers()

	for _, u := range users {
		for _, p := range u.Results {
			if p.Firstname == searchedUser || p.Lastname == searchedUser {
				fmt.Println("Found user: ", p.Firstname, p.Lastname)
				fmt.Println(p.Id, "\n")
			}
		}
	}

	fmt.Println("Copy the ID number beneath the staff name to use in the termination process.")
	apiKey = ""
}

func JumpCloudInit(apiKey string) *JumpCloudConfig {
	// Instantiate the API client
	client := jcapiv1.NewAPIClient(jcapiv1.NewConfiguration())

	// Set up the API key via context
	auth := context.WithValue(context.TODO(), jcapiv1.ContextAPIKey, jcapiv1.APIKey{
		Key: apiKey,
	})

	config := JumpCloudConfig{
		Client:      client,
		Auth:        auth,
		ContentType: "application/json",
		Accept:      "application/json",
	}
	return &config
}

func (j *JumpCloudConfig) getLdapUsers() []jcapiv1.Systemuserslist {
	var userList []jcapiv1.Systemuserslist
	totalCount := 0
	currentCount := 0

	for paginate := true; paginate; paginate = (currentCount != totalCount) {
		ldapUsers, _, err := j.Client.SystemusersApi.SystemusersList(j.Auth, j.ContentType, j.Accept, map[string]interface{}{"skip": int32(currentCount)})

		if err != nil {
			log.Error(err)
		}

		userList = append(userList, ldapUsers)

		totalCount = int(ldapUsers.TotalCount)
		currentCount += len(ldapUsers.Results)
	}

	return userList
}