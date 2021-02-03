package caesarcipher

func encode(input string, shift int, flip bool) string {

	var output string

	for _, char := range input {
		newChar := char

		switch flip {
		case true:
			if (int(char) >= 'a' && char <= 'z') || char >= 'A' && char <= 'Z' {
				newChar = rune(int(char) + shift)
				if int(newChar) > 'z' || int(newChar) > 'Z' && int(newChar) < 'a' {
					newChar -= 26
				}
			}

		case false: // default case - use to encrypt

			if (int(char) >= 'a' && char <= 'z') || char >= 'A' && char <= 'Z' {
				newChar = rune(int(char) - shift)
				if int(newChar) < 'A' || int(newChar) > 'Z' && int(newChar) < 'a' {
					newChar += 26
				}
			}
		}

		output += string(newChar)
	}
	return output
}

// Encrypt encrypts the input
func Encrypt(input string) string {
	return encode(input, 3, false)
}

// Decrypt decrypts the input
func Decrypt(input string) string {
	return encode(input, 3, true)
}
