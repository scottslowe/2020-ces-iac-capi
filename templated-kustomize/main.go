package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pulumi/pulumi/sdk/v2/go/x/auto"
	"github.com/pulumi/pulumi/sdk/v2/go/x/auto/optup"
)

// IaaSData stores data about the IaaS that Cluster API needs
type IaaSData struct {
	VpcID           string
	SubnetIds       []string
	SecurityGroupID string
}

func main() {
	// Set up context and stack and get outputs
	ctx := context.Background()
	sourceStackName := auto.FullyQualifiedStackName("scottslowe", "capa-full-byoi", "cesdemo")
	sourceStack, _ := auto.SelectStackLocalSource(ctx, sourceStackName, filepath.Join("../../../../Sync/Projects", "capa-full-byoi"))
	values, _ := sourceStack.Outputs(ctx)

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
	templateFiles := []string{
		"awscluster-vpc-spec.go.tmpl",
		"cp-machinetemplate.go.tmpl",
		"wkr-machinetemplate.go.tmpl",
	}

	// Iterate over the list of template files and render a template
	for _, v := range templateFiles {
		o := strings.Split(v, ".")[0]
		f, err := os.Create(fmt.Sprintf("k8s/%s.yaml", o))
		if err != nil {
			log.Panicf("error creating output file: %s", err.Error())
		}
		t := template.Must(template.New(v).ParseFiles(v))
		err = t.Execute(f, info)
		if err != nil {
			log.Panicf("error rendering template: %s", err.Error())
		}
	}

	// Use the kustomize provider to apply the templated overlay files
	destStackName := auto.FullyQualifiedStackName("scottslowe", "k8s", "cesdemo")
	workDir := filepath.Join(".", "k8s")
	destStack, err := auto.UpsertStackLocalSource(ctx, destStackName, workDir)
	if err != nil {
		log.Printf("failed to create or select stack: %s", err.Error())
	}

	destStack.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: "us-west-2"})

	_, err = destStack.Refresh(ctx)
	if err != nil {
		log.Printf("failed to refresh stack: %s", err.Error())
	}

	stdoutStreamer := optup.ProgressStreams(os.Stdout)

	_, err = destStack.Up(ctx, stdoutStreamer)
	if err != nil {
		log.Printf("failed to refresh stack: %s", err.Error())
	}
}
