package cache

import (
	"github.com/jlegrand/golang/Mail/mail"
)

type Provider interface {
	Get(uint64 uint64) (*mail.Message, bool)
}

type Mail struct{
	cache map[uint64]*mail.Message
	provider Provider
}

func NewCache(p Provider) *Mail {
	var m *Mail = new(Mail)
	m.cache = make(map[uint64]*mail.Message)
	m.provider = p
	return m
}

func (m *Mail) GetMail(id uint64) *mail.Message {

	if msg, ok := m.cache[id]; ok {
		return msg
	} else {
		if msg, ok := m.provider.Get(id); ok {
			m.SetMail(id, msg)
			return msg
		} else {
			m.SetMail(id, nil)
			return nil
		}
	}
}

func (m *Mail) SetMail(id uint64, mail *mail.Message) {
	m.cache[id] = mail
}

func (m *Mail) DelMail(id uint64) {
	delete(m.cache, id)
}
