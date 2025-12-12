/**
 * Wails API wrapper functions
 */
import { ConvertHex, ConvertInt, ConvertBinary } from '../../wailsjs/go/main/App.js'

/**
 * Convert hex string to all formats
 */
export async function convertHex(hexInput) {
  return await ConvertHex(hexInput)
}

/**
 * Convert integer to all formats
 */
export async function convertInt(intInput, intType) {
  return await ConvertInt(intInput, intType)
}

/**
 * Convert binary string to all formats
 */
export async function convertBinary(binaryInput) {
  return await ConvertBinary(binaryInput)
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
      return await convertInt(inputValue, intType)
    default:
      throw new Error(`Unknown input mode: ${inputMode}`)
  }
}
