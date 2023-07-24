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
	"terraform-provider-prafekt-zwei/client"
	"fmt"
	"context"
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

	ctx := context.Background()

	var client client.Client
	client = client.PrefectClient(ctx, os.Getenv("PREFECT_API_KEY"))

	var create_flow_response client.CreateFlowResponse
	create_flow_response = client.CreateFlow(ctx, client, prefect_account_id, prefect_workspace_id, "go-with-the-flow")
	
	var read_flow_response client.ReadFlowResponse
	read_flow_response = client.ReadFlow(ctx, client, prefect_account_id, prefect_workspace_id, create_flow_response.Id)
	fmt.Println("Read flow name:", read_flow_response.Name)

	var read_flow_by_name_response client.ReadFlowResponse
	read_flow_by_name_response = client.ReadFlowByName(ctx, client, prefect_account_id, prefect_workspace_id, read_flow_response.Name)
	fmt.Println("Read flow by name response:", read_flow_by_name_response)

	var create_work_queue_response client.CreateWorkQueueResponse
	create_work_queue_response = client.CreateWorkQueue(ctx, client, prefect_account_id, prefect_workspace_id, "such-a-cool-work-queue", "Work queue created via api", "false", 0, []string{}, []string{})
	fmt.Println("Create work queue response:", create_work_queue_response)

	var read_work_queue_response client.ReadWorkQueueResponse
	read_work_queue_response = client.ReadWorkQueue(ctx, client, prefect_account_id, prefect_workspace_id, create_work_queue_response.Id)
	fmt.Println("Read work queue response:", read_work_queue_response)
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