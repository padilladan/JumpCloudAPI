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


	// Searched for the staff
	fmt.Print("\nEnter the month: ")
	reader := bufio.NewReader(os.Stdin)
	const inputdelimiter = '\n'
	searchedMonth, err := reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert CRLF to LF
	searchedMonth = strings.Replace(searchedMonth, "\n", "", -1)


	fmt.Println("The current API key used: ", apiKey)
	fmt.Println("Searching for: ", searchedMonth, "\n")
	if apiKey == "" {
		log.Error("No API key provided, exiting.")
		os.Exit(1)
	}
	if searchedMonth == "" {
		log.Error("No searched user existing.")
		os.Exit(1)
	}

	config := JumpCloudInit(apiKey)
	users := config.getLdapUsers()

	fmt.Println(users)

	/*for _, x := range users {
		for _, y := range x.Results {
			if y.password_expiration_date == searchedMonth {
				fmt.Println("FOUND")
			}
		}
	}*/
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