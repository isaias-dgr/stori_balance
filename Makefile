LAMBDA_TARGET := empty
SHELL := /bin/zsh

default: init

.PHONY: init
init:
	@echo "🤖 stori ingest Notify app"


.PHONY: build-lambda-log
build-lambda-logs: LAMBDA_TARGET=$(APP_INGEST_SHORT)
build-lambda-logs: build-lambda 


.PHONY: build-lambda-pdf
build-lambda-pdf: LAMBDA_TARGET=$(APP_NOTIFY_SHORT)
build-lambda-pdf: build-lambda 


.PHONY: build-lambda
build-lambda:
	@echo "🤖 Build Lamdba package $(LAMBDA_TARGET)"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(LAMBDA_TARGET)/main src/cmd/$(LAMBDA_TARGET)/main.go 
	zip -j bin/$(LAMBDA_TARGET)/main.zip bin/$(LAMBDA_TARGET)/main
	ls -lh bin/$(LAMBDA_TARGET)/main.zip


.PHONY: localstack
localstack:
	@echo "🐳 Runing Docker compose"
	docker compose up -d


.PHONY: infra-init
infra-init:
	@echo "🪨 Destroy infra on localstack"
	tflocal -chdir=infra/ init


.PHONY: infra-plan
infra-plan:
	@echo "🏗 Destroy infra on localstack"
	tflocal -chdir=infra/ fmt
	tflocal -chdir=infra/ plan 


.PHONY: infra-apply
infra-apply:
	@echo "🏛️ Destroy infra on localstack $(APP_INGEST_SHORT)"

	tflocal -chdir=infra/ apply -auto-approve

.PHONY: infra-destroy
infra-destroy:
	@echo "🏚 Destroy infra on localstack"
	tflocal -chdir=infra/ destroy -auto-approve

.PHONY: infra
infra:  infra-init infra-plan infra-apply
	@echo "🏘 Destroy infra on localstack"
