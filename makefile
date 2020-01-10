UNAMES = `uname -s | awk '{print tolower($0)}'`
UNAMEM = `uname -m | awk '{print tolower($0)}'`
BINFILE = guppy-$(UNAMES)-$(UNAMEM)
GOBUILD = go build -a -o bin/$(BINFILE) cmd/cmd.go

all: test build install

test:
	go test ./...

build:
	@if [ $(UNAMEM) == "x86_64" ]; then\
		GOOS=$(UNAMES) GOARCH=amd64 $(GOBUILD);\
	elif [ $(UNAMEM) == "i386" ]; then\
		GOOS=$(UNAMES) GOARCH=386 $(GOBUILD);\
	else\
		GOOS=$(UNAMES) GOARCH=$(UNAMEM) $(GOBUILD);\
	fi

install:
	cp bin/$(BINFILE) /usr/local/bin/$(BINFILE)