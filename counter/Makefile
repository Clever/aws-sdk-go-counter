.PHONY: all generate test rewrite-imports

all: generate test

generate:
	go generate .

test:
	go test ./...

deps:
	# for any package we modify in this fork, rewrite package imports to correctly reference our forked package
	find ../private/model/api/         -name \*.go | xargs sed -i.bak -e 's/^	"github.com\/aws\/aws-sdk-go\//	"github.com\/Clever\/aws-sdk-go-counter\//g'
	find ../private/model/cli/gen-api/ -name \*.go | xargs sed -i.bak -e 's/^	"github.com\/aws\/aws-sdk-go\//	"github.com\/Clever\/aws-sdk-go-counter\//g'
	# for everything else reference latest upstream
	go get -u github.com/aws/aws-sdk-go

clean:
	find -mindepth 1 -maxdepth 1 -type d | xargs rm -r
