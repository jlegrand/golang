package main

import (
	"fmt"
	"github.com/jlegrand/golang/Mail/mail"
	"github.com/jlegrand/golang/Mail/cache"
	"time"
)

func main() {

	r := mail.NewRepository()
	c := cache.NewCache(r)

	printMail(2, c)
	printMail(10, c)
	printMail(2, c)
	printMail(10, c)
	printMail(10, c)
	printMail(2, c)
}

func printMail(id uint64, c *cache.MailCache) {
	t1 := time.Now()
	m, ok := c.GetMail(id)
	if ok {
		fmt.Printf("%+v\n", m)
	} else {
		fmt.Println("Can't get mail")
	}
	fmt.Println("printMail execution time: ", time.Now().Sub(t1))
}
