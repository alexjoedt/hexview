package convert

import (
	"testing"
)

// TestHexDisplayFormatting tests that hex display always shows big-endian representation
// of values, regardless of which endianness was used to parse the original bytes.
// This is the fix for: "Incorrect hex value in expert mode" issue.
func TestHexDisplayFormatting(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		int32BE    string
		int32LE    string
		int32BADC  string
		int32CDAB  string
		uint32BE   string
		uint32LE   string
		uint32BADC string
		uint32CDAB string
	}{
		{
			name:       "4-byte example from issue: 11 22 33 44",
			input:      "11223344",
			int32BE:    "11223344",
			int32LE:    "44332211",
			int32BADC:  "22114433",
			int32CDAB:  "33441122",
			uint32BE:   "11223344",
			uint32LE:   "44332211",
			uint32BADC: "22114433",
			uint32CDAB: "33441122",
		},
		{
			name:       "sequential bytes",
			input:      "01020304",
			int32BE:    "01020304",
			int32LE:    "04030201",
			int32BADC:  "02010403",
			int32CDAB:  "03040102",
			uint32BE:   "01020304",
			uint32LE:   "04030201",
			uint32BADC: "02010403",
			uint32CDAB: "03040102",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test signed int32
			valBE, _ := HexToInt32(tt.input)
			if got := Int32ToHex(valBE); got != tt.int32BE {
				t.Errorf("Int32BE hex = %v, want %v", got, tt.int32BE)
			}

			valLE, _ := HexToInt32LE(tt.input)
			if got := Int32ToHexLE(valLE); got != tt.int32LE {
				t.Errorf("Int32LE hex = %v, want %v", got, tt.int32LE)
			}

			valBADC, _ := HexToInt32BADC(tt.input)
			if got := Int32ToHexBADC(valBADC); got != tt.int32BADC {
				t.Errorf("Int32BADC hex = %v, want %v", got, tt.int32BADC)
			}

			valCDAB, _ := HexToInt32CDAB(tt.input)
			if got := Int32ToHexCDAB(valCDAB); got != tt.int32CDAB {
				t.Errorf("Int32CDAB hex = %v, want %v", got, tt.int32CDAB)
			}

			// Test unsigned uint32
			uvalBE, _ := HexToUint32(tt.input)
			if got := Uint32ToHex(uvalBE); got != tt.uint32BE {
				t.Errorf("Uint32BE hex = %v, want %v", got, tt.uint32BE)
			}

			uvalLE, _ := HexToUint32LE(tt.input)
			if got := Uint32ToHexLE(uvalLE); got != tt.uint32LE {
				t.Errorf("Uint32LE hex = %v, want %v", got, tt.uint32LE)
			}

			uvalBADC, _ := HexToUint32BADC(tt.input)
			if got := Uint32ToHexBADC(uvalBADC); got != tt.uint32BADC {
				t.Errorf("Uint32BADC hex = %v, want %v", got, tt.uint32BADC)
			}

			uvalCDAB, _ := HexToUint32CDAB(tt.input)
			if got := Uint32ToHexCDAB(uvalCDAB); got != tt.uint32CDAB {
				t.Errorf("Uint32CDAB hex = %v, want %v", got, tt.uint32CDAB)
			}
		})
	}
}

// TestHexDisplayFormatting64 tests 64-bit values
func TestHexDisplayFormatting64(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		int64BE    string
		int64LE    string
		int64BADC  string
		int64CDAB  string
	}{
		{
			name:      "8-byte example from issue: 11 22 33 44 55 66 77 88",
			input:     "1122334455667788",
			int64BE:   "1122334455667788",
			int64LE:   "8877665544332211",
			int64BADC: "2211443366558877",
			int64CDAB: "3344112277885566",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valBE, _ := HexToInt64(tt.input)
			if got := Int64ToHex(valBE); got != tt.int64BE {
				t.Errorf("Int64BE hex = %v, want %v", got, tt.int64BE)
			}

			valLE, _ := HexToInt64LE(tt.input)
			if got := Int64ToHexLE(valLE); got != tt.int64LE {
				t.Errorf("Int64LE hex = %v, want %v", got, tt.int64LE)
			}

			valBADC, _ := HexToInt64BADC(tt.input)
			if got := Int64ToHexBADC(valBADC); got != tt.int64BADC {
				t.Errorf("Int64BADC hex = %v, want %v", got, tt.int64BADC)
			}

			valCDAB, _ := HexToInt64CDAB(tt.input)
			if got := Int64ToHexCDAB(valCDAB); got != tt.int64CDAB {
				t.Errorf("Int64CDAB hex = %v, want %v", got, tt.int64CDAB)
			}
		})
	}
}
