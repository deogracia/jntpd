OUTPUT="./output"
SRC=$(shell find . -name "*.go")

.PHONY: all clean build test check_fmt fmt vet

all: clean install_deps test build

output:
	$(info ***************** Create "output" directory ***********************************)
	mkdir $(OUTPUT)

clean:
	$(info ***************** Clean ***********************************)
	rm -rf $(OUTPUT)

build: output
	make -f build.make all

test: check_fmt vet
	$(info ***************** Run tests ***********************************)
	echo do some tests!

install_deps:
	$(info ***************** Install dependancies ***********************************)
	go get ./...

check_fmt:
	$(info ***************** Check formatting ***********************************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

fmt:
	$(info ***************** Do the formatting ***********************************)
	gofmt -w $(SRC)

vet:
	$(info ***************** Run go vet ***********************************)
	go vet ./...
