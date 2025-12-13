// Package service provides business logic for hex/int/binary/float conversions.
// It acts as an intermediary between the Wails app layer and the convert package,
// organizing conversion operations and building comprehensive result structures.
package service

import (
	"fmt"
	"math"
	"strings"

	"hexview/convert"
	"hexview/models"
)

// Converter provides methods for converting between hex, integer, binary, and float formats.
type Converter struct{}

// NewConverter creates a new Converter instance.
func NewConverter() *Converter {
	return &Converter{}
}

// ConvertHex performs all possible conversions on hex input.
func (c *Converter) ConvertHex(hexInput string) (*models.ConversionResult, error) {
	if hexInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &models.ConversionResult{}

	// Convert to bytes first to get binary representation
	bytes, err := convert.HexToBytes(hexInput)
	if err != nil {
		return nil, fmt.Errorf("invalid hex input: %w", err)
	}

	result.Binary = convert.BytesToBinary(bytes)
	result.Bytes = convert.BytesToHex(bytes)
	result.ASCII = bytesToASCII(bytes)

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

// ConvertInt performs conversions from integer input to hex and binary.
func (c *Converter) ConvertInt(intInput string, intType string) (*models.ConversionResult, error) {
	if intInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &models.ConversionResult{}

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
		result.ASCII = bytesToASCII(bytes)
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
		result.ASCII = bytesToASCII(bytes)
		result.Int16BE = &val
		result.Int16BEHex = hexStrBE
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
		result.ASCII = bytesToASCII(bytes)
		result.Int32BE = &val
		result.Int32BEHex = hexStrBE
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
		result.ASCII = bytesToASCII(bytes)
		result.Int64BE = &val
		result.Int64BEHex = hexStrBE
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
		result.ASCII = bytesToASCII(bytes)
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
		result.ASCII = bytesToASCII(bytes)
		result.Uint16BE = &val
		result.Uint16BEHex = hexStrBE
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
		result.ASCII = bytesToASCII(bytes)
		result.Uint32BE = &val
		result.Uint32BEHex = hexStrBE
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
		result.ASCII = bytesToASCII(bytes)
		result.Uint64BE = &val
		result.Uint64BEHex = hexStrBE
		if vLE, err := convert.HexToUint64LE(hexStrLE); err == nil {
			result.Uint64LE = &vLE
			result.Uint64LEHex = hexStrLE
		}
		return result, nil

	default:
		return nil, fmt.Errorf("unsupported integer type: %s", intType)
	}
}

