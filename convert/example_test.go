package convert_test

import (
	"fmt"

	"hexview/convert"
)

// Example demonstrates basic hex string parsing with various formats.
func Example() {
	// Parse hex with 0x prefix
	bytes1, _ := convert.HexToBytes("0x48656c6c6f")
	fmt.Printf("With prefix: %s\n", string(bytes1))

	// Parse hex with spaces
	bytes2, _ := convert.HexToBytes("48 65 6c 6c 6f")
	fmt.Printf("With spaces: %s\n", string(bytes2))

	// Parse hex continuous
	bytes3, _ := convert.HexToBytes("48656c6c6f")
	fmt.Printf("Continuous: %s\n", string(bytes3))

	// Output:
	// With prefix: Hello
	// With spaces: Hello
	// Continuous: Hello
}

// ExampleHexToInt32 shows how to convert hex strings to 32-bit integers.
func ExampleHexToInt32() {
	// Big-endian conversion
	val, _ := convert.HexToInt32("7fffffff")
	fmt.Printf("Max int32: %d\n", val)

	val, _ = convert.HexToInt32("ffffffff")
	fmt.Printf("Negative: %d\n", val)

	// Output:
	// Max int32: 2147483647
	// Negative: -1
}

// ExampleHexToInt32LE shows little-endian conversion.
func ExampleHexToInt32LE() {
	// Same bytes, different byte order
	be, _ := convert.HexToInt32("12345678")
	le, _ := convert.HexToInt32LE("12345678")

	fmt.Printf("Big-endian: 0x%08x\n", be)
	fmt.Printf("Little-endian: 0x%08x\n", le)

	// Output:
	// Big-endian: 0x12345678
	// Little-endian: 0x78563412
}

// ExampleInt32ToHex shows converting integers to hex strings.
func ExampleInt32ToHex() {
	hex := convert.Int32ToHex(2147483647)
	fmt.Printf("Max int32 as hex: %s\n", hex)

	hex = convert.Int32ToHex(-1)
	fmt.Printf("Negative as hex: %s\n", hex)

	// Output:
	// Max int32 as hex: 7fffffff
	// Negative as hex: ffffffff
}

// ExampleHexToFloat32 shows converting hex to floating point.
func ExampleHexToFloat32() {
	val, _ := convert.HexToFloat32("3f800000")
	fmt.Printf("Float value: %.1f\n", val)

	val, _ = convert.HexToFloat32("40490fdb")
	fmt.Printf("Pi approximation: %.5f\n", val)

	// Output:
	// Float value: 1.0
	// Pi approximation: 3.14159
}

// ExampleFloat32ToHex shows converting floats to hex.
func ExampleFloat32ToHex() {
	hex := convert.Float32ToHex(1.0)
	fmt.Printf("1.0 as hex: %s\n", hex)

	hex = convert.Float32ToHex(-1.0)
	fmt.Printf("-1.0 as hex: %s\n", hex)

	// Output:
	// 1.0 as hex: 3f800000
	// -1.0 as hex: bf800000
}

// ExampleBinaryToInt8 shows converting binary strings to integers.
func ExampleBinaryToInt8() {
	val, _ := convert.BinaryToInt8("01111111")
	fmt.Printf("Max int8: %d\n", val)

	val, _ = convert.BinaryToInt8("11111111")
	fmt.Printf("Negative: %d\n", val)

	// Binary with spaces
	val, _ = convert.BinaryToInt8("0111 1111")
	fmt.Printf("With spaces: %d\n", val)

	// Output:
	// Max int8: 127
	// Negative: -1
	// With spaces: 127
}

// ExampleInt8ToBinary shows converting integers to binary strings.
func ExampleInt8ToBinary() {
	bin := convert.Int8ToBinary(127)
	fmt.Printf("127 as binary: %s\n", bin)

	bin = convert.Int8ToBinary(-1)
	fmt.Printf("-1 as binary: %s\n", bin)

	// Output:
	// 127 as binary: 01111111
	// -1 as binary: 11111111
}

// ExampleBytesToHex shows converting byte slices to hex strings.
func ExampleBytesToHex() {
	hex := convert.BytesToHex([]byte("Hello"))
	fmt.Printf("Hello as hex: %s\n", hex)

	hex = convert.BytesToHex([]byte{0xff, 0x00, 0xaa})
	fmt.Printf("Bytes as hex: %s\n", hex)

	// Output:
	// Hello as hex: 48656c6c6f
	// Bytes as hex: ff00aa
}

// ExampleBytesToBinary shows converting bytes to binary representation.
func ExampleBytesToBinary() {
	bin := convert.BytesToBinary([]byte{0x0a})
	fmt.Printf("0x0a as binary: %s\n", bin)

	bin = convert.BytesToBinary([]byte{0xff, 0x00})
	fmt.Printf("Multiple bytes: %s\n", bin)

	// Output:
	// 0x0a as binary: 00001010
	// Multiple bytes: 11111111 00000000
}

// ExampleHexToUint64 shows working with large unsigned integers.
func ExampleHexToUint64() {
	val, _ := convert.HexToUint64("ffffffffffffffff")
	fmt.Printf("Max uint64: %d\n", val)

	val, _ = convert.HexToUint64("0000000000000064")
	fmt.Printf("Decimal 100: %d\n", val)

	// Output:
	// Max uint64: 18446744073709551615
	// Decimal 100: 100
}

// ExampleParseHex shows the flexible hex parser with various formats.
func ExampleParseHex() {
	// Different separator styles
	formats := []string{
		"0xAB 0xCD 0xEF",
		"AB:CD:EF",
		"AB,CD,EF",
		"ABCDEF",
		"xAB xCD xEF",
	}

	for _, format := range formats {
		bytes, _ := convert.ParseHex(format)
		fmt.Printf("%s\n", convert.BytesToHex(bytes))
	}

	// Output:
	// abcdef
	// abcdef
	// abcdef
	// abcdef
	// abcdef
}
