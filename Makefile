# Common properties
GOBASE			= ${shell pwd}
GOBIN			= ${GOBASE}/bin
COMMIT_SHA		= ${shell git log -1 --pretty=%h}
IMG_TAG			= ${IMAGENAME}:${COMMIT_SHA}
LATEST_TAG		= ${IMAGENAME}:latest

# Management service properties
PROJECT_NAME_MANAGEMENT		= management
MANAGEMENT_API_PATH 		= api/management/v1
MANAGEMENT_PROTO_API_DIR 	= api/management/v1
MANAGEMENT_PROTO_OUT_DIR 	= pkg/management/v1

# Passport service properties
PROJECT_NAME_PASSPORT		= passport
PASSPORT_API_PATH 		= api/passport/v1
PASSPORT_PROTO_API_DIR 	= api/passport/v1
PASSPORT_PROTO_OUT_DIR 	= pkg/passport/v1

# Commands

.PHONY: gen-proto
gen-proto: gen-proto-management

.PHONY: gen-proto-management
gen-proto-management:
	mkdir -p ${MANAGEMENT_PROTO_OUT_DIR}
	protoc \
		-I ${MANAGEMENT_API_PATH} \
		--go_out=$(MANAGEMENT_PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(MANAGEMENT_PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
		./${MANAGEMENT_PROTO_API_DIR}/*.proto

.PHONY: gen-proto-passport
gen-proto-passport:
	mkdir -p ${PASSPORT_PROTO_OUT_DIR}
	protoc \
		-I ${PASSPORT_API_PATH} \
		--go_out=$(PASSPORT_PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(PASSPORT_PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
		./${PASSPORT_PROTO_API_DIR}/*.proto

.PHONY: build
build: build_mgmt build_passport

.PHONY: build_mgmt
build_mgmt:
	GOOS=linux GOARCH=amd64 go build -o ${GOBIN}/${PROJECT_NAME_MANAGEMENT} ./cmd/${PROJECT_NAME_MANAGEMENT}/main.go || exit 1

.PHONY: build_passport
build_passport:
	GOOS=linux GOARCH=amd64 go build -o ${GOBIN}/${PROJECT_NAME_PASSPORT} ./cmd/${PROJECT_NAME_PASSPORT}/main.go || exit 1

.PHONY: lint
lint: go/lint lint-proto

.PHONY: go/lint
go/lint:
	golangci-lint run $(GOLINT_FLAGS) --config=.golangci.yml --timeout=180s ./...

.PHONY: proto/lint
proto/lint:
	protolint lint $(PROTOLINT_FLAGS) $(PROTO_API_DIR)/*

.PHONY: test
test:
	go test ./internal/...

.PHONY: govet
govet:
	go vet $$( go list ./internal/...)

.PHONY: lint-proto
lint-proto:
	protolint lint -fix ./api/
