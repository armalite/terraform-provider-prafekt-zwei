package rest


import (
	"os"
	"net/http"
)

type HTTPTransport struct {
	RTripper 	http.RoundTripper
	Header		http.Header
}

// Prefect 2 workspace
type Workspace struct { 
	Id		string
	name	string
}

// Prefect 2 HTTP Client with Workspace included
type Client struct {
	HTTPClient http.Client
	Workspace Workspace
}

var prefect_api_key = os.Getenv("PREFECT_API_KEY")
var client = &http.Client{}
var prefect_base_url = "https://api.prefect.cloud/api/" 