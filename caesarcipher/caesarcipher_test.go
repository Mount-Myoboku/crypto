package caesarcipher

import (
	"testing"

	"github.com/jaswdr/faker"
)

var cb = Init(5)

func TestEncrypt(t *testing.T) {

	output := cb.Encrypt("HELLO")

	if output != "MJQQT" {
		t.Error("expected MJQQT, but got ", output)
	}

}

func TestDecrypt(t *testing.T) {

	output := cb.Decrypt("MJQQT")

	if output != "HELLO" {
		t.Error("expected HELLO, but got ", output)
	}
}

func BenchmarkEncrypt(b *testing.B) {

	f := faker.New()

	for i := 0; i < b.N; i++ {
		cb.Encrypt(f.Lorem().Paragraph(10))
	}
}