// ConvertIntAuto performs auto-detection of integer and float types based on the input value.
// It determines which integer types (int8/16/32/64, uint8/16/32/64) can represent
// the given decimal value and populates the result with all compatible representations.
// If the input contains a decimal point or comma, it's treated as a float.
// Negative values automatically exclude unsigned types.
func (c *Converter) ConvertIntAuto(intInput string) (*models.ConversionResult, error) {
	if intInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	// Normalize comma to dot for float parsing (support both European and US notation)
	normalizedInput := strings.ReplaceAll(intInput, ",", ".")

	// Check if input contains a decimal point (float)
	if strings.Contains(normalizedInput, ".") {
		return c.convertFloatAuto(normalizedInput)
	}

	result := &models.ConversionResult{}

	// Parse as int64 to determine value range
	var val64 int64
	_, err := fmt.Sscanf(intInput, "%d", &val64)
	if err != nil {
		return nil, fmt.Errorf("invalid decimal value: %w", err)
	}

	// Helper function to set binary/bytes/ASCII from hex string (use first valid representation)
	setCommonFields := func(hexStr string) {
		if result.Binary == "" {
			bytes, _ := convert.HexToBytes(hexStr)
			result.Binary = convert.BytesToBinary(bytes)
			result.Bytes = hexStr
			result.ASCII = bytesToASCII(bytes)
		}
	}

	// Check int8 (-128 to 127)
	if val64 >= -128 && val64 <= 127 {
		val := int8(val64)
		hexStr := convert.Int8ToHex(val)
		setCommonFields(hexStr)
		result.Int8BE = &val
		result.Int8BEHex = hexStr
	}

	// Check uint8 (0 to 255)
	if val64 >= 0 && val64 <= 255 {
		val := uint8(val64)
		hexStr := convert.Uint8ToHex(val)
		if result.Binary == "" {
			setCommonFields(hexStr)
		}
		result.Uint8BE = &val
		result.Uint8BEHex = hexStr
	}

	// Check int16 (-32768 to 32767)
	if val64 >= -32768 && val64 <= 32767 {
		val := int16(val64)
		hexStrBE := convert.Int16ToHex(val)
		hexStrLE := convert.Int16ToHexLE(val)
		setCommonFields(hexStrBE)
		result.Int16BE = &val
		result.Int16BEHex = hexStrBE
		if vLE, err := convert.HexToInt16LE(hexStrLE); err == nil {
			result.Int16LE = &vLE
			result.Int16LEHex = hexStrLE
		}
	}

	// Check uint16 (0 to 65535)
	if val64 >= 0 && val64 <= 65535 {
		val := uint16(val64)
		hexStrBE := convert.Uint16ToHex(val)
		hexStrLE := convert.Uint16ToHexLE(val)
		if result.Binary == "" {
			setCommonFields(hexStrBE)
		}
		result.Uint16BE = &val
		result.Uint16BEHex = hexStrBE
		if vLE, err := convert.HexToUint16LE(hexStrLE); err == nil {
			result.Uint16LE = &vLE
			result.Uint16LEHex = hexStrLE
		}
	}

	// Check int32 (-2147483648 to 2147483647)
	if val64 >= -2147483648 && val64 <= 2147483647 {
		val := int32(val64)
		hexStrBE := convert.Int32ToHex(val)
		hexStrLE := convert.Int32ToHexLE(val)
		setCommonFields(hexStrBE)
		result.Int32BE = &val
		result.Int32BEHex = hexStrBE
		if vLE, err := convert.HexToInt32LE(hexStrLE); err == nil {
			result.Int32LE = &vLE
			result.Int32LEHex = hexStrLE
		}
	}

	// Check uint32 (0 to 4294967295)
	if val64 >= 0 && val64 <= 4294967295 {
		val := uint32(val64)
		hexStrBE := convert.Uint32ToHex(val)
		hexStrLE := convert.Uint32ToHexLE(val)
		if result.Binary == "" {
			setCommonFields(hexStrBE)
		}
		result.Uint32BE = &val
		result.Uint32BEHex = hexStrBE
		if vLE, err := convert.HexToUint32LE(hexStrLE); err == nil {
			result.Uint32LE = &vLE
			result.Uint32LEHex = hexStrLE
		}
	}

	// Always include int64 (if parsed successfully)
	result.Int64BE = &val64
	hexStrBE := convert.Int64ToHex(val64)
	hexStrLE := convert.Int64ToHexLE(val64)
	setCommonFields(hexStrBE)
	result.Int64BEHex = hexStrBE
	if vLE, err := convert.HexToInt64LE(hexStrLE); err == nil {
		result.Int64LE = &vLE
		result.Int64LEHex = hexStrLE
	}

	// Check uint64 (0 to 9223372036854775807 - limited by int64 parsing)
	// For values beyond int64 max, we'd need different parsing approach
	if val64 >= 0 {
		val := uint64(val64)
		hexStrBE := convert.Uint64ToHex(val)
		hexStrLE := convert.Uint64ToHexLE(val)
		result.Uint64BE = &val
		result.Uint64BEHex = hexStrBE
		if vLE, err := convert.HexToUint64LE(hexStrLE); err == nil {
			result.Uint64LE = &vLE
			result.Uint64LEHex = hexStrLE
		}
	}

	return result, nil
}

