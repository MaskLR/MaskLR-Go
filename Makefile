# 项目配置
APP_NAME = masklr
MAIN = cmd/server/main.go
OUTPUT_DIR = bin

# Go 编译设置
GO = go
GOFLAGS =
BUILD_FLAGS = -ldflags="-s -w"

# 默认本地构建
build:
	$(GO) build $(BUILD_FLAGS) -o $(OUTPUT_DIR)/$(APP_NAME) $(MAIN)

# 构建 Linux ARMv7 可执行文件
build-armv7:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 $(GO) build $(BUILD_FLAGS) -o $(OUTPUT_DIR)/$(APP_NAME)-armv7 $(MAIN)

# 运行所有测试
test:
	$(GO) test ./... -v

# 清理构建产物
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: build build-armv7 test clean
