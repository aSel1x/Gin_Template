.PHONY: help
help:
	@echo "USAGE"
	@echo "  make <command>"
	@echo ""
	@echo "AVAILABLE COMMANDS"
	@echo "  go		Start the app"


.PHONY: go
go:
	set -a; source .env; set +a; \
	go run main.go
