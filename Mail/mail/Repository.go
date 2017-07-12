package mail

import (
	"sync"
)

type Repository struct {
	mails map[uint64]*Message
	rwMutex sync.RWMutex
}

func NewRepository() *Repository {
	var r *Repository = new(Repository)
	r.mails = make(map[uint64]*Message)
	return r
}

func (r *Repository) Get(index uint64) (*Message, bool) {
	r.rwMutex.RLock()
	m, ok := r.mails[index]
	r.rwMutex.RUnlock()
	return m, ok
}

func (r *Repository) Set(id uint64, msg *Message) {
	r.rwMutex.Lock()
	r.mails[id] = msg
	r.rwMutex.Unlock()
}

func (r *Repository) Del(id uint64) {
	r.rwMutex.Lock()
	delete(r.mails, id)
	r.rwMutex.Unlock()
}

