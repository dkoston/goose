build:
	cwd=$(pwd 2>&1) && cd cmd/goose && go get . && go build -ldflags=-s . && cd ${cwd}

install: build
	cp cmd/goose/goose ${GOPATH}/bin/

test:
	cwd=$(pwd 2>&1) && cd lib/goose && go get . && go test && cd ${cwd}

.PHONY: clean test all
clean:
	rm -f cmd/goose/goose