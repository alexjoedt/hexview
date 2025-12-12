# Hexview Frontend Implementation Plan

## Project Overview
Build a single-view Svelte frontend for Hexview that provides real-time hexadecimal conversion with an excellent UX for network engineers. The interface features a single input field with comprehensive output display of all conversion types.

## Design Principles
- **Real-time conversion**: Auto-convert as user types with debounce (300ms)
- **Forgiving input**: Accept multiple hex formats (0x prefix, spaces, colons, commas)
- **Minimalist design**: Clean, focused interface with dark/light theme support
- **Copy-friendly**: One-click copy for any output value
- **Clear feedback**: Highlight different data types, show errors gracefully

## Step-by-Step Implementation Plan

### Step 1: Backend Integration - Add Conversion Methods to App
**File**: `app.go`

Add methods to expose all conversion functions to the frontend:

```go
// ConversionResult holds all conversion outputs
type ConversionResult struct {
    // Signed Integers - Big Endian
    Int8BE    *int8    `json:"int8BE,omitempty"`
    Int16BE   *int16   `json:"int16BE,omitempty"`
    Int32BE   *int32   `json:"int32BE,omitempty"`
    Int64BE   *int64   `json:"int64BE,omitempty"`
    
    // Signed Integers - Little Endian
    Int16LE   *int16   `json:"int16LE,omitempty"`
    Int32LE   *int32   `json:"int32LE,omitempty"`
    Int64LE   *int64   `json:"int64LE,omitempty"`
    
    // Unsigned Integers - Big Endian
    Uint8BE   *uint8   `json:"uint8BE,omitempty"`
    Uint16BE  *uint16  `json:"uint16BE,omitempty"`
    Uint32BE  *uint32  `json:"uint32BE,omitempty"`
    Uint64BE  *uint64  `json:"uint64BE,omitempty"`
    
    // Unsigned Integers - Little Endian
    Uint16LE  *uint16  `json:"uint16LE,omitempty"`
    Uint32LE  *uint32  `json:"uint32LE,omitempty"`
    Uint64LE  *uint64  `json:"uint64LE,omitempty"`
    
    // Floating Point
    Float32BE *float32 `json:"float32BE,omitempty"`
    Float64BE *float64 `json:"float64BE,omitempty"`
    Float32LE *float32 `json:"float32LE,omitempty"`
    Float64LE *float64 `json:"float64LE,omitempty"`
    
    // Binary Representations
    Binary    string   `json:"binary,omitempty"`
    Bytes     string   `json:"bytes,omitempty"`
}

// ConvertHex performs all possible conversions on the hex input
func (a *App) ConvertHex(hexInput string) (*ConversionResult, error) {
    // Validate input is not empty
    if hexInput == "" {
        return nil, fmt.Errorf("empty input")
    }
    
    result := &ConversionResult{}
    
    // Convert to bytes first to get binary representation
    bytes, err := convert.HexToBytes(hexInput)
    if err != nil {
        return nil, fmt.Errorf("invalid hex input: %w", err)
    }
    
    result.Binary = convert.BytesToBinary(bytes)
    result.Bytes = convert.BytesToHex(bytes)
    
    // Try all signed integer conversions (Big Endian)
    if v, err := convert.HexToInt8(hexInput); err == nil {
        result.Int8BE = &v
    }
    if v, err := convert.HexToInt16(hexInput); err == nil {
        result.Int16BE = &v
    }
    if v, err := convert.HexToInt32(hexInput); err == nil {
        result.Int32BE = &v
    }
    if v, err := convert.HexToInt64(hexInput); err == nil {
        result.Int64BE = &v
    }
    
    // Try all signed integer conversions (Little Endian)
    if v, err := convert.HexToInt16LE(hexInput); err == nil {
        result.Int16LE = &v
    }
    if v, err := convert.HexToInt32LE(hexInput); err == nil {
        result.Int32LE = &v
    }
    if v, err := convert.HexToInt64LE(hexInput); err == nil {
        result.Int64LE = &v
    }
    
    // Try all unsigned integer conversions (Big Endian)
    if v, err := convert.HexToUint8(hexInput); err == nil {
        result.Uint8BE = &v
    }
    if v, err := convert.HexToUint16(hexInput); err == nil {
        result.Uint16BE = &v
    }
    if v, err := convert.HexToUint32(hexInput); err == nil {
        result.Uint32BE = &v
    }
    if v, err := convert.HexToUint64(hexInput); err == nil {
        result.Uint64BE = &v
    }
    
    // Try all unsigned integer conversions (Little Endian)
    if v, err := convert.HexToUint16LE(hexInput); err == nil {
        result.Uint16LE = &v
    }
    if v, err := convert.HexToUint32LE(hexInput); err == nil {
        result.Uint32LE = &v
    }
    if v, err := convert.HexToUint64LE(hexInput); err == nil {
        result.Uint64LE = &v
    }
    
    // Try float conversions
    if v, err := convert.HexToFloat32(hexInput); err == nil {
        result.Float32BE = &v
    }
    if v, err := convert.HexToFloat64(hexInput); err == nil {
        result.Float64BE = &v
    }
    if v, err := convert.HexToFloat32LE(hexInput); err == nil {
        result.Float32LE = &v
    }
    if v, err := convert.HexToFloat64LE(hexInput); err == nil {
        result.Float64LE = &v
    }
    
    return result, nil
}
```

