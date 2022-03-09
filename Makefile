build:
	@echo "\033[93mBuilding docker container ibanvalidator \033[0m"
	docker build -t ibanvalidator:latest -f Dockerfile.validator .
	docker build -t ibanvalidatorservice:latest -f Dockerfile.service .
	@echo "\033[92mBuild finished\033[0m"
.PHONY: build

runservice-docker:
	@echo "\033[92mStarting IBAN Validation in docker\033[0m"
	docker run -d -p 5000:5000 ibanvalidatorservice
.PHONY: runservice-docker

runservice:
	@echo "\033[92mRunning validation service on port 5000. To run the service in docker, type: make runservice-docker\033[0m"
	go run ibanService/ibanService.go
.PHONY: runservice

test:
	@echo "\033[92mTest of successful IBAN Validation\033[0m"
	docker run ibanvalidator GB82WEST12345698765432
	@echo "\033[92mTest of unsuccessfull IBAN Validation\033[0m"
	docker run ibanvalidator GB82WEST12335698765432
.PHONY: test

