package caesarcipher

import "strings"

const plain =  "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const cipher =   "XYZABCDEFGHIJKLMNOPQRSTUVW"

func Encrypt(input string) string {

	var output string

	for _, char := range input {
		i := strings.Index(plain, string(char))
		newI := cipher[i]
		output = output + string(newI)
	}

	return output
}

func Decrypt(input string) string {

	var output string

	for _, char := range input {
		i := strings.Index(cipher, string(char))
		newI := plain[i]
		output = output + string(newI)
	}

	return output
}