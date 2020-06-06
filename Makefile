OUTPUT="./output"
SRC=$(shell find . -name "*.go")

.PHONY: all clean build test check_fmt fmt vet security

all: clean install_deps security test build

output:
	$(info ***************** Create "output" directory ***********************************)
	mkdir $(OUTPUT)

clean:
	$(info ***************** Clean ***********************************)
	rm -rf $(OUTPUT)
	go clean -cache -testcache -modcache

security:
	$(info ***************** Security ***********************************)
	@~/go/bin/gosec ./...
	@echo "[OK] Go security check is done!"

build: output
	make -f build.make all

test: check_fmt vet
	$(info ***************** Run tests ***********************************)
	echo do some tests!

install_deps:
	$(info ***************** Install dependancies ***********************************)
	@go get ./...
	@echo "[OK] Go dependancies is get!"

check_fmt:
	$(info ***************** Check formatting ***********************************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)
	@echo "[OK] Go code formating"

fmt:
	$(info ***************** Do the formatting ***********************************)
	gofmt -w $(SRC)

vet:
	$(info ***************** Run go vet ***********************************)
	@go vet ./...
	@echo "[OK] Go vet is done!"

# vi:syntax=make
