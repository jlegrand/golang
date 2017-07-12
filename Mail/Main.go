package main

import (
	"github.com/jlegrand/golang/Mail/mail"
	"fmt"
	"github.com/jlegrand/golang/Mail/cache"
)

func main() {

	/*msg := mail.New()
	msg.From = "julien.legrand@protonmail.com"
	msg.To[0] = "julien.legrand@protonmail.com"
	msg.Subject = "1st mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"

	msg.Send()

	m, ok := mail.Get(0)
	if ok {
		fmt.Printf("%+v\n", m)
	}*/

	r := mail.NewRepository()
	r.Init()

	c := cache.NewCache(r)
	m, ok := c.GetMail(2)
	if ok {
		fmt.Printf("%+v\n", m)
	} else {
		panic("Can't get mail")
	}

	m, ok = c.GetMail(10)
	if ok {
		fmt.Printf("%+v\n", m)
	} else {
		panic("Can't get mail")
	}



}
