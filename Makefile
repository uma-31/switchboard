usage: ## このメッセージを表示します。
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@grep -P '^[a-zA-Z_-]+:\s+##.*$$' Makefile \
		| perl -ne '/^([a-zA-Z_-]+):\s+\#\#\s+(.*)$$/ && print "\t$$1\t$$2\n"'

format: ## コードのフォーマットを実行します。
	go fmt ./...
	golangci-lint run --fix ./...

lint: ## Go の静的解析を実行します。
	golangci-lint run ./...