**Acceptance Criteria**:
- `ConvertHex` method successfully exports to frontend
- Returns all possible conversions based on input length
- Handles invalid input gracefully with clear error messages
- Uses pointers for optional values (nil when conversion not applicable)

---

### Step 2: Frontend State Management
**File**: `frontend/src/App.svelte`

Set up reactive state and import the conversion function:

```javascript
<script>
  import { ConvertHex } from '../wailsjs/go/main/App.js'
  
  // State
  let hexInput = ''
  let result = null
  let error = null
  let isLoading = false
  let debounceTimer = null
  
  // Theme state (default to system preference)
  let darkMode = window.matchMedia('(prefers-color-scheme: dark)').matches
  
  // Reactive conversion with debounce
  $: {
    if (hexInput.trim() === '') {
      result = null
      error = null
    } else {
      debouncedConvert(hexInput)
    }
  }
  
  function debouncedConvert(input) {
    clearTimeout(debounceTimer)
    isLoading = true
    
    debounceTimer = setTimeout(() => {
      ConvertHex(input)
        .then(res => {
          result = res
          error = null
          isLoading = false
        })
        .catch(err => {
          error = err
          result = null
          isLoading = false
        })
    }, 300) // 300ms debounce
  }
  
  function toggleTheme() {
    darkMode = !darkMode
  }
  
  async function copyToClipboard(text) {
    try {
      await navigator.clipboard.writeText(String(text))
      // Optional: Show toast notification
    } catch (err) {
      console.error('Failed to copy:', err)
    }
  }
</script>
```

**Acceptance Criteria**:
- Input changes trigger debounced conversion (300ms delay)
- Empty input clears results immediately
- Loading state shows during conversion
- Theme preference persists in session
- Copy function handles all data types correctly

---

### Step 3: Input Component with Format Hints
**File**: `frontend/src/App.svelte` (template section)

Create the input area with helpful placeholders:

