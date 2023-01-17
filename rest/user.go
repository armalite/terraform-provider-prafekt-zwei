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

func ReadUserById(posturl string, userid string) ReadUserResponse { 
	post_url := posturl + fmt.Sprintf("users/%s", userid)
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

//func CreateUserAPIKey(posturl)