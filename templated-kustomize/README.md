# Templated IaC-Cluster API Integration via Kustomize

This directory provides a second example of using Go templating as a mechanism for integrating infrastructure-as-code implemented by Pulumi and Cluster API (CAPI). This example uses the Automation API (https://github.com/pulumi/pulumi/tree/master/sdk/go/x/auto) and standard Go templating functionality to create `kustomize` overlays for a base CAPI manifest.

Since this example leverages code from the "templating" example, much of the information in [the README for that directory](../templated/README.md) also applies here.

To run this example:

1. Make sure you have an active Kubeconfig pointing to a properly configured CAPI management cluster.
2. Make sure you've made the necessary changes to `main.go` to reflect your own Pulumi project, stack, and output variables.
3. Run `go run main.go` to execute the code.
