.PHONY: unix-build
unix-build: ## Build Unix binary
	@$(MAKE) build GOOS=darwin GOARCH=arm64

.PHONY: win-build
win-build: ## Build Windows binary
	@$(MAKE) build GOOS=windows GOARCH=386

build: ## Build binary
	go build -o ./bin/encryptor

.PHONY: licenses
licenses: ## Generate LICENSE-DEPENDENCIES.md using golicense binary
	@touch LICENSE-DEPENDENCIES.md
	@golicense -plain ./bin/prevensys_alerts_center > LICENSE-DEPENDENCIES.md
