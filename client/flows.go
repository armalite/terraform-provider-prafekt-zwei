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

type CreateFlowResponse struct {
	Id 			string
	Created		string
	Updated		string
	Name 		string
}

type ReadFlowResponse struct {
	Id 			string
	Created		string
	Updated		string
	Name 		string
	Tags		[]string
}


func CreateFlow(ctx context.Context, client Client, accountid string, workspaceid string, flowName string) CreateFlowResponse { 
	post_url := prefect_base_url + "accounts/" + accountid + "/workspaces/" + workspaceid + "/flows/"
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s"}`, flowName))
    req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

	var response CreateFlowResponse

    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}

func ReadFlow(ctx context.Context, client Client, accountid string, workspaceid string, flowId string) ReadFlowResponse {
	post_url := prefect_base_url + "accounts/" + accountid + "/workspaces/" + workspaceid + fmt.Sprintf("/flows/%s", flowId)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response ReadFlowResponse

    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}

func ReadFlowByName(ctx context.Context, client Client, accountid string, workspaceid string, name string) ReadFlowResponse {
	post_url := prefect_base_url + "accounts/" + accountid + "/workspaces/" + workspaceid + fmt.Sprintf("/flows/name/%s", name)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Content-Type", "application/json")
	
    resp, err := client.HTTPClient.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response ReadFlowResponse

    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	return response
}
