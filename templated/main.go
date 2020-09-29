package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pulumi/pulumi/sdk/v2/go/x/auto"
)

// IaaSData stores data about the IaaS that Cluster API needs
type IaaSData struct {
	VpcID           string
	SubnetIds       []string
	SecurityGroupID string
	Hack            string
}

func main() {
	// Set up context and stack and get outputs
	ctx := context.Background()
	stackName := auto.FullyQualifiedStackName("scottslowe", "capa-full-byoi", "cesdemo")
	stack, _ := auto.SelectStackLocalSource(ctx, stackName, filepath.Join("../../../../Sync/Projects", "capa-full-byoi"))
	values, _ := stack.Outputs(ctx)

	// Set up a variable of type IaaSData
	var info IaaSData

	// Put the VPC ID from the stack into the info struct
	info.VpcID = values["vpcId"].Value.(string)

	// Put the bastion security group ID from the stack into the info struct
	info.SecurityGroupID = values["bastionSecGrpId"].Value.(string)

	// Get the public subnet IDs as a string slice and get how many
	var pubSubnetsOutput []interface{} = values["pubSubnetIds"].Value.([]interface{})
	var pubSubnetIds []string
	if pubSubnetsOutput != nil {
		for _, v := range pubSubnetsOutput {
			pubSubnetIds = append(pubSubnetIds, v.(string))
		}
	}

	// Get the private subnet IDs as a string slice and get how many
	var privSubnetsOutput []interface{} = values["privSubnetIds"].Value.([]interface{})
	var privSubnetIds []string
	if privSubnetsOutput != nil {
		for _, v := range privSubnetsOutput {
			privSubnetIds = append(privSubnetIds, v.(string))
		}
	}

	// Store a combined slice of all subnet IDs into the info struct
	info.SubnetIds = append(pubSubnetIds, privSubnetIds...)

	// Define the template file
	templateFiles := []string{"cesdemo.tmpl"}

	// Add the hack/workaround for subsequent Go templating by CAPI
	info.Hack = "'{{ ds.meta_data.local_hostname }}'"

	// Create and open a file for the rendered template
	f, err := os.Create("cesdemo.yaml")
	if err != nil {
		log.Panicf("error creating output file: %s", err.Error())
	}

	// Render the template
	tmpl := template.Must(template.New("cesdemo.tmpl").ParseFiles(templateFiles...))
	err = tmpl.Execute(f, info)
	if err != nil {
		log.Panicf("error rendering template: %s", err.Error())
	}
}
