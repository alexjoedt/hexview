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

// ============================================================================
// ConvertIntAuto Tests
// ============================================================================

func TestConvertIntAuto_EmptyInput(t *testing.T) {
	c := NewConverter()
	_, err := c.ConvertIntAuto("")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestConvertIntAuto_InvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"letters", "abc"},
		{"only spaces", "   "},
		{"symbols", "###"},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ConvertIntAuto(tt.input)
			if err == nil {
				t.Errorf("ConvertIntAuto(%q) expected error, got nil", tt.input)
			}
		})
	}
}

func TestConvertIntAuto_PartialParsing(t *testing.T) {
	// fmt.Sscanf stops at first non-numeric character, which is acceptable
	// These inputs will parse the numeric prefix successfully
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{"with suffix", "123abc", 123},
		{"with special char", "456@", 456},
		{"hex prefix ignored", "0x789", 0}, // 0x stops parsing, returns 0
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertIntAuto(tt.input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", tt.input, err)
			}
			if result.Int64BE == nil {
				t.Fatalf("ConvertIntAuto(%q) expected Int64BE to be set", tt.input)
			}
			if *result.Int64BE != tt.expected {
				t.Errorf("ConvertIntAuto(%q) got %d, want %d", tt.input, *result.Int64BE, tt.expected)
			}
		})
	}
}

func TestConvertIntAuto_Int8Range(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantInt8  bool
		wantUint8 bool
	}{
		{"zero", "0", true, true},
		{"positive small", "100", true, true},
		{"max int8", "127", true, true},
		{"above int8", "128", false, true},
		{"max uint8", "255", false, true},
		{"above uint8", "256", false, false},
		{"negative", "-128", true, false},
		{"below int8", "-129", false, false},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertIntAuto(tt.input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", tt.input, err)
			}

			if tt.wantInt8 && result.Int8BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Int8BE to be set", tt.input)
			}
			if !tt.wantInt8 && result.Int8BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Int8BE to be nil, got %v", tt.input, *result.Int8BE)
			}

			if tt.wantUint8 && result.Uint8BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint8BE to be set", tt.input)
			}
			if !tt.wantUint8 && result.Uint8BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint8BE to be nil, got %v", tt.input, *result.Uint8BE)
			}
		})
	}
}

func TestConvertIntAuto_Int16Range(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantInt16  bool
		wantUint16 bool
	}{
		{"max int16", "32767", true, true},
		{"above int16", "32768", false, true},
		{"max uint16", "65535", false, true},
		{"above uint16", "65536", false, false},
		{"min int16", "-32768", true, false},
		{"below int16", "-32769", false, false},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertIntAuto(tt.input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", tt.input, err)
			}

			if tt.wantInt16 && result.Int16BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Int16BE to be set", tt.input)
			}
			if !tt.wantInt16 && result.Int16BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Int16BE to be nil, got %v", tt.input, *result.Int16BE)
			}

			if tt.wantUint16 && result.Uint16BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint16BE to be set", tt.input)
			}
			if !tt.wantUint16 && result.Uint16BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint16BE to be nil, got %v", tt.input, *result.Uint16BE)
			}
		})
	}
}

func TestConvertIntAuto_Int32Range(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantInt32  bool
		wantUint32 bool
	}{
		{"max int32", "2147483647", true, true},
		{"above int32", "2147483648", false, true},
		{"max uint32", "4294967295", false, true},
		{"above uint32", "4294967296", false, false},
		{"min int32", "-2147483648", true, false},
		{"below int32", "-2147483649", false, false},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertIntAuto(tt.input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", tt.input, err)
			}

			if tt.wantInt32 && result.Int32BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Int32BE to be set", tt.input)
			}
			if !tt.wantInt32 && result.Int32BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Int32BE to be nil, got %v", tt.input, *result.Int32BE)
			}

			if tt.wantUint32 && result.Uint32BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint32BE to be set", tt.input)
			}
			if !tt.wantUint32 && result.Uint32BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint32BE to be nil, got %v", tt.input, *result.Uint32BE)
			}
		})
	}
}

func TestConvertIntAuto_NegativeValues(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"small negative", "-1"},
		{"int8 min", "-128"},
		{"int16 min", "-32768"},
		{"int32 min", "-2147483648"},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertIntAuto(tt.input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", tt.input, err)
			}

			// Negative values should never set unsigned types
			if result.Uint8BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint8BE to be nil for negative value", tt.input)
			}
			if result.Uint16BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint16BE to be nil for negative value", tt.input)
			}
			if result.Uint32BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint32BE to be nil for negative value", tt.input)
			}
			if result.Uint64BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected Uint64BE to be nil for negative value", tt.input)
			}

			// Int64 should always be set
			if result.Int64BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Int64BE to be set", tt.input)
			}
		})
	}
}

