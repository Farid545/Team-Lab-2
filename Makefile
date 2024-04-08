default: out/example

clean:
	rm -r ./out

test: *.go
	go test .

out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
