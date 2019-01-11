BINARY_NAME = clogserver
TARGET_ARCH = "amd64"
TARGET_OS = "linux"

ensure:
	dep ensure

build: ensure
	@GOARCH=$(TARGET_ARCH) GOOS=$(TARGET_OS) go build -o $(BINARY_NAME) *.go
