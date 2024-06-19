APP_NAME=log-server
VERSION_FILE=VERSION

# Read the version from the VERSION file
VERSION=$(shell cat $(VERSION_FILE) 2>/dev/null || echo "1.0.0") 

# Increment the version
define increment_version
	awk 'BEGIN { FS = OFS = "." } {$$NF++; print}' $(VERSION_FILE) > $(VERSION_FILE).tmp && mv $(VERSION_FILE).tmp $(VERSION_FILE)
endef

build:
	$(call increment_version)
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o bin/$(APP_NAME)-linux-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.version=$(VERSION)" -o bin/$(APP_NAME)-darwin-arm64 main.go

clean:
	rm -rf bin/

.PHONY: build clean