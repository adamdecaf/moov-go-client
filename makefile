.PHONY: client
client:
# Download [latest] OpenAPI spec
	rm -rf openapi.yaml
	@wget -q -O openapi.yaml https://raw.githubusercontent.com/moov-io/api/master/openapi.yaml
# Generate client
# Checkout https://github.com/OpenAPITools/openapi-generator/releases
	export OPENAPI_GENERATOR_VERSION=v4.0.0-SNAPSHOT
	chmod +x ./openapi-generator
	./openapi-generator generate -i openapi.yaml -g go -o ./client
	go build github.com/moov-io/go-client/client
	go test ./...
