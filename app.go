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
	Int8BE  *int8  `json:"int8BE,omitempty"`
	Int16BE *int16 `json:"int16BE,omitempty"`
	Int32BE *int32 `json:"int32BE,omitempty"`
	Int64BE *int64 `json:"int64BE,omitempty"`

	// Signed Integers - Little Endian
	Int16LE *int16 `json:"int16LE,omitempty"`
	Int32LE *int32 `json:"int32LE,omitempty"`
	Int64LE *int64 `json:"int64LE,omitempty"`

	// Unsigned Integers - Big Endian
	Uint8BE  *uint8  `json:"uint8BE,omitempty"`
	Uint16BE *uint16 `json:"uint16BE,omitempty"`
	Uint32BE *uint32 `json:"uint32BE,omitempty"`
	Uint64BE *uint64 `json:"uint64BE,omitempty"`

	// Unsigned Integers - Little Endian
	Uint16LE *uint16 `json:"uint16LE,omitempty"`
	Uint32LE *uint32 `json:"uint32LE,omitempty"`
	Uint64LE *uint64 `json:"uint64LE,omitempty"`

	// Floating Point (stored as strings to support NaN/Inf)
	Float32BE *string `json:"float32BE,omitempty"`
	Float64BE *string `json:"float64BE,omitempty"`
	Float32LE *string `json:"float32LE,omitempty"`
	Float64LE *string `json:"float64LE,omitempty"`

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
	}
	if v, err := convert.HexToInt16(hexInput); err == nil {
		result.Int16BE = &v
	}
	if v, err := convert.HexToInt32(hexInput); err == nil {
		result.Int32BE = &v
	}
	if v, err := convert.HexToInt64(hexInput); err == nil {
		result.Int64BE = &v
	}

	// Try all signed integer conversions (Little Endian)
	if v, err := convert.HexToInt16LE(hexInput); err == nil {
		result.Int16LE = &v
	}
	if v, err := convert.HexToInt32LE(hexInput); err == nil {
		result.Int32LE = &v
	}
	if v, err := convert.HexToInt64LE(hexInput); err == nil {
		result.Int64LE = &v
	}

	// Try all unsigned integer conversions (Big Endian)
	if v, err := convert.HexToUint8(hexInput); err == nil {
		result.Uint8BE = &v
	}
	if v, err := convert.HexToUint16(hexInput); err == nil {
		result.Uint16BE = &v
	}
	if v, err := convert.HexToUint32(hexInput); err == nil {
		result.Uint32BE = &v
	}
	if v, err := convert.HexToUint64(hexInput); err == nil {
		result.Uint64BE = &v
	}

	// Try all unsigned integer conversions (Little Endian)
	if v, err := convert.HexToUint16LE(hexInput); err == nil {
		result.Uint16LE = &v
	}
	if v, err := convert.HexToUint32LE(hexInput); err == nil {
		result.Uint32LE = &v
	}
	if v, err := convert.HexToUint64LE(hexInput); err == nil {
		result.Uint64LE = &v
	}

	// Try float conversions
	if v, err := convert.HexToFloat32(hexInput); err == nil {
		formatted := formatFloat32(v)
		result.Float32BE = &formatted
	}
	if v, err := convert.HexToFloat64(hexInput); err == nil {
		formatted := formatFloat64(v)
		result.Float64BE = &formatted
	}
	if v, err := convert.HexToFloat32LE(hexInput); err == nil {
		formatted := formatFloat32(v)
		result.Float32LE = &formatted
	}
	if v, err := convert.HexToFloat64LE(hexInput); err == nil {
		formatted := formatFloat64(v)
		result.Float64LE = &formatted
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
		// Also populate LE conversion
		if vLE, err := convert.HexToInt16LE(hexStrLE); err == nil {
			result.Int16LE = &vLE
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
		// Also populate LE conversion
		if vLE, err := convert.HexToInt32LE(hexStrLE); err == nil {
			result.Int32LE = &vLE
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
		// Also populate LE conversion
		if vLE, err := convert.HexToInt64LE(hexStrLE); err == nil {
			result.Int64LE = &vLE
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
		// Also populate LE conversion
		if vLE, err := convert.HexToUint16LE(hexStrLE); err == nil {
			result.Uint16LE = &vLE
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
		// Also populate LE conversion
		if vLE, err := convert.HexToUint32LE(hexStrLE); err == nil {
			result.Uint32LE = &vLE
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
		// Also populate LE conversion
		if vLE, err := convert.HexToUint64LE(hexStrLE); err == nil {
			result.Uint64LE = &vLE
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
	}
	if v, err := convert.HexToInt16(hexStr); err == nil {
		result.Int16BE = &v
	}
	if v, err := convert.HexToInt32(hexStr); err == nil {
		result.Int32BE = &v
	}
	if v, err := convert.HexToInt64(hexStr); err == nil {
		result.Int64BE = &v
	}

	// Try all signed integer conversions (Little Endian)
	if v, err := convert.HexToInt16LE(hexStr); err == nil {
		result.Int16LE = &v
	}
	if v, err := convert.HexToInt32LE(hexStr); err == nil {
		result.Int32LE = &v
	}
	if v, err := convert.HexToInt64LE(hexStr); err == nil {
		result.Int64LE = &v
	}

	// Try all unsigned integer conversions (Big Endian)
	if v, err := convert.HexToUint8(hexStr); err == nil {
		result.Uint8BE = &v
	}
	if v, err := convert.HexToUint16(hexStr); err == nil {
		result.Uint16BE = &v
	}
	if v, err := convert.HexToUint32(hexStr); err == nil {
		result.Uint32BE = &v
	}
	if v, err := convert.HexToUint64(hexStr); err == nil {
		result.Uint64BE = &v
	}

	// Try all unsigned integer conversions (Little Endian)
	if v, err := convert.HexToUint16LE(hexStr); err == nil {
		result.Uint16LE = &v
	}
	if v, err := convert.HexToUint32LE(hexStr); err == nil {
		result.Uint32LE = &v
	}
	if v, err := convert.HexToUint64LE(hexStr); err == nil {
		result.Uint64LE = &v
	}

	// Try float conversions
	if v, err := convert.HexToFloat32(hexStr); err == nil {
		formatted := formatFloat32(v)
		result.Float32BE = &formatted
	}
	if v, err := convert.HexToFloat64(hexStr); err == nil {
		formatted := formatFloat64(v)
		result.Float64BE = &formatted
	}
	if v, err := convert.HexToFloat32LE(hexStr); err == nil {
		formatted := formatFloat32(v)
		result.Float32LE = &formatted
	}
	if v, err := convert.HexToFloat64LE(hexStr); err == nil {
		formatted := formatFloat64(v)
		result.Float64LE = &formatted
	}

	return result, nil
}
