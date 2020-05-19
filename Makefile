export LISTEN_ADDRESS=:12345
export STORAGE_ROOT=$(PWD)/tmp

TARGETS := object-storage-service
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

build: $(TARGETS)

$(TARGETS):$(SRC)
	go build -o $(TARGETS) cmd/*.go

run:$(TARGETS)
	rm -rf $(TARGETS)
	go build -o $(TARGETS) cmd/*.go
	./$(TARGETS)

clear:
	rm -rf $(TARGETS)