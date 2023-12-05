package obfuscation

import (
	"reflect"
	"testing"
)

func TestObfuscate(t *testing.T) {
	tests := []struct {
		input  []byte
		shift  int
		output []byte
	}{
		{[]byte{2, 3, 4}, 2, []byte{4, 5, 6}},
		{[]byte{5, 15, 25}, 5, []byte{10, 20, 30}},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		result := Obfuscate(tt.input, tt.shift)
		if !reflect.DeepEqual(result, tt.output) {
			t.Errorf("Obfuscate(%v, %d) = %v; want %v", tt.input, tt.shift, result, tt.output)
		}
	}
}
