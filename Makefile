
ci-test:
	go test  ./...
go:
	rm -rf /docs \
	&& rm model_*
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli:v7.1.0 generate \
      -i https://raw.githubusercontent.com/conekta/openapi/main/_build/api.yaml \
      -g go \
      -o /local \
      -c /local/config-go.json \
      --global-property modelTests=false \
      --additional-properties=hideGenerationTimestamp=true
