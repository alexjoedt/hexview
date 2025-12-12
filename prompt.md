## Implementation Plan - Step by Step

### Phase 1: Foundation (Core Utilities)
**Files:** `convert.go`, `parse.go`

1. **Create package structure**
   - Initialize `convert/convert.go` with package documentation
   - Define error variables
   - Add license header if required

2. **Implement hex string parser**
   - Function: `ParseHex(input string) ([]byte, error)`
   - Strip prefixes (0x, 0X, x, X)
   - Remove whitespace and separators
   - Validate hex characters
   - Convert to byte slice
   - Write comprehensive tests for all input formats

3. **Create internal generic helpers**
   - Generic integer conversion helper
   - Generic float conversion helper
   - Endianness parameter support

### Phase 2: Byte Conversions
**Files:** `convert.go` or `bytes.go`

4. **Implement byte functions**
   - `HexToBytes(hex string) ([]byte, error)`
   - `BytesToHex(b []byte) string`
   - Unit tests with various input formats
   - Edge cases: empty input, single byte, large arrays

### Phase 3: Integer Conversions (Signed)
**Files:** `convert.go` or `int.go`

5. **Implement signed integer functions (big-endian)**
   - Int8, Int16, Int32, Int64 conversions
   - Validate input length matches type size
   - Handle overflow cases
   - Unit tests for each type

6. **Implement signed integer functions (little-endian)**
   - LE variants for all signed types
   - Unit tests

7. **Implement reverse conversions (int to hex)**
   - Both big-endian and little-endian
   - Unit tests

### Phase 4: Integer Conversions (Unsigned)
**Files:** `convert.go` or `uint.go`

8. **Implement unsigned integer functions (big-endian)**
   - Uint8, Uint16, Uint32, Uint64 conversions
   - Unit tests for each type

9. **Implement unsigned integer functions (little-endian)**
   - LE variants for all unsigned types
   - Unit tests

10. **Implement reverse conversions (uint to hex)**
    - Both big-endian and little-endian
    - Unit tests

### Phase 5: Float Conversions
**Files:** `convert.go` or `float.go`

11. **Implement float32 functions**
    - Both endianness variants
    - Use `math.Float32bits` and `math.Float32frombits`
    - Unit tests including special values (NaN, Inf, -Inf, 0.0, -0.0)

12. **Implement float64 functions**
    - Both endianness variants
    - Use `math.Float64bits` and `math.Float64frombits`
    - Unit tests including special values

### Phase 6: Testing & Quality Assurance
**Files:** `convert_test.go`, `*_test.go`

13. **Achieve 80%+ test coverage**
    - Table-driven tests for all functions
    - Edge cases: empty strings, invalid input, overflow
    - Boundary values for each type (min, max values)
    - Round-trip tests (hex → value → hex)

14. **Add example tests**
    - Example usage for each major function
    - Demonstrate various input formats
    - Show error handling patterns

15. **Code quality checks**
    - Run `go fmt`
    - Run `go vet`
    - Run `golint` or `staticcheck`
    - Fix any issues

### Phase 7: Documentation
**Files:** All Go files, `README.md`

16. **Add comprehensive documentation**
    - Package-level documentation
    - Function documentation with examples
    - Document all error conditions
    - Add usage examples in doc comments

17. **Create package README** (optional but recommended)
    - Installation instructions
    - Quick start guide
    - API overview
    - Example usage
    - Performance considerations

### Phase 8: Final Polish
**Files:** All files

18. **Performance considerations**
    - Review allocations
    - Optimize hot paths if needed
    - Consider adding benchmark tests

19. **API review**
    - Ensure consistent naming
    - Verify all functions follow Go conventions
    - Check that public API is minimal and clear

20. **Integration testing**
    - Test package in hexview application
    - Verify it works as external import
    - Validate versioning strategy (if publishing)

## Testing Requirements

### Coverage Target: 80%+

**Test Categories:**

