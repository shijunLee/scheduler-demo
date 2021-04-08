VERSION ?= $(shell cat ./pkg/version/version.go|grep "version ="|awk '{print $$3}'| sed 's/\"//g' | tr  "\n" " " | tr "\n" " " | sed 's/[[:space:]]//g' | tr "\n" " ")
TAG ?= $(shell git describe --abbrev=0 --tags)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_ID = $(shell git rev-parse --short HEAD)
BUILD_TIME= $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
FLAGS="-X github.com/shijunLee/scheduler-demo/pkg/version.CommitId=$(COMMIT_ID) -X github.com/shijunLee/scheduler-demo/pkg/version.Branch=$(BRANCH) -X github.com/shijunLee/scheduler-demo/pkg/version.Tag=$(TAG) -X github.com/shijunLee/scheduler-demo/pkg/version.BuildTime=$(BUILD_TIME)"
IMG ?= docker.shijunlee.local/library/kube-scheduler-demo:$(VERSION)

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...

.PHONY: manager
manager: fmt vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build  -ldflags $(FLAGS) -o bin/kube-scheduler main.go && chmod +x bin/kube-scheduler
 
.PHONY: build-docker
build-docker: manager
	docker build -t $(IMG) .
