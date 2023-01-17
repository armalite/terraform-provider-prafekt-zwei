package rest


import (
	"os"
	"net/http"
)

type HTTPTransport struct {
	RTripper 	http.RoundTripper
	Header		http.Header
}

var prefect_api_key = os.Getenv("PREFECT_API_KEY")
var client = &http.Client{}
var prefect_base_url = "https://api.prefect.cloud/api/" 