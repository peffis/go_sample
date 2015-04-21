TARGET = hello
GOPATH := $(GOPATH):$(shell pwd)


all:
	rm dummy.go
	go get ./...
	go install ./...
	cp /opt/apcera/go/bin/$(TARGET) /opt/apcera/go/bin/app
