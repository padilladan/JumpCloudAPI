package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type bodystruct struct {
	TotalCount int `json:"totalCount"`
	Results    []struct {
		AccountLocked               bool          `json:"account_locked"`
		Activated                   bool          `json:"activated"`
		Addresses                   []interface{} `json:"addresses"`
		AllowPublicKey              bool          `json:"allow_public_key"`
		Attributes                  []interface{} `json:"attributes"`
		Email                       string        `json:"email"`
		EnableManagedUID            bool          `json:"enable_managed_uid"`
		EnableUserPortalMultifactor bool          `json:"enable_user_portal_multifactor"`
		ExternallyManaged           bool          `json:"externally_managed"`
		Firstname                   string        `json:"firstname"`
		Lastname                    string        `json:"lastname"`
		LdapBindingUser             bool          `json:"ldap_binding_user"`
		Mfa                         struct {
			Exclusion  bool `json:"exclusion"`
			Configured bool `json:"configured"`
		} `json:"mfa,omitempty"`
		PasswordNeverExpires   bool          `json:"password_never_expires"`
		PasswordlessSudo       bool          `json:"passwordless_sudo"`
		PhoneNumbers           []interface{} `json:"phoneNumbers"`
		SambaServiceUser       bool          `json:"samba_service_user"`
		SSHKeys                []interface{} `json:"ssh_keys"`
		Sudo                   bool          `json:"sudo"`
		Suspended              bool          `json:"suspended"`
		UnixGUID               int           `json:"unix_guid"`
		UnixUID                int           `json:"unix_uid"`
		Username               string        `json:"username"`
		Created                time.Time     `json:"created"`
		PasswordExpirationDate time.Time     `json:"password_expiration_date,omitempty"`
		PasswordExpired        bool          `json:"password_expired"`
		TotpEnabled            bool          `json:"totp_enabled"`
		_ID                     string        `json:"_id"`
		ID                     string        `json:"id"`
		Displayname            string        `json:"displayname,omitempty"`
		Description            string        `json:"description,omitempty"`
		Middlename         string      `json:"middlename,omitempty"`
		Company            string      `json:"company,omitempty"`
		CostCenter         string      `json:"costCenter,omitempty"`
		Department         string      `json:"department,omitempty"`
		EmployeeIdentifier interface{} `json:"employeeIdentifier,omitempty"`
		EmployeeType       string      `json:"employeeType,omitempty"`
		ExternalDn         string      `json:"external_dn,omitempty"`
		ExternalSourceType string      `json:"external_source_type,omitempty"`
		JobTitle           string      `json:"jobTitle,omitempty"`
		Location           string      `json:"location,omitempty"`
	} `json:"results"`
}

type User struct {
	Name string
	ExpirationDate time.Time
	Email string
	ID string
}

func main() {

	apiKey := os.Getenv("api")
	if apiKey == "" {
		fmt.Println("API Key missing")
		os.Exit(1)
	}
	fmt.Println("The current API key used: ", apiKey)

	url := "https://console.jumpcloud.com/api/systemusers"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	jcdata := bodystruct{}

	err = json.Unmarshal(body, &jcdata)
	if err != nil {
		fmt.Println(err)
	}

	count := 0
	today := time.Now()
	var passwordsExpiringSoon []User

	for _ = range jcdata.Results{
		t := jcdata.Results[count].PasswordExpirationDate
		firstLast := jcdata.Results[count].Firstname + " " + jcdata.Results[count].Lastname

		// Creates a struct of required user data - flush out unneeded info
		staffMember := User{
			Name: firstLast,
			ExpirationDate: t,
			Email: jcdata.Results[count].Email,
			ID: jcdata.Results[count].ID,
		}
		//////////////////////////////////////////////////////

		expirationDate := today.AddDate(0,1,0)

		// Only inserts users who have expiration dates within the next 30 days from today
		if t.Before(expirationDate) {
			if t.After(today) {
				passwordsExpiringSoon = append(passwordsExpiringSoon, staffMember)
			}
		}
		/////////////////////////////////////////////////////////////////////////////////
		count += 1
	}


	// Loop to organize the users by expiration date
	var orderedList []User
	a := len(passwordsExpiringSoon)
	complete := a * a

	for i := 0 ; i < a ; i++ {
		temp := i + 1
		if temp == a {
			orderedList = append(orderedList, passwordsExpiringSoon[i])
			break
		} else {
			if passwordsExpiringSoon[i].ExpirationDate.Before(passwordsExpiringSoon[temp].ExpirationDate) {
				orderedList = append(orderedList, passwordsExpiringSoon[i])
			} else {
				orderedList = append(orderedList, passwordsExpiringSoon[temp])
			}
		}
	}

	for i := 0 ; i == complete ; i ++ {
		temp := i + 1
		if temp == a {
			orderedList = append(orderedList, orderedList[i])
			break
		} else {
			if orderedList[i].ExpirationDate.Before(orderedList[temp].ExpirationDate) {
				orderedList = append(orderedList, orderedList[i])
			} else {
				orderedList = append(orderedList, orderedList[temp])
			}
		}
	}
	/////////////////////////////////////////////////



	fmt.Println(orderedList)
	fmt.Println(len(orderedList))
}
