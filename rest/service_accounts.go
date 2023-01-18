package rest

import (
	"io/ioutil"
	"net/http"
	"bytes"
    "fmt"
	"log"
	"encoding/json"
	"context"
)

type ServiceAccountAPIKey struct {
	Id 			string
	Created 	string
	Name 		string
	Expiration 	string
	key 		string
}

type CreateServiceAccountResponse struct {
	Id 			string
	Created		string
	Updated		string
	AccountId	string
	Name 		string
	AccountRoleName string
	APIKey		ServiceAccountAPIKey
}


func CreateServiceAccount(ctx context.Context, client Client, accountid string, name string, api_key_expiration string, account_role_id string, workspace_role_id string) CreateServiceAccountResponse { 
	post_url := prefect_base_url + "accounts/" + accountid + "/bots/"
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s", "api_key_expiration": "%s", "account_role_id": "%s"}`, name, api_key_expiration, account_role_id))
    req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")


	// TODO: if workspace_role_id is specified, then generate a second
	// http request against the bots_access endpoint (not documented) to
	// assign a workspace role to the service account
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

	var response CreateServiceAccountResponse

    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}