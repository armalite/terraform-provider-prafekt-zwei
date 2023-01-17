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
	"fmt"
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

var prefect_account_id = os.Getenv("PREFECT_ACCOUNT_ID")
var prefect_workspace_id = os.Getenv("PREFECT_WORKSPACE_ID")

func main() {

	//post_url := prefect_base_url + "accounts/" + prefect_account_id + "/workspaces/" + prefect_workspace_id + "/"
	var create_flow_response rest.CreateFlowResponse
	create_flow_response = rest.CreateFlow(prefect_account_id, prefect_workspace_id, "go-with-no-flow")
	
	var read_flow_response rest.ReadFlowResponse
	read_flow_response = rest.ReadFlow(prefect_account_id, prefect_workspace_id, create_flow_response.Id)
	fmt.Println("Read flow name:", read_flow_response.Name)

	var read_flow_by_name_response rest.ReadFlowResponse
	read_flow_by_name_response = rest.ReadFlowByName(prefect_account_id, prefect_workspace_id, read_flow_response.Name)
	fmt.Println("Read flow by name response:", read_flow_by_name_response)
	//rest.read

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