# 输出文件名
BINARY_NAME = MaskLR-Go
VERSION ?= dev

# 默认目标：构建本地版本
all: build

# 本地构建
build:
	go build -o dist/$(BINARY_NAME)-$(VERSION) .

# 交叉编译：为 Termux（Linux ARMv7）构建
build-armv7:
	GOOS=linux GOARCH=arm GOARM=7 go build -o dist/$(BINARY_NAME)-armv7-$(VERSION) .

# 清理编译输出
clean:
	rm -rf dist
