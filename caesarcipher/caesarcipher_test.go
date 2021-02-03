package caesarcipher

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
)

type testCase struct {
	input  string
	output string
}

func TestEncrypt(t *testing.T) {

	cases := []testCase{
		{input: "HELLO", output: "EBIIL"},
		{input: "!*@{,", output: "!*@{,"},
		{input: "ZAZZY", output: "WXWWV"},
	}

	for _, test := range cases {
		output := Encrypt(test.input)
		if output != test.output {
			t.Errorf("expected %s, but got %s", test.output, output)
		}
	}

}

func TestDecrypt(t *testing.T) {

	cases := []testCase{
		{output: "HELLO", input: "EBIIL"},
		{output: "!*@{,", input: "!*@{,"},
		{output: "ZAZZY", input: "WXWWV"},
	}

	for _, test := range cases {
		output := Decrypt(test.input)
		if output != test.output {
			t.Errorf("expected %s, but got %s", test.output, output)
		}
	}
}

func BenchmarkEncrypt(b *testing.B) {
	f := faker.New()
	for i := 0; i < b.N; i++ {
		Encrypt(f.Lorem().Paragraph(10))
	}
}

func BenchmarkDecrypt(b *testing.B) {
	f := faker.New()
	for i := 0; i < b.N; i++ {
		Decrypt(f.Lorem().Paragraph(10))
	}
}

func ExampleEncrypt() {
	output := Encrypt("HELLO!")
	fmt.Println(output)
	// "EBIIL!"
}

func ExampleDecrypt() {
	output := Encrypt("EBIIL!")
	fmt.Println(output)
	// "HELLO!"
}
