APP_NAME := masklr
SRC := main.go
BUILD_DIR := build
OUTPUT := $(BUILD_DIR)/$(APP_NAME)-linux-armv7

build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(OUTPUT) $(SRC)

clean:
	rm -rf $(BUILD_DIR)

.PHONY: build clean
