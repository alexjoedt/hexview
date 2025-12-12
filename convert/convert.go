// Package convert provides comprehensive hexadecimal conversion utilities
// for converting between hex strings and various numeric types.
//
// The package supports:
//   - Flexible hex string parsing (with/without prefixes, various separators)
//   - Multiple numeric types (int8-64, uint8-64, float32/64, bytes)
//   - Both big-endian (default) and little-endian conversions
//   - Bidirectional conversions (hex ↔ numeric)
//   - Binary string conversions (e.g., "0001" ↔ numeric)
//
// Example usage:
//
//	// Parse various hex formats
//	bytes, _ := convert.HexToBytes("0x48 65 6c 6c 6f")
//
//	// Convert to integers
//	val, _ := convert.HexToInt32("0x7FFFFFFF")
//
//	// Convert back to hex
//	hex := convert.Int32ToHex(2147483647) // "7fffffff"
package convert

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Error definitions for conversion operations
var (
	// ErrInvalidHexChar indicates an invalid hexadecimal character was encountered
	ErrInvalidHexChar = errors.New("invalid hexadecimal character")

	// ErrInvalidLength indicates the hex string length doesn't match the target type
	ErrInvalidLength = errors.New("invalid hex string length for type")

	// ErrOverflow indicates the value would overflow the target type
	ErrOverflow = errors.New("value overflow for target type")

	// ErrEmptyInput indicates an empty hex string was provided
	ErrEmptyInput = errors.New("empty hex string")

	// ErrInvalidBinaryChar indicates an invalid binary character was encountered
	ErrInvalidBinaryChar = errors.New("invalid binary character")
)

// ParseHex parses a hex string in various formats and returns the byte representation.
// Supported formats include:
//   - "0x123456" (standard prefix)
//   - "04 ab cd" (space-separated)
//   - "11abcd" (continuous)
//   - "0xab 0xff" (multiple prefixed values)
//   - "xAB xCF" (x prefix without 0)
//   - Mixed case and various separators (spaces, commas, colons)
func ParseHex(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrEmptyInput
	}

	// Remove common separators and whitespace
	cleaned := strings.Builder{}
	cleaned.Grow(len(input))

	i := 0
	for i < len(input) {
		ch := input[i]

		// Skip whitespace and common separators
		if unicode.IsSpace(rune(ch)) || ch == ',' || ch == ':' || ch == '-' {
			i++
			continue
		}

		// Handle prefixes: 0x, 0X, x, X
		if ch == '0' && i+1 < len(input) && (input[i+1] == 'x' || input[i+1] == 'X') {
			i += 2
			continue
		}
		if ch == 'x' || ch == 'X' {
			i++
			continue
		}

		// Validate hex character
		if !isHexChar(ch) {
			return nil, fmt.Errorf("%w: '%c' at position %d", ErrInvalidHexChar, ch, i)
		}

		cleaned.WriteByte(ch)
		i++
	}

	hexStr := cleaned.String()
	if len(hexStr) == 0 {
		return nil, ErrEmptyInput
	}

	// Ensure even length for proper byte decoding
	if len(hexStr)%2 != 0 {
		hexStr = "0" + hexStr
	}

	// Decode hex string to bytes
	result, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidHexChar, err)
	}

	return result, nil
}

// isHexChar checks if a byte represents a valid hexadecimal character
func isHexChar(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

// ParseBinary parses a binary string (e.g., "00001010") and returns the byte representation.
// Supports formats like "1010", "0000 1111", etc.
func ParseBinary(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrEmptyInput
	}

	// Remove whitespace and separators
	cleaned := strings.Builder{}
	cleaned.Grow(len(input))

	for _, ch := range input {
		if unicode.IsSpace(ch) || ch == ',' || ch == ':' || ch == '-' || ch == '_' {
			continue
		}
		if ch != '0' && ch != '1' {
			return nil, fmt.Errorf("%w: '%c'", ErrInvalidBinaryChar, ch)
		}
		cleaned.WriteRune(ch)
	}

	binStr := cleaned.String()
	if len(binStr) == 0 {
		return nil, ErrEmptyInput
	}

	// Pad to multiple of 8
	if rem := len(binStr) % 8; rem != 0 {
		binStr = strings.Repeat("0", 8-rem) + binStr
	}

	// Convert to bytes
	result := make([]byte, len(binStr)/8)
	for i := 0; i < len(result); i++ {
		byteStr := binStr[i*8 : (i+1)*8]
		val, err := strconv.ParseUint(byteStr, 2, 8)
		if err != nil {
			return nil, fmt.Errorf("failed to parse binary: %w", err)
		}
		result[i] = byte(val)
	}

	return result, nil
}

// HexToBytes converts a hex string to a byte slice.
func HexToBytes(hexStr string) ([]byte, error) {
	return ParseHex(hexStr)
}

// BytesToHex converts a byte slice to a lowercase hex string without prefix.
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

// BytesToBinary converts a byte slice to a binary string representation.
func BytesToBinary(b []byte) string {
	var result strings.Builder
	result.Grow(len(b) * 8)

	for i, bt := range b {
		if i > 0 {
			result.WriteByte(' ')
		}
		result.WriteString(fmt.Sprintf("%08b", bt))
	}

	return result.String()
}