1. **Unit Tests**
   - Each function must have tests
   - Cover happy paths and error cases
   - Use table-driven test pattern

2. **Edge Cases**
   - Empty input
   - Invalid hex characters
   - Wrong length for type
   - Overflow/underflow values
   - Special float values (NaN, Inf)

3. **Format Compatibility**
   - Test all supported hex input formats
   - Mixed case
   - Various separators
   - Different prefix styles

4. **Round-Trip Tests**
   - Value → Hex → Value should be identical
   - Test for all types and endianness

5. **Example Tests**
   - Runnable examples in documentation
   - Show real-world usage patterns

## Code Organization

```
convert/
├── convert.go         # Main package file with documentation
├── parse.go           # Hex string parsing utilities
├── bytes.go           # Byte conversion functions (optional split)
├── int.go             # Integer conversion functions (optional split)
├── uint.go            # Unsigned integer functions (optional split)
├── float.go           # Float conversion functions (optional split)
├── errors.go          # Error definitions
├── convert_test.go    # Main test file
├── parse_test.go      # Parser tests
├── example_test.go    # Example tests for documentation
└── README.md          # Package documentation (optional)
```

*Note: File splitting is optional - can keep all in `convert.go` if preferred for simplicity*

## Success Criteria

✅ All specified hex input formats are parsed correctly  
✅ All numeric types (int8-64, uint8-64, float32/64) are supported  
✅ Both big-endian (default) and little-endian conversions work  
✅ Bidirectional conversions (hex↔numeric) are implemented  
✅ Idiomatic Go error handling throughout  
✅ 80%+ test coverage achieved  
✅ All tests pass (`go test ./...`)  
✅ Code passes `go vet` and linting tools  
✅ Comprehensive documentation with examples  
✅ Package works both internally and as external library  

## Additional Considerations

### Performance
- Minimize allocations in hot paths
- Reuse buffers where possible
- Consider pre-allocating slices with known sizes

### Security
- Validate all input thoroughly
- Prevent panic conditions
- Handle malformed input gracefully

### Maintainability
- Clear function names
- Consistent patterns across similar functions
- Well-organized code structure
- Comprehensive comments

### Future Extensibility
- Consider adding octal/binary support later
- Could add format options (uppercase/lowercase output)
- Might add prefix control for output format

## Implementation Order Priority

**High Priority (Core Functionality):**
1. Hex string parser
2. Byte conversions
3. Integer conversions (signed & unsigned)
4. Basic test coverage

**Medium Priority (Complete Feature Set):**
5. Float conversions
6. Little-endian variants
7. 80% test coverage
8. Documentation

**Low Priority (Polish):**
9. Performance optimization
10. Extended examples
11. Package README

---

## Getting Started

Begin with Phase 1 by creating the foundational parsing functionality. This is the most critical component as all other conversions depend on reliable hex string parsing. Write tests immediately after implementing each function to ensure correctness before moving forward.
EOF
cat /tmp/convert_package_prompt.md

Let me create the prompt in a better way and copy it to clipboard: 

Ran terminal command: cat << 'PROMPTEOF' > /tmp/convert_prompt.md
# Production-Ready Go Convert Package Implementation Plan

## Project Overview

Create a robust, production-ready Go package (`github.com/alexjoedt/hexview/convert`) that provides comprehensive hexadecimal conversion utilities with support for multiple numeric types, endianness configurations, and flexible input parsing.

## Package Scope & Requirements

### Target Audience
- External library users (public API package)
- Internal hexview application components
- Developers needing reliable hex-to-numeric conversions

### Core Capabilities
1. **Flexible Hex Input Parsing**: Parse various hex string formats
2. **Multi-Type Conversions**: Support bytes, integers (signed/unsigned), and floats
3. **Bidirectional Operations**: Convert to and from hex representations
4. **Endianness Control**: Big-endian (default) and little-endian support
5. **Idiomatic Go Design**: Error handling, clear APIs, comprehensive documentation

## Detailed Requirements