```svelte
<main class:dark={darkMode}>
  <div class="container">
    <!-- Header -->
    <header>
      <h1>Hexview</h1>
      <button class="theme-toggle" on:click={toggleTheme} aria-label="Toggle theme">
        {darkMode ? '☀️' : '🌙'}
      </button>
    </header>
    
    <!-- Input Section -->
    <section class="input-section">
      <label for="hex-input">Hexadecimal Input</label>
      <input
        id="hex-input"
        type="text"
        bind:value={hexInput}
        placeholder="Enter hex (e.g., 0x1A2B, 1A 2B 3C, 1A:2B:3C)"
        autocomplete="off"
        autofocus
        class:error={error}
      />
      <div class="input-hint">
        Accepts: 0x prefix, spaces, colons, commas • Case insensitive
      </div>
      
      {#if error}
        <div class="error-message">
          ⚠️ {error}
        </div>
      {/if}
      
      {#if !hexInput && !result}
        <div class="help-text">
          <h3>Quick Start</h3>
          <p>Enter a hexadecimal value to see all possible conversions.</p>
          <div class="examples">
            <strong>Examples:</strong>
            <code>0xFF</code> <code>7F FF FF FF</code> <code>1A:2B:3C:4D</code>
          </div>
        </div>
      {/if}
    </section>
    
    <!-- Results Section -->
    {#if result && !error}
      <section class="results">
        <div class="loading-indicator" class:visible={isLoading}>
          Converting...
        </div>
        
        <!-- Results will be added in next steps -->
      </section>
    {/if}
  </div>
</main>
```

**Acceptance Criteria**:
- Input accepts all hex formats without validation errors
- Placeholder shows example formats
- Hint text explains accepted formats
- Help text appears when no input is present
- Error messages display clearly below input
- Auto-focus on input for immediate use
- Theme toggle button visible and functional

---

### Step 4: Results Display - Signed Integers Group
**File**: `frontend/src/App.svelte` (add to results section)

Create organized output for signed integers:

```svelte
<!-- Signed Integers - Big Endian -->
{#if result.int8BE !== null || result.int16BE !== null || result.int32BE !== null || result.int64BE !== null}
  <div class="result-group signed-integers">
    <h3 class="group-title">Signed Integers (Big Endian)</h3>
    <div class="result-grid">
      {#if result.int8BE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT8</span>
          <span class="value">{result.int8BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int8BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.int16BE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT16</span>
          <span class="value">{result.int16BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int16BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.int32BE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT32</span>
          <span class="value">{result.int32BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int32BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.int64BE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT64</span>
          <span class="value">{result.int64BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int64BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}

<!-- Signed Integers - Little Endian -->
{#if result.int16LE !== null || result.int32LE !== null || result.int64LE !== null}
  <div class="result-group signed-integers">
    <h3 class="group-title">Signed Integers (Little Endian)</h3>
    <div class="result-grid">
      {#if result.int16LE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT16 LE</span>
          <span class="value">{result.int16LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int16LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.int32LE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT32 LE</span>
          <span class="value">{result.int32LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int32LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.int64LE !== null}
        <div class="result-item">
          <span class="type-label int-signed">INT64 LE</span>
          <span class="value">{result.int64LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.int64LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}
```

**Acceptance Criteria**:
- Groups appear only when values exist for that group
- Copy button works for each value
- Type labels clearly identify data type
- Visual distinction for signed integers (color coding)

---

### Step 5: Results Display - Unsigned Integers Group
**File**: `frontend/src/App.svelte` (add to results section)

Add unsigned integer outputs:

```svelte
<!-- Unsigned Integers - Big Endian -->
{#if result.uint8BE !== null || result.uint16BE !== null || result.uint32BE !== null || result.uint64BE !== null}
  <div class="result-group unsigned-integers">
    <h3 class="group-title">Unsigned Integers (Big Endian)</h3>
    <div class="result-grid">
      {#if result.uint8BE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT8</span>
          <span class="value">{result.uint8BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint8BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.uint16BE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT16</span>
          <span class="value">{result.uint16BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint16BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.uint32BE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT32</span>
          <span class="value">{result.uint32BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint32BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.uint64BE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT64</span>
          <span class="value">{result.uint64BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint64BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}

<!-- Unsigned Integers - Little Endian -->
{#if result.uint16LE !== null || result.uint32LE !== null || result.uint64LE !== null}
  <div class="result-group unsigned-integers">
    <h3 class="group-title">Unsigned Integers (Little Endian)</h3>
    <div class="result-grid">
      {#if result.uint16LE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT16 LE</span>
          <span class="value">{result.uint16LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint16LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.uint32LE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT32 LE</span>
          <span class="value">{result.uint32LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint32LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.uint64LE !== null}
        <div class="result-item">
          <span class="type-label int-unsigned">UINT64 LE</span>
          <span class="value">{result.uint64LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.uint64LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}
```

