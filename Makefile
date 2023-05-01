# Ensure go modules are enabled:
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Disable CGO so that we always generate static binaries:
export CGO_ENABLED=0

BINARY=Web_Crawler
all: build

build:
	go build -o ${GOPATH}/bin/Web_Crawler

install:
	go install

run:
	go build 
	$(BINARY)

test:
	go test -v

mod:
	go mod tidy 

clean:
	rm ${BINARY}