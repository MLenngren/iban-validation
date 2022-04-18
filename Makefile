build: ## build validator and validatorservice images
	@echo "\033[93mBuilding docker container ibanvalidator \033[0m"
	docker build -t pfcmale/ibanvalidator:latestlocal -f Dockerfile.validator . 
	docker build -t pfcmale/ibanvalidatorservice:latestlocal -f Dockerfile.service .
	docker push pfcmale/ibanvalidatorservice:latestlocal
	docker push pfcmale/ibanvalidator:latestlocal
	docker build --platform linux/amd64 -t pfcmale/ibanvalidator:latest -f Dockerfile.validator . 
	docker build --platform linux/amd64 -t pfcmale/ibanvalidatorservice:latest -f Dockerfile.service .
	docker push pfcmale/ibanvalidatorservice:latest
	docker push pfcmale/ibanvalidator:latest
	@echo "\033[92mBuild finished\033[0m"
.PHONY: build

runs: ## Running validation service on port 5000. To run the service in docker, type: make runservice-docker
	@echo "\033[92mRunning validation service on port 5000. To run the service in docker, type: make runservice-docker\033[0m"
	go run ibanService/ibanService.go
.PHONY: runservice

test: ## run tests from docker
	@echo "\033[92mRunning tests\033[0m"
	docker run  -v "$$PWD":/usr/src/myapp -w /usr/src/myapp  golang:1.17 go test ibanValidator/ibanValidator_test.go -v
	docker run --rm -v "$$PWD":/usr/src/myapp --env GIN_MODE=release -w /usr/src/myapp  golang:1.17 go test ibanService/ibanService_test.go -v
.PHONY: test

help: ## Help target
	@echo "\033[36mtest\t\t\033[0m run existing tests in go"
	@echo "\033[36mbuild\t\t\033[0m build validator and validatorservice images"
	@echo "\033[36mruns\t\t\033[0m Running validation service on port 5000"

.DEFAULT_GOAL := help