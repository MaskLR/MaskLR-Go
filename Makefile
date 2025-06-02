# 项目配置
APP_NAME := MaskLR_Go
SRC := cmd/server/main.go
OUTPUT_DIR := build
OUTPUT := $(BUILD_DIR)/$(APP_NAME)-linux-armv7

# 默认本地构建
build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o $(OUTPUT) $(SRC)
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(OUTPUT)-debug $(SRC)

# 清理构建产物
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build clean
