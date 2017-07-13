package cache

import (
	"github.com/jlegrand/golang/Mail/mail"
	"sync"
	"time"
	"sync/atomic"
)

var _ts uint64 = 0

type Provider interface {
	Get(uint64) (*mail.Message, bool)
	Set(uint64, *mail.Message)
	Del(uint64)
}

type Cache struct{
	cache    map[uint64]*MailCache
	provider Provider
	rwMutex  sync.RWMutex
	ticker *time.Ticker
}

type MailCache struct {
	message *mail.Message
	timestamp uint64
}

func NewMailCache(mail *mail.Message) *MailCache {
	x:= new(MailCache)
	x.message = mail
	x.timestamp = atomic.AddUint64(&_ts, 10)
	return x
}

func (mc *MailCache) update() {
	// update timestamp
}

func NewCache(p Provider) *Cache {
	var c *Cache = new(Cache)
	c.cache = make(map[uint64]*MailCache)
	c.provider = p
	c.ticker = time.NewTicker(10*time.Second)

	go func(c *Cache) {
		for range c.ticker.C {
			c.rwMutex.Lock()

			for id, mc := range c.cache {
				if mc.timestamp < _ts {
					delete(c.cache, id)
				}
			}
			c.rwMutex.Unlock()

			// inceremente _ts de façon atomic
			atomic.AddUint64(&_ts, 1)
		}
	}(c)

	return c
}


func (m *Cache) GetMail(id uint64) (*mail.Message, bool) {

	m.rwMutex.RLock()
	mc, ok := m.cache[id]
	m.rwMutex.RUnlock()
	if ok {
		mc.update()
		return mc.message, ok
	} else {
		if msg, ok := m.provider.Get(id); ok {
			mc := NewMailCache(msg)
			m.SetMail(id, mc)
			return msg, true
		} else {
			m.SetMail(id, nil)
			return nil, false
		}
	}
}

func (m *Cache) SetMail(id uint64, mail *MailCache) {
	m.rwMutex.Lock()
	m.cache[id] = mail
	m.rwMutex.Unlock()
}

func (m *Cache) DelMail(id uint64) {
	m.rwMutex.Lock()
	delete(m.cache, id)
	m.rwMutex.Unlock()
}
