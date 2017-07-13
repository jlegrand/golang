package main

import (
	"fmt"
	"github.com/jlegrand/golang/Mail/mail"
	"github.com/jlegrand/golang/Mail/cache"
	"time"
	"strconv"
	"math/rand"
	"github.com/jlegrand/golang/Mail/server"
)

//var _execNum int = 0
const NB_MAILS int = 5
const NB_READ int = 20

func main() {

	r := mail.NewRepository()
	c := cache.NewCache(r)

	//testSync(r, c)
	testFileStorage(r, c)

}

func testFileStorage(r *mail.Repository, c *cache.Cache) {

	p := server.NewProvider("C:/Users/A454023/.babun/cygwin/home/a454023/golang/src/github.com/jlegrand/golang/Mail/server/_store")

	var msg *mail.Message
	msg = mail.New()
	msg.From = strconv.Itoa(1) + "@protonmail.com"
	msg.To[0] = strconv.Itoa(1+1) + "@protonmail.com"
	msg.Subject = strconv.Itoa(1) + "th mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"

	p.Set(1, msg)

	fmt.Printf("%+v\n", p.Get(1))

}

func testSync(r *mail.Repository, c *cache.Cache) {

	var msg *mail.Message
	for i := 0; i < NB_MAILS; i++ {
		msg = mail.New()
		msg.From = strconv.Itoa(i) + "@protonmail.com"
		msg.To[0] = strconv.Itoa(i+1) + "@protonmail.com"
		msg.Subject = strconv.Itoa(i) + "th mail"
		msg.SetHeader("lang", "en")
		msg.Body = "Hello"
		go func(i int, msg *mail.Message) {
			r.Set(uint64(i), msg)
		}(i, msg)
	}

	for i := 0; i < NB_READ; i++ {
		e := random(1, NB_MAILS)
		go func(e uint64) {
			printMail(e, c)
		}(uint64(e))
	}

}

func printMail(id uint64, c *cache.Cache) {
	fmt.Println("***************")
	fmt.Println("mail id ", id)
	t1 := time.Now()
	m, ok := c.GetMail(id)
	if ok {
		fmt.Printf("%+v\n", m)
	} else {
		fmt.Println("Can't get mail")
	}
	//fmt.Println("exec num ", _execNum)
	fmt.Println("exec time: ", time.Now().Sub(t1))
	//_execNum++
}

func random(min, max int) int {
	rand.Seed(rand.Int63n(1e9))
	return rand.Intn(max-min) + min
}
