package server

import (
	"testing"
	"os"
	"github.com/jlegrand/golang/Mail/mail"
	"crypto/sha256"
	"io"
	"bytes"
)

func TestMain(m *testing.M) {

	//setup

	os.Exit(m.Run())
}

func TestProvider_Get(t *testing.T) {

	p := NewProvider("C:/Users/A454023/.babun/cygwin/home/a454023/golang/src/github.com/jlegrand/golang/Mail/server/test/_store")
	msg := p.Get(1)

	switch {
	case msg.From != "1@protonmail.com":
		t.Fatal("Invalid mail structure")
	case msg.To[0] != "2@protonmail.com":
		t.Fatal("Invalid mail structure")
	case msg.Body != "Hello":
		t.Fatal("Invalid mail structure")
	case msg.Subject != "1th mail":
		t.Fatal("Invalid mail structure")
	}
}

func TestProvider_Set(t *testing.T)  {

	testing_store_path := "C:/Users/A454023/.babun/cygwin/home/a454023/golang/src/github.com/jlegrand/golang/Mail/server/test/_store"

	p := NewProvider(testing_store_path)

	var msg *mail.Message
	msg = mail.New()
	msg.From = "1@protonmail.com"
	msg.To[0] = "2@protonmail.com"
	msg.Subject ="1th mail"
	msg.SetHeader("lang", "en")
	msg.Body = "Hello"

	p.Set(2, msg)


	if bytes.Compare(sha(testing_store_path + "/1.json"), sha(testing_store_path + "/2.json")) != 0 {
		t.Fatal("Invalid provider.Set")
	}



}

func sha(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}

	return h.Sum(nil)
}
