/**
 * Constants and data structures for Hexview
 */

// Integer type configurations
export const INT_TYPES = [
  { name: 'INT8', key: 'int8', signed: true, size: 8, endianness: ['BE'] },
  { name: 'INT16', key: 'int16', signed: true, size: 16, endianness: ['BE', 'LE', 'BADC', 'CDAB'] },
  { name: 'INT32', key: 'int32', signed: true, size: 32, endianness: ['BE', 'LE', 'BADC', 'CDAB'] },
  { name: 'INT64', key: 'int64', signed: true, size: 64, endianness: ['BE', 'LE', 'BADC', 'CDAB'] },
  { name: 'UINT8', key: 'uint8', signed: false, size: 8, endianness: ['BE'] },
  { name: 'UINT16', key: 'uint16', signed: false, size: 16, endianness: ['BE', 'LE', 'BADC', 'CDAB'] },
  { name: 'UINT32', key: 'uint32', signed: false, size: 32, endianness: ['BE', 'LE', 'BADC', 'CDAB'] },
  { name: 'UINT64', key: 'uint64', signed: false, size: 64, endianness: ['BE', 'LE', 'BADC', 'CDAB'] }
]

// Float type configurations
export const FLOAT_TYPES = [
  { name: 'FLOAT32', key: 'float32', size: 32, endianness: ['BE', 'LE', 'BADC', 'CDAB'] },
  { name: 'FLOAT64', key: 'float64', size: 64, endianness: ['BE', 'LE', 'BADC', 'CDAB'] }
]

// Integer type options for dropdown
export const INT_TYPE_OPTIONS = [
  { value: 'int8', label: 'INT8' },
  { value: 'int16', label: 'INT16' },
  { value: 'int32', label: 'INT32' },
  { value: 'int64', label: 'INT64' },
  { value: 'uint8', label: 'UINT8' },
  { value: 'uint16', label: 'UINT16' },
  { value: 'uint32', label: 'UINT32' },
  { value: 'uint64', label: 'UINT64' }
]

// Input mode options
export const INPUT_MODES = [
  { value: 'hex', label: 'Hex' },
  { value: 'int', label: 'Integer' },
  { value: 'binary', label: 'Binary' }
]

/**
 * Get result value for a specific type and endianness
 */
export function getResultValue(result, key, endianness) {
  if (!result) return null
  
  const propertyName = endianness === 'BE' 
    ? `${key}${endianness}`
    : `${key}${endianness}`
  
  return result[propertyName]
}

/**
 * Get hex value for a specific type and endianness
 */
export function getResultHex(result, key, endianness) {
  if (!result) return null
  
  const propertyName = `${key}${endianness}Hex`
  return result[propertyName]
}
