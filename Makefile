PROJECT_NAME=login

.DEFAULT_GOAL := test

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

i18n-extract:
	goi18n extract -outdir language/locales

i18n-merge:
	goi18n merge -outdir language/locales language/locales/active.*.toml language/locales/translate.*.toml

i18n-translations:
	goi18n merge -outdir language/locales language/locales/active.*.toml

lint:
	@echo linting
	@golint $(shell go list ./... | grep -v /vendor/)

test: tidy fmt lint
	go test -cover ./...

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

.PHONY: fmt i18n-extract i18n-merge i18n-translations lint test tidy vendor