package main

import (
	"context"
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
	apiKey := os.Getenv("api")
	fmt.Println("The current API key used: ", apiKey)
	if apiKey == "" {
		log.Error("No API key provided, exiting.")
		os.Exit(1)
	}

	config := JumpCloudInit(apiKey)
	users := config.getLdapUsers()

	for _, u := range users {
		fmt.Println(u)
	}


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

func (j *JumpCloudConfig) getLdapUsers() []string {
	var userList []string
	totalCount := 0
	currentCount := 0

	for paginate := true; paginate; paginate = (currentCount != totalCount) {
		ldapUsers, _, err := j.Client.SystemusersApi.SystemusersList(j.Auth, j.ContentType, j.Accept, map[string]interface{}{"skip": int32(currentCount)})

		if err != nil {
			log.Error(err)
		}

		for _, l := range ldapUsers.Results {
			userName := fmt.Sprintf("%s %s", l.Firstname, l.Lastname)
			userID := fmt.Sprintf("%s", l.Id)
			userList = append(userList, userName, userID)
		}

		totalCount = int(ldapUsers.TotalCount)
		currentCount += len(ldapUsers.Results)
	}

	return userList
}
