.PHONY: coverage
coverage:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -html="profile.cov"
	rm profile.cov

.PHONY: e2e
e2e:
	e2e/test.sh

.PHONY:view-e2e-coverage 
view-e2e-coverage:
	go tool covdata textfmt -i=covdatafiles -o profile.cov
	go tool cover -func="profile.cov"
	go tool cover -html="profile.cov"    
	rm profile.cov
