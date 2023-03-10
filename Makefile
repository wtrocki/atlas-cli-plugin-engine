

## Example of orgs cli
.PHONY: orgs
orgs:
	go build -o "${GOPATH}/bin/orgs" ./cmd/orgs/*