### 1. Hex String Parsing

The package must parse and normalize various hex input formats:

**Supported Input Formats:**
- `"0x123456789ABCDEF"` - Standard 0x prefix
- `"04 ab"` - Space-separated bytes
- `"11abcd"` - Continuous hex string
- `"0xab 0xff"` - Multiple 0x-prefixed values
- `"xAB xCF"` - 'x' prefix without '0'
- Mixed case (both uppercase and lowercase)
- With or without separators (spaces, commas, colons)

**Parsing Requirements:**
- Strip common prefixes: `0x`, `0X`, `x`, `X`
- Handle whitespace and common separators
- Validate hex characters (0-9, a-f, A-F)
- Return descriptive errors for invalid input
- Internal helper: `ParseHex(input string) ([]byte, error)`

### 2. Supported Numeric Types

#### Integers (Signed & Unsigned)
- `int8` / `uint8` (8-bit)
- `int16` / `uint16` (16-bit)
- `int32` / `uint32` (32-bit)
- `int64` / `uint64` (64-bit)

#### Floating Point
- `float32` (32-bit IEEE 754)
- `float64` (64-bit IEEE 754)

#### Raw Bytes
- `[]byte` - byte slice representation

### 3. Endianness Support

**Default Behavior:**
- Big-endian as default for all conversions
- Explicit functions for little-endian conversions

**Function Naming Convention:**
- `HexToInt32(hex string) (int32, error)` - Big-endian
- `HexToInt32LE(hex string) (int32, error)` - Little-endian
- Same pattern for all numeric types

### 4. API Design

#### Hex to Numeric Conversions

```go
// Bytes
func HexToBytes(hex string) ([]byte, error)
func BytesToHex(b []byte) string

// Signed Integers (Big-Endian)
func HexToInt8(hex string) (int8, error)
func HexToInt16(hex string) (int16, error)
func HexToInt32(hex string) (int32, error)
func HexToInt64(hex string) (int64, error)

// Signed Integers (Little-Endian)
func HexToInt8LE(hex string) (int8, error)
func HexToInt16LE(hex string) (int16, error)
func HexToInt32LE(hex string) (int32, error)
func HexToInt64LE(hex string) (int64, error)

// Unsigned Integers (Big-Endian)
func HexToUint8(hex string) (uint8, error)
func HexToUint16(hex string) (uint16, error)
func HexToUint32(hex string) (uint32, error)
func HexToUint64(hex string) (uint64, error)

// Unsigned Integers (Little-Endian)
func HexToUint8LE(hex string) (uint8, error)
func HexToUint16LE(hex string) (uint16, error)
func HexToUint32LE(hex string) (uint32, error)
func HexToUint64LE(hex string) (uint64, error)

// Floats (Big-Endian)
func HexToFloat32(hex string) (float32, error)
func HexToFloat64(hex string) (float64, error)

// Floats (Little-Endian)
func HexToFloat32LE(hex string) (float32, error)
func HexToFloat64LE(hex string) (float64, error)
```

#### Numeric to Hex Conversions

```go
// Signed Integers (Big-Endian)
func Int8ToHex(n int8) string
func Int16ToHex(n int16) string
func Int32ToHex(n int32) string
func Int64ToHex(n int64) string

// Signed Integers (Little-Endian)
func Int8ToHexLE(n int8) string
func Int16ToHexLE(n int16) string
func Int32ToHexLE(n int32) string
func Int64ToHexLE(n int64) string

// Unsigned Integers (Big-Endian)
func Uint8ToHex(n uint8) string
func Uint16ToHex(n uint16) string
func Uint32ToHex(n uint32) string
func Uint64ToHex(n uint64) string

// Unsigned Integers (Little-Endian)
func Uint8ToHexLE(n uint8) string
func Uint16ToHexLE(n uint16) string
func Uint32ToHexLE(n uint32) string
func Uint64ToHexLE(n uint64) string

// Floats (Big-Endian)
func Float32ToHex(f float32) string
func Float64ToHex(f float64) string

// Floats (Little-Endian)
func Float32ToHexLE(f float32) string
func Float64ToHexLE(f float64) string
```

