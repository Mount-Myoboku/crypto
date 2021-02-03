package caesarcipher

import (
	"testing"

	"github.com/jaswdr/faker"
)

func TestEncrypt(t *testing.T) {
	output := Encrypt("HELLO")
	if output != "EBIIL" {
		t.Error("expected EBIIL, but got ", output)
	}

}

func TestDecrypt(t *testing.T) {
	output := Decrypt("EBIIL")
	if output != "HELLO" {
		t.Error("expected HELLO, but got ", output)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	f := faker.New()
	for i := 0; i < b.N; i++ {
		Encrypt(f.Lorem().Paragraph(10))
	}
}
