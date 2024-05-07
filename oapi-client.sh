#!/bin/bash

curl -sL https://api.hellman.oxygen.dfds.cloud/ssu/api/swagger/v1/swagger.json -o api.json


if [ -d "openapiclient" ]; then
  rm -rf "openapiclient"
fi

mkdir "openapiclient"

oapi-codegen -package openapiclient -generate types,client api.json > openapiclient/openapiclient.gen.go

echo "oapi-codegen completed."

