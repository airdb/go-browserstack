run: setenv
	go run main.go

setenv:
	@echo export GOPATH=`pwd`
	@echo export GOBIN=`pwd`/bin

indent:
	gofmt -w main.go
