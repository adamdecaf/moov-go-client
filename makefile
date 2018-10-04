.PHONY: client
client:
# Download [latest] OpenAPI spec
	@wget -q -O openapi.yaml https://raw.githubusercontent.com/moov-io/api/master/openapi.yaml
# Generate client
# Checkout https://github.com/OpenAPITools/openapi-generator/releases
	export OPENAPI_GENERATOR_VERSION=v3.3.0
	chmod +x ./openapi-generator
	./openapi-generator generate -i openapi.yaml -g go -o ./client
	go build github.com/moov-io/go-client/client
	go test ./...
