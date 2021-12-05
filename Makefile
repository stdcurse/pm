build:
	CGO_ENABLED=0 go build -v -ldflags="-w -s -extldflags=-static"
	upx --best --lzma ./pm
