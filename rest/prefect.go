package rest


import (
	"os"
	"net/http"
	"fmt"
	"context"
)

type HTTPTransport struct {
	RTripper 	http.RoundTripper
	Header		http.Header
}

// Prefect 2 workspace
type Workspace struct { 
	Id		string
	Name	string
}

// Prefect 2 HTTP Client with Workspace included
type Client struct {
	HTTPClient http.Client
	//Workspace Workspace
}

var prefect_api_key = os.Getenv("PREFECT_API_KEY")
var client = &http.Client{}
const prefect_base_url = "https://api.prefect.cloud/api/" 


// implement the RoundTrip function - needed when generating Prefect Client
func (tr *HTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, values := range tr.Header {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	return tr.RTripper.RoundTrip(req)
}


func PrefectClient(ctx context.Context, apiKey string) (Client, error) {

	tr := HTTPTransport{
		RTripper: http.DefaultTransport,
		Header:       make(http.Header),
	}

	tr.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	http_client := http.Client{Transport: &tr}

	// TODO: Get current user
	// TODO: Get workspace info for current user

	var client Client
	client.HTTPClient = http_client

	return client, nil
}