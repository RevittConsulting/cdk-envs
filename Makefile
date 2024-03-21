default: all

all: cdu

.PHONY: cdu
cdu:
	@echo "Building Chain Dev Utils..."
	@export DATA_DIR_HOST=$(data) ; \
	docker compose build && docker compose up -d