// Generic constraint for integer types
type integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// hexToInt is a generic helper for converting hex strings to integer types.
func hexToInt[T integer](hexStr string, byteSize int, endian binary.ByteOrder) (T, error) {
	bytes, err := ParseHex(hexStr)
	if err != nil {
		return 0, err
	}

	if len(bytes) != byteSize {
		return 0, fmt.Errorf("%w: expected %d bytes, got %d", ErrInvalidLength, byteSize, len(bytes))
	}

	var result T
	switch byteSize {
	case 1:
		result = T(bytes[0])
	case 2:
		result = T(endian.Uint16(bytes))
	case 4:
		result = T(endian.Uint32(bytes))
	case 8:
		result = T(endian.Uint64(bytes))
	}

	return result, nil
}

// intToHex is a generic helper for converting integer types to hex strings.
func intToHex[T integer](n T, byteSize int, endian binary.ByteOrder) string {
	bytes := make([]byte, byteSize)

	switch byteSize {
	case 1:
		bytes[0] = byte(n)
	case 2:
		endian.PutUint16(bytes, uint16(n))
	case 4:
		endian.PutUint32(bytes, uint32(n))
	case 8:
		endian.PutUint64(bytes, uint64(n))
	}

	return hex.EncodeToString(bytes)
}

// binaryToInt converts a binary string to an integer type.
func binaryToInt[T integer](binStr string, byteSize int, endian binary.ByteOrder) (T, error) {
	bytes, err := ParseBinary(binStr)
	if err != nil {
		return 0, err
	}

	if len(bytes) != byteSize {
		return 0, fmt.Errorf("%w: expected %d bytes, got %d", ErrInvalidLength, byteSize, len(bytes))
	}

	var result T
	switch byteSize {
	case 1:
		result = T(bytes[0])
	case 2:
		result = T(endian.Uint16(bytes))
	case 4:
		result = T(endian.Uint32(bytes))
	case 8:
		result = T(endian.Uint64(bytes))
	}

	return result, nil
}

// intToBinary converts an integer type to a binary string.
func intToBinary[T integer](n T, byteSize int, endian binary.ByteOrder) string {
	bytes := make([]byte, byteSize)

	switch byteSize {
	case 1:
		bytes[0] = byte(n)
	case 2:
		endian.PutUint16(bytes, uint16(n))
	case 4:
		endian.PutUint32(bytes, uint32(n))
	case 8:
		endian.PutUint64(bytes, uint64(n))
	}

	return BytesToBinary(bytes)
}

// swapToBADC swaps bytes to Mid-Big Endian (BADC) byte order.
// For 2-byte values: equivalent to big-endian (no swap needed)
// For 4-byte values: swap word pairs [A,B,C,D] → [B,A,D,C]
// For 8-byte values: swap 32-bit halves [A,B,C,D,E,F,G,H] → [E,F,G,H,A,B,C,D]
func swapToBADC(bytes []byte) []byte {
	result := make([]byte, len(bytes))
	copy(result, bytes)

	switch len(bytes) {
	case 2:
		// No swap for 16-bit (equivalent to big-endian)
		return result
	case 4:
		// Swap bytes within each 16-bit word
		result[0], result[1] = bytes[1], bytes[0]
		result[2], result[3] = bytes[3], bytes[2]
	case 8:
		// Swap 32-bit words
		result[0], result[1], result[2], result[3] = bytes[4], bytes[5], bytes[6], bytes[7]
		result[4], result[5], result[6], result[7] = bytes[0], bytes[1], bytes[2], bytes[3]
	}

	return result
}

// swapToCDAB swaps bytes to Mid-Little Endian (CDAB) byte order.
// For 2-byte values: equivalent to little-endian (reverse bytes)
// For 4-byte values: swap 16-bit words [A,B,C,D] → [C,D,A,B]
// For 8-byte values: swap and reverse 32-bit halves [A,B,C,D,E,F,G,H] → [C,D,A,B,G,H,E,F]
func swapToCDAB(bytes []byte) []byte {
	result := make([]byte, len(bytes))
	copy(result, bytes)

	switch len(bytes) {
	case 2:
		// Equivalent to little-endian
		result[0], result[1] = bytes[1], bytes[0]
	case 4:
		// Swap 16-bit word pairs
		result[0], result[1] = bytes[2], bytes[3]
		result[2], result[3] = bytes[0], bytes[1]
	case 8:
		// Swap 16-bit words within each 32-bit half
		result[0], result[1] = bytes[2], bytes[3]
		result[2], result[3] = bytes[0], bytes[1]
		result[4], result[5] = bytes[6], bytes[7]
		result[6], result[7] = bytes[4], bytes[5]
	}

	return result
}

// hexToIntBADC is a helper for converting hex strings to integer types using BADC byte order.
func hexToIntBADC[T integer](hexStr string, byteSize int) (T, error) {
	bytes, err := ParseHex(hexStr)
	if err != nil {
		return 0, err
	}

	if len(bytes) != byteSize {
		return 0, fmt.Errorf("%w: expected %d bytes, got %d", ErrInvalidLength, byteSize, len(bytes))
	}

	// Convert to big-endian first, then swap to BADC
	swapped := swapToBADC(bytes)

	var result T
	switch byteSize {
	case 1:
		result = T(swapped[0])
	case 2:
		result = T(binary.BigEndian.Uint16(swapped))
	case 4:
		result = T(binary.BigEndian.Uint32(swapped))
	case 8:
		result = T(binary.BigEndian.Uint64(swapped))
	}

	return result, nil
}

