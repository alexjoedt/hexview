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
		{"positive max", "7f", 127, false},
		{"negative", "ff", -1, false},
		{"negative min", "80", -128, false},
		{"zero", "00", 0, false},
		{"with prefix", "0x7f", 127, false},
		{"overflow - too many bytes", "1234", 0, true},
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

func TestHexToInt16(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    int16
		wantErr bool
	}{
		{"positive max", "7fff", 32767, false},
		{"negative", "ffff", -1, false},
		{"negative min", "8000", -32768, false},
		{"zero", "0000", 0, false},
		{"auto-pad 1 byte", "12", 0x0012, false},
		{"auto-pad 1 byte max", "ff", 0x00ff, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToInt16(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToInt16LE(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    int16
		wantErr bool
	}{
		{"LE order", "ff7f", 32767, false},
		{"LE negative", "ffff", -1, false},
		{"LE zero", "0000", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToInt16LE(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToInt16LE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToInt16LE() = %v, want %v", got, tt.want)
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
		{"positive max", "7fffffff", 2147483647, false},
		{"negative", "ffffffff", -1, false},
		{"negative min", "80000000", -2147483648, false},
		{"zero", "00000000", 0, false},
		{"auto-pad 3 bytes", "2a 22 2b", 0x002a222b, false},
		{"auto-pad 1 byte", "ff", 0x000000ff, false},
		{"auto-pad 2 bytes", "1234", 0x00001234, false},
		{"overflow - too many bytes", "1234567890", 0, true},
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

func TestHexToInt64(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    int64
		wantErr bool
	}{
		{"positive max", "7fffffffffffffff", 9223372036854775807, false},
		{"negative", "ffffffffffffffff", -1, false},
		{"negative min", "8000000000000000", -9223372036854775808, false},
		{"zero", "0000000000000000", 0, false},
		{"auto-pad 4 bytes", "12345678", 0x0000000012345678, false},
		{"auto-pad 5 bytes", "1234567890", 0x0000001234567890, false},
		{"auto-pad 1 byte", "ff", 0x00000000000000ff, false},
		{"overflow - too many bytes", "12345678901234567890", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToInt64(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToInt64() = %v, want %v", got, tt.want)
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
		{"positive max", 2147483647, "7fffffff"},
		{"negative", -1, "ffffffff"},
		{"negative min", -2147483648, "80000000"},
		{"zero", 0, "00000000"},
		{"small positive", 100, "00000064"},
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

func TestInt64ToHex(t *testing.T) {
	tests := []struct {
		name string
		val  int64
		want string
	}{
		{"positive max", 9223372036854775807, "7fffffffffffffff"},
		{"negative", -1, "ffffffffffffffff"},
		{"zero", 0, "0000000000000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64ToHex(tt.val)
			if got != tt.want {
				t.Errorf("Int64ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToUint8(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    uint8
		wantErr bool
	}{
		{"max", "ff", 255, false},
		{"zero", "00", 0, false},
		{"mid", "7f", 127, false},
		{"overflow - too many bytes", "1234", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToUint8(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToUint16(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    uint16
		wantErr bool
	}{
		{"max", "ffff", 65535, false},
		{"zero", "0000", 0, false},
		{"mid", "7fff", 32767, false},
		{"auto-pad 1 byte", "ff", 0x00ff, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToUint16(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToUint16() = %v, want %v", got, tt.want)
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
		{"mid", "7fffffff", 2147483647, false},
		{"auto-pad 3 bytes", "2a222b", 0x002a222b, false},
		{"auto-pad 1 byte", "ff", 0x000000ff, false},
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

func TestHexToUint64(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    uint64
		wantErr bool
	}{
		{"max", "ffffffffffffffff", 18446744073709551615, false},
		{"zero", "0000000000000000", 0, false},
		{"mid", "7fffffffffffffff", 9223372036854775807, false},
		{"auto-pad 5 bytes", "1234567890", 0x0000001234567890, false},
		{"auto-pad 1 byte", "ff", 0x00000000000000ff, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToUint64(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("HexToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16ToHex(t *testing.T) {
	tests := []struct {
		name string
		val  uint16
		want string
	}{
		{"max", 65535, "ffff"},
		{"zero", 0, "0000"},
		{"mid", 32767, "7fff"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToHex(tt.val)
			if got != tt.want {
				t.Errorf("Uint16ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64ToHex(t *testing.T) {
	tests := []struct {
		name string
		val  uint64
		want string
	}{
		{"max", 18446744073709551615, "ffffffffffffffff"},
		{"zero", 0, "0000000000000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint64ToHex(tt.val)
			if got != tt.want {
				t.Errorf("Uint64ToHex() = %v, want %v", got, tt.want)
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
		isNaN   bool
		isInf   bool
		infSign int
	}{
		{"zero", "00000000", 0.0, false, false, false, 0},
		{"one", "3f800000", 1.0, false, false, false, 0},
		{"negative one", "bf800000", -1.0, false, false, false, 0},
		{"pi", "40490fdb", 3.14159265, false, false, false, 0},
		{"NaN", "7fc00000", 0, false, true, false, 0},
		{"positive infinity", "7f800000", 0, false, false, true, 1},
		{"negative infinity", "ff800000", 0, false, false, true, -1},
		{"auto-pad 3 bytes", "800000", 0.0, false, false, false, 0},
		{"overflow - too many bytes", "1234567890", 0, true, false, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToFloat32(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			// Check special float values
			if tt.isNaN {
				if got == got { // NaN != NaN
					t.Errorf("HexToFloat32() = %v, want NaN", got)
				}
				return
			}
			if tt.isInf {
				if (tt.infSign > 0 && got <= 0) || (tt.infSign < 0 && got >= 0) {
					t.Errorf("HexToFloat32() = %v, want infinity with sign %d", got, tt.infSign)
				}
				return
			}

			// Normal comparison with tolerance for floating point
			if diff := got - tt.want; diff < -0.00001 || diff > 0.00001 {
				t.Errorf("HexToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToFloat64(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    float64
		wantErr bool
	}{
		{"zero", "0000000000000000", 0.0, false},
		{"one", "3ff0000000000000", 1.0, false},
		{"negative one", "bff0000000000000", -1.0, false},
		{"pi", "400921fb54442d18", 3.141592653589793, false},
		{"large number", "40f86a0000000000", 100000.0, false},
		{"auto-pad 4 bytes", "00000000", 0.0, false},
		{"overflow - too many bytes", "12345678901234567890", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToFloat64(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := got - tt.want; diff < -0.0000000001 || diff > 0.0000000001 {
					t.Errorf("HexToFloat64() = %.15f, want %.15f", got, tt.want)
				}
			}
		})
	}
}

func TestFloat64ToHex(t *testing.T) {
	tests := []struct {
		name string
		val  float64
		want string
	}{
		{"zero", 0.0, "0000000000000000"},
		{"one", 1.0, "3ff0000000000000"},
		{"negative one", -1.0, "bff0000000000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64ToHex(tt.val)
			if got != tt.want {
				t.Errorf("Float64ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBinary(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []byte
		wantErr bool
	}{
		{"simple", "00001010", []byte{0x0a}, false},
		{"two bytes", "00001010 11110000", []byte{0x0a, 0xf0}, false},
		{"four bytes", "11111111 00000000 10101010 01010101", []byte{0xff, 0x00, 0xaa, 0x55}, false},
		{"no spaces", "1111111100000000", []byte{0xff, 0x00}, false},
		{"short input", "1010", []byte{0x0a}, false},
		{"single bit", "1", []byte{0x01}, false},
		{"empty", "", nil, true},
		{"invalid char", "0012", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBinary(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !bytesEqual(got, tt.want) {
				t.Errorf("ParseBinary() = %v, want %v", got, tt.want)
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
		{"positive max", "01111111", 127, false},
		{"negative", "11111111", -1, false},
		{"negative min", "10000000", -128, false},
		{"zero", "00000000", 0, false},
		{"with spaces", "0111 1111", 127, false},
		{"overflow - too many bits", "111111111111111111", 0, true},
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

func TestBinaryToInt32(t *testing.T) {
	tests := []struct {
		name    string
		bin     string
		want    int32
		wantErr bool
	}{
		{"positive", "01111111111111111111111111111111", 2147483647, false},
		{"negative", "11111111111111111111111111111111", -1, false},
		{"zero", "00000000000000000000000000000000", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryToInt32(tt.bin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("BinaryToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryToUint8(t *testing.T) {
	tests := []struct {
		name    string
		bin     string
		want    uint8
		wantErr bool
	}{
		{"max", "11111111", 255, false},
		{"zero", "00000000", 0, false},
		{"mid", "01111111", 127, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryToUint8(tt.bin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryToUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("BinaryToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryToUint16(t *testing.T) {
	tests := []struct {
		name    string
		bin     string
		want    uint16
		wantErr bool
	}{
		{"max", "1111111111111111", 65535, false},
		{"zero", "0000000000000000", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryToUint16(tt.bin)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryToUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("BinaryToUint16() = %v, want %v", got, tt.want)
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
		{"positive max", 127, "01111111"},
		{"negative", -1, "11111111"},
		{"negative min", -128, "10000000"},
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

func TestInt16ToBinary(t *testing.T) {
	tests := []struct {
		name string
		val  int16
		want string
	}{
		{"positive", 32767, "01111111 11111111"},
		{"negative", -1, "11111111 11111111"},
		{"zero", 0, "00000000 00000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int16ToBinary(tt.val)
			if got != tt.want {
				t.Errorf("Int16ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8ToBinary(t *testing.T) {
	tests := []struct {
		name string
		val  uint8
		want string
	}{
		{"max", 255, "11111111"},
		{"zero", 0, "00000000"},
		{"mid", 127, "01111111"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint8ToBinary(tt.val)
			if got != tt.want {
				t.Errorf("Uint8ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16ToBinary(t *testing.T) {
	tests := []struct {
		name string
		val  uint16
		want string
	}{
		{"max", 65535, "11111111 11111111"},
		{"zero", 0, "00000000 00000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint16ToBinary(tt.val)
			if got != tt.want {
				t.Errorf("Uint16ToBinary() = %v, want %v", got, tt.want)
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
		{"empty", []byte{}, ""},
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

// Additional coverage tests for all remaining functions
func TestAllIntToHexFunctions(t *testing.T) {
	// Test all Int8 variants
	if Int8ToHexLE(127) != "7f" {
		t.Error("Int8ToHexLE failed")
	}

	// Test Int16 variants
	if Int16ToHex(1000) != "03e8" {
		t.Error("Int16ToHex failed")
	}
	if Int16ToHexLE(1000) != "e803" {
		t.Error("Int16ToHexLE failed")
	}

	// Test Int32 variants
	if Int32ToHexLE(1000) != "e8030000" {
		t.Error("Int32ToHexLE failed")
	}

	// Test Int64 variants
	if Int64ToHex(1000) != "00000000000003e8" {
		t.Error("Int64ToHex failed")
	}
	if Int64ToHexLE(1000) != "e803000000000000" {
		t.Error("Int64ToHexLE failed")
	}
}

func TestAllUintToHexFunctions(t *testing.T) {
	// Test Uint8
	if Uint8ToHex(255) != "ff" {
		t.Error("Uint8ToHex failed")
	}
	if Uint8ToHexLE(255) != "ff" {
		t.Error("Uint8ToHexLE failed")
	}

	// Test Uint16
	if Uint16ToHexLE(1000) != "e803" {
		t.Error("Uint16ToHexLE failed")
	}

	// Test Uint32
	if Uint32ToHex(1000) != "000003e8" {
		t.Error("Uint32ToHex failed")
	}
	if Uint32ToHexLE(1000) != "e8030000" {
		t.Error("Uint32ToHexLE failed")
	}

	// Test Uint64
	if Uint64ToHexLE(1000) != "e803000000000000" {
		t.Error("Uint64ToHexLE failed")
	}
}

func TestAllLittleEndianIntConversions(t *testing.T) {
	// Int8 (no endianness difference)
	val8, _ := HexToInt8LE("7f")
	if val8 != 127 {
		t.Error("HexToInt8LE failed")
	}

	// Int32LE
	val32, _ := HexToInt32LE("e8030000")
	if val32 != 1000 {
		t.Error("HexToInt32LE failed")
	}

	// Int64LE
	val64, _ := HexToInt64LE("e803000000000000")
	if val64 != 1000 {
		t.Error("HexToInt64LE failed")
	}
}

func TestAllLittleEndianUintConversions(t *testing.T) {
	// Uint8LE (no endianness difference)
	val8, _ := HexToUint8LE("ff")
	if val8 != 255 {
		t.Error("HexToUint8LE failed")
	}

	// Uint16LE
	val16, _ := HexToUint16LE("e803")
	if val16 != 1000 {
		t.Error("HexToUint16LE failed")
	}

	// Uint32LE
	val32, _ := HexToUint32LE("e8030000")
	if val32 != 1000 {
		t.Error("HexToUint32LE failed")
	}

	// Uint64LE
	val64, _ := HexToUint64LE("e803000000000000")
	if val64 != 1000 {
		t.Error("HexToUint64LE failed")
	}
}

func TestFloatLittleEndian(t *testing.T) {
	// Float32LE
	val32, _ := HexToFloat32LE("0000803f")
	if val32 != 1.0 {
		t.Errorf("HexToFloat32LE failed: got %v, want 1.0", val32)
	}

	hex32 := Float32ToHexLE(1.0)
	if hex32 != "0000803f" {
		t.Errorf("Float32ToHexLE failed: got %v, want 0000803f", hex32)
	}

	// Float64LE
	val64, _ := HexToFloat64LE("000000000000f03f")
	if val64 != 1.0 {
		t.Errorf("HexToFloat64LE failed: got %v, want 1.0", val64)
	}

	hex64 := Float64ToHexLE(1.0)
	if hex64 != "000000000000f03f" {
		t.Errorf("Float64ToHexLE failed: got %v, want 000000000000f03f", hex64)
	}
}

func TestAllBinaryIntConversions(t *testing.T) {
	// Int16
	val16, _ := BinaryToInt16("0000001111101000")
	if val16 != 1000 {
		t.Error("BinaryToInt16 failed")
	}

	// Int16LE
	val16le, _ := BinaryToInt16LE("1110100000000011")
	if val16le != 1000 {
		t.Error("BinaryToInt16LE failed")
	}

	// Int32LE
	val32le, _ := BinaryToInt32LE("11101000000000110000000000000000")
	if val32le != 1000 {
		t.Error("BinaryToInt32LE failed")
	}

	// Int64
	val64, _ := BinaryToInt64("0000000000000000000000000000000000000000000000000000001111101000")
	if val64 != 1000 {
		t.Error("BinaryToInt64 failed")
	}

	// Int64LE
	val64le, _ := BinaryToInt64LE("1110100000000011000000000000000000000000000000000000000000000000")
	if val64le != 1000 {
		t.Error("BinaryToInt64LE failed")
	}
}

func TestAllBinaryUintConversions(t *testing.T) {
	// Uint32
	val32, _ := BinaryToUint32("00000000000000000000001111101000")
	if val32 != 1000 {
		t.Error("BinaryToUint32 failed")
	}

	// Uint32LE
	val32le, _ := BinaryToUint32LE("11101000000000110000000000000000")
	if val32le != 1000 {
		t.Error("BinaryToUint32LE failed")
	}

	// Uint64
	val64, _ := BinaryToUint64("0000000000000000000000000000000000000000000000000000001111101000")
	if val64 != 1000 {
		t.Error("BinaryToUint64 failed")
	}

	// Uint64LE
	val64le, _ := BinaryToUint64LE("1110100000000011000000000000000000000000000000000000000000000000")
	if val64le != 1000 {
		t.Error("BinaryToUint64LE failed")
	}
}

func TestAllBinaryToStringConversions(t *testing.T) {
	// Int32ToBinary
	if Int32ToBinary(1000) != "00000000 00000000 00000011 11101000" {
		t.Error("Int32ToBinary failed")
	}

	// Int32ToBinaryLE
	if Int32ToBinaryLE(1000) != "11101000 00000011 00000000 00000000" {
		t.Error("Int32ToBinaryLE failed")
	}

	// Int64ToBinary
	bin64 := Int64ToBinary(1000)
	if len(bin64) == 0 {
		t.Error("Int64ToBinary failed")
	}

	// Int64ToBinaryLE
	bin64le := Int64ToBinaryLE(1000)
	if len(bin64le) == 0 {
		t.Error("Int64ToBinaryLE failed")
	}

	// Uint32ToBinary
	if Uint32ToBinary(1000) != "00000000 00000000 00000011 11101000" {
		t.Error("Uint32ToBinary failed")
	}

	// Uint32ToBinaryLE
	if Uint32ToBinaryLE(1000) != "11101000 00000011 00000000 00000000" {
		t.Error("Uint32ToBinaryLE failed")
	}

	// Uint64ToBinary
	binU64 := Uint64ToBinary(1000)
	if len(binU64) == 0 {
		t.Error("Uint64ToBinary failed")
	}

	// Uint64ToBinaryLE
	binU64le := Uint64ToBinaryLE(1000)
	if len(binU64le) == 0 {
		t.Error("Uint64ToBinaryLE failed")
	}
}

func TestHexToBytes(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    []byte
		wantErr bool
	}{
		{"valid hex", "48656c6c6f", []byte("Hello"), false},
		{"with prefix", "0x48656c6c6f", []byte("Hello"), false},
		{"empty", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToBytes(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !bytesEqual(got, tt.want) {
				t.Errorf("HexToBytes() = %v, want %v", got, tt.want)
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
