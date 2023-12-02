build:
	go build -o ts cmd/ts-cli/*.go

test: build
	@exec cmd/ts-cli/ts 你好
	@exec cmd/ts-cli/ts hello

clean:
	@rm -rf cmd/ts-cli/ts

.PHONY:
	build test clean