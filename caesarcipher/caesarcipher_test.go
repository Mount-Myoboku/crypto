package caesarcipher

import "testing"

func TestEncrypt(t *testing.T) {

	cb := Init(5)

	output := cb.Encrypt("HELLO")

	if output != "MJQQT" {
		t.Error("expected MJQQT, but got ", output)
	}

	output = cb.Decrypt("MJQQT")

	if output != "HELLO" {
		t.Error("expected HELLO, but got ", output)
	}

}