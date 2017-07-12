package basics

import (
	"fmt"
	"math"
)

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


////////////////////////////////////////

type Circle struct{
	x int
	y int
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func UseArea() {
	circle := Circle{radius:10}
	rectangle := Rectangle{width:10, height:20}

	fmt.Printf("Circle area : %f\n", getArea(circle))
	fmt.Printf("Rectangle area : %f\n", getArea(rectangle))
}