### 5. Internal Implementation Guidelines

**Use Generics Internally:**
- Create private generic helper functions to reduce code duplication
- Example: `func hexToInt[T Integer](hex string, endian binary.ByteOrder) (T, error)`
- Keep public API simple with explicit function names

**Error Handling:**
- Return `error` for all operations that can fail
- Use descriptive error messages (e.g., "invalid hex character 'g' at position 5")
- Wrap errors with context using `fmt.Errorf`

**Standard Library Usage:**
- Use `encoding/binary` for byte order operations
- Use `encoding/hex` for base conversions where appropriate
- Use `math` for float bit conversions

### 6. Error Types & Messages

Define clear error categories:

```go
var (
    ErrInvalidHexChar   = errors.New("invalid hexadecimal character")
    ErrInvalidLength    = errors.New("invalid hex string length for type")
    ErrOverflow         = errors.New("value overflow for target type")
    ErrEmptyInput       = errors.New("empty hex string")
)
```

## Implementation Plan - Step by Step

### Phase 1: Foundation (Core Utilities)
**Files:** `convert.go`, `parse.go`

1. **Create package structure**
   - Initialize `convert/convert.go` with package documentation
   - Define error variables
   - Add license header if required

2. **Implement hex string parser**
   - Function: `ParseHex(input string) ([]byte, error)`
   - Strip prefixes (0x, 0X, x, X)
   - Remove whitespace and separators
   - Validate hex characters
   - Convert to byte slice
   - Write comprehensive tests for all input formats

3. **Create internal generic helpers**
   - Generic integer conversion helper
   - Generic float conversion helper
   - Endianness parameter support

### Phase 2: Byte Conversions
**Files:** `convert.go` or `bytes.go`

4. **Implement byte functions**
   - `HexToBytes(hex string) ([]byte, error)`
   - `BytesToHex(b []byte) string`
   - Unit tests with various input formats
   - Edge cases: empty input, single byte, large arrays

### Phase 3: Integer Conversions (Signed)
**Files:** `convert.go` or `int.go`

5. **Implement signed integer functions (big-endian)**
   - Int8, Int16, Int32, Int64 conversions
   - Validate input length matches type size
   - Handle overflow cases
   - Unit tests for each type

6. **Implement signed integer functions (little-endian)**
   - LE variants for all signed types
   - Unit tests

7. **Implement reverse conversions (int to hex)**
   - Both big-endian and little-endian
   - Unit tests

### Phase 4: Integer Conversions (Unsigned)
**Files:** `convert.go` or `uint.go`

8. **Implement unsigned integer functions (big-endian)**
   - Uint8, Uint16, Uint32, Uint64 conversions
   - Unit tests for each type

9. **Implement unsigned integer functions (little-endian)**
   - LE variants for all unsigned types
   - Unit tests

10. **Implement reverse conversions (uint to hex)**
    - Both big-endian and little-endian
    - Unit tests

### Phase 5: Float Conversions
**Files:** `convert.go` or `float.go`

11. **Implement float32 functions**
    - Both endianness variants
    - Use `math.Float32bits` and `math.Float32frombits`
    - Unit tests including special values (NaN, Inf, -Inf, 0.0, -0.0)

12. **Implement float64 functions**
    - Both endianness variants
    - Use `math.Float64bits` and `math.Float64frombits`
    - Unit tests including special values

### Phase 6: Testing & Quality Assurance
**Files:** `convert_test.go`, `*_test.go`

13. **Achieve 80%+ test coverage**
    - Table-driven tests for all functions
    - Edge cases: empty strings, invalid input, overflow
    - Boundary values for each type (min, max values)
    - Round-trip tests (hex → value → hex)

14. **Add example tests**
    - Example usage for each major function
    - Demonstrate various input formats
    - Show error handling patterns

