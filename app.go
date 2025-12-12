package main

import (
	"context"
	"fmt"
	"math"

	"hexview/convert"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ConversionResult holds all conversion outputs
type ConversionResult struct {
	// Signed Integers - Big Endian
	Int8BE     *int8  `json:"int8BE,omitempty"`
	Int8BEHex  string `json:"int8BEHex,omitempty"`
	Int16BE    *int16 `json:"int16BE,omitempty"`
	Int16BEHex string `json:"int16BEHex,omitempty"`
	Int32BE    *int32 `json:"int32BE,omitempty"`
	Int32BEHex string `json:"int32BEHex,omitempty"`
	Int64BE    *int64 `json:"int64BE,omitempty"`
	Int64BEHex string `json:"int64BEHex,omitempty"`

	// Signed Integers - Little Endian
	Int16LE    *int16 `json:"int16LE,omitempty"`
	Int16LEHex string `json:"int16LEHex,omitempty"`
	Int32LE    *int32 `json:"int32LE,omitempty"`
	Int32LEHex string `json:"int32LEHex,omitempty"`
	Int64LE    *int64 `json:"int64LE,omitempty"`
	Int64LEHex string `json:"int64LEHex,omitempty"`

	// Signed Integers - Mid-Big Endian (BADC)
	Int16BADC    *int16 `json:"int16BADC,omitempty"`
	Int16BADCHex string `json:"int16BADCHex,omitempty"`
	Int32BADC    *int32 `json:"int32BADC,omitempty"`
	Int32BADCHex string `json:"int32BADCHex,omitempty"`
	Int64BADC    *int64 `json:"int64BADC,omitempty"`
	Int64BADCHex string `json:"int64BADCHex,omitempty"`

	// Signed Integers - Mid-Little Endian (CDAB)
	Int16CDAB    *int16 `json:"int16CDAB,omitempty"`
	Int16CDABHex string `json:"int16CDABHex,omitempty"`
	Int32CDAB    *int32 `json:"int32CDAB,omitempty"`
	Int32CDABHex string `json:"int32CDABHex,omitempty"`
	Int64CDAB    *int64 `json:"int64CDAB,omitempty"`
	Int64CDABHex string `json:"int64CDABHex,omitempty"`

	// Unsigned Integers - Big Endian
	Uint8BE     *uint8  `json:"uint8BE,omitempty"`
	Uint8BEHex  string  `json:"uint8BEHex,omitempty"`
	Uint16BE    *uint16 `json:"uint16BE,omitempty"`
	Uint16BEHex string  `json:"uint16BEHex,omitempty"`
	Uint32BE    *uint32 `json:"uint32BE,omitempty"`
	Uint32BEHex string  `json:"uint32BEHex,omitempty"`
	Uint64BE    *uint64 `json:"uint64BE,omitempty"`
	Uint64BEHex string  `json:"uint64BEHex,omitempty"`

	// Unsigned Integers - Little Endian
	Uint16LE    *uint16 `json:"uint16LE,omitempty"`
	Uint16LEHex string  `json:"uint16LEHex,omitempty"`
	Uint32LE    *uint32 `json:"uint32LE,omitempty"`
	Uint32LEHex string  `json:"uint32LEHex,omitempty"`
	Uint64LE    *uint64 `json:"uint64LE,omitempty"`
	Uint64LEHex string  `json:"uint64LEHex,omitempty"`

	// Unsigned Integers - Mid-Big Endian (BADC)
	Uint16BADC    *uint16 `json:"uint16BADC,omitempty"`
	Uint16BADCHex string  `json:"uint16BADCHex,omitempty"`
	Uint32BADC    *uint32 `json:"uint32BADC,omitempty"`
	Uint32BADCHex string  `json:"uint32BADCHex,omitempty"`
	Uint64BADC    *uint64 `json:"uint64BADC,omitempty"`
	Uint64BADCHex string  `json:"uint64BADCHex,omitempty"`

	// Unsigned Integers - Mid-Little Endian (CDAB)
	Uint16CDAB    *uint16 `json:"uint16CDAB,omitempty"`
	Uint16CDABHex string  `json:"uint16CDABHex,omitempty"`
	Uint32CDAB    *uint32 `json:"uint32CDAB,omitempty"`
	Uint32CDABHex string  `json:"uint32CDABHex,omitempty"`
	Uint64CDAB    *uint64 `json:"uint64CDAB,omitempty"`
	Uint64CDABHex string  `json:"uint64CDABHex,omitempty"`

	// Floating Point (stored as strings to support NaN/Inf)
	Float32BE    *string `json:"float32BE,omitempty"`
	Float32BEHex string  `json:"float32BEHex,omitempty"`
	Float64BE    *string `json:"float64BE,omitempty"`
	Float64BEHex string  `json:"float64BEHex,omitempty"`
	Float32LE    *string `json:"float32LE,omitempty"`
	Float32LEHex string  `json:"float32LEHex,omitempty"`
	Float64LE    *string `json:"float64LE,omitempty"`
	Float64LEHex string  `json:"float64LEHex,omitempty"`

	// Floating Point - Mid-Big Endian (BADC)
	Float32BADC    *string `json:"float32BADC,omitempty"`
	Float32BADCHex string  `json:"float32BADCHex,omitempty"`
	Float64BADC    *string `json:"float64BADC,omitempty"`
	Float64BADCHex string  `json:"float64BADCHex,omitempty"`

	// Floating Point - Mid-Little Endian (CDAB)
	Float32CDAB    *string `json:"float32CDAB,omitempty"`
	Float32CDABHex string  `json:"float32CDABHex,omitempty"`
	Float64CDAB    *string `json:"float64CDAB,omitempty"`
	Float64CDABHex string  `json:"float64CDABHex,omitempty"`

	// Binary Representations
	Binary string `json:"binary,omitempty"`
	Bytes  string `json:"bytes,omitempty"`

	// ASCII representation (printable chars, '.' for non-printable)
	ASCII string `json:"ascii,omitempty"`
}

// ModbusRegister represents a single 16-bit Modbus register
type ModbusRegister struct {
	Index    int    `json:"index"`
	Hex      string `json:"hex"`
	Unsigned uint16 `json:"unsigned"`
	Signed   int16  `json:"signed"`
	Binary   string `json:"binary"`
}

// ModbusCombined32 represents a 32-bit value from two registers
type ModbusCombined32 struct {
	RegisterStart int     `json:"registerStart"`
	Hex           string  `json:"hex"`
	Uint32BE      uint32  `json:"uint32BE"`
	Uint32LE      uint32  `json:"uint32LE"`
	Uint32BADC    uint32  `json:"uint32BADC"`
	Uint32CDAB    uint32  `json:"uint32CDAB"`
	Int32BE       int32   `json:"int32BE"`
	Int32LE       int32   `json:"int32LE"`
	Int32BADC     int32   `json:"int32BADC"`
	Int32CDAB     int32   `json:"int32CDAB"`
	Float32BE     string  `json:"float32BE"`
	Float32LE     string  `json:"float32LE"`
	Float32BADC   string  `json:"float32BADC"`
	Float32CDAB   string  `json:"float32CDAB"`
}

// ModbusCombined64 represents a 64-bit value from four registers
type ModbusCombined64 struct {
	RegisterStart int     `json:"registerStart"`
	Hex           string  `json:"hex"`
	Uint64BE      uint64  `json:"uint64BE"`
	Uint64LE      uint64  `json:"uint64LE"`
	Int64BE       int64   `json:"int64BE"`
	Int64LE       int64   `json:"int64LE"`
	Float64BE     string  `json:"float64BE"`
	Float64LE     string  `json:"float64LE"`
}

// ModbusResult holds the conversion results for Modbus registers
type ModbusResult struct {
	Registers   []ModbusRegister   `json:"registers"`
	Combined32  []ModbusCombined32 `json:"combined32"`
	Combined64  []ModbusCombined64 `json:"combined64"`
	RawHex      string             `json:"rawHex"`
	ASCII       string             `json:"ascii"`
}

// formatFloat converts float values to strings, handling NaN and Inf
func formatFloat32(v float32) string {
	if math.IsNaN(float64(v)) {
		return "NaN"
	}
	if math.IsInf(float64(v), 1) {
		return "+Inf"
	}
	if math.IsInf(float64(v), -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%g", v)
}

func formatFloat64(v float64) string {
	if math.IsNaN(v) {
		return "NaN"
	}
	if math.IsInf(v, 1) {
		return "+Inf"
	}
	if math.IsInf(v, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%g", v)
}

// ConvertHex performs all possible conversions on the hex input
func (a *App) ConvertHex(hexInput string) (*ConversionResult, error) {
	if hexInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &ConversionResult{}

	// Convert to bytes first to get binary representation
	bytes, err := convert.HexToBytes(hexInput)
	if err != nil {
		return nil, fmt.Errorf("invalid hex input: %w", err)
	}

	result.Binary = convert.BytesToBinary(bytes)
	result.Bytes = convert.BytesToHex(bytes)

	// Try all signed integer conversions (Big Endian)
	if v, err := convert.HexToInt8(hexInput); err == nil {
		result.Int8BE = &v
		result.Int8BEHex = convert.Int8ToHex(v)
	}
	if v, err := convert.HexToInt16(hexInput); err == nil {
		result.Int16BE = &v
		result.Int16BEHex = convert.Int16ToHex(v)
	}
	if v, err := convert.HexToInt32(hexInput); err == nil {
		result.Int32BE = &v
		result.Int32BEHex = convert.Int32ToHex(v)
	}
	if v, err := convert.HexToInt64(hexInput); err == nil {
		result.Int64BE = &v
		result.Int64BEHex = convert.Int64ToHex(v)
	}

	// Try all signed integer conversions (Little Endian)
	if v, err := convert.HexToInt16LE(hexInput); err == nil {
		result.Int16LE = &v
		result.Int16LEHex = convert.Int16ToHexLE(v)
	}
	if v, err := convert.HexToInt32LE(hexInput); err == nil {
		result.Int32LE = &v
		result.Int32LEHex = convert.Int32ToHexLE(v)
	}
	if v, err := convert.HexToInt64LE(hexInput); err == nil {
		result.Int64LE = &v
		result.Int64LEHex = convert.Int64ToHexLE(v)
	}

	// Try all signed integer conversions (Mid-Big Endian / BADC)
	if v, err := convert.HexToInt16BADC(hexInput); err == nil {
		result.Int16BADC = &v
		result.Int16BADCHex = convert.Int16ToHexBADC(v)
	}
	if v, err := convert.HexToInt32BADC(hexInput); err == nil {
		result.Int32BADC = &v
		result.Int32BADCHex = convert.Int32ToHexBADC(v)
	}
	if v, err := convert.HexToInt64BADC(hexInput); err == nil {
		result.Int64BADC = &v
		result.Int64BADCHex = convert.Int64ToHexBADC(v)
	}

	// Try all signed integer conversions (Mid-Little Endian / CDAB)
	if v, err := convert.HexToInt16CDAB(hexInput); err == nil {
		result.Int16CDAB = &v
		result.Int16CDABHex = convert.Int16ToHexCDAB(v)
	}
	if v, err := convert.HexToInt32CDAB(hexInput); err == nil {
		result.Int32CDAB = &v
		result.Int32CDABHex = convert.Int32ToHexCDAB(v)
	}
	if v, err := convert.HexToInt64CDAB(hexInput); err == nil {
		result.Int64CDAB = &v
		result.Int64CDABHex = convert.Int64ToHexCDAB(v)
	}

	// Try all unsigned integer conversions (Big Endian)
	if v, err := convert.HexToUint8(hexInput); err == nil {
		result.Uint8BE = &v
		result.Uint8BEHex = convert.Uint8ToHex(v)
	}
	if v, err := convert.HexToUint16(hexInput); err == nil {
		result.Uint16BE = &v
		result.Uint16BEHex = convert.Uint16ToHex(v)
	}
	if v, err := convert.HexToUint32(hexInput); err == nil {
		result.Uint32BE = &v
		result.Uint32BEHex = convert.Uint32ToHex(v)
	}
	if v, err := convert.HexToUint64(hexInput); err == nil {
		result.Uint64BE = &v
		result.Uint64BEHex = convert.Uint64ToHex(v)
	}

	// Try all unsigned integer conversions (Little Endian)
	if v, err := convert.HexToUint16LE(hexInput); err == nil {
		result.Uint16LE = &v
		result.Uint16LEHex = convert.Uint16ToHexLE(v)
	}
	if v, err := convert.HexToUint32LE(hexInput); err == nil {
		result.Uint32LE = &v
		result.Uint32LEHex = convert.Uint32ToHexLE(v)
	}
	if v, err := convert.HexToUint64LE(hexInput); err == nil {
		result.Uint64LE = &v
		result.Uint64LEHex = convert.Uint64ToHexLE(v)
	}

	// Try all unsigned integer conversions (Mid-Big Endian / BADC)
	if v, err := convert.HexToUint16BADC(hexInput); err == nil {
		result.Uint16BADC = &v
		result.Uint16BADCHex = convert.Uint16ToHexBADC(v)
	}
	if v, err := convert.HexToUint32BADC(hexInput); err == nil {
		result.Uint32BADC = &v
		result.Uint32BADCHex = convert.Uint32ToHexBADC(v)
	}
	if v, err := convert.HexToUint64BADC(hexInput); err == nil {
		result.Uint64BADC = &v
		result.Uint64BADCHex = convert.Uint64ToHexBADC(v)
	}

	// Try all unsigned integer conversions (Mid-Little Endian / CDAB)
	if v, err := convert.HexToUint16CDAB(hexInput); err == nil {
		result.Uint16CDAB = &v
		result.Uint16CDABHex = convert.Uint16ToHexCDAB(v)
	}
	if v, err := convert.HexToUint32CDAB(hexInput); err == nil {
		result.Uint32CDAB = &v
		result.Uint32CDABHex = convert.Uint32ToHexCDAB(v)
	}
	if v, err := convert.HexToUint64CDAB(hexInput); err == nil {
		result.Uint64CDAB = &v
		result.Uint64CDABHex = convert.Uint64ToHexCDAB(v)
	}

	// Try float conversions (Big Endian)
	if v, err := convert.HexToFloat32(hexInput); err == nil {
		formatted := formatFloat32(v)
		result.Float32BE = &formatted
		result.Float32BEHex = convert.Float32ToHex(v)
	}
	if v, err := convert.HexToFloat64(hexInput); err == nil {
		formatted := formatFloat64(v)
		result.Float64BE = &formatted
		result.Float64BEHex = convert.Float64ToHex(v)
	}

	// Try float conversions (Little Endian)
	if v, err := convert.HexToFloat32LE(hexInput); err == nil {
		formatted := formatFloat32(v)
		result.Float32LE = &formatted
		result.Float32LEHex = convert.Float32ToHexLE(v)
	}
	if v, err := convert.HexToFloat64LE(hexInput); err == nil {
		formatted := formatFloat64(v)
		result.Float64LE = &formatted
		result.Float64LEHex = convert.Float64ToHexLE(v)
	}

	// Try float conversions (Mid-Big Endian / BADC)
	if v, err := convert.HexToFloat32BADC(hexInput); err == nil {
		formatted := formatFloat32(v)
		result.Float32BADC = &formatted
		result.Float32BADCHex = convert.Float32ToHexBADC(v)
	}
	if v, err := convert.HexToFloat64BADC(hexInput); err == nil {
		formatted := formatFloat64(v)
		result.Float64BADC = &formatted
		result.Float64BADCHex = convert.Float64ToHexBADC(v)
	}

	// Try float conversions (Mid-Little Endian / CDAB)
	if v, err := convert.HexToFloat32CDAB(hexInput); err == nil {
		formatted := formatFloat32(v)
		result.Float32CDAB = &formatted
		result.Float32CDABHex = convert.Float32ToHexCDAB(v)
	}
	if v, err := convert.HexToFloat64CDAB(hexInput); err == nil {
		formatted := formatFloat64(v)
		result.Float64CDAB = &formatted
		result.Float64CDABHex = convert.Float64ToHexCDAB(v)
	}

	return result, nil
}

// ConvertInt performs conversions from integer input to hex and binary
func (a *App) ConvertInt(intInput string, intType string) (*ConversionResult, error) {
	if intInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &ConversionResult{}

	// Parse the integer based on selected type
	switch intType {
	case "int8":
		var val int8
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid int8 value: %w", err)
		}
		hexStr := convert.Int8ToHex(val)
		bytes, _ := convert.HexToBytes(hexStr)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStr
		result.Int8BE = &val
		result.Int8BEHex = hexStr
		return result, nil

	case "int16":
		var val int16
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid int16 value: %w", err)
		}
		hexStrBE := convert.Int16ToHex(val)
		hexStrLE := convert.Int16ToHexLE(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.Int16BE = &val
		result.Int16BEHex = hexStrBE
		// Also populate LE conversion
		if vLE, err := convert.HexToInt16LE(hexStrLE); err == nil {
			result.Int16LE = &vLE
			result.Int16LEHex = hexStrLE
		}
		return result, nil

	case "int32":
		var val int32
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid int32 value: %w", err)
		}
		hexStrBE := convert.Int32ToHex(val)
		hexStrLE := convert.Int32ToHexLE(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.Int32BE = &val
		result.Int32BEHex = hexStrBE
		// Also populate LE conversion
		if vLE, err := convert.HexToInt32LE(hexStrLE); err == nil {
			result.Int32LE = &vLE
			result.Int32LEHex = hexStrLE
		}
		return result, nil

	case "int64":
		var val int64
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid int64 value: %w", err)
		}
		hexStrBE := convert.Int64ToHex(val)
		hexStrLE := convert.Int64ToHexLE(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.Int64BE = &val
		result.Int64BEHex = hexStrBE
		// Also populate LE conversion
		if vLE, err := convert.HexToInt64LE(hexStrLE); err == nil {
			result.Int64LE = &vLE
			result.Int64LEHex = hexStrLE
		}
		return result, nil

	case "uint8":
		var val uint8
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid uint8 value: %w", err)
		}
		hexStr := convert.Uint8ToHex(val)
		bytes, _ := convert.HexToBytes(hexStr)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStr
		result.Uint8BE = &val
		result.Uint8BEHex = hexStr
		return result, nil

	case "uint16":
		var val uint16
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid uint16 value: %w", err)
		}
		hexStrBE := convert.Uint16ToHex(val)
		hexStrLE := convert.Uint16ToHexLE(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.Uint16BE = &val
		result.Uint16BEHex = hexStrBE
		// Also populate LE conversion
		if vLE, err := convert.HexToUint16LE(hexStrLE); err == nil {
			result.Uint16LE = &vLE
			result.Uint16LEHex = hexStrLE
		}
		return result, nil

	case "uint32":
		var val uint32
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid uint32 value: %w", err)
		}
		hexStrBE := convert.Uint32ToHex(val)
		hexStrLE := convert.Uint32ToHexLE(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.Uint32BE = &val
		result.Uint32BEHex = hexStrBE
		// Also populate LE conversion
		if vLE, err := convert.HexToUint32LE(hexStrLE); err == nil {
			result.Uint32LE = &vLE
			result.Uint32LEHex = hexStrLE
		}
		return result, nil

	case "uint64":
		var val uint64
		_, err := fmt.Sscanf(intInput, "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid uint64 value: %w", err)
		}
		hexStrBE := convert.Uint64ToHex(val)
		hexStrLE := convert.Uint64ToHexLE(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.Uint64BE = &val
		result.Uint64BEHex = hexStrBE
		// Also populate LE conversion
		if vLE, err := convert.HexToUint64LE(hexStrLE); err == nil {
			result.Uint64LE = &vLE
			result.Uint64LEHex = hexStrLE
		}
		return result, nil

	default:
		return nil, fmt.Errorf("unsupported integer type: %s", intType)
	}
}

// ConvertBinary performs all possible conversions on the binary input
func (a *App) ConvertBinary(binaryInput string) (*ConversionResult, error) {
	if binaryInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &ConversionResult{}

	// Convert to bytes first to get hex representation
	bytes, err := convert.ParseBinary(binaryInput)
	if err != nil {
		return nil, fmt.Errorf("invalid binary input: %w", err)
	}

	result.Binary = convert.BytesToBinary(bytes)
	result.Bytes = convert.BytesToHex(bytes)

	// Get the hex representation for further conversions
	hexStr := convert.BytesToHex(bytes)

	// Try all signed integer conversions (Big Endian)
	if v, err := convert.HexToInt8(hexStr); err == nil {
		result.Int8BE = &v
		result.Int8BEHex = convert.Int8ToHex(v)
	}
	if v, err := convert.HexToInt16(hexStr); err == nil {
		result.Int16BE = &v
		result.Int16BEHex = convert.Int16ToHex(v)
	}
	if v, err := convert.HexToInt32(hexStr); err == nil {
		result.Int32BE = &v
		result.Int32BEHex = convert.Int32ToHex(v)
	}
	if v, err := convert.HexToInt64(hexStr); err == nil {
		result.Int64BE = &v
		result.Int64BEHex = convert.Int64ToHex(v)
	}

	// Try all signed integer conversions (Little Endian)
	if v, err := convert.HexToInt16LE(hexStr); err == nil {
		result.Int16LE = &v
		result.Int16LEHex = convert.Int16ToHexLE(v)
	}
	if v, err := convert.HexToInt32LE(hexStr); err == nil {
		result.Int32LE = &v
		result.Int32LEHex = convert.Int32ToHexLE(v)
	}
	if v, err := convert.HexToInt64LE(hexStr); err == nil {
		result.Int64LE = &v
		result.Int64LEHex = convert.Int64ToHexLE(v)
	}

	// Try all signed integer conversions (Mid-Big Endian / BADC)
	if v, err := convert.HexToInt16BADC(hexStr); err == nil {
		result.Int16BADC = &v
		result.Int16BADCHex = convert.Int16ToHexBADC(v)
	}
	if v, err := convert.HexToInt32BADC(hexStr); err == nil {
		result.Int32BADC = &v
		result.Int32BADCHex = convert.Int32ToHexBADC(v)
	}
	if v, err := convert.HexToInt64BADC(hexStr); err == nil {
		result.Int64BADC = &v
		result.Int64BADCHex = convert.Int64ToHexBADC(v)
	}

	// Try all signed integer conversions (Mid-Little Endian / CDAB)
	if v, err := convert.HexToInt16CDAB(hexStr); err == nil {
		result.Int16CDAB = &v
		result.Int16CDABHex = convert.Int16ToHexCDAB(v)
	}
	if v, err := convert.HexToInt32CDAB(hexStr); err == nil {
		result.Int32CDAB = &v
		result.Int32CDABHex = convert.Int32ToHexCDAB(v)
	}
	if v, err := convert.HexToInt64CDAB(hexStr); err == nil {
		result.Int64CDAB = &v
		result.Int64CDABHex = convert.Int64ToHexCDAB(v)
	}

	// Try all unsigned integer conversions (Big Endian)
	if v, err := convert.HexToUint8(hexStr); err == nil {
		result.Uint8BE = &v
		result.Uint8BEHex = convert.Uint8ToHex(v)
	}
	if v, err := convert.HexToUint16(hexStr); err == nil {
		result.Uint16BE = &v
		result.Uint16BEHex = convert.Uint16ToHex(v)
	}
	if v, err := convert.HexToUint32(hexStr); err == nil {
		result.Uint32BE = &v
		result.Uint32BEHex = convert.Uint32ToHex(v)
	}
	if v, err := convert.HexToUint64(hexStr); err == nil {
		result.Uint64BE = &v
		result.Uint64BEHex = convert.Uint64ToHex(v)
	}

	// Try all unsigned integer conversions (Little Endian)
	if v, err := convert.HexToUint16LE(hexStr); err == nil {
		result.Uint16LE = &v
		result.Uint16LEHex = convert.Uint16ToHexLE(v)
	}
	if v, err := convert.HexToUint32LE(hexStr); err == nil {
		result.Uint32LE = &v
		result.Uint32LEHex = convert.Uint32ToHexLE(v)
	}
	if v, err := convert.HexToUint64LE(hexStr); err == nil {
		result.Uint64LE = &v
		result.Uint64LEHex = convert.Uint64ToHexLE(v)
	}

	// Try all unsigned integer conversions (Mid-Big Endian / BADC)
	if v, err := convert.HexToUint16BADC(hexStr); err == nil {
		result.Uint16BADC = &v
		result.Uint16BADCHex = convert.Uint16ToHexBADC(v)
	}
	if v, err := convert.HexToUint32BADC(hexStr); err == nil {
		result.Uint32BADC = &v
		result.Uint32BADCHex = convert.Uint32ToHexBADC(v)
	}
	if v, err := convert.HexToUint64BADC(hexStr); err == nil {
		result.Uint64BADC = &v
		result.Uint64BADCHex = convert.Uint64ToHexBADC(v)
	}

	// Try all unsigned integer conversions (Mid-Little Endian / CDAB)
	if v, err := convert.HexToUint16CDAB(hexStr); err == nil {
		result.Uint16CDAB = &v
		result.Uint16CDABHex = convert.Uint16ToHexCDAB(v)
	}
	if v, err := convert.HexToUint32CDAB(hexStr); err == nil {
		result.Uint32CDAB = &v
		result.Uint32CDABHex = convert.Uint32ToHexCDAB(v)
	}
	if v, err := convert.HexToUint64CDAB(hexStr); err == nil {
		result.Uint64CDAB = &v
		result.Uint64CDABHex = convert.Uint64ToHexCDAB(v)
	}

	// Try float conversions (Big Endian)
	if v, err := convert.HexToFloat32(hexStr); err == nil {
		formatted := formatFloat32(v)
		result.Float32BE = &formatted
		result.Float32BEHex = convert.Float32ToHex(v)
	}
	if v, err := convert.HexToFloat64(hexStr); err == nil {
		formatted := formatFloat64(v)
		result.Float64BE = &formatted
		result.Float64BEHex = convert.Float64ToHex(v)
	}

	// Try float conversions (Little Endian)
	if v, err := convert.HexToFloat32LE(hexStr); err == nil {
		formatted := formatFloat32(v)
		result.Float32LE = &formatted
		result.Float32LEHex = convert.Float32ToHexLE(v)
	}
	if v, err := convert.HexToFloat64LE(hexStr); err == nil {
		formatted := formatFloat64(v)
		result.Float64LE = &formatted
		result.Float64LEHex = convert.Float64ToHexLE(v)
	}

	// Try float conversions (Mid-Big Endian / BADC)
	if v, err := convert.HexToFloat32BADC(hexStr); err == nil {
		formatted := formatFloat32(v)
		result.Float32BADC = &formatted
		result.Float32BADCHex = convert.Float32ToHexBADC(v)
	}
	if v, err := convert.HexToFloat64BADC(hexStr); err == nil {
		formatted := formatFloat64(v)
		result.Float64BADC = &formatted
		result.Float64BADCHex = convert.Float64ToHexBADC(v)
	}

	// Try float conversions (Mid-Little Endian / CDAB)
	if v, err := convert.HexToFloat32CDAB(hexStr); err == nil {
		formatted := formatFloat32(v)
		result.Float32CDAB = &formatted
		result.Float32CDABHex = convert.Float32ToHexCDAB(v)
	}
	if v, err := convert.HexToFloat64CDAB(hexStr); err == nil {
		formatted := formatFloat64(v)
		result.Float64CDAB = &formatted
		result.Float64CDABHex = convert.Float64ToHexCDAB(v)
	}

	return result, nil
}

// ConvertModbusRegisters converts an array of 16-bit register values
// Input can be space/comma separated hex values (e.g., "1234 5678" or "0x1234, 0x5678")
// or decimal values with 'd' prefix (e.g., "d1000 d2000")
func (a *App) ConvertModbusRegisters(input string) (*ModbusResult, error) {
	if input == "" {
		return nil, fmt.Errorf("empty input")
	}

	// Parse the input into individual register values
	registers, err := parseModbusInput(input)
	if err != nil {
		return nil, err
	}

	if len(registers) == 0 {
		return nil, fmt.Errorf("no valid register values found")
	}

	result := &ModbusResult{
		Registers:  make([]ModbusRegister, len(registers)),
		Combined32: make([]ModbusCombined32, 0),
		Combined64: make([]ModbusCombined64, 0),
	}

	// Build raw hex and bytes for ASCII
	var allBytes []byte
	var hexParts []string

	for i, val := range registers {
		regHex := convert.Uint16ToHex(val)
		hexParts = append(hexParts, regHex)

		// Get bytes for ASCII
		regBytes, _ := convert.HexToBytes(regHex)
		allBytes = append(allBytes, regBytes...)

		result.Registers[i] = ModbusRegister{
			Index:    i + 1,
			Hex:      regHex,
			Unsigned: val,
			Signed:   int16(val),
			Binary:   convert.Uint16ToBinary(val),
		}
	}

	result.RawHex = joinStrings(hexParts, " ")
	result.ASCII = bytesToASCII(allBytes)

	// Generate 32-bit combinations (overlapping pairs)
	for i := 0; i <= len(registers)-2; i++ {
		hexStr := convert.Uint16ToHex(registers[i]) + convert.Uint16ToHex(registers[i+1])

		combined := ModbusCombined32{
			RegisterStart: i + 1,
			Hex:           hexStr,
		}

		// Unsigned integers
		if v, err := convert.HexToUint32(hexStr); err == nil {
			combined.Uint32BE = v
		}
		if v, err := convert.HexToUint32LE(hexStr); err == nil {
			combined.Uint32LE = v
		}
		if v, err := convert.HexToUint32BADC(hexStr); err == nil {
			combined.Uint32BADC = v
		}
		if v, err := convert.HexToUint32CDAB(hexStr); err == nil {
			combined.Uint32CDAB = v
		}

		// Signed integers
		if v, err := convert.HexToInt32(hexStr); err == nil {
			combined.Int32BE = v
		}
		if v, err := convert.HexToInt32LE(hexStr); err == nil {
			combined.Int32LE = v
		}
		if v, err := convert.HexToInt32BADC(hexStr); err == nil {
			combined.Int32BADC = v
		}
		if v, err := convert.HexToInt32CDAB(hexStr); err == nil {
			combined.Int32CDAB = v
		}

		// Floats
		if v, err := convert.HexToFloat32(hexStr); err == nil {
			combined.Float32BE = formatFloat32(v)
		}
		if v, err := convert.HexToFloat32LE(hexStr); err == nil {
			combined.Float32LE = formatFloat32(v)
		}
		if v, err := convert.HexToFloat32BADC(hexStr); err == nil {
			combined.Float32BADC = formatFloat32(v)
		}
		if v, err := convert.HexToFloat32CDAB(hexStr); err == nil {
			combined.Float32CDAB = formatFloat32(v)
		}

		result.Combined32 = append(result.Combined32, combined)
	}

	// Generate 64-bit combinations (overlapping quads)
	for i := 0; i <= len(registers)-4; i++ {
		hexStr := convert.Uint16ToHex(registers[i]) +
			convert.Uint16ToHex(registers[i+1]) +
			convert.Uint16ToHex(registers[i+2]) +
			convert.Uint16ToHex(registers[i+3])

		combined := ModbusCombined64{
			RegisterStart: i + 1,
			Hex:           hexStr,
		}

		// Unsigned integers
		if v, err := convert.HexToUint64(hexStr); err == nil {
			combined.Uint64BE = v
		}
		if v, err := convert.HexToUint64LE(hexStr); err == nil {
			combined.Uint64LE = v
		}

		// Signed integers
		if v, err := convert.HexToInt64(hexStr); err == nil {
			combined.Int64BE = v
		}
		if v, err := convert.HexToInt64LE(hexStr); err == nil {
			combined.Int64LE = v
		}

		// Floats
		if v, err := convert.HexToFloat64(hexStr); err == nil {
			combined.Float64BE = formatFloat64(v)
		}
		if v, err := convert.HexToFloat64LE(hexStr); err == nil {
			combined.Float64LE = formatFloat64(v)
		}

		result.Combined64 = append(result.Combined64, combined)
	}

	return result, nil
}

// parseModbusInput parses space/comma separated register values
// Supports hex (0x1234, 1234) and decimal with 'd' prefix (d1000)
func parseModbusInput(input string) ([]uint16, error) {
	// Replace common separators with spaces
	normalized := input
	for _, sep := range []string{",", ";", "\t", "\n", ":"} {
		normalized = replaceAll(normalized, sep, " ")
	}

	// Split by spaces
	parts := splitAndTrim(normalized)

	registers := make([]uint16, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		var val uint64
		var err error

		// Check for decimal prefix
		if len(part) > 1 && (part[0] == 'd' || part[0] == 'D') {
			_, err = fmt.Sscanf(part[1:], "%d", &val)
			if err != nil {
				return nil, fmt.Errorf("invalid decimal value: %s", part)
			}
		} else {
			// Try hex (with or without 0x prefix)
			cleanHex := part
			if len(part) > 2 && (part[:2] == "0x" || part[:2] == "0X") {
				cleanHex = part[2:]
			}
			_, err = fmt.Sscanf(cleanHex, "%x", &val)
			if err != nil {
				return nil, fmt.Errorf("invalid hex value: %s", part)
			}
		}

		if val > 0xFFFF {
			return nil, fmt.Errorf("value exceeds 16-bit range: %s", part)
		}

		registers = append(registers, uint16(val))
	}

	return registers, nil
}

// Helper functions
func replaceAll(s, old, new string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		found := false
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			result += new
			i += len(old) - 1
			found = true
		}
		if !found {
			result += string(s[i])
		}
	}
	return result
}

