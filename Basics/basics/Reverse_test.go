package basics

import (
	"testing"
	"strings"
	"bytes"
)

func TestReverse(t *testing.T) {

	actualResult := Reverse("Hello")
	var expectedResult = "olleH"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

/*func BenchmarkReverse(b *testing.B) {

	Reverse("Hello")
	b.StopTimer()

}*/

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}
func BenchmarkJoin(b *testing.B) {
	var str string
	var a []string
	for n := 0; n < b.N; n++ {
		a = append(a, "x")
	}

	str = strings.Join(a, "")

	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s{
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}
func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}
