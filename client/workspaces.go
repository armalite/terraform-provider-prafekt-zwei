package rest


import (
	"io/ioutil"
	"net/http"
    "fmt"
	"log"
	"encoding/json"
	"context"
	"bytes"
)

type CreateWorkspaceResponse struct {
	Id 						string
	Created					string
	Updated					string
	AccountId				string
	Name					string
	Description				string
	Handle					string
	DefaultWorkspaceRoleId  string
}


func CreateWorkspace(ctx context.Context, client Client, accountid, name string, description string, handle string) CreateWorkspaceResponse { 
	post_url := prefect_base_url + "accounts/" + accountid + "/workspaces/" 
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s", "description": "%s", "handle": "%s"}`, name, description, handle))
    req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response CreateWorkspaceResponse
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}

