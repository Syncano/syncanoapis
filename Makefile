.PHONY: help proto proto-go proto-python
.DEFAULT_GOAL := help
$(VERBOSE).SILENT:

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

require-%:
	if ! hash ${*} 2>/dev/null; then \
		echo "! ${*} not installed"; \
		exit 1; \
	fi

proto-go: require-protoc ## Run protobuf compiler on all .proto files for go
	mkdir -p gen/go
	for f in $$(find . -name \*.proto -type f); do \
		protoc -I. \
			--go_out gen/go \
			--go_opt paths=source_relative,plugins=grpc \
			$$f; \
	done

proto-python: ## Run protobuf compiler on all .proto files for python
	mkdir -p gen/python
	for f in $$(find . -name \*.proto -type f); do \
		python3 -m grpc_tools.protoc -I. \
			--python_out gen/python \
			--grpc_python_out gen/python \
			$$f; \
	done

proto: proto-python proto-go ## Run protobuf compiler on all .proto files for both python and go
