GOBASE			= ${shell pwd}
GOBIN			= ${GOBASE}/bin
PROJECTNAME		= online_shop
COMMIT_SHA		= ${shell git log -1 --pretty=%h}
IMG_TAG			= ${IMAGENAME}:${COMMIT_SHA}
LATEST_TAG		= ${IMAGENAME}:latest
MANAGEMENT_API_PATH 		= api/management/v1
MANAGEMENT_PROTO_API_DIR 	= api/management/v1
MANAGEMENT_PROTO_OUT_DIR 	= pkg/management/v1
MANAGEMENT_PROTO_API_OUT_DIR = ${PROTO_OUT_DIR}
PASSPORT_API_PATH 		= api/passport/v1
PASSPORT_PROTO_API_DIR 	= api/passport/v1
PASSPORT_PROTO_OUT_DIR 	= pkg/passport/v1
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
build:
	GOOS=linux GOARCH=amd64 go build -o ${GOBIN}/${PROJECTNAME} ./cmd/${PROJECTNAME}/main.go || exit 1

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
	go vet $$( go list ./... | grep -v vendor)

.PHONY: lint-proto
lint-proto:
	protolint lint -fix ./api/
