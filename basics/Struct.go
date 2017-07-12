package basics

import "fmt"

type Book struct{
	title string
	author string
	subject string
	book_id int
}

func Struct() {

	var Book1 Book
	// new(Type) = malloc structure
	var Book2 *Book = new(Book)

	Book1.title = "Go Programming"
	Book1.author = "MK"
	Book1.subject = "CS"
	Book1.book_id = 62582

	Book2.title = "dfff"
	Book2.author = "TF"
	Book2.subject = "CS"
	Book2.book_id = 84158
}

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