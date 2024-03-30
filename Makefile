default: out/example

clean:
	rm .\out

test: *.go
	go test ./...
	if make "$@"; then
		echo BUILD SUCCESSFUL
	else
		echo BUILD FAILED
	fi

out/example: implementation.go cmd/example/main.go
	mkdir out
	go build -o out/example ./cmd/example
