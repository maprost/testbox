FULL_FOLDER_PATH=$(shell go list -f '{{.Dir}}' ./... | grep -v /vendor)
FOLDER_PATH=$(shell go list ./... | grep -v /vendor)

GOFMT=$(shell go fmt)

fmt:
	$(GOFMT)

fmtcheck:
ifeq ($(GOFMT),)
	echo "go fmt failure"
	exit 1
endif

dep:
	go get -u github.com/Masterminds/glide
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update
	glide install

lint:
	gometalinter --enable-all --disable=lll --disable=golint --vendor $(FULL_FOLDER_PATH)

test:
	go test -cover $(FOLDER_PATH)

build:
	CGO_ENABLED=0 \
	go build

graph:
	godepgraph -s github.com/NatPro/Kassenbons | dot -Tpng -o godepgraph.png

all: fmtcheck lint dep test build