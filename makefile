VERSION=v$(shell date -u +"%Y.%m.%d").1

.PHONY: client
client:
	@mkdir -p ./data/
	@rm -rf openapi.yaml
#	wget -q -O openapi.yaml https://raw.githubusercontent.com/moov-io/api/master/openapi.yaml
#
# curl -H 'Authorization: token INSERTACCESSTOKENHERE' \
#   -H 'Accept: application/vnd.github.v3.raw' \
#   -O \
#   -L https://api.github.com/repos/owner/repo/contents/path

	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=4.3.1 ./openapi-generator generate --package-name client -i openapi.yaml -g go -o ./client
	rm -f client/go.mod client/go.sum
	go fmt ./...
	go build github.com/moov-io/go-client/client
	go test ./...

.PHONY: clean
clean:
	@rm -rf client/ openapi-generator-cli-*.jar

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)
