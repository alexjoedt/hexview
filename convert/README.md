# Convert Package

A comprehensive Go package for hexadecimal and binary string conversions with support for multiple numeric types and endianness configurations.

## Features

- **Flexible Input Parsing**: Parse hex strings in various formats (with/without prefixes, various separators)
- **Multiple Numeric Types**: Support for int8-64, uint8-64, float32/64, and byte slices
- **Bidirectional Conversions**: Convert both to and from hex/binary representations
- **Endianness Control**: Big-endian (default) and little-endian support for all multi-byte types
- **Binary String Support**: Convert between binary strings (e.g., "00001010") and numeric types
- **Comprehensive Testing**: 92% test coverage with extensive edge case handling

## Installation

```bash
go get hexview/convert
```

## Quick Start

```go
package main

import (
    "fmt"
    "hexview/convert"
)

func main() {
    // Parse hex strings in various formats
    bytes, _ := convert.HexToBytes("0x48 65 6c 6c 6f")
    fmt.Println(string(bytes)) // Output: Hello
    
    // Convert integers
    val, _ := convert.HexToInt32("7fffffff")
    fmt.Println(val) // Output: 2147483647
    
    // Convert back to hex
    hex := convert.Int32ToHex(2147483647)
    fmt.Println(hex) // Output: 7fffffff
    
    // Binary conversions
    binVal, _ := convert.BinaryToInt8("01111111")
    fmt.Println(binVal) // Output: 127
}
```

## Supported Hex Input Formats

The package accepts hex strings in multiple formats:

```go
"0x123456"        // Standard 0x prefix
"0X123456"        // Uppercase prefix
"x123456"         // x prefix without 0
"12 34 56"        // Space-separated
"12:34:56"        // Colon-separated
"12,34,56"        // Comma-separated
"123456"          // Continuous
"0xAB 0xCD"       // Multiple prefixes
"aAbBcC"          // Mixed case
```

## API Overview

### Byte Conversions

```go
func HexToBytes(hex string) ([]byte, error)
func BytesToHex(b []byte) string
func BytesToBinary(b []byte) string
```

### Integer Conversions (Signed)

**Big-Endian (Default):**
```go
func HexToInt8(hex string) (int8, error)
func HexToInt16(hex string) (int16, error)
func HexToInt32(hex string) (int32, error)
func HexToInt64(hex string) (int64, error)

func Int8ToHex(n int8) string
func Int16ToHex(n int16) string
func Int32ToHex(n int32) string
func Int64ToHex(n int64) string
```

**Little-Endian:**
```go
func HexToInt16LE(hex string) (int16, error)
func HexToInt32LE(hex string) (int32, error)
func HexToInt64LE(hex string) (int64, error)

func Int16ToHexLE(n int16) string
func Int32ToHexLE(n int32) string
func Int64ToHexLE(n int64) string
```

### Integer Conversions (Unsigned)

Similar functions available for `Uint8`, `Uint16`, `Uint32`, and `Uint64` with both big-endian and little-endian variants.

### Float Conversions

```go
func HexToFloat32(hex string) (float32, error)
func HexToFloat64(hex string) (float64, error)
func Float32ToHex(f float32) string
func Float64ToHex(f float64) string

// Little-endian variants
func HexToFloat32LE(hex string) (float32, error)
func HexToFloat64LE(hex string) (float64, error)
func Float32ToHexLE(f float32) string
func Float64ToHexLE(f float64) string
```

### Binary String Conversions

```go
func BinaryToInt8(bin string) (int8, error)
func BinaryToInt16(bin string) (int16, error)
func BinaryToInt32(bin string) (int32, error)
func BinaryToInt64(bin string) (int64, error)

func Int8ToBinary(n int8) string
func Int16ToBinary(n int16) string
func Int32ToBinary(n int32) string
func Int64ToBinary(n int64) string
```

Similar functions available for unsigned integers and little-endian variants.

## Examples

### Hex Parsing

```go
// Parse various hex formats
bytes, _ := convert.ParseHex("0xAB CD EF")
fmt.Println(convert.BytesToHex(bytes)) // Output: abcdef
```

### Endianness

```go
// Big-endian (default)
be, _ := convert.HexToInt32("12345678")
fmt.Printf("0x%08x\n", be) // Output: 0x12345678

// Little-endian
le, _ := convert.HexToInt32LE("12345678")
fmt.Printf("0x%08x\n", le) // Output: 0x78563412
```

### Float Conversions

```go
// Convert hex to float
val, _ := convert.HexToFloat32("40490fdb")
fmt.Printf("%.5f\n", val) // Output: 3.14159

// Convert float to hex
hex := convert.Float32ToHex(3.14159265)
fmt.Println(hex) // Output: 40490fdb
```

### Binary Strings

```go
// Binary to integer
val, _ := convert.BinaryToInt8("01111111")
fmt.Println(val) // Output: 127

// Integer to binary
bin := convert.Int8ToBinary(127)
fmt.Println(bin) // Output: 01111111

// With spaces for readability
bin = convert.BinaryToInt8("0111 1111")
```

## Error Handling

The package returns descriptive errors for invalid input:

```go
// Invalid hex character
_, err := convert.HexToInt32("0xGGGG")
// Error: invalid hexadecimal character: 'G' at position 2

// Wrong length
_, err := convert.HexToInt32("12")
// Error: invalid hex string length for type: expected 4 bytes, got 1

// Empty input
_, err := convert.HexToBytes("")
// Error: empty hex string
```

## Performance Considerations

- Uses generic internal helpers to reduce code duplication
- Minimizes allocations by pre-allocating buffers where possible
- Efficient string parsing with single-pass algorithms
- Uses standard library's `encoding/binary` and `encoding/hex` for optimal performance

## Testing

The package has 92% test coverage with comprehensive tests for:

- All supported hex input formats
- Edge cases (empty strings, invalid characters, wrong lengths)
- Boundary values for all numeric types (min, max values)
- Round-trip conversions (value → hex → value)
- Endianness differences
- Special float values (NaN, Infinity)
- Binary string conversions

Run tests:
```bash
go test ./convert/... -v
go test ./convert/... -cover
```

## License

See LICENSE file in the root directory.

## Contributing

Contributions are welcome! Please ensure:
- All tests pass
- Test coverage remains above 80%
- Code follows Go conventions (`go fmt`, `go vet`)
- New features include tests and documentation
