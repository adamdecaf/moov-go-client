SWAGGER_VERSION := v3.0.0

.PHONY: client
client:
	mkdir -p .tmp/
# Download swagger dependencies
	if [ ! -f .tmp/swagger-codegen-$(SWAGGER_VERSION).zip ]; then \
		wget -O .tmp/swagger-codegen-$(SWAGGER_VERSION).zip https://github.com/swagger-api/swagger-codegen/archive/$(SWAGGER_VERSION).zip; \
	fi
# Extract swagger-codegen zip 
	if [ ! -d .tmp/swagger-codegen-$(shell echo $(SWAGGER_VERSION) | tr -d 'v') ]; then \
		unzip -d .tmp/ -o -qq .tmp/swagger-codegen-$(SWAGGER_VERSION).zip; \
	fi
# Download [latest] OpenAPI spec
	wget -O openapi.yml https://raw.githubusercontent.com/adamdecaf/moov-api/auth-docs/openapi.yaml
# Generate client
	cd .tmp/swagger-codegen-$(shell echo $(SWAGGER_VERSION) | tr -d 'v') && \
	mvn clean package && \
	java -jar modules/swagger-codegen-cli/target/swagger-codegen-cli.jar generate -i openapi.yaml -l go -o ../../client/
