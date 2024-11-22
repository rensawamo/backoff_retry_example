buf-gen: ## buf generate code
	buf generate --path proto && go mod tidy