// hexToIntCDAB is a helper for converting hex strings to integer types using CDAB byte order.
func hexToIntCDAB[T integer](hexStr string, byteSize int) (T, error) {
	bytes, err := ParseHex(hexStr)
	if err != nil {
		return 0, err
	}

	if len(bytes) != byteSize {
		return 0, fmt.Errorf("%w: expected %d bytes, got %d", ErrInvalidLength, byteSize, len(bytes))
	}

	// Convert to big-endian first, then swap to CDAB
	swapped := swapToCDAB(bytes)

	var result T
	switch byteSize {
	case 1:
		result = T(swapped[0])
	case 2:
		result = T(binary.BigEndian.Uint16(swapped))
	case 4:
		result = T(binary.BigEndian.Uint32(swapped))
	case 8:
		result = T(binary.BigEndian.Uint64(swapped))
	}

	return result, nil
}

// intToHexBADC is a helper for converting integer types to hex strings using BADC byte order.
func intToHexBADC[T integer](n T, byteSize int) string {
	bytes := make([]byte, byteSize)

	switch byteSize {
	case 1:
		bytes[0] = byte(n)
	case 2:
		binary.BigEndian.PutUint16(bytes, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(bytes, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(bytes, uint64(n))
	}

	// Swap to BADC order
	swapped := swapToBADC(bytes)
	return hex.EncodeToString(swapped)
}

// intToHexCDAB is a helper for converting integer types to hex strings using CDAB byte order.
func intToHexCDAB[T integer](n T, byteSize int) string {
	bytes := make([]byte, byteSize)

	switch byteSize {
	case 1:
		bytes[0] = byte(n)
	case 2:
		binary.BigEndian.PutUint16(bytes, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(bytes, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(bytes, uint64(n))
	}

	// Swap to CDAB order
	swapped := swapToCDAB(bytes)
	return hex.EncodeToString(swapped)
}

// binaryToIntBADC converts a binary string to an integer type using BADC byte order.
func binaryToIntBADC[T integer](binStr string, byteSize int) (T, error) {
	bytes, err := ParseBinary(binStr)
	if err != nil {
		return 0, err
	}

	if len(bytes) != byteSize {
		return 0, fmt.Errorf("%w: expected %d bytes, got %d", ErrInvalidLength, byteSize, len(bytes))
	}

	swapped := swapToBADC(bytes)

	var result T
	switch byteSize {
	case 1:
		result = T(swapped[0])
	case 2:
		result = T(binary.BigEndian.Uint16(swapped))
	case 4:
		result = T(binary.BigEndian.Uint32(swapped))
	case 8:
		result = T(binary.BigEndian.Uint64(swapped))
	}

	return result, nil
}

// binaryToIntCDAB converts a binary string to an integer type using CDAB byte order.
func binaryToIntCDAB[T integer](binStr string, byteSize int) (T, error) {
	bytes, err := ParseBinary(binStr)
	if err != nil {
		return 0, err
	}

	if len(bytes) != byteSize {
		return 0, fmt.Errorf("%w: expected %d bytes, got %d", ErrInvalidLength, byteSize, len(bytes))
	}

	swapped := swapToCDAB(bytes)

	var result T
	switch byteSize {
	case 1:
		result = T(swapped[0])
	case 2:
		result = T(binary.BigEndian.Uint16(swapped))
	case 4:
		result = T(binary.BigEndian.Uint32(swapped))
	case 8:
		result = T(binary.BigEndian.Uint64(swapped))
	}

	return result, nil
}

// intToBinaryBADC converts an integer type to a binary string using BADC byte order.
func intToBinaryBADC[T integer](n T, byteSize int) string {
	bytes := make([]byte, byteSize)

	switch byteSize {
	case 1:
		bytes[0] = byte(n)
	case 2:
		binary.BigEndian.PutUint16(bytes, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(bytes, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(bytes, uint64(n))
	}

	swapped := swapToBADC(bytes)
	return BytesToBinary(swapped)
}

// intToBinaryCDAB converts an integer type to a binary string using CDAB byte order.
func intToBinaryCDAB[T integer](n T, byteSize int) string {
	bytes := make([]byte, byteSize)

	switch byteSize {
	case 1:
		bytes[0] = byte(n)
	case 2:
		binary.BigEndian.PutUint16(bytes, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(bytes, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(bytes, uint64(n))
	}

	swapped := swapToCDAB(bytes)
	return BytesToBinary(swapped)
}

// ============================================================================
// Signed Integer Conversions
// ============================================================================

// HexToInt8 converts a hex string to an int8 (big-endian).
func HexToInt8(hexStr string) (int8, error) {
	return hexToInt[int8](hexStr, 1, binary.BigEndian)
}

// HexToInt16 converts a hex string to an int16 (big-endian).
func HexToInt16(hexStr string) (int16, error) {
	return hexToInt[int16](hexStr, 2, binary.BigEndian)
}

// HexToInt32 converts a hex string to an int32 (big-endian).
func HexToInt32(hexStr string) (int32, error) {
	return hexToInt[int32](hexStr, 4, binary.BigEndian)
}

// HexToInt64 converts a hex string to an int64 (big-endian).
func HexToInt64(hexStr string) (int64, error) {
	return hexToInt[int64](hexStr, 8, binary.BigEndian)
}

// HexToInt8LE converts a hex string to an int8 (little-endian).
func HexToInt8LE(hexStr string) (int8, error) {
	return hexToInt[int8](hexStr, 1, binary.LittleEndian)
}

// HexToInt16LE converts a hex string to an int16 (little-endian).
func HexToInt16LE(hexStr string) (int16, error) {
	return hexToInt[int16](hexStr, 2, binary.LittleEndian)
}

// HexToInt32LE converts a hex string to an int32 (little-endian).
func HexToInt32LE(hexStr string) (int32, error) {
	return hexToInt[int32](hexStr, 4, binary.LittleEndian)
}

// HexToInt64LE converts a hex string to an int64 (little-endian).
func HexToInt64LE(hexStr string) (int64, error) {
	return hexToInt[int64](hexStr, 8, binary.LittleEndian)
}

// Int8ToHex converts an int8 to a hex string (big-endian).
func Int8ToHex(n int8) string {
	return intToHex(n, 1, binary.BigEndian)
}

// Int16ToHex converts an int16 to a hex string (big-endian).
func Int16ToHex(n int16) string {
	return intToHex(n, 2, binary.BigEndian)
}

// Int32ToHex converts an int32 to a hex string (big-endian).
func Int32ToHex(n int32) string {
	return intToHex(n, 4, binary.BigEndian)
}

// Int64ToHex converts an int64 to a hex string (big-endian).
func Int64ToHex(n int64) string {
	return intToHex(n, 8, binary.BigEndian)
}

// Int8ToHexLE converts an int8 to a hex string (little-endian).
func Int8ToHexLE(n int8) string {
	return intToHex(n, 1, binary.LittleEndian)
}

// Int16ToHexLE converts an int16 to a hex string (little-endian).
func Int16ToHexLE(n int16) string {
	return intToHex(n, 2, binary.LittleEndian)
}

// Int32ToHexLE converts an int32 to a hex string (little-endian).
func Int32ToHexLE(n int32) string {
	return intToHex(n, 4, binary.LittleEndian)
}

// Int64ToHexLE converts an int64 to a hex string (little-endian).
func Int64ToHexLE(n int64) string {
	return intToHex(n, 8, binary.LittleEndian)
}

// HexToInt16BADC converts a hex string to an int16 (mid-big-endian/BADC).
func HexToInt16BADC(hexStr string) (int16, error) {
	return hexToIntBADC[int16](hexStr, 2)
}

// HexToInt32BADC converts a hex string to an int32 (mid-big-endian/BADC).
func HexToInt32BADC(hexStr string) (int32, error) {
	return hexToIntBADC[int32](hexStr, 4)
}

// HexToInt64BADC converts a hex string to an int64 (mid-big-endian/BADC).
func HexToInt64BADC(hexStr string) (int64, error) {
	return hexToIntBADC[int64](hexStr, 8)
}

// HexToInt16CDAB converts a hex string to an int16 (mid-little-endian/CDAB).
func HexToInt16CDAB(hexStr string) (int16, error) {
	return hexToIntCDAB[int16](hexStr, 2)
}

// HexToInt32CDAB converts a hex string to an int32 (mid-little-endian/CDAB).
func HexToInt32CDAB(hexStr string) (int32, error) {
	return hexToIntCDAB[int32](hexStr, 4)
}

// HexToInt64CDAB converts a hex string to an int64 (mid-little-endian/CDAB).
func HexToInt64CDAB(hexStr string) (int64, error) {
	return hexToIntCDAB[int64](hexStr, 8)
}

// Int16ToHexBADC converts an int16 to a hex string (mid-big-endian/BADC).
func Int16ToHexBADC(n int16) string {
	return intToHexBADC(n, 2)
}

// Int32ToHexBADC converts an int32 to a hex string (mid-big-endian/BADC).
func Int32ToHexBADC(n int32) string {
	return intToHexBADC(n, 4)
}

// Int64ToHexBADC converts an int64 to a hex string (mid-big-endian/BADC).
func Int64ToHexBADC(n int64) string {
	return intToHexBADC(n, 8)
}

// Int16ToHexCDAB converts an int16 to a hex string (mid-little-endian/CDAB).
func Int16ToHexCDAB(n int16) string {
	return intToHexCDAB(n, 2)
}

// Int32ToHexCDAB converts an int32 to a hex string (mid-little-endian/CDAB).
func Int32ToHexCDAB(n int32) string {
	return intToHexCDAB(n, 4)
}

// Int64ToHexCDAB converts an int64 to a hex string (mid-little-endian/CDAB).
func Int64ToHexCDAB(n int64) string {
	return intToHexCDAB(n, 8)
}

// ============================================================================
// Unsigned Integer Conversions
// ============================================================================

// HexToUint8 converts a hex string to a uint8 (big-endian).
func HexToUint8(hexStr string) (uint8, error) {
	return hexToInt[uint8](hexStr, 1, binary.BigEndian)
}

// HexToUint16 converts a hex string to a uint16 (big-endian).
func HexToUint16(hexStr string) (uint16, error) {
	return hexToInt[uint16](hexStr, 2, binary.BigEndian)
}

// HexToUint32 converts a hex string to a uint32 (big-endian).
func HexToUint32(hexStr string) (uint32, error) {
	return hexToInt[uint32](hexStr, 4, binary.BigEndian)
}

// HexToUint64 converts a hex string to a uint64 (big-endian).
func HexToUint64(hexStr string) (uint64, error) {
	return hexToInt[uint64](hexStr, 8, binary.BigEndian)
}

// HexToUint8LE converts a hex string to a uint8 (little-endian).
func HexToUint8LE(hexStr string) (uint8, error) {
	return hexToInt[uint8](hexStr, 1, binary.LittleEndian)
}

// HexToUint16LE converts a hex string to a uint16 (little-endian).
func HexToUint16LE(hexStr string) (uint16, error) {
	return hexToInt[uint16](hexStr, 2, binary.LittleEndian)
}

// HexToUint32LE converts a hex string to a uint32 (little-endian).
func HexToUint32LE(hexStr string) (uint32, error) {
	return hexToInt[uint32](hexStr, 4, binary.LittleEndian)
}

// HexToUint64LE converts a hex string to a uint64 (little-endian).
func HexToUint64LE(hexStr string) (uint64, error) {
	return hexToInt[uint64](hexStr, 8, binary.LittleEndian)
}

// Uint8ToHex converts a uint8 to a hex string (big-endian).
func Uint8ToHex(n uint8) string {
	return intToHex(n, 1, binary.BigEndian)
}

// Uint16ToHex converts a uint16 to a hex string (big-endian).
func Uint16ToHex(n uint16) string {
	return intToHex(n, 2, binary.BigEndian)
}

// Uint32ToHex converts a uint32 to a hex string (big-endian).
func Uint32ToHex(n uint32) string {
	return intToHex(n, 4, binary.BigEndian)
}

// Uint64ToHex converts a uint64 to a hex string (big-endian).
func Uint64ToHex(n uint64) string {
	return intToHex(n, 8, binary.BigEndian)
}

// Uint8ToHexLE converts a uint8 to a hex string (little-endian).
func Uint8ToHexLE(n uint8) string {
	return intToHex(n, 1, binary.LittleEndian)
}

// Uint16ToHexLE converts a uint16 to a hex string (little-endian).
func Uint16ToHexLE(n uint16) string {
	return intToHex(n, 2, binary.LittleEndian)
}

// Uint32ToHexLE converts a uint32 to a hex string (little-endian).
func Uint32ToHexLE(n uint32) string {
	return intToHex(n, 4, binary.LittleEndian)
}

// Uint64ToHexLE converts a uint64 to a hex string (little-endian).
func Uint64ToHexLE(n uint64) string {
	return intToHex(n, 8, binary.LittleEndian)
}

// HexToUint16BADC converts a hex string to a uint16 (mid-big-endian/BADC).
func HexToUint16BADC(hexStr string) (uint16, error) {
	return hexToIntBADC[uint16](hexStr, 2)
}

// HexToUint32BADC converts a hex string to a uint32 (mid-big-endian/BADC).
func HexToUint32BADC(hexStr string) (uint32, error) {
	return hexToIntBADC[uint32](hexStr, 4)
}

// HexToUint64BADC converts a hex string to a uint64 (mid-big-endian/BADC).
func HexToUint64BADC(hexStr string) (uint64, error) {
	return hexToIntBADC[uint64](hexStr, 8)
}

// HexToUint16CDAB converts a hex string to a uint16 (mid-little-endian/CDAB).
func HexToUint16CDAB(hexStr string) (uint16, error) {
	return hexToIntCDAB[uint16](hexStr, 2)
}

// HexToUint32CDAB converts a hex string to a uint32 (mid-little-endian/CDAB).
func HexToUint32CDAB(hexStr string) (uint32, error) {
	return hexToIntCDAB[uint32](hexStr, 4)
}

// HexToUint64CDAB converts a hex string to a uint64 (mid-little-endian/CDAB).
func HexToUint64CDAB(hexStr string) (uint64, error) {
	return hexToIntCDAB[uint64](hexStr, 8)
}

// Uint16ToHexBADC converts a uint16 to a hex string (mid-big-endian/BADC).
func Uint16ToHexBADC(n uint16) string {
	return intToHexBADC(n, 2)
}

// Uint32ToHexBADC converts a uint32 to a hex string (mid-big-endian/BADC).
func Uint32ToHexBADC(n uint32) string {
	return intToHexBADC(n, 4)
}

// Uint64ToHexBADC converts a uint64 to a hex string (mid-big-endian/BADC).
func Uint64ToHexBADC(n uint64) string {
	return intToHexBADC(n, 8)
}

// Uint16ToHexCDAB converts a uint16 to a hex string (mid-little-endian/CDAB).
func Uint16ToHexCDAB(n uint16) string {
	return intToHexCDAB(n, 2)
}

// Uint32ToHexCDAB converts a uint32 to a hex string (mid-little-endian/CDAB).
func Uint32ToHexCDAB(n uint32) string {
	return intToHexCDAB(n, 4)
}

// Uint64ToHexCDAB converts a uint64 to a hex string (mid-little-endian/CDAB).
func Uint64ToHexCDAB(n uint64) string {
	return intToHexCDAB(n, 8)
}

// ============================================================================
// Float Conversions
// ============================================================================

// HexToFloat32 converts a hex string to a float32 (big-endian).
func HexToFloat32(hexStr string) (float32, error) {
	bits, err := hexToInt[uint32](hexStr, 4, binary.BigEndian)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(bits), nil
}

// HexToFloat64 converts a hex string to a float64 (big-endian).
func HexToFloat64(hexStr string) (float64, error) {
	bits, err := hexToInt[uint64](hexStr, 8, binary.BigEndian)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(bits), nil
}

// HexToFloat32LE converts a hex string to a float32 (little-endian).
func HexToFloat32LE(hexStr string) (float32, error) {
	bits, err := hexToInt[uint32](hexStr, 4, binary.LittleEndian)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(bits), nil
}

// HexToFloat64LE converts a hex string to a float64 (little-endian).
func HexToFloat64LE(hexStr string) (float64, error) {
	bits, err := hexToInt[uint64](hexStr, 8, binary.LittleEndian)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(bits), nil
}

// Float32ToHex converts a float32 to a hex string (big-endian).
func Float32ToHex(f float32) string {
	bits := math.Float32bits(f)
	return intToHex(bits, 4, binary.BigEndian)
}

// Float64ToHex converts a float64 to a hex string (big-endian).
func Float64ToHex(f float64) string {
	bits := math.Float64bits(f)
	return intToHex(bits, 8, binary.BigEndian)
}

// Float32ToHexLE converts a float32 to a hex string (little-endian).
func Float32ToHexLE(f float32) string {
	bits := math.Float32bits(f)
	return intToHex(bits, 4, binary.LittleEndian)
}

// Float64ToHexLE converts a float64 to a hex string (little-endian).
func Float64ToHexLE(f float64) string {
	bits := math.Float64bits(f)
	return intToHex(bits, 8, binary.LittleEndian)
}

// HexToFloat32BADC converts a hex string to a float32 (mid-big-endian/BADC).
func HexToFloat32BADC(hexStr string) (float32, error) {
	bits, err := hexToIntBADC[uint32](hexStr, 4)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(bits), nil
}

// HexToFloat64BADC converts a hex string to a float64 (mid-big-endian/BADC).
func HexToFloat64BADC(hexStr string) (float64, error) {
	bits, err := hexToIntBADC[uint64](hexStr, 8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(bits), nil
}

// HexToFloat32CDAB converts a hex string to a float32 (mid-little-endian/CDAB).
func HexToFloat32CDAB(hexStr string) (float32, error) {
	bits, err := hexToIntCDAB[uint32](hexStr, 4)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(bits), nil
}

// HexToFloat64CDAB converts a hex string to a float64 (mid-little-endian/CDAB).
func HexToFloat64CDAB(hexStr string) (float64, error) {
	bits, err := hexToIntCDAB[uint64](hexStr, 8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(bits), nil
}

// Float32ToHexBADC converts a float32 to a hex string (mid-big-endian/BADC).
func Float32ToHexBADC(f float32) string {
	bits := math.Float32bits(f)
	return intToHexBADC(bits, 4)
}

// Float64ToHexBADC converts a float64 to a hex string (mid-big-endian/BADC).
func Float64ToHexBADC(f float64) string {
	bits := math.Float64bits(f)
	return intToHexBADC(bits, 8)
}

// Float32ToHexCDAB converts a float32 to a hex string (mid-little-endian/CDAB).
func Float32ToHexCDAB(f float32) string {
	bits := math.Float32bits(f)
	return intToHexCDAB(bits, 4)
}

// Float64ToHexCDAB converts a float64 to a hex string (mid-little-endian/CDAB).
func Float64ToHexCDAB(f float64) string {
	bits := math.Float64bits(f)
	return intToHexCDAB(bits, 8)
}

// ============================================================================
// Binary String Conversions (Signed Integers)
// ============================================================================

// BinaryToInt8 converts a binary string to an int8 (big-endian).
func BinaryToInt8(binStr string) (int8, error) {
	return binaryToInt[int8](binStr, 1, binary.BigEndian)
}

// BinaryToInt16 converts a binary string to an int16 (big-endian).
func BinaryToInt16(binStr string) (int16, error) {
	return binaryToInt[int16](binStr, 2, binary.BigEndian)
}

// BinaryToInt32 converts a binary string to an int32 (big-endian).
func BinaryToInt32(binStr string) (int32, error) {
	return binaryToInt[int32](binStr, 4, binary.BigEndian)
}

// BinaryToInt64 converts a binary string to an int64 (big-endian).
func BinaryToInt64(binStr string) (int64, error) {
	return binaryToInt[int64](binStr, 8, binary.BigEndian)
}

// BinaryToInt8LE converts a binary string to an int8 (little-endian).
func BinaryToInt8LE(binStr string) (int8, error) {
	return binaryToInt[int8](binStr, 1, binary.LittleEndian)
}

// BinaryToInt16LE converts a binary string to an int16 (little-endian).
func BinaryToInt16LE(binStr string) (int16, error) {
	return binaryToInt[int16](binStr, 2, binary.LittleEndian)
}

// BinaryToInt32LE converts a binary string to an int32 (little-endian).
func BinaryToInt32LE(binStr string) (int32, error) {
	return binaryToInt[int32](binStr, 4, binary.LittleEndian)
}

// BinaryToInt64LE converts a binary string to an int64 (little-endian).
func BinaryToInt64LE(binStr string) (int64, error) {
	return binaryToInt[int64](binStr, 8, binary.LittleEndian)
}

// Int8ToBinary converts an int8 to a binary string (big-endian).
func Int8ToBinary(n int8) string {
	return intToBinary(n, 1, binary.BigEndian)
}

// Int16ToBinary converts an int16 to a binary string (big-endian).
func Int16ToBinary(n int16) string {
	return intToBinary(n, 2, binary.BigEndian)
}

// Int32ToBinary converts an int32 to a binary string (big-endian).
func Int32ToBinary(n int32) string {
	return intToBinary(n, 4, binary.BigEndian)
}

// Int64ToBinary converts an int64 to a binary string (big-endian).
func Int64ToBinary(n int64) string {
	return intToBinary(n, 8, binary.BigEndian)
}

// Int8ToBinaryLE converts an int8 to a binary string (little-endian).
func Int8ToBinaryLE(n int8) string {
	return intToBinary(n, 1, binary.LittleEndian)
}

// Int16ToBinaryLE converts an int16 to a binary string (little-endian).
func Int16ToBinaryLE(n int16) string {
	return intToBinary(n, 2, binary.LittleEndian)
}

// Int32ToBinaryLE converts an int32 to a binary string (little-endian).
func Int32ToBinaryLE(n int32) string {
	return intToBinary(n, 4, binary.LittleEndian)
}

// Int64ToBinaryLE converts an int64 to a binary string (little-endian).
func Int64ToBinaryLE(n int64) string {
	return intToBinary(n, 8, binary.LittleEndian)
}

// BinaryToInt16BADC converts a binary string to an int16 (mid-big-endian/BADC).
func BinaryToInt16BADC(binStr string) (int16, error) {
	return binaryToIntBADC[int16](binStr, 2)
}

// BinaryToInt32BADC converts a binary string to an int32 (mid-big-endian/BADC).
func BinaryToInt32BADC(binStr string) (int32, error) {
	return binaryToIntBADC[int32](binStr, 4)
}

// BinaryToInt64BADC converts a binary string to an int64 (mid-big-endian/BADC).
func BinaryToInt64BADC(binStr string) (int64, error) {
	return binaryToIntBADC[int64](binStr, 8)
}

// BinaryToInt16CDAB converts a binary string to an int16 (mid-little-endian/CDAB).
func BinaryToInt16CDAB(binStr string) (int16, error) {
	return binaryToIntCDAB[int16](binStr, 2)
}

// BinaryToInt32CDAB converts a binary string to an int32 (mid-little-endian/CDAB).
func BinaryToInt32CDAB(binStr string) (int32, error) {
	return binaryToIntCDAB[int32](binStr, 4)
}

// BinaryToInt64CDAB converts a binary string to an int64 (mid-little-endian/CDAB).
func BinaryToInt64CDAB(binStr string) (int64, error) {
	return binaryToIntCDAB[int64](binStr, 8)
}

// Int16ToBinaryBADC converts an int16 to a binary string (mid-big-endian/BADC).
func Int16ToBinaryBADC(n int16) string {
	return intToBinaryBADC(n, 2)
}

// Int32ToBinaryBADC converts an int32 to a binary string (mid-big-endian/BADC).
func Int32ToBinaryBADC(n int32) string {
	return intToBinaryBADC(n, 4)
}

// Int64ToBinaryBADC converts an int64 to a binary string (mid-big-endian/BADC).
func Int64ToBinaryBADC(n int64) string {
	return intToBinaryBADC(n, 8)
}

// Int16ToBinaryCDAB converts an int16 to a binary string (mid-little-endian/CDAB).
func Int16ToBinaryCDAB(n int16) string {
	return intToBinaryCDAB(n, 2)
}

// Int32ToBinaryCDAB converts an int32 to a binary string (mid-little-endian/CDAB).
func Int32ToBinaryCDAB(n int32) string {
	return intToBinaryCDAB(n, 4)
}

// Int64ToBinaryCDAB converts an int64 to a binary string (mid-little-endian/CDAB).
func Int64ToBinaryCDAB(n int64) string {
	return intToBinaryCDAB(n, 8)
}

// ============================================================================
// Binary String Conversions (Unsigned Integers)
// ============================================================================

// BinaryToUint8 converts a binary string to a uint8 (big-endian).
func BinaryToUint8(binStr string) (uint8, error) {
	return binaryToInt[uint8](binStr, 1, binary.BigEndian)
}

// BinaryToUint16 converts a binary string to a uint16 (big-endian).
func BinaryToUint16(binStr string) (uint16, error) {
	return binaryToInt[uint16](binStr, 2, binary.BigEndian)
}

// BinaryToUint32 converts a binary string to a uint32 (big-endian).
func BinaryToUint32(binStr string) (uint32, error) {
	return binaryToInt[uint32](binStr, 4, binary.BigEndian)
}

// BinaryToUint64 converts a binary string to a uint64 (big-endian).
func BinaryToUint64(binStr string) (uint64, error) {
	return binaryToInt[uint64](binStr, 8, binary.BigEndian)
}

// BinaryToUint8LE converts a binary string to a uint8 (little-endian).
func BinaryToUint8LE(binStr string) (uint8, error) {
	return binaryToInt[uint8](binStr, 1, binary.LittleEndian)
}

// BinaryToUint16LE converts a binary string to a uint16 (little-endian).
func BinaryToUint16LE(binStr string) (uint16, error) {
	return binaryToInt[uint16](binStr, 2, binary.LittleEndian)
}

// BinaryToUint32LE converts a binary string to a uint32 (little-endian).
func BinaryToUint32LE(binStr string) (uint32, error) {
	return binaryToInt[uint32](binStr, 4, binary.LittleEndian)
}

// BinaryToUint64LE converts a binary string to a uint64 (little-endian).
func BinaryToUint64LE(binStr string) (uint64, error) {
	return binaryToInt[uint64](binStr, 8, binary.LittleEndian)
}

// Uint8ToBinary converts a uint8 to a binary string (big-endian).
func Uint8ToBinary(n uint8) string {
	return intToBinary(n, 1, binary.BigEndian)
}

// Uint16ToBinary converts a uint16 to a binary string (big-endian).
func Uint16ToBinary(n uint16) string {
	return intToBinary(n, 2, binary.BigEndian)
}

// Uint32ToBinary converts a uint32 to a binary string (big-endian).
func Uint32ToBinary(n uint32) string {
	return intToBinary(n, 4, binary.BigEndian)
}

// Uint64ToBinary converts a uint64 to a binary string (big-endian).
func Uint64ToBinary(n uint64) string {
	return intToBinary(n, 8, binary.BigEndian)
}

// Uint8ToBinaryLE converts a uint8 to a binary string (little-endian).
func Uint8ToBinaryLE(n uint8) string {
	return intToBinary(n, 1, binary.LittleEndian)
}

// Uint16ToBinaryLE converts a uint16 to a binary string (little-endian).
func Uint16ToBinaryLE(n uint16) string {
	return intToBinary(n, 2, binary.LittleEndian)
}

// Uint32ToBinaryLE converts a uint32 to a binary string (little-endian).
func Uint32ToBinaryLE(n uint32) string {
	return intToBinary(n, 4, binary.LittleEndian)
}

// Uint64ToBinaryLE converts a uint64 to a binary string (little-endian).
func Uint64ToBinaryLE(n uint64) string {
	return intToBinary(n, 8, binary.LittleEndian)
}

// BinaryToUint16BADC converts a binary string to a uint16 (mid-big-endian/BADC).
func BinaryToUint16BADC(binStr string) (uint16, error) {
	return binaryToIntBADC[uint16](binStr, 2)
}

// BinaryToUint32BADC converts a binary string to a uint32 (mid-big-endian/BADC).
func BinaryToUint32BADC(binStr string) (uint32, error) {
	return binaryToIntBADC[uint32](binStr, 4)
}

// BinaryToUint64BADC converts a binary string to a uint64 (mid-big-endian/BADC).
func BinaryToUint64BADC(binStr string) (uint64, error) {
	return binaryToIntBADC[uint64](binStr, 8)
}

// BinaryToUint16CDAB converts a binary string to a uint16 (mid-little-endian/CDAB).
func BinaryToUint16CDAB(binStr string) (uint16, error) {
	return binaryToIntCDAB[uint16](binStr, 2)
}

// BinaryToUint32CDAB converts a binary string to a uint32 (mid-little-endian/CDAB).
func BinaryToUint32CDAB(binStr string) (uint32, error) {
	return binaryToIntCDAB[uint32](binStr, 4)
}

// BinaryToUint64CDAB converts a binary string to a uint64 (mid-little-endian/CDAB).
func BinaryToUint64CDAB(binStr string) (uint64, error) {
	return binaryToIntCDAB[uint64](binStr, 8)
}

// Uint16ToBinaryBADC converts a uint16 to a binary string (mid-big-endian/BADC).
func Uint16ToBinaryBADC(n uint16) string {
	return intToBinaryBADC(n, 2)
}

// Uint32ToBinaryBADC converts a uint32 to a binary string (mid-big-endian/BADC).
func Uint32ToBinaryBADC(n uint32) string {
	return intToBinaryBADC(n, 4)
}

// Uint64ToBinaryBADC converts a uint64 to a binary string (mid-big-endian/BADC).
func Uint64ToBinaryBADC(n uint64) string {
	return intToBinaryBADC(n, 8)
}

// Uint16ToBinaryCDAB converts a uint16 to a binary string (mid-little-endian/CDAB).
func Uint16ToBinaryCDAB(n uint16) string {
	return intToBinaryCDAB(n, 2)
}

// Uint32ToBinaryCDAB converts a uint32 to a binary string (mid-little-endian/CDAB).
func Uint32ToBinaryCDAB(n uint32) string {
	return intToBinaryCDAB(n, 4)
}

// Uint64ToBinaryCDAB converts a uint64 to a binary string (mid-little-endian/CDAB).
func Uint64ToBinaryCDAB(n uint64) string {
	return intToBinaryCDAB(n, 8)
}