// convertFloatAuto is a helper function that handles float value auto-detection.
// It populates float32 and float64 representations in all endianness variants.
func (c *Converter) convertFloatAuto(floatInput string) (*models.ConversionResult, error) {
	result := &models.ConversionResult{}

	// Parse as float64 first
	var val64 float64
	_, err := fmt.Sscanf(floatInput, "%f", &val64)
	if err != nil {
		return nil, fmt.Errorf("invalid float value: %w", err)
	}

	// Convert to float32 to check if it fits without precision loss
	val32 := float32(val64)

	// Helper function to set binary/bytes/ASCII from hex string (use first valid representation)
	setCommonFields := func(hexStr string) {
		if result.Binary == "" {
			bytes, _ := convert.HexToBytes(hexStr)
			result.Binary = convert.BytesToBinary(bytes)
			result.Bytes = hexStr
			result.ASCII = bytesToASCII(bytes)
		}
	}

	// Float32 conversions (all endianness variants)
	hexStrBE32 := convert.Float32ToHex(val32)
	setCommonFields(hexStrBE32)
	formatted32 := formatFloat32(val32)
	result.Float32BE = &formatted32
	result.Float32BEHex = hexStrBE32

	hexStrLE32 := convert.Float32ToHexLE(val32)
	if vLE, err := convert.HexToFloat32LE(hexStrLE32); err == nil {
		formattedLE := formatFloat32(vLE)
		result.Float32LE = &formattedLE
		result.Float32LEHex = hexStrLE32
	}

	hexStrBADC32 := convert.Float32ToHexBADC(val32)
	if vBADC, err := convert.HexToFloat32BADC(hexStrBADC32); err == nil {
		formattedBADC := formatFloat32(vBADC)
		result.Float32BADC = &formattedBADC
		result.Float32BADCHex = hexStrBADC32
	}

	hexStrCDAB32 := convert.Float32ToHexCDAB(val32)
	if vCDAB, err := convert.HexToFloat32CDAB(hexStrCDAB32); err == nil {
		formattedCDAB := formatFloat32(vCDAB)
		result.Float32CDAB = &formattedCDAB
		result.Float32CDABHex = hexStrCDAB32
	}

	// Float64 conversions (all endianness variants)
	hexStrBE64 := convert.Float64ToHex(val64)
	if result.Binary == "" {
		setCommonFields(hexStrBE64)
	}
	formatted64 := formatFloat64(val64)
	result.Float64BE = &formatted64
	result.Float64BEHex = hexStrBE64

	hexStrLE64 := convert.Float64ToHexLE(val64)
	if vLE, err := convert.HexToFloat64LE(hexStrLE64); err == nil {
		formattedLE := formatFloat64(vLE)
		result.Float64LE = &formattedLE
		result.Float64LEHex = hexStrLE64
	}

	hexStrBADC64 := convert.Float64ToHexBADC(val64)
	if vBADC, err := convert.HexToFloat64BADC(hexStrBADC64); err == nil {
		formattedBADC := formatFloat64(vBADC)
		result.Float64BADC = &formattedBADC
		result.Float64BADCHex = hexStrBADC64
	}

	hexStrCDAB64 := convert.Float64ToHexCDAB(val64)
	if vCDAB, err := convert.HexToFloat64CDAB(hexStrCDAB64); err == nil {
		formattedCDAB := formatFloat64(vCDAB)
		result.Float64CDAB = &formattedCDAB
		result.Float64CDABHex = hexStrCDAB64
	}

	return result, nil
}

