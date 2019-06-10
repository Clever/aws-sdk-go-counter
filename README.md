# aws-sdk-go-counter

Session handler that counts calls made to AWS APIs.

## Motivation

AWS APIs have rate limits, so monitoring the rate at which code is using different AWS APIs is useful.

## Usage

Take normal aws-sdk-go client usage like this:

``` go
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	_, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

and add a handler to the session:

``` go
package main

import (
	"fmt"
	"log"

	"github.com/Clever/aws-sdk-go-counter"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.Must(session.NewSession())
	counter := counter.New()
	sess.Handlers.Send.PushFront(counter.SessionHandler)
	svc := s3.New(sess)
	_, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", counter.Counters())
}
```

Output:

```
[{s3 ListBuckets 1}]

```

You could also print the values periodically in the background e.g.

```go
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for _ = range ticker.C {
			fmt.Printf("%v\n", counter.Counters())
		}
	}()
```

## Example

See `example/` for a an example.
