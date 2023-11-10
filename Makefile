.PHONY: mac-build
mac-build: ## Build Unix binary
	@export GOOS=darwin GOARCH=arm64
	@go build -o ./bin/darwin/encryptor
	@echo "Done"

.PHONY: win-build
win-build: ## Build Windows binary
	@export GOOS=windows GOARCH=amd64
	@go build -o ./bin/windows/encryptor.exe
	@echo "Done"
