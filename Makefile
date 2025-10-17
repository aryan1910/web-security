.PHONY: help csrf xss sqli cmd-injection sync clean

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

sync: ## Sync all workspace modules
	go work sync

csrf: ## Run CSRF protection demo
	@echo "🔒 Starting CSRF demo on http://localhost:8080"
	cd csrf && go run main.go

xss: ## Run XSS vulnerability demo  
	@echo "⚠️  Starting XSS demo on http://localhost:8080"
	cd xss && go run main.go

sqli: ## Run SQL injection demo
	@echo "💉 Starting SQL injection demo on http://localhost:8080"
	cd sqli && go run main.go

cmd-injection: ## Run command injection demo
	@echo "💻 Starting command injection demo on http://localhost:8080"
	cd cmd-injection && go run main.go

clean: ## Clean all module caches
	go clean -modcache

test: ## Test all modules
	@echo "Testing all modules..."
	@for dir in csrf xss sqli cmd-injection; do \
		echo "Testing $$dir..."; \
		cd $$dir && go build . && cd ..; \
	done
	@echo "All modules build successfully!"