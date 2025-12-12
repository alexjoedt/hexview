package main

import (
	"context"
	"fmt"

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

	// Floating Point
	Float32BE *float32 `json:"float32BE,omitempty"`
	Float64BE *float64 `json:"float64BE,omitempty"`
	Float32LE *float32 `json:"float32LE,omitempty"`
	Float64LE *float64 `json:"float64LE,omitempty"`

	// Binary Representations
	Binary string `json:"binary,omitempty"`
	Bytes  string `json:"bytes,omitempty"`
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
		result.Float32BE = &v
	}
	if v, err := convert.HexToFloat64(hexInput); err == nil {
		result.Float64BE = &v
	}
	if v, err := convert.HexToFloat32LE(hexInput); err == nil {
		result.Float32LE = &v
	}
	if v, err := convert.HexToFloat64LE(hexInput); err == nil {
		result.Float64LE = &v
	}

	return result, nil
}
