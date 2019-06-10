package main

import (
	"fmt"
	"time"

	counter "github.com/Clever/aws-sdk-go-counter"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	sess := session.Must(session.NewSession())
	counter := counter.New()
	sess.Handlers.Send.PushFront(counter.SessionHandler)
	stsAPI := sts.New(sess)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			fmt.Printf("%v\n", counter.Counters())
		}
	}()

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			stsAPI.GetCallerIdentity(&sts.GetCallerIdentityInput{})
		}
	}()

	select {}
}
