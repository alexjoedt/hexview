// Package models contains data structures for conversion results
// that are shared between the service layer and Wails frontend bindings.
package models

// ConversionResult holds all conversion outputs for hex, integer, binary, and float inputs.
// It provides multiple representations across different data types and endianness formats.
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

// ModbusCombined32 represents a 32-bit value from two consecutive Modbus registers
type ModbusCombined32 struct {
	RegisterStart int    `json:"registerStart"`
	Hex           string `json:"hex"`
	Uint32BE      uint32 `json:"uint32BE"`
	Uint32LE      uint32 `json:"uint32LE"`
	Uint32BADC    uint32 `json:"uint32BADC"`
	Uint32CDAB    uint32 `json:"uint32CDAB"`
	Int32BE       int32  `json:"int32BE"`
	Int32LE       int32  `json:"int32LE"`
	Int32BADC     int32  `json:"int32BADC"`
	Int32CDAB     int32  `json:"int32CDAB"`
	Float32BE     string `json:"float32BE"`
	Float32LE     string `json:"float32LE"`
	Float32BADC   string `json:"float32BADC"`
	Float32CDAB   string `json:"float32CDAB"`
}

// ModbusCombined64 represents a 64-bit value from four consecutive Modbus registers
type ModbusCombined64 struct {
	RegisterStart int    `json:"registerStart"`
	Hex           string `json:"hex"`
	Uint64BE      uint64 `json:"uint64BE"`
	Uint64LE      uint64 `json:"uint64LE"`
	Int64BE       int64  `json:"int64BE"`
	Int64LE       int64  `json:"int64LE"`
	Float64BE     string `json:"float64BE"`
	Float64LE     string `json:"float64LE"`
}

// ModbusResult holds the conversion results for Modbus registers
type ModbusResult struct {
	Registers  []ModbusRegister   `json:"registers"`
	Combined32 []ModbusCombined32 `json:"combined32"`
	Combined64 []ModbusCombined64 `json:"combined64"`
	RawHex     string             `json:"rawHex"`
	ASCII      string             `json:"ascii"`
}
