package twiliogo

import (
	"encoding/json"
	"net/url"
)


type Account struct {
	Sid           		string `json:"sid"`
	OwnerAccountSid  	string `json:"owner_account_sid"`
	FriendlyName    	string `json:"friendly_name"`
	Status    			string `json:"status"`
	DateCreated     	string `json:"date_created"`
	DateUpdated         string `json:"date_updated"`
	AuthToken           string `json:"auth_token"`
	Type 				string `json:"type"`
}

func GetAccount(client Client, sid string) (*Account, error) {
	var account *Account

	res, err := client.get(url.Values{}, ".json")

	if err != nil {
		return nil, err
	}

	account = new(Account)
	err = json.Unmarshal(res, account)

	return account, err
}