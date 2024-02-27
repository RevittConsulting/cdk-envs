default: all

all: app

.PHONY: app
app:
	@echo "Building APP"
	docker compose up -d