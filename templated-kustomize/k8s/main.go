package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/kustomize"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new Kubernetes provider to render YAML to directory
		// kubeyaml, err := k8s.NewProvider(ctx, "kubeyaml", &k8s.ProviderArgs{
		// 	RenderYamlToDirectory: pulumi.String("rendered"),
		// })
		// if err != nil {
		// 	return err
		// }

		// Use Kustomize support to modify base YAML
		_, err2 := kustomize.NewDirectory(ctx, "tk", kustomize.DirectoryArgs{
			Directory: pulumi.String("."),
			//}, pulumi.Provider(kubeyaml))
		})
		if err2 != nil {
			return err2
		}

		return nil
	})
}