**Acceptance Criteria**:
- Unsigned integers visually distinct from signed (different color)
- All copy buttons functional
- Groups collapse when no values present

---

### Step 6: Results Display - Floating Point & Binary Groups
**File**: `frontend/src/App.svelte` (add to results section)

Add float and binary representations:

```svelte
<!-- Floating Point Numbers -->
{#if result.float32BE !== null || result.float64BE !== null || result.float32LE !== null || result.float64LE !== null}
  <div class="result-group float-numbers">
    <h3 class="group-title">Floating Point</h3>
    <div class="result-grid">
      {#if result.float32BE !== null}
        <div class="result-item">
          <span class="type-label float">FLOAT32 BE</span>
          <span class="value">{result.float32BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.float32BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.float64BE !== null}
        <div class="result-item">
          <span class="type-label float">FLOAT64 BE</span>
          <span class="value">{result.float64BE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.float64BE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.float32LE !== null}
        <div class="result-item">
          <span class="type-label float">FLOAT32 LE</span>
          <span class="value">{result.float32LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.float32LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.float64LE !== null}
        <div class="result-item">
          <span class="type-label float">FLOAT64 LE</span>
          <span class="value">{result.float64LE}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.float64LE)} title="Copy">
            📋
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}

<!-- Binary & Byte Representations -->
{#if result.binary || result.bytes}
  <div class="result-group binary-group">
    <h3 class="group-title">Binary Representations</h3>
    <div class="result-grid">
      {#if result.bytes}
        <div class="result-item full-width">
          <span class="type-label bytes">HEX BYTES</span>
          <span class="value mono">{result.bytes}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.bytes)} title="Copy">
            📋
          </button>
        </div>
      {/if}
      
      {#if result.binary}
        <div class="result-item full-width">
          <span class="type-label binary">BINARY</span>
          <span class="value mono">{result.binary}</span>
          <button class="copy-btn" on:click={() => copyToClipboard(result.binary)} title="Copy">
            📋
          </button>
        </div>
      {/if}
    </div>
  </div>
{/if}
```

**Acceptance Criteria**:
- Float values display with appropriate precision
- Binary strings are monospaced for readability
- Full-width items for long binary/hex strings
- All values copyable

---

### Step 7: Styling - Base & Theme Variables
**File**: `frontend/src/App.svelte` (style section)

Create CSS custom properties for theming:

```css
<style>
  :root {
    /* Light Theme Colors */
    --bg-primary: #ffffff;
    --bg-secondary: #f5f5f5;
    --bg-tertiary: #e8e8e8;
    --text-primary: #1a1a1a;
    --text-secondary: #666666;
    --text-tertiary: #999999;
    --border-color: #d0d0d0;
    --shadow: rgba(0, 0, 0, 0.1);
    
    /* Type Colors - Light */
    --color-int-signed: #2563eb;
    --color-int-unsigned: #059669;
    --color-float: #dc2626;
    --color-binary: #7c3aed;
    --color-bytes: #ea580c;
    
    /* Status Colors */
    --color-error: #ef4444;
    --color-success: #10b981;
    --color-warning: #f59e0b;
    
    /* Spacing */
    --spacing-xs: 0.25rem;
    --spacing-sm: 0.5rem;
    --spacing-md: 1rem;
    --spacing-lg: 1.5rem;
    --spacing-xl: 2rem;
    
    /* Border Radius */
    --radius-sm: 0.25rem;
    --radius-md: 0.5rem;
    --radius-lg: 0.75rem;
    
    /* Typography */
    --font-mono: 'SF Mono', 'Monaco', 'Consolas', monospace;
    --font-sans: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  }
  
  /* Dark Theme */
  .dark {
    --bg-primary: #1a1a1a;
    --bg-secondary: #2a2a2a;
    --bg-tertiary: #3a3a3a;
    --text-primary: #f0f0f0;
    --text-secondary: #b0b0b0;
    --text-tertiary: #808080;
    --border-color: #404040;
    --shadow: rgba(0, 0, 0, 0.4);
    
    /* Type Colors - Dark */
    --color-int-signed: #60a5fa;
    --color-int-unsigned: #34d399;
    --color-float: #f87171;
    --color-binary: #a78bfa;
    --color-bytes: #fb923c;
  }
  
  /* Reset & Base */
  * {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
  }
  
  main {
    min-height: 100vh;
    background: var(--bg-primary);
    color: var(--text-primary);
    font-family: var(--font-sans);
    transition: background 0.3s, color 0.3s;
  }
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: var(--spacing-xl);
  }
</style>
```

