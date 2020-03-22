
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/order-$(GOOS)-$(GOARCH) ./cmd/main.go )))
