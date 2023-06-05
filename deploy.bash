#!/bin/bash

# get endpoint
endpoint=$1

if [ -z "$endpoint" ]; then
    echo "required endpoint argument not passed"
    exit 1
fi

# generate a seed
seed=$(openssl rand -hex 16)

# register DID
url="http://dev.bcovrin.vonx.io/register"
json_data='{"role": "ENDORSER", "alias": null, "did": null, "seed": "'$seed'"}'
response=$(curl -s -X POST -H "Content-Type: application/json" -d "$json_data" "$url")

# loggin
echo "Seed: $seed"
echo "Endpoint:$endpoint"

# Deploy agent
export WALLET_SEED=$seed
export ENDPOINT=$endpoint
docker compose up -d