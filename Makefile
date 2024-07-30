.PHONY: e2e
e2e:
	e2e/test.sh

.PHONY:view-e2e-coverage 
view-e2e-coverage:
	go tool covdata textfmt -i=covdatafiles -o profile.cov
	go tool cover -func profile.cov
	go tool cover -html profile.cov
	rm profile.cov

.PHONY:unitcover
unitcover:
	go test ./src/... -cover -coverprofile=_profile.cov
	cat _profile.cov | grep -v "src/repository" | grep -v "src/handler/openapigenmodel" > profile.cov
	go tool cover -func profile.cov
	@rm _profile.cov profile.cov


.PHONY:srcloc
srcloc:
	find ./src -name '*.go' | xargs wc -l


.PHONY:generateFromOpenApi
generateFromOpenApi:
	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/swagger.yaml \
    -g go-server \
	--global-property models \
	--type-mappings=integer=int \
    -o /local/src/handler/openapigenmodel
	# see:https://openapi-generator.tech/docs/generators/go-server/