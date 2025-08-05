SERVICE_NAME=volumez-openapi-client
PACKAGE_NAME=openapi
OUTPUT_OPENAPI_CLIENT_PATH=pkg/$(PACKAGE_NAME)
VOLUMEZ_OPENAPI_SPEC=openapi.yaml

.PHONY: generate-openapi-client clean validate-openapi-client generate-example

clean:
	rm -rf $(OUTPUT_OPENAPI_CLIENT_PATH)/*[!.openapi-generator-ignore]

generate-openapi-client: clean
	squire openapi generate-client --service-name $(SERVICE_NAME) --spec ${VOLUMEZ_OPENAPI_SPEC} --output $(OUTPUT_OPENAPI_CLIENT_PATH) --repo-id $(SERVICE_NAME) --skip-validate-spec --package-name $(PACKAGE_NAME)
	make validate-openapi-client

validate-openapi-client:
	go mod tidy
	go build ./...
	go vet ./...
