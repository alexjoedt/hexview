/**
 * Wails API wrapper functions
 */
import { ConvertHex, ConvertInt, ConvertIntAuto, ConvertBinary, ConvertModbusRegisters } from '../../wailsjs/go/main/App.js'

/**
 * Convert hex string to all formats
 */
export async function convertHex(hexInput) {
  return await ConvertHex(hexInput)
}

/**
 * Convert integer to all formats (legacy - explicit type)
 */
export async function convertInt(intInput, intType) {
  return await ConvertInt(intInput, intType)
}

/**
 * Convert integer with auto-detection of compatible types
 */
export async function convertIntAuto(intInput) {
  return await ConvertIntAuto(intInput)
}

/**
 * Convert binary string to all formats
 */
export async function convertBinary(binaryInput) {
  return await ConvertBinary(binaryInput)
}

/**
 * Convert Modbus registers (batch input)
 */
export async function convertModbus(input) {
  return await ConvertModbusRegisters(input)
}

/**
 * Perform conversion based on input mode
 */
export async function convert(inputValue, inputMode, intType) {
  switch (inputMode) {
    case 'hex':
      return await convertHex(inputValue)
    case 'binary':
      return await convertBinary(inputValue)
    case 'int':
      // Use auto-detection for integer mode (intType parameter ignored)
      return await convertIntAuto(inputValue)
    case 'modbus':
      return await convertModbus(inputValue)
    default:
      throw new Error(`Unknown input mode: ${inputMode}`)
  }
}
