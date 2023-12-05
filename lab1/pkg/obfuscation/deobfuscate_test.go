package obfuscation

import (
	"reflect"
	"testing"
)

func TestDeobfuscate(t *testing.T) {
	tests := []struct {
		input  []byte
		shift  int
		output []byte
	}{
		{[]byte{4, 5, 6}, 2, []byte{2, 3, 4}},
		{[]byte{10, 20, 30}, 5, []byte{5, 15, 25}},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		result := Deobfuscate(tt.input, tt.shift)
		if !reflect.DeepEqual(result, tt.output) {
			t.Errorf("Deobfuscate(%v, %d) = %v; want %v", tt.input, tt.shift, result, tt.output)
		}
	}
}
