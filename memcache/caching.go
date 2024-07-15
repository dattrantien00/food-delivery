package memcache

import (
	"food-delivery/common"
	"sync"
	"time"
)

type Caching interface {
	Write(k string, value interface{})
	Read(k string) interface{}
	WriteTTL(k string, value interface{}, exp int)
}

type caching struct {
	store  map[string]interface{}
	locker *sync.RWMutex
}

func NewCache() Caching {
	return &caching{
		store:  make(map[string]interface{}),
		locker: new(sync.RWMutex),
	}
}

func (c *caching) Write(k string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.store[k] = value
}

func (c *caching) Read(k string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.store[k]
}

func (c *caching) WriteTTL(k string, value interface{}, exp int) {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.store[k] = value

	go func() {
		defer common.AppRecover()
		<-time.NewTimer(time.Duration(exp) * time.Second).C
		c.Write(k, value)
	}()
}


type requestCounter struct{
	Url string
	Count int
}

type limitRateEngine struct{
	store Caching
}

func (c *limitRateEngine) WriteTTL(k string, value interface{}, exp int) {
	cached,ok := c.store.Read(k).(requestCounter)
	if !ok{
		c.store.Write(k,requestCounter{
			Url: k,
			Count: 1,
		})
		go func() {
			defer common.AppRecover()
			<-time.NewTimer(time.Duration(exp) * time.Second).C
			c.store.Write(k, value)
		}()
	}
	
	cached.Count+=1
	c.store.Write(k,cached)
}
