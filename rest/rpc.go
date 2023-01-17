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
	id 			string
	created		string
	updated		string
	name 		string
}

var prefect_api_key = os.Getenv("PREFECT_API_KEY")

func CreateFlow(posturl string, flowName string) []byte { 
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
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)
    fmt.Println("response Body:", string(body))
	fmt.Println("response parsed:", response)
	return body
}

func ReadFlow(posturl string, flowId string) []byte {
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

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

	return body
}