**Acceptance Criteria**:
- Theme switches smoothly between light and dark
- Color variables clearly differentiate data types
- Responsive container with max-width
- CSS custom properties used consistently

---

### Step 8: Styling - Header & Input Section
**File**: `frontend/src/App.svelte` (add to style section)

Style the header and input area:

```css
/* Header */
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-xl);
  padding-bottom: var(--spacing-lg);
  border-bottom: 2px solid var(--border-color);
}

h1 {
  font-size: 2rem;
  font-weight: 600;
  letter-spacing: -0.02em;
}

.theme-toggle {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: var(--spacing-sm) var(--spacing-md);
  font-size: 1.25rem;
  cursor: pointer;
  transition: all 0.2s;
}

.theme-toggle:hover {
  background: var(--bg-tertiary);
  transform: scale(1.05);
}

/* Input Section */
.input-section {
  margin-bottom: var(--spacing-xl);
}

.input-section label {
  display: block;
  font-weight: 600;
  font-size: 0.875rem;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--text-secondary);
  margin-bottom: var(--spacing-sm);
}

.input-section input {
  width: 100%;
  padding: var(--spacing-md);
  font-size: 1.125rem;
  font-family: var(--font-mono);
  background: var(--bg-secondary);
  border: 2px solid var(--border-color);
  border-radius: var(--radius-md);
  color: var(--text-primary);
  transition: all 0.2s;
}

.input-section input:focus {
  outline: none;
  border-color: var(--color-int-signed);
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.1);
}

.input-section input.error {
  border-color: var(--color-error);
}

.input-hint {
  margin-top: var(--spacing-sm);
  font-size: 0.75rem;
  color: var(--text-tertiary);
  font-style: italic;
}

.error-message {
  margin-top: var(--spacing-md);
  padding: var(--spacing-md);
  background: rgba(239, 68, 68, 0.1);
  border-left: 4px solid var(--color-error);
  border-radius: var(--radius-sm);
  color: var(--color-error);
  font-size: 0.875rem;
}

/* Help Text */
.help-text {
  margin-top: var(--spacing-xl);
  padding: var(--spacing-lg);
  background: var(--bg-secondary);
  border-radius: var(--radius-lg);
  border: 1px dashed var(--border-color);
}

.help-text h3 {
  font-size: 1.25rem;
  margin-bottom: var(--spacing-md);
}

.help-text p {
  color: var(--text-secondary);
  margin-bottom: var(--spacing-md);
}

.examples {
  display: flex;
  gap: var(--spacing-sm);
  flex-wrap: wrap;
  align-items: center;
}

.examples strong {
  color: var(--text-secondary);
  font-size: 0.875rem;
}

.examples code {
  padding: var(--spacing-xs) var(--spacing-sm);
  background: var(--bg-tertiary);
  border-radius: var(--radius-sm);
  font-family: var(--font-mono);
  font-size: 0.875rem;
  color: var(--color-int-signed);
}

/* Loading Indicator */
.loading-indicator {
  opacity: 0;
  transition: opacity 0.2s;
  font-size: 0.875rem;
  color: var(--text-tertiary);
  margin-bottom: var(--spacing-md);
}

.loading-indicator.visible {
  opacity: 1;
}
```

**Acceptance Criteria**:
- Input field is prominent and easy to use
- Focus states clearly visible
- Error states distinguished visually
- Help text styled attractively
- Loading indicator fades in smoothly

