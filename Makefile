export LISTEN_ADDRESS=:12345
export STORAGE_ROOT=$(PWD)/tmp

TARGETS := object-storage-api object-storage-data
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

build: $(TARGETS)

$(TARGETS):$(SRC)
	go build $(PWD)/cmd/$@

clear: $(TARGETS)
	rm -rf $(TARGETS)

run: $(TARGETS)
	./$(TARGETS)

