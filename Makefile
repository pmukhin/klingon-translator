all:
	go build -ldflags "-X main.AppVersion=0.0.1" -o "bin/klingon-translator" "cmd/klingon-translator/main.go"
