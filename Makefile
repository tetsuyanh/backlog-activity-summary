
setup:
	go get github.com/golang/dep
	go install github.com/golang/dep/cmd/dep
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/golang/lint/golint

install:
	dep ensure

update:
	dep ensure -update

test:
	go test ./...

fmt:
	goimports -w ./model ./nulab ./renderer ./service

lint: fmt
	go vet ./...
	golint ./model ./nulab ./renderer ./service
