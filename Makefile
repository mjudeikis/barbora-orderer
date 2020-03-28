
PLATFORMS=darwin linux
ARCHITECTURES=386 amd64

build_all: build_win
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/order-$(GOOS)-$(GOARCH) ./cmd/main.go )))

build_win:
	$(foreach GOOS, windows,\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/order-$(GOOS)-$(GOARCH).exe ./cmd/main.go )))
