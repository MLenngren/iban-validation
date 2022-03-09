build: ## build validator and validatorservice images
	@echo "\033[93mBuilding docker container ibanvalidator \033[0m"
	docker build -t ibanvalidator:latest -f Dockerfile.validator .
	docker build -t ibanvalidatorservice:latest -f Dockerfile.service .
	@echo "\033[92mBuild finished\033[0m"
.PHONY: build

runds: ## Starting IBAN Validation in docker
	@echo "\033[92mStarting IBAN Validation in docker\033[0m"
	docker run -d -p 5000:5000 ibanvalidatorservice
.PHONY: runservice-docker

runs: ## Running validation service on port 5000. To run the service in docker, type: make runservice-docker
	@echo "\033[92mRunning validation service on port 5000. To run the service in docker, type: make runservice-docker\033[0m"
	go run ibanService/ibanService.go
.PHONY: runservice

test: ## run docker checks against 1 valid and 1 unvalid IBAN
	@echo "\033[92mTest of successful IBAN Validation\033[0m"
	docker run ibanvalidator GB82WEST12345698765432
	@echo "\033[92mTest of unsuccessfull IBAN Validation\033[0m"
	docker run ibanvalidator GB82WEST12335698765432
.PHONY: test

help: ## Help target
	@echo "\033[36mtest\t\t\033[0m run docker checks against 1 valid and 1 unvalid IBAN"
	@echo "\033[36mbuild\t\t\033[0m build validator and validatorservice images"
	@echo "\033[36mruns\t\t\033[0m Running validation service on port 5000. To run the service in docker, type: make runds"
	@echo "\033[36mrunds\t\t\033[0m Starting IBAN Validation in docker"


.DEFAULT_GOAL := help