# Tools
GOLANG := go

# Build
BUILD_FILES := $(shell find . -type f -name '*.go' -not -name '*_test.go')
BUILD_FILES := ${BUILD_FILES} go.mod go.sum

.PHONY: ls
ls:
	@echo "BUILD_FILES: ${BUILD_FILES}"

GO_LAMBDA_ARCH ?= arm64
GO_LAMBDA_OS ?= linux
GO_LAMBDA_BUILD_FLAGS ?= -ldflags="-s -w" -tags "lambda.norpc"
GO_LAMBDA_BUILD ?= GOOS=${GO_LAMBDA_OS} GOARCH=${GO_LAMBDA_ARCH} ${GOLANG} build ${GO_LAMBDA_BUILD_FLAGS}

build/lambda/${GO_LAMBDA_OS}/${GO_LAMBDA_ARCH}/%/bootstrap: ${BUILD_FILES}
	${GO_LAMBDA_BUILD} -o $@ ./lambda/$(patsubst build/lambda/${GO_LAMBDA_OS}/${GO_LAMBDA_ARCH}/%/bootstrap,%,$@)

.PHONY: build
build: build/lambda/${GO_LAMBDA_OS}/${GO_LAMBDA_ARCH}/request-handler/bootstrap

# Test
GO_TEST_FLAGS ?= -v -race
GO_TEST_PATH ?= ./...
GO_TEST_RUN ?= .
GO_TEST_TIMEOUT ?= 30s

.PHONY: test
test:
	${GOLANG} test ${GO_TEST_FLAGS} -run ${GO_TEST_RUN} ${GO_TEST_PATH} -timeout ${GO_TEST_TIMEOUT}
