package main

import (
	"fmt"
	"github.com/jlegrand/golang/Mail/mail"
	"github.com/jlegrand/golang/Mail/cache"
)

func main() {

	r := mail.NewRepository()
	c := cache.NewCache(r)

	printMail(2, c)
	printMail(10, c)
	printMail(2, c)
}

func printMail(id uint64, c *cache.Mail) {
	m, ok := c.GetMail(id)
	if ok {
		fmt.Printf("%+v\n", m)
	} else {
		fmt.Println("Can't get mail")
	}
}
