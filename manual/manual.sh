#!/usr/bin/env bash

# Assign some variables for easier readability/modification
PROJDIR="$HOME/Sync/Projects/capa-full-byoi"
STACK="cesdemo"
BASECMD="pulumi stack output -C ${PROJDIR} -s ${STACK}"
VPCFIELD="vpcId"
PRVFIELD="privSubnetIds"
PUBFIELD="pubSubnetIds"

# Make a copy of the original file to work on
cp cesdemo.yaml modified.yaml

# Add the VPC ID
yq w -i -d1 modified.yaml 'spec.networkSpec.vpc.id' $(${BASECMD} ${VPCFIELD})

# Add the private subnets
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PRVFIELD} | jq -r '.[0]')
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PRVFIELD} | jq -r '.[1]')
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PRVFIELD} | jq -r '.[2]')
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PRVFIELD} | jq -r '.[3]')

# Add the public subnets
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PUBFIELD} | jq -r '.[0]')
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PUBFIELD} | jq -r '.[1]')
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PUBFIELD} | jq -r '.[2]')
yq w -i -d1 modified.yaml 'spec.networkSpec.subnets[+].id' \
$(${BASECMD} ${PUBFIELD} | jq -r '.[3]')
