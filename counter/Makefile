.PHONY: all generate test

all: generate test

generate:
	go generate .

test:
	go test ./...

clean:
	find -mindepth 1 -maxdepth 1 -type d | xargs rm -r
