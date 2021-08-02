.PHONY: intigration-test

intigration-test: 
	@echo "======================= intigration-test ======================="
	@go test -v -tags=integration .
