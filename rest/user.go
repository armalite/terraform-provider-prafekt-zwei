package rest


import (
	"io/ioutil"
	"net/http"
    "fmt"
	"log"
	"encoding/json"
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


func ReadUserById(userid string) ReadUserResponse { 
	post_url := prefect_base_url + fmt.Sprintf("users/%s", userid)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", prefect_api_key))
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response ReadUserResponse

    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}

func CreateUserAPIKey(userid string) CreateUserAPIKeyResponse {

	post_url := prefect_base_url + fmt.Sprintf("users/%s/api_keys", userid)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", prefect_api_key))
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response CreateUserAPIKeyResponse

    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response

}