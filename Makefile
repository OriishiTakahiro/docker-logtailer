BINARY_NAME = tail-clog
TARGET_ARCH = "amd64"
TARGET_OS = "linux"

ensure:
	dep ensure -update

build: ensure
	@GOARCH=$(TARGET_ARCH) GOOS=$(TARGET_OS) go build -o ./output/$(BINARY_NAME) *.go

test: ./.circleci/config.yml
	echo $<
	circleci config validate -c $<
	circleci build --job test $<
