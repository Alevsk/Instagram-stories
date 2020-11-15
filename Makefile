PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

default: instagram-stories

.PHONY: instagram-stories

instagram-stories:
	@echo "Building instagram-stories binary to './instagram-stories'"
	@(GO111MODULE=on CGO_ENABLED=0 go build -trimpath --tags=kqueue --ldflags "-s -w" -o instagram-stories ./cmd/instagram-stories)

force:
	@echo "Force building instagram-stories binary to './instagram-stories'"
	@(go build -o instagram-stories ./cmd/instagram-stories)

getdeps:
	@mkdir -p ${GOPATH}/bin
	@which golangci-lint 1>/dev/null || (echo "Installing golangci-lint" && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.27.0)

verifiers: getdeps fmt lint

fmt:
	@echo "Running $@ check"
	@GO111MODULE=on gofmt -d cmd/
	@GO111MODULE=on gofmt -d pkg/

lint:
	@echo "Running $@ check"
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint cache clean
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=5m --config ./.golangci.yml

swagger-gen:
	@echo "Generating swagger server code from yaml"
	@swagger generate server -A instagram_stories --main-package=instagram_stories --exclude-main -P models.Principal -f ./swagger.yml -r NOTICE

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@find . -name '*~' | xargs rm -fv
	@rm -vf instagram-stories
