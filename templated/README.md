# Templated IaC-Cluster API Integration

This directory provides an example of using Go templating as a mechanism for integrating infrastructure-as-code implemented by Pulumi and Cluster API (CAPI).The example uses the new Automation API (https://github.com/pulumi/pulumi/tree/master/sdk/go/x/auto) and standard Go templating functionality.

The `cesdemo.tmpl` file is a templated CAPI manifest; it is based on the YAML manifest generated in [the manual integration example](../manual/README.md).

In order to use this example in your own environment, you would have to modify `main.go` to reference your own infrastructure stack. This means changes to the values specified in lines 24 and 25 of `main.go`. Depending on how the code for your stack is written, changes to lines 32, 35, 38-44, and 47-53 may also be necessary (these are the sections that retrieve outputs from the referenced stack).

To run this example, use `go run main.go` from this directory.
