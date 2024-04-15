all: go-test go-build
# go
go-test:
	go test -v ./...

go-build:
	cd cmd && go build -o helmut
	mv cmd/helmut .

IMG ?= helmut:latest

# build image and push to dockerhub
.PHONY: docker-build
docker-build:
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}