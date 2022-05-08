SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'

PLATFORM := $(shell uname -o)

NAME := starship.exe
OS := windows

ifeq ($(PLATFORM), Msys)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else ifeq ($(PLATFORM), Cygwin)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else
	INCLUDE := $(GOPATH)
	NAME=starship
	OS=linux
endif

# enable second expansion
.SECONDEXPANSION:

.PHONY: all
.PHONY: pbs
.PHONY: test
.PHONY: contract

BINDIR=$(INCLUDE)/bin

all: pbs contract build

build:
	GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o $(BINDIR)/$(NAME)

pbs:
	cd pbs/ && $(MAKE)

sol:
	cd contract/ && $(MAKE)

target:=mac

mac:
	GOOS=darwin go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).mac
arm:
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(NAME).lnx
win:
	GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).exe

clean:
	rm $(BINDIR)/$(NAME)
