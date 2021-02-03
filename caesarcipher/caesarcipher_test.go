package caesarcipher

import (
	"fmt"
	"testing"

	// "github.com/jaswdr/faker"
	"github.com/jaswdr/faker"
	"gopkg.in/check.v1"
)

// type testCase struct {
// 	input  string
// 	output string
// }

// func TestEncrypt(t *testing.T) {

// 	cases := []testCase{
// 		{input: "HELLO", output: "EBIIL"},
// 		{input: "!*@{,", output: "!*@{,"},
// 		{input: "ZAZZY", output: "WXWWV"},
// 	}

// 	for _, test := range cases {
// 		output := Encrypt(test.input)
// 		if output != test.output {
// 			t.Errorf("expected %s, but got %s", test.output, output)
// 		}
// 	}

// }

// func TestDecrypt(t *testing.T) {

// 	cases := []testCase{
// 		{output: "HELLO", input: "EBIIL"},
// 		{output: "!*@{,", input: "!*@{,"},
// 		{output: "ZAZZY", input: "WXWWV"},
// 	}

// 	for _, test := range cases {
// 		output := Decrypt(test.input)
// 		if output != test.output {
// 			t.Errorf("expected %s, but got %s", test.output, output)
// 		}
// 	}
// }

// func BenchmarkEncrypt(b *testing.B) {
// 	f := faker.New()
// 	for i := 0; i < b.N; i++ {
// 		Encrypt(f.Lorem().Paragraph(10))
// 	}
// }

// func BenchmarkDecrypt(b *testing.B) {
// 	f := faker.New()
// 	for i := 0; i < b.N; i++ {
// 		Decrypt(f.Lorem().Paragraph(10))
// 	}
// }

// func ExampleEncrypt() {
// 	output := Encrypt("HELLO!")
// 	fmt.Println(output)
// 	// "EBIIL!"
// }

// func ExampleDecrypt() {
// 	output := Encrypt("EBIIL!")
// 	fmt.Println(output)
// 	// "HELLO!"
// }

func Test(t *testing.T) { check.TestingT(t) }

type TestSuite struct {
	cases []TestCase
}

type TestCase struct {
	plain  string
	cipher string
}

var _ = check.Suite(&TestSuite{})

func (s *TestSuite) SetUpSuite(c *check.C) {
	fmt.Println("HELLO")
	s.cases = []TestCase{
		{plain: "HELLO", cipher: "EBIIL"},
		{plain: "!*@{,", cipher: "!*@{,"},
		{plain: "ZAZZY", cipher: "WXWWV"},
	}
}

func (s *TestSuite) TestEncrypt(c *check.C) {

	for _, test := range s.cases {
		output := Encrypt(test.plain)
		c.Assert(output, check.Equals, test.cipher, check.Commentf("expected %s, but got %s", test.cipher, output))
	}

}

func (s *TestSuite) TestDecrypt(c *check.C) {

	for _, test := range s.cases {
		output := Decrypt(test.cipher)
		c.Assert(output, check.Equals, test.plain, check.Commentf("expected %s, but got %s", test.plain, output))
	}
}

func (s *TestSuite) BenchmarkEncrypt(c *check.C) {
	f := faker.New()
	for i := 0; i < c.N; i++ {
		Encrypt(f.Lorem().Paragraph(10))
	}
}

func (s *TestSuite) BenchmarkDecrypt(c *check.C) {
	f := faker.New()
	for i := 0; i < c.N; i++ {
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
