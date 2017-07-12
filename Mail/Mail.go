package mail

import "fmt"

type Mail struct {
	from string
	to []string
	header map[string]string
}

// fonction contextualisé à la structure Mail (m *Mail)
func (m *Mail) toFile(path string){
	fmt.Println(m.from)
}

func SendMail() {
	var m *Mail = new(Mail)
	m.to = make([]string,0)
	m.header = make(map[string]string)

	m.toFile("/tmp/out.mail")
}
