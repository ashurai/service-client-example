.PHONY: test

test: 
	@go test -v -tags=integration .
