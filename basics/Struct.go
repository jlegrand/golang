package basics

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

