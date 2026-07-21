
ci-test:
	go test  ./...
go:
	rm -rf ./docs && rm -rf /test \
	&& (rm model_* api_*.go || true) \
      && docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli:v7.24.0 generate \
      -i https://raw.githubusercontent.com/conekta/openapi/refs/heads/release/v2.3.0/_build/api.yaml \
      -g go \
      -o /local \
      -c /local/config-go.json \
      --global-property modelTests=false \
      --additional-properties=hideGenerationTimestamp=true
