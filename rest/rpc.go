package rest

import (
	"io/ioutil"
	"net/http"
	"bytes"
    "fmt"
	"os"
	"log"
	"encoding/json"
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

var prefect_api_key = os.Getenv("PREFECT_API_KEY")

func CreateFlow(posturl string, flowName string) CreateFlowResponse { 
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s"}`, flowName))
	post_url := posturl + "flows/"
    req, err := http.NewRequest("POST", post_url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", prefect_api_key))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

	var response CreateFlowResponse

    fmt.Println("response Status:", resp.Status)
	
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)
	fmt.Println("Created flow id:", response.Id)
	fmt.Println("Created flow name:", response.Name)
	return response
}

func ReadFlow(posturl string, flowId string) ReadFlowResponse {
	post_url := posturl + fmt.Sprintf("/flows/%s", flowId)
	req, err := http.NewRequest("GET", post_url, nil)
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", prefect_api_key))
    req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		log.Fatal(err.Error())
    }
    defer resp.Body.Close()

    var response ReadFlowResponse

    fmt.Println("response Status:", resp.Status)
	
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)
	fmt.Println("Created flow id:", response.Id)
	fmt.Println("Created flow name:", response.Name)

	return response
}