---

### Step 9: Styling - Results Groups & Items
**File**: `frontend/src/App.svelte` (add to style section)

Style the results display:

```css
/* Results Section */
.results {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
}

.result-group {
  background: var(--bg-secondary);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  border: 1px solid var(--border-color);
}

.group-title {
  font-size: 1rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--text-secondary);
  margin-bottom: var(--spacing-md);
  padding-bottom: var(--spacing-sm);
  border-bottom: 1px solid var(--border-color);
}

.result-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--spacing-md);
}

.result-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-md);
  background: var(--bg-primary);
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color);
}

.result-item.full-width {
  grid-column: 1 / -1;
}

.type-label {
  flex-shrink: 0;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: 0.75rem;
  font-weight: 700;
  font-family: var(--font-mono);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.type-label.int-signed {
  background: rgba(37, 99, 235, 0.15);
  color: var(--color-int-signed);
}

.type-label.int-unsigned {
  background: rgba(5, 150, 105, 0.15);
  color: var(--color-int-unsigned);
}

.type-label.float {
  background: rgba(220, 38, 38, 0.15);
  color: var(--color-float);
}

.type-label.binary {
  background: rgba(124, 58, 237, 0.15);
  color: var(--color-binary);
}

.type-label.bytes {
  background: rgba(234, 88, 12, 0.15);
  color: var(--color-bytes);
}

.value {
  flex: 1;
  font-size: 1rem;
  font-weight: 500;
  color: var(--text-primary);
  word-break: break-all;
}

.value.mono {
  font-family: var(--font-mono);
  font-size: 0.875rem;
}

.copy-btn {
  flex-shrink: 0;
  background: transparent;
  border: none;
  padding: var(--spacing-xs);
  cursor: pointer;
  font-size: 1rem;
  opacity: 0.5;
  transition: opacity 0.2s, transform 0.2s;
}

.copy-btn:hover {
  opacity: 1;
  transform: scale(1.1);
}

.copy-btn:active {
  transform: scale(0.95);
}
```

**Acceptance Criteria**:
- Results grouped logically with clear headers
- Color-coded type labels for quick identification
- Responsive grid layout adapts to screen size
- Copy buttons hover states feel interactive
- Monospace font for binary/hex strings

---

### Step 10: Responsive Design & Polish
**File**: `frontend/src/App.svelte` (add to style section)

Add responsive breakpoints and final polish:

```css
/* Responsive Design */
@media (max-width: 768px) {
  .container {
    padding: var(--spacing-md);
  }
  
  h1 {
    font-size: 1.5rem;
  }
  
  .result-grid {
    grid-template-columns: 1fr;
  }
  
  .input-section input {
    font-size: 1rem;
  }
  
  .theme-toggle {
    font-size: 1rem;
  }
}

@media (max-width: 480px) {
  .container {
    padding: var(--spacing-sm);
  }
  
  header {
    flex-direction: column;
    gap: var(--spacing-md);
    align-items: flex-start;
  }
  
  .theme-toggle {
    align-self: flex-end;
  }
  
  .examples {
    flex-direction: column;
    align-items: flex-start;
  }
}

/* Animations */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.result-group {
  animation: fadeIn 0.3s ease-out;
}

/* Scrollbar Styling */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: var(--bg-secondary);
}

::-webkit-scrollbar-thumb {
  background: var(--border-color);
  border-radius: var(--radius-sm);
}

::-webkit-scrollbar-thumb:hover {
  background: var(--text-tertiary);
}

/* Selection */
::selection {
  background: rgba(37, 99, 235, 0.3);
}
```

**Acceptance Criteria**:
- Mobile responsive down to 320px width
- Smooth animations for result appearance
- Styled scrollbars match theme
- Touch-friendly button sizes on mobile
- Text selection styled consistently

---

### Step 11: Testing & Validation
**Tasks**:

