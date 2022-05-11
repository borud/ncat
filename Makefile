ifeq ($(OS),Windows_NT)
	EXTENSION=.exe
else
	EXTENSION=""
endif

all: vet build

build: ncat

ncat:
	@cd cmd/$@ && go build -o ../../bin/$@$(EXTENSIION) --trimpath -tags osusergo,netgo -ldflags="-s -w"
	@cd cmd/$@ && GOARCH=arm GOOS=linux go build -o ../../bin/$@-arm --trimpath -tags osusergo,netgo -ldflags="-s -w"

vet:
	@echo "*** $@"
	@go vet ./...

clean:
	@rm -rf bin