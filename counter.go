package counter

import (
	"strings"
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/request"
)

// Counter counts requests made to AWS APIs.
type Counter struct {
	counts sync.Map
}

// New returns a new Counter.
func New() *Counter {
	return &Counter{}
}

// SessionHandler should be used in conjunction with the Session object's support
// for Handlers, e.g.:
//   	sess.Handlers.Send.PushFront(counter.SessionHandler)
func (c *Counter) SessionHandler(r *request.Request) {
	c.inc(r.ClientInfo.ServiceName + "|" + r.Operation.Name)
}

// ServiceCount represents an AWS service and the count of calls to it.
type ServiceCount struct {
	Service   string
	Operation string
	Count     int64
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *Counter) Counters() []ServiceCount {
	counters := []ServiceCount{}
	c.counts.Range(func(key, value interface{}) bool {
		k := key.(string)
		ps := strings.Split(k, "|")
		counters = append(counters, ServiceCount{
			Service:   ps[0],
			Operation: ps[1],
			Count:     value.(*atomicCounter).count(),
		})
		return true
	})
	return counters
}

// inc increments the counter for the operation by 1.
func (c *Counter) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &atomicCounter{})
	cnt.(*atomicCounter).inc()
}

// atomicCounter is a threadsafe cumulative counter.
type atomicCounter struct {
	c int64
}

// count returns the current count.
func (c *atomicCounter) count() int64 {
	return atomic.LoadInt64(&c.c)
}

// inc increments the counter by one.
func (c *atomicCounter) inc() {
	atomic.AddInt64(&c.c, 1)
}
