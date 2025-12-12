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
