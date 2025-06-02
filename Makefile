APP_NAME := MaskLR_Go
SRC := cmd/server/main.go
OUTPUT_DIR := build
OUTPUT := $(OUTPUT_DIR)/$(APP_NAME)-linux-armv7

# 默认本地构建
build:
	mkdir -p $(OUTPUT_DIR)
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o $(OUTPUT) $(SRC)
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(OUTPUT_DIR)/$(APP_NAME)-linux-armv7-debug $(SRC)

# 清理构建产物
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: build clean
