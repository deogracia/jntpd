XC_OS="linux darwin windows"
XC_ARCH="amd64 386"
XC_PARALLEL="2"
OUTPUT="./output"
NAME="jntpdn"
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif

.PHONY: all build

all: build
	@echo "[OK] Build done!"

build:
	$(info ***************** Build the artefacts ***********************************)
	gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(OUTPUT)/$(NAME)_{{.OS}}_{{.Arch}} \
		;

# vi:syntax=make
