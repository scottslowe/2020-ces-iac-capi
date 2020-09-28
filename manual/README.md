# Manual Cluster API-IaC Integration

This directory provides a rudimentary example of manual integration between infrastructure-as-code implemented with Pulumi and Cluster API (CAPI).

The `cesdemo.yaml` file in this directory was generated using `clusterctl` with the following command:

    clusterctl config cluster cesdemo --kubernetes-version=v1.18.2 \
    --control-plane-machine-count=3 \
    --worker-machine-count=3 > cesdemo.yaml

Creating your own `cesdemo.yaml` file using this command will generate a file that is similar to the one here but not identical; your file will use your own AWS region (which _may_ not be the same as the one I'm using) and your own SSH key (which _definitely_ will not be the same one I'm using).

The `manual.sh` script illustrates one way to automate this level of integration. It relies upon the `yq` command (https://github.com/mikefarah/yq). You will have to modify this script in order for it work on your system; all the variables are defined at the top of the script:

* The `PROJDIR` variable should point to the directory where your Pulumi project is stored.
* The `STACK` variable should contain the name of the Pulumi stack from which you want to pull information.
* The `BASECMD` variable creates the base `pulumi` command necessary to retrieve information from the stack. You shouldn't need to modify this command.
* The `VPCFIELD` variable should contain the name of the exported field from your Pulumi stack that contains the ID of the VPC.
* Similarly, the `PUBFIELD` and `PRIVFIELD` variables should contain the names of the exported variables that contain the subnet IDs you want to use.

The script assumes that the Pulumi stack created subnets in four AZs. If your region or Pulumi stack contains more or less AZs, you'll need to adjust the script accordingly.
