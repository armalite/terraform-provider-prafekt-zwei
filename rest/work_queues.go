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

type CreateWorkQueueResponse struct {
	Id 					string
	Created 			string
	Updated				string
	Name 				string
	Description 		string
	Is_Paused 			bool
	Concurrency_Limit 	int
}


func CreateWorkQueue(ctx context.Context, client Client, accountid string, workspaceid string, name string, description string, is_paused string, concurrency_limit int, deploymentids []string, tags []string) CreateServiceAccountResponse { 
	post_url := prefect_base_url + "accounts/" + accountid + "/workspaces/" + workspaceid + "/work_queues/"

	// TODO: Inject array of deployment ids and tags where provided
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s", "description": "%s", "is_paused": "%s", "concurrency_limit": "%d"}`, name, description, is_paused, concurrency_limit))
    req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

	
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