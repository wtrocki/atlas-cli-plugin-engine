#!/usr/bin/env bash
set -o errexit
set -o nounset

######################################################
# Generate client using OpenAPI generator
# Environment variables:
#   OPENAPI_FOLDER - folder containing openapi file
#   OPENAPI_FILE_NAME - openapi file name
#   SDK_FOLDER - folder location for generated client
######################################################

OPENAPI_FOLDER=${OPENAPI_FOLDER:-./openapi}
OPENAPI_FILE_NAME=${OPENAPI_FILE_NAME:-atlas-api.yaml}
SDK_FOLDER=${SDK_FOLDER:-./}

transformed_file="atlas-api-transformed.yaml"
client_package="atlas"
openapiFileLocation="$OPENAPI_FOLDER/$transformed_file"

echo "# Running generation pipeline"
echo "# Running transformation from $transformed_file OpenAPI from $OPENAPI_FILE_NAME"
cp "$OPENAPI_FOLDER/$OPENAPI_FILE_NAME" "$openapiFileLocation"

npm install
npm run sdk:transform -- "$openapiFileLocation"

openapi-cli-generator generate $openapiFileLocation
mv ./atlas-api-transformed.go ../cmd/orgs/generated.go 
