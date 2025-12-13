package service

import (
	"testing"
)

func TestNewConverter(t *testing.T) {
	c := NewConverter()
	if c == nil {
		t.Error("NewConverter() returned nil")
	}
}

func TestConvertHex_EmptyInput(t *testing.T) {
	c := NewConverter()
	_, err := c.ConvertHex("")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestConvertHex_ValidInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"single byte", "FF", false},
		{"two bytes", "AABB", false},
		{"four bytes", "DEADBEEF", false},
		{"with 0x prefix", "0xDEADBEEF", false},
		{"lowercase", "deadbeef", false},
		{"with spaces", "DE AD BE EF", false},
		{"invalid hex", "GHIJ", true},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertHex(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertHex(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("ConvertHex(%q) returned nil result", tt.input)
			}
		})
	}
}

func TestConvertHex_IntegerConversions(t *testing.T) {
	c := NewConverter()

	result, err := c.ConvertHex("7F")
	if err != nil {
		t.Fatalf("ConvertHex(7F) error: %v", err)
	}
	if result.Int8BE == nil || *result.Int8BE != 127 {
		t.Errorf("Expected Int8BE=127, got %v", result.Int8BE)
	}
	if result.Uint8BE == nil || *result.Uint8BE != 127 {
		t.Errorf("Expected Uint8BE=127, got %v", result.Uint8BE)
	}

	result, err = c.ConvertHex("0100")
	if err != nil {
		t.Fatalf("ConvertHex(0100) error: %v", err)
	}
	if result.Int16BE == nil || *result.Int16BE != 256 {
		t.Errorf("Expected Int16BE=256, got %v", result.Int16BE)
	}

	result, err = c.ConvertHex("00000100")
	if err != nil {
		t.Fatalf("ConvertHex(00000100) error: %v", err)
	}
	if result.Int32BE == nil || *result.Int32BE != 256 {
		t.Errorf("Expected Int32BE=256, got %v", result.Int32BE)
	}
}

func TestConvertHex_ASCII(t *testing.T) {
	c := NewConverter()
	result, err := c.ConvertHex("4869")
	if err != nil {
		t.Fatalf("ConvertHex(4869) error: %v", err)
	}
	if result.ASCII != "Hi" {
		t.Errorf("Expected ASCII='Hi', got '%s'", result.ASCII)
	}
}

func TestConvertInt_EmptyInput(t *testing.T) {
	c := NewConverter()
	_, err := c.ConvertInt("", "int8")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestConvertInt_ValidTypes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		intType string
		wantErr bool
	}{
		{"int8 positive", "127", "int8", false},
		{"int8 negative", "-128", "int8", false},
		{"int16 positive", "32767", "int16", false},
		{"int32 positive", "2147483647", "int32", false},
		{"uint8", "255", "uint8", false},
		{"uint16", "65535", "uint16", false},
		{"uint32", "4294967295", "uint32", false},
		{"invalid type", "123", "int128", true},
		{"invalid value", "abc", "int8", true},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertInt(tt.input, tt.intType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertInt(%q, %q) error = %v, wantErr %v", tt.input, tt.intType, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("ConvertInt(%q, %q) returned nil result", tt.input, tt.intType)
			}
		})
	}
}

func TestConvertBinary_EmptyInput(t *testing.T) {
	c := NewConverter()
	_, err := c.ConvertBinary("")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestConvertBinary_ValidInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"8 bits", "11111111", false},
		{"16 bits", "1111111100000000", false},
		{"with spaces", "1111 1111 0000 0000", false},
		{"invalid chars", "12345678", true},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertBinary(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertBinary(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("ConvertBinary(%q) returned nil result", tt.input)
			}
		})
	}
}

