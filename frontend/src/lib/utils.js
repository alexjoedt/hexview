/**
 * Utility functions for Hexview
 */

/**
 * Format value for display, showing placeholder for null/undefined
 */
export function formatValue(value) {
  return value !== null && value !== undefined ? String(value) : '—'
}

/**
 * Check if value exists (not null/undefined)
 */
export function hasValue(value) {
  return value !== null && value !== undefined
}

/**
 * Copy text to clipboard
 * @returns {Promise<void>}
 */
export async function copyToClipboard(text) {
  await navigator.clipboard.writeText(String(text))
}
