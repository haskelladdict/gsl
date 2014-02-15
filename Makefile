
PACKAGES = ./src/stats ./src/util


all:
	go build runner.go


test: $(PACKAGES)
	go test ./$^


format:
	gofmt -w=true -tabs=false -tabwidth=2 $(PACKAGES)