func TestConvertFloat_EmptyInput(t *testing.T) {
	c := NewConverter()
	_, err := c.ConvertFloat("", "float32")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestConvertFloat_ValidInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		floatType string
		wantErr   bool
	}{
		{"float32 positive", "3.14159", "float32", false},
		{"float32 negative", "-3.14159", "float32", false},
		{"float64 positive", "3.14159265358979", "float64", false},
		{"invalid type", "3.14", "float128", true},
		{"invalid value", "abc", "float32", true},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertFloat(tt.input, tt.floatType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertFloat(%q, %q) error = %v, wantErr %v", tt.input, tt.floatType, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result == nil {
				t.Errorf("ConvertFloat(%q, %q) returned nil result", tt.input, tt.floatType)
			}
		})
	}
}

func TestConvertModbusRegisters_EmptyInput(t *testing.T) {
	c := NewConverter()
	_, err := c.ConvertModbusRegisters("")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestConvertModbusRegisters_ValidInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		regCnt  int
	}{
		{"single hex", "1234", false, 1},
		{"two hex", "1234 5678", false, 2},
		{"with 0x prefix", "0x1234 0x5678", false, 2},
		{"comma separated", "1234,5678", false, 2},
		{"decimal with d prefix", "d1000 d2000", false, 2},
		{"four registers", "0x1234 0x5678 0x9ABC 0xDEF0", false, 4},
		{"invalid hex", "GHIJ", true, 0},
		{"value too large", "FFFFF", true, 0},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertModbusRegisters(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertModbusRegisters(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result == nil {
					t.Errorf("ConvertModbusRegisters(%q) returned nil result", tt.input)
					return
				}
				if len(result.Registers) != tt.regCnt {
					t.Errorf("Expected %d registers, got %d", tt.regCnt, len(result.Registers))
				}
			}
		})
	}
}

func TestConvertModbusRegisters_Combined32(t *testing.T) {
	c := NewConverter()
	result, err := c.ConvertModbusRegisters("0001 0000")
	if err != nil {
		t.Fatalf("ConvertModbusRegisters error: %v", err)
	}
	if len(result.Combined32) != 1 {
		t.Fatalf("Expected 1 Combined32, got %d", len(result.Combined32))
	}
	if result.Combined32[0].Uint32BE != 65536 {
		t.Errorf("Expected Uint32BE=65536, got %d", result.Combined32[0].Uint32BE)
	}
}

func TestFormatFloat32(t *testing.T) {
	tests := []struct {
		val  float32
		want string
	}{
		{3.14, "3.14"},
		{-3.14, "-3.14"},
		{0, "0"},
	}
	for _, tt := range tests {
		got := formatFloat32(tt.val)
		if got != tt.want {
			t.Errorf("formatFloat32(%v) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestFormatFloat64(t *testing.T) {
	tests := []struct {
		val  float64
		want string
	}{
		{3.14, "3.14"},
		{-3.14, "-3.14"},
		{0, "0"},
	}
	for _, tt := range tests {
		got := formatFloat64(tt.val)
		if got != tt.want {
			t.Errorf("formatFloat64(%v) = %q, want %q", tt.val, got, tt.want)
		}
	}
}

func TestBytesToASCII(t *testing.T) {
	tests := []struct {
		bytes []byte
		want  string
	}{
		{[]byte("Hello"), "Hello"},
		{[]byte{0x48, 0x00, 0x69}, "H.i"},
		{[]byte{}, ""},
	}
	for _, tt := range tests {
		got := bytesToASCII(tt.bytes)
		if got != tt.want {
			t.Errorf("bytesToASCII(%v) = %q, want %q", tt.bytes, got, tt.want)
		}
	}
}

func TestParseModbusInput(t *testing.T) {
	tests := []struct {
		input   string
		want    []uint16
		wantErr bool
	}{
		{"1234", []uint16{0x1234}, false},
		{"1234 5678", []uint16{0x1234, 0x5678}, false},
		{"0x1234", []uint16{0x1234}, false},
		{"d1000", []uint16{1000}, false},
		{"GHIJ", nil, true},
	}
	for _, tt := range tests {
		got, err := parseModbusInput(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseModbusInput(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && len(got) != len(tt.want) {
			t.Errorf("parseModbusInput(%q) len = %d, want %d", tt.input, len(got), len(tt.want))
		}
	}
}