1. **Unit Testing** (Manual):
   - Test all hex input formats: `0x`, spaces, colons, commas, mixed case
   - Verify empty input clears results immediately
   - Test invalid hex characters show error
   - Verify debounce delay (should wait 300ms)
   - Test copy-to-clipboard for each data type

2. **Cross-browser Testing**:
   - Chrome/Edge (Chromium)
   - Firefox
   - Safari
   - Verify theme switching works in all browsers
   - Test clipboard API compatibility

3. **Responsive Testing**:
   - Desktop: 1920px, 1366px, 1024px
   - Tablet: 768px, 600px
   - Mobile: 480px, 375px, 320px
   - Verify layout doesn't break at any size

4. **Performance Testing**:
   - Measure debounce timing accuracy
   - Verify no memory leaks with repeated conversions
   - Test with very long hex strings (64+ bytes)
   - Ensure smooth theme transitions

5. **Accessibility Testing**:
   - Keyboard navigation works (Tab, Enter, Escape)
   - Screen reader announces errors and results
   - Focus indicators visible
   - Color contrast meets WCAG AA standards
   - ARIA labels present where needed

**Acceptance Criteria**:
- All conversion types work correctly
- No console errors in any browser
- Smooth performance with <100ms conversion time
- Passes WCAG 2.1 Level AA accessibility standards
- Works offline (after initial load)

---

### Step 12: Build & Deploy
**Commands**:

```bash
# Development mode with hot reload
wails dev

# Build production binary
wails build

# Build for specific platforms
wails build -platform darwin/amd64
wails build -platform darwin/arm64
wails build -platform windows/amd64
wails build -platform linux/amd64
```

**Pre-deployment Checklist**:
- [ ] Remove console.log statements
- [ ] Verify all imports are used
- [ ] Check bundle size is reasonable
- [ ] Test production build locally
- [ ] Verify theme preference persists
- [ ] All error messages are user-friendly
- [ ] Help text is clear and accurate

**Acceptance Criteria**:
- Production build creates working binary
- Binary size < 50MB (reasonable for desktop app)
- App launches in < 2 seconds
- No runtime errors in production mode

---

## Additional Enhancements (Optional Future Steps)

### Copy Feedback Toast
Add visual confirmation when copying values:
```javascript
let toastMessage = ''
let showToast = false

async function copyToClipboard(text) {
  try {
    await navigator.clipboard.writeText(String(text))
    toastMessage = 'Copied to clipboard!'
    showToast = true
    setTimeout(() => showToast = false, 2000)
  } catch (err) {
    toastMessage = 'Failed to copy'
    showToast = true
    setTimeout(() => showToast = false, 2000)
  }
}
```

### Keyboard Shortcuts
- `Cmd/Ctrl + K`: Clear input
- `Cmd/Ctrl + ,`: Toggle theme
- `Escape`: Clear input and focus

### Input History
Store last 10 conversions in localStorage for quick access

### Export Results
Add "Export as JSON" button to save all conversions

---

## Success Metrics

- **Conversion Speed**: < 100ms from input to display
- **First Contentful Paint**: < 500ms
- **Accessibility Score**: 95+ (Lighthouse)
- **Bundle Size**: < 2MB (frontend assets)
- **User Actions to Convert**: 0 (real-time, no button click needed)

---

## Development Timeline Estimate

- **Step 1-2**: Backend + State Management - 1-2 hours
- **Step 3**: Input Component - 1 hour
- **Step 4-6**: Results Display - 2-3 hours
- **Step 7-10**: Styling & Polish - 2-3 hours
- **Step 11**: Testing - 2-3 hours
- **Step 12**: Build & Deploy - 1 hour

**Total**: ~10-14 hours for complete implementation

---

## Notes for Developer

- Use Wails v2 CLI for development: `wails dev` provides hot reload
- TypeScript bindings auto-generate on build - don't edit them manually
- Follow the project's Copilot instructions in `.github/copilot-instructions.md`
- Test with actual network engineering hex values (MAC addresses, IP headers, etc.)
- Prioritize UX: every click saved is valuable for rapid troubleshooting workflows
