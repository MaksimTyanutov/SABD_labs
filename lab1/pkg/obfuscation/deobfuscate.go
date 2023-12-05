package obfuscation

func Deobfuscate(data []byte, shift int) []byte {
	deobfuscated := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		deobfuscated[i] = data[i] - byte(shift)
	}
	return deobfuscated
}
