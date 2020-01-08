package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AutoGenerated struct {
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
		PasswordNeverExpires        bool          `json:"password_never_expires"`
		PasswordlessSudo            bool          `json:"passwordless_sudo"`
		PhoneNumbers                []interface{} `json:"phoneNumbers"`
		SambaServiceUser            bool          `json:"samba_service_user"`
		SSHKeys                     []interface{} `json:"ssh_keys"`
		Sudo                        bool          `json:"sudo"`
		Suspended                   bool          `json:"suspended"`
		UnixGUID                    int           `json:"unix_guid"`
		UnixUID                     int           `json:"unix_uid"`
		Username                    string        `json:"username"`
		Created                     time.Time     `json:"created"`
		PasswordExpirationDate      time.Time     `json:"password_expiration_date,omitempty"`
		PasswordExpired             bool          `json:"password_expired"`
		TotpEnabled                 bool          `json:"totp_enabled"`
		_ID                         string        `json:"_id"`
		ID                          string        `json:"id"`
		Displayname                 string        `json:"displayname,omitempty"`
		Description                 string        `json:"description,omitempty"`
		Middlename                  string        `json:"middlename,omitempty"`
		Company                     string        `json:"company,omitempty"`
		CostCenter                  string        `json:"costCenter,omitempty"`
		Department                  string        `json:"department,omitempty"`
		EmployeeIdentifier          interface{}   `json:"employeeIdentifier,omitempty"`
		EmployeeType                string        `json:"employeeType,omitempty"`
		ExternalDn                  string        `json:"external_dn,omitempty"`
		ExternalSourceType          string        `json:"external_source_type,omitempty"`
		JobTitle                    string        `json:"jobTitle,omitempty"`
		Location                    string        `json:"location,omitempty"`
		Mfa                         struct {
			Exclusion      bool        `json:"exclusion"`
			ExclusionUntil interface{} `json:"exclusionUntil"`
			Configured     bool        `json:"configured"`
		} `json:"mfa,omitempty"`
	} `json:"results"`
}

func main() {

	url := "https://console.jumpcloud.com/api/systemusers"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", "8699c0037f165deb108ca913fcc8a8ea853f5997")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}