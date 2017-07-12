package main

import (
	"github.com/jlegrand/golang/Mail/mail"
)

func main() {

	msg := mail.New()
	msg.From = "julien.legrand@protonmail.com"
	msg.To[0] = "julien.legrand@protonmail.com"
	msg.Subject = "1st mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"

	msg.Send()

}
