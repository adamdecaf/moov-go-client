VERSION=v$(shell date -u +"%Y.%m.%d").1

.PHONY: client
client:
# Download [latest] OpenAPI spec
	rm -rf openapi.yaml
#	wget -q -O openapi.yaml https://github.com/moov-io/api/releases/download/master/openapi.yaml
	wget -q -O openapi.yaml https://raw.githubusercontent.com/moov-io/api/master/openapi.yaml
# Generate client
# Checkout https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=4.3.0 ./openapi-generator generate --package-name client -i openapi.yaml -g go -o ./client
	rm -f client/go.mod client/go.sum
	go fmt ./...
	go build github.com/moov-io/go-client/client
	go test ./...

.PHONY: clean
clean:
	@rm -rf client/
	@rm -f openapi-generator-cli-*.jar

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)