// ConvertBinary performs all possible conversions on binary input.
func (c *Converter) ConvertBinary(binaryInput string) (*models.ConversionResult, error) {
	if binaryInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &models.ConversionResult{}

	bytes, err := convert.ParseBinary(binaryInput)
	if err != nil {
		return nil, fmt.Errorf("invalid binary input: %w", err)
	}

	result.Binary = convert.BytesToBinary(bytes)
	result.Bytes = convert.BytesToHex(bytes)
	result.ASCII = bytesToASCII(bytes)

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

// ConvertFloat performs conversions from float input to hex and binary.
func (c *Converter) ConvertFloat(floatInput string, floatType string) (*models.ConversionResult, error) {
	if floatInput == "" {
		return nil, fmt.Errorf("empty input")
	}

	result := &models.ConversionResult{}

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
		result.ASCII = bytesToASCII(bytes)

		formatted := formatFloat32(val)
		result.Float32BE = &formatted
		result.Float32BEHex = hexStrBE

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
		_, err := fmt.Sscanf(floatInput, "%f", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid float64 value: %w", err)
		}
		hexStrBE := convert.Float64ToHex(val)
		bytes, _ := convert.HexToBytes(hexStrBE)
		result.Binary = convert.BytesToBinary(bytes)
		result.Bytes = hexStrBE
		result.ASCII = bytesToASCII(bytes)

		formatted := formatFloat64(val)
		result.Float64BE = &formatted
		result.Float64BEHex = hexStrBE

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

// ConvertModbusRegisters converts an array of 16-bit register values.
func (c *Converter) ConvertModbusRegisters(input string) (*models.ModbusResult, error) {
	if input == "" {
		return nil, fmt.Errorf("empty input")
	}

	registers, err := parseModbusInput(input)
	if err != nil {
		return nil, err
	}

	if len(registers) == 0 {
		return nil, fmt.Errorf("no valid register values found")
	}

	result := &models.ModbusResult{
		Registers:  make([]models.ModbusRegister, len(registers)),
		Combined32: make([]models.ModbusCombined32, 0),
		Combined64: make([]models.ModbusCombined64, 0),
	}

	var allBytes []byte
	var hexParts []string

	for i, val := range registers {
		regHex := convert.Uint16ToHex(val)
		hexParts = append(hexParts, regHex)

		regBytes, _ := convert.HexToBytes(regHex)
		allBytes = append(allBytes, regBytes...)

		result.Registers[i] = models.ModbusRegister{
			Index:    i + 1,
			Hex:      regHex,
			Unsigned: val,
			Signed:   int16(val),
			Binary:   convert.Uint16ToBinary(val),
		}
	}

	result.RawHex = strings.Join(hexParts, " ")
	result.ASCII = bytesToASCII(allBytes)

	// Generate 32-bit combinations
	for i := 0; i <= len(registers)-2; i++ {
		hexStr := convert.Uint16ToHex(registers[i]) + convert.Uint16ToHex(registers[i+1])

		combined := models.ModbusCombined32{
			RegisterStart: i + 1,
			Hex:           hexStr,
		}

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

	// Generate 64-bit combinations
	for i := 0; i <= len(registers)-4; i++ {
		hexStr := convert.Uint16ToHex(registers[i]) +
			convert.Uint16ToHex(registers[i+1]) +
			convert.Uint16ToHex(registers[i+2]) +
			convert.Uint16ToHex(registers[i+3])

		combined := models.ModbusCombined64{
			RegisterStart: i + 1,
			Hex:           hexStr,
		}

		if v, err := convert.HexToUint64(hexStr); err == nil {
			combined.Uint64BE = v
		}
		if v, err := convert.HexToUint64LE(hexStr); err == nil {
			combined.Uint64LE = v
		}
		if v, err := convert.HexToInt64(hexStr); err == nil {
			combined.Int64BE = v
		}
		if v, err := convert.HexToInt64LE(hexStr); err == nil {
			combined.Int64LE = v
		}
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

// Helper functions

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

func bytesToASCII(bytes []byte) string {
	var sb strings.Builder
	for _, b := range bytes {
		if b >= 32 && b <= 126 {
			sb.WriteByte(b)
		} else {
			sb.WriteByte('.')
		}
	}
	return sb.String()
}

func parseModbusInput(input string) ([]uint16, error) {
	// Replace common separators with spaces
	normalized := strings.ReplaceAll(input, ",", " ")
	normalized = strings.ReplaceAll(normalized, ";", " ")
	normalized = strings.ReplaceAll(normalized, "\t", " ")
	normalized = strings.ReplaceAll(normalized, "\n", " ")
	normalized = strings.ReplaceAll(normalized, ":", " ")

	parts := strings.Fields(normalized)
	registers := make([]uint16, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		var val uint64
		var err error

		if len(part) > 1 && (part[0] == 'd' || part[0] == 'D') {
			_, err = fmt.Sscanf(part[1:], "%d", &val)
			if err != nil {
				return nil, fmt.Errorf("invalid decimal value: %s", part)
			}
		} else {
			cleanHex := strings.TrimPrefix(part, "0x")
			cleanHex = strings.TrimPrefix(cleanHex, "0X")
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
