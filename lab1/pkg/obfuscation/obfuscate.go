package obfuscation

func Obfuscate(data []byte, shift int) []byte {
	obfuscated := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		obfuscated[i] = data[i] + byte(shift)
	}
	return obfuscated
}
