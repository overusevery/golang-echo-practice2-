.PHONY: coverage
coverage:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -html="profile.cov"
	rm profile.cov