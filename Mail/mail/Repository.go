package mail

import "time"

type Repository struct {
	mails map[uint64]*Message
}

func NewRepository() *Repository {
	var r *Repository = new(Repository)
	r.mails = make(map[uint64]*Message)

	var msg *Message
	r.mails = make(map[uint64]*Message)

	msg = New()
	msg.From = "jle@protonmail.com"
	msg.To[0] = "phd@protonmail.com"
	msg.Subject = "1st mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"
	r.mails[msg.id] = msg

	msg = New()
	msg.From = "danesa@protonmail.com"
	msg.To[0] = "vic@protonmail.com"
	msg.Subject = "2nd mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"
	r.mails[msg.id] = msg

	msg = New()
	msg.From = "mga@protonmail.com"
	msg.To[0] = "igor@protonmail.com"
	msg.Subject = "2nd mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"
	r.mails[msg.id] = msg

	msg = New()
	msg.From = "phd@protonmail.com"
	msg.To[0] = "trex@protonmail.com"
	msg.Subject = "2nd mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"
	r.mails[msg.id] = msg

	return r
}

func (r *Repository) Get(index uint64) (*Message, bool) {
	m, ok := r.mails[index]
	time.Sleep(500 * time.Millisecond)
	return m, ok
}

