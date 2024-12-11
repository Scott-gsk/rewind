.PHONY: setup docs vendor

setup:
	go install github.com/swaggo/swag/cmd/swag@latest

docs:
	cd incubation/node-manager && \
	swag init --parseDependency

vendor:
	go mod tidy && go mod vendor

push:
	git add . && codegpt commit . && git push
