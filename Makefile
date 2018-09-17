GO=go
OUTPUT_DIR=bin/klingon-translator
MAIN_FILE=cmd/klingon-translator/main.go

all:
	$(GO) build -ldflags "-X main.AppVersion=0.0.1" -o $(OUTPUT_DIR) $(MAIN_FILE)

test:
	$(GO) test github.com/pmukhin/klingon-translator/pkg/klingon/parser
	$(GO) test github.com/pmukhin/klingon-translator/pkg/klingon/stapi/character

