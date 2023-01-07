build: setup
	go build -trimpath -ldflags="-s -w" -o out/code-letter-cli

test: setup
	go test -race -coverprofile=./out/coverprofile.out -covermode=atomic ./...

report: test
	go tool cover -html=./out/coverprofile.out

setup: clean
	mkdir ./out

clean:
	rm -rf ./out
	go clean
