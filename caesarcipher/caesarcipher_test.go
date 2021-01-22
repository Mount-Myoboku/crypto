package caesarcipher

import "testing"

func TestEncrypt(t *testing.T) {
	output := Decrypt("HELLO")

	if output != "EBIIL" {
		t.Error("expected EBIIL, but got ", output)
	}

	output = Encrypt("EBIIL")

	if output != "HELLO" {
		t.Error("expected HELLO, but got ", output)
	}

}