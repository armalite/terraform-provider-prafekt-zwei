package main

/*
import (
	"context"
	"log"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)
*/
import (
	"os"
	"terraform-provider-prafekt-zwei/rest"
)
// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

var prefect_base_url = "https://api.prefect.cloud/api/" 
var prefect_account_id = os.Getenv("PREFECT_ACCOUNT_ID")
var prefect_workspace_id = os.Getenv("PREFECT_WORKSPACE_ID")

func main() {

	post_url := prefect_base_url + "accounts/" + prefect_account_id + "/workspaces/" + prefect_workspace_id + "/flows/"
	rest.CreateFlow("go-with-no-flow", post_url)

	/* 
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/Armalite/prafekt-zwei",
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
	*/

}