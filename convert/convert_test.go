package convert

import (
	"testing"
)

// ============================================================================
// ParseHex Tests
// ============================================================================

func TestParseHex(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []byte
		wantErr bool
	}{
		{"standard prefix", "0x48656c6c6f", []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}, false},
		{"uppercase prefix", "0X48656C6C6F", []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}, false},
		{"x prefix", "x48656c6c6f", []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}, false},
		{"space separated", "04 ab cd", []byte{0x04, 0xab, 0xcd}, false},
		{"continuous", "11abcd", []byte{0x11, 0xab, 0xcd}, false},
		{"multiple prefixes", "0xab 0xff", []byte{0xab, 0xff}, false},
		{"comma separated", "0x12,0x34,0x56", []byte{0x12, 0x34, 0x56}, false},
		{"colon separated", "AA:BB:CC", []byte{0xaa, 0xbb, 0xcc}, false},
		{"mixed case", "aAbBcC", []byte{0xaa, 0xbb, 0xcc}, false},
		{"single byte", "ff", []byte{0xff}, false},
		{"odd length", "123", []byte{0x01, 0x23}, false},
		{"empty", "", nil, true},
		{"invalid char", "0xGG", nil, true},
		{"only prefix", "0x", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHex(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !bytesEqual(got, tt.want) {
				t.Errorf("ParseHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToInt8(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    int8
		wantErr bool
	}{
		{"positive", "7f", 127, false},
		{"negative", "ff", -1, false},
		{"zero", "00", 0, false},
		{"wrong length", "1234", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToInt8(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToInt32(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    int32
		wantErr bool
	}{
		{"positive BE", "7fffffff", 2147483647, false},
		{"negative BE", "ffffffff", -1, false},
		{"zero", "00000000", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToInt32(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32ToHex(t *testing.T) {
	tests := []struct {
		name string
		val  int32
		want string
	}{
		{"positive", 2147483647, "7fffffff"},
		{"negative", -1, "ffffffff"},
		{"zero", 0, "00000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32ToHex(tt.val)
			if got != tt.want {
				t.Errorf("Int32ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToUint32(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    uint32
		wantErr bool
	}{
		{"max", "ffffffff", 4294967295, false},
		{"zero", "00000000", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToUint32(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToFloat32(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    float32
		wantErr bool
	}{
		{"zero", "00000000", 0.0, false},
		{"one", "3f800000", 1.0, false},
		{"negative one", "bf800000", -1.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToFloat32(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryToInt8(t *testing.T) {
	tests := []struct {
		name    string
		bin     string
		want    int8
		wantErr bool
	}{
		{"positive", "01111111", 127, false},
		{"negative", "11111111", -1, false},
		{"zero", "00000000", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryToInt8(tt.bin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryToInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("BinaryToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8ToBinary(t *testing.T) {
	tests := []struct {
		name string
		val  int8
		want string
	}{
		{"positive", 127, "01111111"},
		{"negative", -1, "11111111"},
		{"zero", 0, "00000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int8ToBinary(tt.val)
			if got != tt.want {
				t.Errorf("Int8ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndianness(t *testing.T) {
	hex := "12345678"

	be, _ := HexToInt32(hex)
	le, _ := HexToInt32LE(hex)

	if be == le {
		t.Error("Big-endian and little-endian should produce different results")
	}

	if be != 0x12345678 {
		t.Errorf("Big-endian: got 0x%x, want 0x12345678", be)
	}
	if le != 0x78563412 {
		t.Errorf("Little-endian: got 0x%x, want 0x78563412", le)
	}
}

func TestRoundTripInt32(t *testing.T) {
	tests := []int32{0, 1, -1, 127, -128, 32767, -32768, 2147483647, -2147483648}

	for _, val := range tests {
		hex := Int32ToHex(val)
		got, err := HexToInt32(hex)
		if err != nil {
			t.Errorf("Round trip failed for %d: %v", val, err)
		}
		if got != val {
			t.Errorf("Round trip failed: got %v, want %v", got, val)
		}
	}
}

func TestRoundTripFloat64(t *testing.T) {
	tests := []float64{0.0, 1.0, -1.0, 3.14159, -3.14159}

	for _, val := range tests {
		hex := Float64ToHex(val)
		got, err := HexToFloat64(hex)
		if err != nil {
			t.Errorf("Round trip failed for %f: %v", val, err)
		}
		if got != val {
			t.Errorf("Round trip failed: got %v, want %v", got, val)
		}
	}
}

func TestBytesToHex(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  string
	}{
		{"hello", []byte("Hello"), "48656c6c6f"},
		{"empty", []byte{}, ""},
		{"single byte", []byte{0xff}, "ff"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToHex(tt.bytes)
			if got != tt.want {
				t.Errorf("BytesToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToBinary(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  string
	}{
		{"single byte", []byte{0x0a}, "00001010"},
		{"multiple bytes", []byte{0xff, 0x00}, "11111111 00000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BytesToBinary(tt.bytes)
			if got != tt.want {
				t.Errorf("BytesToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
