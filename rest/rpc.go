package rest

import (
	"io/ioutil"
	"net/http"
	"bytes"
    "fmt"
	"os"
	"log"
)

var prefect_api_key = os.Getenv("PREFECT_API_KEY")

func CreateFlow(flowName string, posturl string) { 
	var jsonStr = []byte(fmt.Sprintf(`{"name":"%s"}`, flowName))
    req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonStr))
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
}
