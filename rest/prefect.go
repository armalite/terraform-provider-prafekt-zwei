package rest


import (
	"os"
)

var prefect_api_key = os.Getenv("PREFECT_API_KEY")