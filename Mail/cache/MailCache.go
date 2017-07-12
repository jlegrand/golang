package cache

import (
	"github.com/jlegrand/golang/Mail/mail"
	"sync"
)

type Provider interface {
	Get(uint64 uint64) (*mail.Message, bool)
}

type MailCache struct{
	cache    map[uint64]*mail.Message
	provider Provider
	rwMutex  sync.RWMutex
}

func NewCache(p Provider) *MailCache {
	var m *MailCache = new(MailCache)
	m.cache = make(map[uint64]*mail.Message)
	m.provider = p
	return m
}

func (m *MailCache) GetMail(id uint64) (*mail.Message, bool) {

	m.rwMutex.RLock()
	msg, ok := m.cache[id]
	m.rwMutex.RUnlock()
	if ok {
		return msg, ok
	} else {
		if msg, ok := m.provider.Get(id); ok {
			m.SetMail(id, msg)
			return msg, true
		} else {
			m.SetMail(id, nil)
			return nil, false
		}
	}
}

func (m *MailCache) SetMail(id uint64, mail *mail.Message) {
	m.rwMutex.Lock()
	m.cache[id] = mail
	m.rwMutex.Unlock()
}

func (m *MailCache) DelMail(id uint64) {
	m.rwMutex.Lock()
	delete(m.cache, id)
	m.rwMutex.Unlock()
}
