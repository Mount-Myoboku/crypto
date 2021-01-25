package caesarcipher

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type caesarCipherBuilder struct {
	plainTextRuneValueMap  map[rune]int
	plainTextValueRuneMap  map[int]rune
	cipherTextValueRuneMap map[int]rune
	cipherTextRuneValueMap map[rune]int
}

func Init(shift int) caesarCipherBuilder {

	var ccb caesarCipherBuilder

	var plainTextRuneValueMap = make(map[rune]int)
	var plainTextValueRuneMap = make(map[int]rune)

	var cipherTextValueRuneMap = make(map[int]rune)
	var cipherTextRuneValueMap = make(map[rune]int)

	for i, v := range alphabet {
		plainTextRuneValueMap[v] = i + 1
		plainTextValueRuneMap[i+1] = v
		newIndex := i + shift
		newRune := rune(alphabet[newIndex%26])
		cipherTextValueRuneMap[i+1] = newRune
		cipherTextRuneValueMap[newRune] = i + 1
	}

	ccb.plainTextRuneValueMap = plainTextRuneValueMap
	ccb.cipherTextValueRuneMap = cipherTextValueRuneMap
	ccb.plainTextValueRuneMap = plainTextValueRuneMap
	ccb.cipherTextRuneValueMap = cipherTextRuneValueMap

	return ccb
}

func (ccb *caesarCipherBuilder) Encrypt(input string) string {

	var output string

	for _, char := range input {
		newRune := ccb.cipherTextValueRuneMap[(ccb.plainTextRuneValueMap[char])]
		output = output + string(newRune)
	}

	return output
}

func (ccb *caesarCipherBuilder) Decrypt(input string) string {

	var output string

	for _, char := range input {
		newRune := ccb.plainTextValueRuneMap[(ccb.cipherTextRuneValueMap[char])]
		output = output + string(newRune)
	}

	return output
}