func splitAndTrim(s string) []string {
	var result []string
	current := ""
	for _, c := range s {
		if c == ' ' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func joinStrings(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += sep + parts[i]
	}
	return result
}

func bytesToASCII(bytes []byte) string {
	result := ""
	for _, b := range bytes {
		if b >= 32 && b <= 126 {
			result += string(b)
		} else {
			result += "."
		}
	}
	return result
}

// ConvertFloat performs conversions from decimal float input to hex and binary
func (a *App) ConvertFloat(floatInput string, floatType string) (*ConversionResult, error) {
	if floatInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &ConversionResult{}

	switch floatType {
	case "float32":
		var val float32
		_, err := fmt.Sscanf(floatInput, "%f", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid float32 value: %w", err)
		}
		hexStrBE := convert.Float32ToHex(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE

		// Set all float32 representations
		formatted := formatFloat32(val)
		result.Float32BE = &formatted
		result.Float32BEHex = hexStrBE

		// Populate other endianness formats
		hexStrLE := convert.Float32ToHexLE(val)
		if vLE, err := convert.HexToFloat32LE(hexStrLE); err == nil {
			fmtLE := formatFloat32(vLE)
			result.Float32LE = &fmtLE
			result.Float32LEHex = hexStrLE
		}
		hexStrBADC := convert.Float32ToHexBADC(val)
		if vBADC, err := convert.HexToFloat32BADC(hexStrBADC); err == nil {
			fmtBADC := formatFloat32(vBADC)
			result.Float32BADC = &fmtBADC
			result.Float32BADCHex = hexStrBADC
		}
		hexStrCDAB := convert.Float32ToHexCDAB(val)
		if vCDAB, err := convert.HexToFloat32CDAB(hexStrCDAB); err == nil {
			fmtCDAB := formatFloat32(vCDAB)
			result.Float32CDAB = &fmtCDAB
			result.Float32CDABHex = hexStrCDAB
		}

		// Also show as integers (reinterpret bits)
		if v, err := convert.HexToUint32(hexStrBE); err == nil {
			result.Uint32BE = &v
			result.Uint32BEHex = hexStrBE
		}
		if v, err := convert.HexToInt32(hexStrBE); err == nil {
			result.Int32BE = &v
			result.Int32BEHex = hexStrBE
		}

		return result, nil

	case "float64":
		var val float64
		_, err := fmt.Sscanf(floatInput, "%lf", &val)
		if err != nil {
			// Try alternative parsing
			_, err = fmt.Sscanf(floatInput, "%f", &val)
			if err != nil {
				return nil, fmt.Errorf("invalid float64 value: %w", err)
			}
		}
		hexStrBE := convert.Float64ToHex(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE

		// Set all float64 representations
		formatted := formatFloat64(val)
		result.Float64BE = &formatted
		result.Float64BEHex = hexStrBE

		// Populate other endianness formats
		hexStrLE := convert.Float64ToHexLE(val)
		if vLE, err := convert.HexToFloat64LE(hexStrLE); err == nil {
			fmtLE := formatFloat64(vLE)
			result.Float64LE = &fmtLE
			result.Float64LEHex = hexStrLE
		}
		hexStrBADC := convert.Float64ToHexBADC(val)
		if vBADC, err := convert.HexToFloat64BADC(hexStrBADC); err == nil {
			fmtBADC := formatFloat64(vBADC)
			result.Float64BADC = &fmtBADC
			result.Float64BADCHex = hexStrBADC
		}
		hexStrCDAB := convert.Float64ToHexCDAB(val)
		if vCDAB, err := convert.HexToFloat64CDAB(hexStrCDAB); err == nil {
			fmtCDAB := formatFloat64(vCDAB)
			result.Float64CDAB = &fmtCDAB
			result.Float64CDABHex = hexStrCDAB
		}

		// Also show as integers (reinterpret bits)
		if v, err := convert.HexToUint64(hexStrBE); err == nil {
			result.Uint64BE = &v
			result.Uint64BEHex = hexStrBE
		}
		if v, err := convert.HexToInt64(hexStrBE); err == nil {
			result.Int64BE = &v
			result.Int64BEHex = hexStrBE
		}

		return result, nil

	default:
		return nil, fmt.Errorf("unsupported float type: %s", floatType)
	}
}
