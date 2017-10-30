# aws-sdk-go-counter

Wrapper around [aws-sdk-go](https://github.com/aws/aws-sdk-go) that counts API requests made to AWS.

## Motivation

AWS APIs have rate limits, so monitoring the rate at which code is using different AWS APIs is useful.

## Example usage

Take normal aws-sdk-go client usage like this:

``` go
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.New()
	svc := s3.New(sess)
	_, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

and wrap it

``` go
package main

import (
	"fmt"
	"log"

	"github.com/Clever/aws-sdk-go-counter/counter/s3counter"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.New()
	svc := s3counter.New(s3.New(sess))
	_, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", svc.Counters())
}
```

Output:

```
map[ListBuckets:1]
```