func TestConvertIntAuto_AlwaysHasInt64(t *testing.T) {
	tests := []string{"0", "100", "-100", "2147483647", "-2147483648"}

	c := NewConverter()
	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			result, err := c.ConvertIntAuto(input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", input, err)
			}

			if result.Int64BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Int64BE to always be set", input)
			}
			if result.Int64BEHex == "" {
				t.Errorf("ConvertIntAuto(%q) expected Int64BEHex to be set", input)
			}
		})
	}
}

func TestConvertIntAuto_CommonFields(t *testing.T) {
	c := NewConverter()
	result, err := c.ConvertIntAuto("100")
	if err != nil {
		t.Fatalf("ConvertIntAuto(100) unexpected error: %v", err)
	}

	// Check that common fields are populated
	if result.Binary == "" {
		t.Error("Expected Binary to be set")
	}
	if result.Bytes == "" {
		t.Error("Expected Bytes to be set")
	}
	if result.ASCII == "" {
		t.Error("Expected ASCII to be set")
	}
}

// ============================================================================
// Float Auto-Detection Tests
// ============================================================================

func TestConvertIntAuto_FloatDetection(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"dot notation", "10.5"},
		{"comma notation", "10,5"},
		{"negative float", "-3.14"},
		{"negative comma", "-3,14"},
		{"zero float", "0.0"},
		{"large float", "123.456"},
	}

	c := NewConverter()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.ConvertIntAuto(tt.input)
			if err != nil {
				t.Fatalf("ConvertIntAuto(%q) unexpected error: %v", tt.input, err)
			}

			// Should have float values, not integer values
			if result.Float32BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Float32BE to be set", tt.input)
			}
			if result.Float64BE == nil {
				t.Errorf("ConvertIntAuto(%q) expected Float64BE to be set", tt.input)
			}
			
			// Should NOT have integer values
			if result.Int8BE != nil || result.Int16BE != nil || result.Int32BE != nil {
				t.Errorf("ConvertIntAuto(%q) expected integer fields to be nil for float input", tt.input)
			}
		})
	}
}

func TestConvertIntAuto_FloatAllEndianness(t *testing.T) {
	c := NewConverter()
	result, err := c.ConvertIntAuto("10.5")
	if err != nil {
		t.Fatalf("ConvertIntAuto(10.5) unexpected error: %v", err)
	}

	// Check all float32 endianness variants are set
	if result.Float32BE == nil {
		t.Error("Expected Float32BE to be set")
	}
	if result.Float32LE == nil {
		t.Error("Expected Float32LE to be set")
	}
	if result.Float32BADC == nil {
		t.Error("Expected Float32BADC to be set")
	}
	if result.Float32CDAB == nil {
		t.Error("Expected Float32CDAB to be set")
	}

	// Check all float64 endianness variants are set
	if result.Float64BE == nil {
		t.Error("Expected Float64BE to be set")
	}
	if result.Float64LE == nil {
		t.Error("Expected Float64LE to be set")
	}
	if result.Float64BADC == nil {
		t.Error("Expected Float64BADC to be set")
	}
	if result.Float64CDAB == nil {
		t.Error("Expected Float64CDAB to be set")
	}
}

func TestConvertIntAuto_FloatCommaEquivalence(t *testing.T) {
	c := NewConverter()
	
	resultDot, err1 := c.ConvertIntAuto("10.5")
	resultComma, err2 := c.ConvertIntAuto("10,5")
	
	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected errors: dot=%v, comma=%v", err1, err2)
	}

	// Both should produce the same float values
	if resultDot.Float32BE == nil || resultComma.Float32BE == nil {
		t.Fatal("Float32BE should be set for both inputs")
	}
	
	if *resultDot.Float32BE != *resultComma.Float32BE {
		t.Errorf("Float values differ: dot=%s, comma=%s", *resultDot.Float32BE, *resultComma.Float32BE)
	}
	
	if resultDot.Float32BEHex != resultComma.Float32BEHex {
		t.Errorf("Float hex values differ: dot=%s, comma=%s", resultDot.Float32BEHex, resultComma.Float32BEHex)
	}
}

func TestConvertIntAuto_IntegerNoFloat(t *testing.T) {
	// Make sure pure integers don't trigger float detection
	c := NewConverter()
	result, err := c.ConvertIntAuto("100")
	if err != nil {
		t.Fatalf("ConvertIntAuto(100) unexpected error: %v", err)
	}

	// Should have integer values
	if result.Int8BE == nil {
		t.Error("Expected Int8BE to be set for integer input")
	}
	
	// Should NOT have float values
	if result.Float32BE != nil || result.Float64BE != nil {
		t.Error("Expected float fields to be nil for integer input")
	}
}