15. **Code quality checks**
    - Run `go fmt`
    - Run `go vet`
    - Run `golint` or `staticcheck`
    - Fix any issues

### Phase 7: Documentation
**Files:** All Go files, `README.md`

16. **Add comprehensive documentation**
    - Package-level documentation
    - Function documentation with examples
    - Document all error conditions
    - Add usage examples in doc comments

17. **Create package README** (optional but recommended)
    - Installation instructions
    - Quick start guide
    - API overview
    - Example usage
    - Performance considerations

### Phase 8: Final Polish
**Files:** All files

18. **Performance considerations**
    - Review allocations
    - Optimize hot paths if needed
    - Consider adding benchmark tests

19. **API review**
    - Ensure consistent naming
    - Verify all functions follow Go conventions
    - Check that public API is minimal and clear

20. **Integration testing**
    - Test package in hexview application
    - Verify it works as external import
    - Validate versioning strategy (if publishing)

## Testing Requirements

### Coverage Target: 80%+

**Test Categories:**

1. **Unit Tests**
   - Each function must have tests
   - Cover happy paths and error cases
   - Use table-driven test pattern

2. **Edge Cases**
   - Empty input
   - Invalid hex characters
   - Wrong length for type
   - Overflow/underflow values
   - Special float values (NaN, Inf)

3. **Format Compatibility**
   - Test all supported hex input formats
   - Mixed case
   - Various separators
   - Different prefix styles

4. **Round-Trip Tests**
   - Value → Hex → Value should be identical
   - Test for all types and endianness

5. **Example Tests**
   - Runnable examples in documentation
   - Show real-world usage patterns

## Code Organization

```
convert/
├── convert.go         # Main package file with documentation
├── parse.go           # Hex string parsing utilities
├── bytes.go           # Byte conversion functions (optional split)
├── int.go             # Integer conversion functions (optional split)
├── uint.go            # Unsigned integer functions (optional split)
├── float.go           # Float conversion functions (optional split)
├── errors.go          # Error definitions
├── convert_test.go    # Main test file
├── parse_test.go      # Parser tests
├── example_test.go    # Example tests for documentation
└── README.md          # Package documentation (optional)
```

*Note: File splitting is optional - can keep all in `convert.go` if preferred for simplicity*

## Success Criteria

✅ All specified hex input formats are parsed correctly  
✅ All numeric types (int8-64, uint8-64, float32/64) are supported  
✅ Both big-endian (default) and little-endian conversions work  
✅ Bidirectional conversions (hex↔numeric) are implemented  
✅ Idiomatic Go error handling throughout  
✅ 80%+ test coverage achieved  
✅ All tests pass (`go test ./...`)  
✅ Code passes `go vet` and linting tools  
✅ Comprehensive documentation with examples  
✅ Package works both internally and as external library  

## Additional Considerations

### Performance
- Minimize allocations in hot paths
- Reuse buffers where possible
- Consider pre-allocating slices with known sizes

### Security
- Validate all input thoroughly
- Prevent panic conditions
- Handle malformed input gracefully

### Maintainability
- Clear function names
- Consistent patterns across similar functions
- Well-organized code structure
- Comprehensive comments

### Future Extensibility
- Consider adding octal/binary support later
- Could add format options (uppercase/lowercase output)
- Might add prefix control for output format

## Implementation Order Priority

**High Priority (Core Functionality):**
1. Hex string parser
2. Byte conversions
3. Integer conversions (signed & unsigned)
4. Basic test coverage

**Medium Priority (Complete Feature Set):**
5. Float conversions
6. Little-endian variants
7. 80% test coverage
8. Documentation

**Low Priority (Polish):**
9. Performance optimization
10. Extended examples
11. Package README

---

## Getting Started

Begin with Phase 1 by creating the foundational parsing functionality. This is the most critical component as all other conversions depend on reliable hex string parsing. Write tests immediately after implementing each function to ensure correctness before moving forward.
PROMPTEOF
