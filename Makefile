default: all

all: cdk-envs

.PHONY: cdk-envs
cdk-envs:
	@echo "Building APP"
	@export DATA_DIR_HOST=$(data) ; \
	docker compose build && docker compose up -d