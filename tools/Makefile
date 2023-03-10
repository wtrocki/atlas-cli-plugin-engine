export OPENAPI_FOLDER=../openapi
export OPENAPI_FILE_NAME=atlas-api.yaml
export SDK_FOLDER=../client

.DEFAULT_GOAL := generate_client

.PHONY: clean_client
clean_client:
	rm -Rf ../client

.PHONY: fetch_openapi
fetch_openapi:
	./scripts/fetch.sh

.PHONY: transform_openapi
transform_openapi:
	./scripts/transform.sh

.PHONY: generate_client
generate_client:
	./scripts/generate.sh

.PHONY: clean_and_generate
clean_and_generate: clean_client generate_client


.PHONY: install_generator
install_generator: 
	go get -u github.com/rigetti/openapi-cli-generator

.PHONY: generate_cli_diff
generate_cli_diff: 
	cp ../openapi/atlas-api.yaml ../openapi/atlas-api-revision.yaml
	git checkout HEAD^ -- ../openapi/atlas-api.yaml
	oasdiff --exclude-endpoints --format=json --base ../openapi/atlas-api.yaml \
		 --revision ../openapi/atlas-api-revision.yaml > ../openapi/openapi.diff 
	cat ../openapi/openapi.diff | jq '.paths.modified | keys'
	


