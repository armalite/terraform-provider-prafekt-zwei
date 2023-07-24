package rest


import (
	"io/ioutil"
	"net/http"
    "fmt"
	"log"
	"encoding/json"
	"context"
)

type ReadUserResponse struct {
	Id 			string
	Created		string
	Updated		string
	ActorId		string
	FirstName	string
	LastName	string
	Email		string
	PersonalAccountId	string
	LastLogin	string
}

type CreateUserAPIKeyResponse struct {
	Id			string
	Created		string
	Name		string
	Expiration	string
	Key			string
}

type ReadCurrentUserResponse struct {
	Id 					string
	Handle				string
	FirstName			string
	LastName			string
	Email				string
	PersonalAccountId	string
}


func ReadUserById(ctx context.Context, client Client, userid string) ReadUserResponse { 
	post_url := prefect_base_url + fmt.Sprintf("users/%s", userid)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response ReadUserResponse
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}

func CreateUserAPIKey(ctx context.Context, client Client, userid string) CreateUserAPIKeyResponse {
	post_url := prefect_base_url + fmt.Sprintf("users/%s/api_keys", userid)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response CreateUserAPIKeyResponse
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}

func ReadCurrentUser(ctx context.Context, client Client) ReadCurrentUserResponse {
	post_url := prefect_base_url + "me/"
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response ReadCurrentUserResponse
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}