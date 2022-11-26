build: setup
	go build -trimpath -ldflags="-s -w" -o out/code-letter-cli

test: setup
	go test -coverprofile=./out/coverprofile.out ./...
	go tool cover -func=./out/coverprofile.out

report: test
	go tool cover -html=./out/coverprofile.out

setup: clean
	mkdir ./out

clean:
	rm -rf ./out
	go clean
