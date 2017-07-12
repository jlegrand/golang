package mail

import "fmt"

type message struct {
	id uint64
	From string
	To []string
	Subject string
	head map[string]string
	Body string
}

var _id uint64 = 0

func New() *message {
	var m *message = new(message)

	m.id = _id
	m.To = make([]string, 1)
	m.head = make(map[string]string)
	_id += 1
	return m
}

func (m *message) GetHeader(key string) (string, bool) {
	val, ok := m.head[key]
	return val, ok
}

func (m *message) SetHeader(key, val string) {
	m.head[key] = val
}

func (m *message) UnsetHeader(key string) {
	delete(m.head, key)
}

func (m *message) Send() {
	fmt.Println("Sending message: ", m.id)
	fmt.Println("  From: ", m.From)
	for _, to := range m.To {
		fmt.Println("  To: ", to)
	}
	fmt.Println("  Subjet:", m.Subject)
	fmt.Println("-- headers --")
	for k, v := range m.head {
		fmt.Println("  ", k, ": ", v)
	}
	fmt.Println("-- body --")
	fmt.Println("  ", m.Body)
}

