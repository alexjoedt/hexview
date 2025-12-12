<script>
  import { ConvertHex, ConvertInt, ConvertBinary } from '../wailsjs/go/main/App.js'
  
  // State
  let inputMode = 'hex' // 'hex', 'int', or 'binary'
  let intType = 'int16' // Default to int16 (common for Modbus)
  let inputValue = ''
  let result = null
  let error = null
  let isLoading = false
  let debounceTimer = null
  
  // Theme state (default to system preference)
  let darkMode = window.matchMedia('(prefers-color-scheme: dark)').matches
  
  // Toast state
  let toastMessage = ''
  let showToast = false
  
  // Reactive conversion with debounce
  $: {
    if (inputValue.trim() === '') {
      result = null
      error = null
    } else {
      debouncedConvert(inputValue, inputMode, intType)
    }
  }
  
  function debouncedConvert(input, mode, type) {
    clearTimeout(debounceTimer)
    isLoading = true
    
    debounceTimer = setTimeout(() => {
      let conversionPromise
      if (mode === 'hex') {
        conversionPromise = ConvertHex(input)
      } else if (mode === 'binary') {
        conversionPromise = ConvertBinary(input)
      } else {
        conversionPromise = ConvertInt(input, type)
      }
      
      conversionPromise
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
      toastMessage = 'Copied!'
      showToast = true
      setTimeout(() => showToast = false, 1500)
    } catch (err) {
      toastMessage = 'Failed to copy'
      showToast = true
      setTimeout(() => showToast = false, 1500)
    }
  }
  
  function clearInput() {
    inputValue = ''
    result = null
    error = null
  }
  
  function handleKeydown(event) {
    if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
      event.preventDefault()
      clearInput()
    }
  }
  
  // Helper to format value or show placeholder
  function formatValue(value) {
    return value !== null && value !== undefined ? String(value) : '—'
  }
  
  // Check if value exists
  function hasValue(value) {
    return value !== null && value !== undefined
  }
</script>

<svelte:window on:keydown={handleKeydown} />

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
      <!-- Mode Selector -->
      <div class="mode-selector">
        <button 
          class="mode-btn" 
          class:active={inputMode === 'hex'}
          on:click={() => inputMode = 'hex'}
        >
          Hex
        </button>
        <button 
          class="mode-btn" 
          class:active={inputMode === 'binary'}
          on:click={() => inputMode = 'binary'}
        >
          Binary
        </button>
        <button 
          class="mode-btn" 
          class:active={inputMode === 'int'}
          on:click={() => inputMode = 'int'}
        >
          Integer
        </button>
        
        {#if inputMode === 'int'}
          <select class="type-selector" bind:value={intType}>
            <option value="int8">INT8</option>
            <option value="int16">INT16</option>
            <option value="int32">INT32</option>
            <option value="int64">INT64</option>
            <option value="uint8">UINT8</option>
            <option value="uint16">UINT16</option>
            <option value="uint32">UINT32</option>
            <option value="uint64">UINT64</option>
          </select>
        {/if}
      </div>
      
      <input
        id="input-field"
        type="text"
        bind:value={inputValue}
        placeholder={inputMode === 'hex' ? 'Enter hex (e.g., 0xFF, 1A 2B 3C, 1A:2B:3C)' : inputMode === 'binary' ? 'Enter binary (e.g., 1001 0000, 10011111)' : 'Enter integer (e.g., 1234, -456)'}
        autocomplete="off"
        autofocus
        class:error={error}
      />
      
      {#if error}
        <div class="error-message">⚠️ {error}</div>
      {/if}
      
      {#if !inputValue && !result}
        <div class="help-text">
          {inputMode === 'hex' ? 'Enter hex to convert' : inputMode === 'binary' ? 'Enter binary to convert' : 'Enter integer to convert'} • <kbd>⌘K</kbd> to clear
        </div>
      {/if}
    </section>
    
    <!-- Results Section -->
    {#if result && !error}
      <section class="results">
        <!-- Integer Conversions Table -->
        <div class="table-wrapper">
          <h3 class="table-title">Integer Conversions</h3>
          <table>
            <thead>
              <tr>
                <th>Type</th>
                <th>Big Endian</th>
                <th>Little Endian</th>
                <th>Mid-Big (BADC)</th>
                <th>Mid-Little (CDAB)</th>
              </tr>
            </thead>
            <tbody>
              <!-- INT8 -->
              <tr class:unavailable={!hasValue(result.int8BE)} class:highlighted={inputMode === 'int' && intType === 'int8'}>
                <td class="type-cell"><span class="type-badge int-signed">INT8</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int8BE)}</span>
                  {#if hasValue(result.int8BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int8BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell na">—</td>
                <td class="value-cell na">—</td>
                <td class="value-cell na">—</td>
              </tr>
              
              <!-- INT16 -->
              <tr class:unavailable={!hasValue(result.int16BE) && !hasValue(result.int16LE)} class:highlighted={inputMode === 'int' && intType === 'int16'}>
                <td class="type-cell"><span class="type-badge int-signed">INT16</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int16BE)}</span>
                  {#if hasValue(result.int16BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int16BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int16LE)}</span>
                  {#if hasValue(result.int16LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int16LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell na">—</td>
                <td class="value-cell na">—</td>
              </tr>
              
              <!-- INT32 -->
              <tr class:unavailable={!hasValue(result.int32BE) && !hasValue(result.int32LE)} class:highlighted={inputMode === 'int' && intType === 'int32'}>
                <td class="type-cell"><span class="type-badge int-signed">INT32</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int32BE)}</span>
                  {#if hasValue(result.int32BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int32BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int32LE)}</span>
                  {#if hasValue(result.int32LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int32LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int32BADC)}</span>
                  {#if hasValue(result.int32BADC)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int32BADC)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int32CDAB)}</span>
                  {#if hasValue(result.int32CDAB)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int32CDAB)} title="Copy">📋</button>
                  {/if}
                </td>
              </tr>
              
              <!-- INT64 -->
              <tr class:unavailable={!hasValue(result.int64BE) && !hasValue(result.int64LE)} class:highlighted={inputMode === 'int' && intType === 'int64'}>
                <td class="type-cell"><span class="type-badge int-signed">INT64</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int64BE)}</span>
                  {#if hasValue(result.int64BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int64BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int64LE)}</span>
                  {#if hasValue(result.int64LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int64LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int64BADC)}</span>
                  {#if hasValue(result.int64BADC)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int64BADC)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.int64CDAB)}</span>
                  {#if hasValue(result.int64CDAB)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.int64CDAB)} title="Copy">📋</button>
                  {/if}
                </td>
              </tr>
              
              <!-- UINT8 -->
              <tr class:unavailable={!hasValue(result.uint8BE)} class:highlighted={inputMode === 'int' && intType === 'uint8'}>
                <td class="type-cell"><span class="type-badge int-unsigned">UINT8</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint8BE)}</span>
                  {#if hasValue(result.uint8BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint8BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell na">—</td>
                <td class="value-cell na">—</td>
                <td class="value-cell na">—</td>
              </tr>
              
              <!-- UINT16 -->
              <tr class:unavailable={!hasValue(result.uint16BE) && !hasValue(result.uint16LE)} class:highlighted={inputMode === 'int' && intType === 'uint16'}>
                <td class="type-cell"><span class="type-badge int-unsigned">UINT16</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint16BE)}</span>
                  {#if hasValue(result.uint16BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint16BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint16LE)}</span>
                  {#if hasValue(result.uint16LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint16LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell na">—</td>
                <td class="value-cell na">—</td>
              </tr>
              
              <!-- UINT32 -->
              <tr class:unavailable={!hasValue(result.uint32BE) && !hasValue(result.uint32LE)} class:highlighted={inputMode === 'int' && intType === 'uint32'}>
                <td class="type-cell"><span class="type-badge int-unsigned">UINT32</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint32BE)}</span>
                  {#if hasValue(result.uint32BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint32BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint32LE)}</span>
                  {#if hasValue(result.uint32LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint32LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint32BADC)}</span>
                  {#if hasValue(result.uint32BADC)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint32BADC)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint32CDAB)}</span>
                  {#if hasValue(result.uint32CDAB)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint32CDAB)} title="Copy">📋</button>
                  {/if}
                </td>
              </tr>
              
              <!-- UINT64 -->
              <tr class:unavailable={!hasValue(result.uint64BE) && !hasValue(result.uint64LE)} class:highlighted={inputMode === 'int' && intType === 'uint64'}>
                <td class="type-cell"><span class="type-badge int-unsigned">UINT64</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint64BE)}</span>
                  {#if hasValue(result.uint64BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint64BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint64LE)}</span>
                  {#if hasValue(result.uint64LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint64LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint64BADC)}</span>
                  {#if hasValue(result.uint64BADC)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint64BADC)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.uint64CDAB)}</span>
                  {#if hasValue(result.uint64CDAB)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.uint64CDAB)} title="Copy">📋</button>
                  {/if}
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Floating Point Table -->
        <div class="table-wrapper">
          <h3 class="table-title">Floating Point</h3>
          <table>
            <thead>
              <tr>
                <th>Type</th>
                <th>Big Endian</th>
                <th>Little Endian</th>
                <th>Mid-Big (BADC)</th>
                <th>Mid-Little (CDAB)</th>
              </tr>
            </thead>
            <tbody>
              <!-- FLOAT32 -->
              <tr class:unavailable={!hasValue(result.float32BE) && !hasValue(result.float32LE)}>
                <td class="type-cell"><span class="type-badge float">FLOAT32</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float32BE)}</span>
                  {#if hasValue(result.float32BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float32BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float32LE)}</span>
                  {#if hasValue(result.float32LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float32LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float32BADC)}</span>
                  {#if hasValue(result.float32BADC)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float32BADC)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float32CDAB)}</span>
                  {#if hasValue(result.float32CDAB)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float32CDAB)} title="Copy">📋</button>
                  {/if}
                </td>
              </tr>
              
              <!-- FLOAT64 -->
              <tr class:unavailable={!hasValue(result.float64BE) && !hasValue(result.float64LE)}>
                <td class="type-cell"><span class="type-badge float">FLOAT64</span></td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float64BE)}</span>
                  {#if hasValue(result.float64BE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float64BE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float64LE)}</span>
                  {#if hasValue(result.float64LE)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float64LE)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float64BADC)}</span>
                  {#if hasValue(result.float64BADC)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float64BADC)} title="Copy">📋</button>
                  {/if}
                </td>
                <td class="value-cell-with-copy">
                  <span class="value-text">{formatValue(result.float64CDAB)}</span>
                  {#if hasValue(result.float64CDAB)}
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.float64CDAB)} title="Copy">📋</button>
                  {/if}
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Binary Representations -->
        <div class="table-wrapper">
          <h3 class="table-title">Binary Representations</h3>
          <table>
            <tbody>
              {#if result.bytes}
                <tr>
                  <td class="type-cell"><span class="type-badge bytes">HEX</span></td>
                  <td class="value-cell-with-copy mono wide">
                    <span class="value-text">{result.bytes}</span>
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.bytes)} title="Copy">📋</button>
                  </td>
                </tr>
              {/if}
              
              {#if result.binary}
                <tr>
                  <td class="type-cell"><span class="type-badge binary">BIN</span></td>
                  <td class="value-cell-with-copy mono wide">
                    <span class="value-text">{result.binary}</span>
                    <button class="copy-btn-inline" on:click={() => copyToClipboard(result.binary)} title="Copy">📋</button>
                  </td>
                </tr>
              {/if}
            </tbody>
          </table>
        </div>
      </section>
    {/if}
  </div>
  
  <!-- Toast notification -->
  {#if showToast}
    <div class="toast">{toastMessage}</div>
  {/if}
</main>

<style>
  :root {
    /* Light Theme Colors */
    --bg-primary: #ffffff;
    --bg-secondary: #f8f9fa;
    --bg-tertiary: #e9ecef;
    --bg-hover: #f1f3f5;
    --text-primary: #1a1a1a;
    --text-secondary: #666666;
    --text-tertiary: #999999;
    --border-color: #dee2e6;
    --shadow: rgba(0, 0, 0, 0.05);
    
    /* Type Colors - Light */
    --color-int-signed: #2563eb;
    --color-int-unsigned: #059669;
    --color-float: #dc2626;
    --color-binary: #7c3aed;
    --color-bytes: #ea580c;
    
    /* Status Colors */
    --color-error: #ef4444;
    --color-success: #10b981;
    
    /* Spacing - Compact */
    --spacing-xs: 0.25rem;
    --spacing-sm: 0.5rem;
    --spacing-md: 0.75rem;
    --spacing-lg: 1rem;
    
    /* Border Radius */
    --radius-sm: 0.25rem;
    --radius-md: 0.375rem;
    
    /* Typography */
    --font-mono: 'SF Mono', 'Monaco', 'Consolas', monospace;
    --font-sans: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  }
  
  /* Dark Theme */
  .dark {
    --bg-primary: #1a1a1a;
    --bg-secondary: #242424;
    --bg-tertiary: #2e2e2e;
    --bg-hover: #333333;
    --text-primary: #f0f0f0;
    --text-secondary: #b0b0b0;
    --text-tertiary: #707070;
    --border-color: #3a3a3a;
    --shadow: rgba(0, 0, 0, 0.3);
    
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
    font-size: 13px;
  }
  
  .container {
    height: 100vh;
    display: flex;
    flex-direction: column;
    padding: var(--spacing-md);
  }

  /* Header */
  header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-md);
    padding-bottom: var(--spacing-md);
    border-bottom: 1px solid var(--border-color);
    flex-shrink: 0;
  }

  h1 {
    font-size: 1.25rem;
    font-weight: 600;
    letter-spacing: -0.01em;
  }

  .theme-toggle {
    background: transparent;
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    padding: var(--spacing-xs) var(--spacing-sm);
    font-size: 1rem;
    cursor: pointer;
    transition: all 0.15s;
  }

  .theme-toggle:hover {
    background: var(--bg-hover);
  }

  /* Input Section */
  .input-section {
    margin-bottom: var(--spacing-md);
    flex-shrink: 0;
  }

  /* Mode Selector */
  .mode-selector {
    display: flex;
    gap: var(--spacing-sm);
    margin-bottom: var(--spacing-sm);
    align-items: center;
  }

  .mode-btn {
    padding: var(--spacing-xs) var(--spacing-md);
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    color: var(--text-secondary);
    font-size: 12px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s;
  }

  .mode-btn:hover {
    background: var(--bg-hover);
    border-color: var(--text-tertiary);
  }

  .mode-btn.active {
    background: var(--color-int-signed);
    color: white;
    border-color: var(--color-int-signed);
  }

  .type-selector {
    padding: var(--spacing-xs) var(--spacing-sm);
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    color: var(--text-primary);
    font-size: 11px;
    font-weight: 600;
    font-family: var(--font-mono);
    cursor: pointer;
    transition: all 0.15s;
    min-width: 80px;
  }

  .type-selector:hover {
    background: var(--bg-hover);
    border-color: var(--text-tertiary);
  }

  .type-selector:focus {
    outline: none;
    border-color: var(--color-int-signed);
  }

  .input-section input {
    width: 100%;
    padding: var(--spacing-sm) var(--spacing-md);
    font-size: 13px;
    font-family: var(--font-mono);
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    color: var(--text-primary);
    transition: all 0.15s;
  }

  .input-section input:focus {
    outline: none;
    border-color: var(--color-int-signed);
    background: var(--bg-primary);
  }

  .input-section input.error {
    border-color: var(--color-error);
  }

  .error-message {
    margin-top: var(--spacing-sm);
    padding: var(--spacing-sm);
    background: rgba(239, 68, 68, 0.1);
    border-left: 3px solid var(--color-error);
    border-radius: var(--radius-sm);
    color: var(--color-error);
    font-size: 12px;
  }

  /* Help Text */
  .help-text {
    margin-top: var(--spacing-sm);
    padding: var(--spacing-sm);
    text-align: center;
    color: var(--text-tertiary);
    font-size: 11px;
  }
  
  .help-text kbd {
    background: var(--bg-tertiary);
    padding: 0.1rem 0.3rem;
    border-radius: var(--radius-sm);
    font-family: var(--font-mono);
    font-size: 10px;
  }

  /* Results Section */
  .results {
    flex: 1;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
  }

  /* Table Wrapper */
  .table-wrapper {
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    overflow-x: auto;
    overflow-y: visible;
  }

  .table-title {
    font-size: 11px;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: var(--text-secondary);
    padding: var(--spacing-sm) var(--spacing-md);
    background: var(--bg-secondary);
    border-bottom: 1px solid var(--border-color);
  }

  /* Table Styles */
  table {
    width: 100%;
    min-width: 600px;
    border-collapse: collapse;
    font-size: 12px;
  }

  thead {
    background: var(--bg-secondary);
    position: sticky;
    top: 0;
    z-index: 10;
  }

  th {
    text-align: left;
    font-weight: 600;
    font-size: 11px;
    text-transform: uppercase;
    letter-spacing: 0.03em;
    color: var(--text-secondary);
    padding: var(--spacing-sm) var(--spacing-md);
    border-bottom: 1px solid var(--border-color);
  }

  tbody tr {
    background: var(--bg-primary);
    transition: background 0.1s;
  }

  tbody tr:hover {
    background: var(--bg-hover);
  }

  tbody tr.unavailable {
    opacity: 0.4;
  }

  tbody tr.highlighted {
    background: rgba(37, 99, 235, 0.08) !important;
    border-left: 3px solid var(--color-int-signed);
  }

  .dark tbody tr.highlighted {
    background: rgba(96, 165, 250, 0.12) !important;
  }

  tbody tr:not(:last-child) td {
    border-bottom: 1px solid var(--border-color);
  }

  td {
    padding: var(--spacing-sm) var(--spacing-md);
  }

  .type-cell {
    width: 90px;
  }

  .value-cell {
    font-family: var(--font-mono);
    font-size: 12px;
    color: var(--text-primary);
  }

  .value-cell.na {
    color: var(--text-tertiary);
    text-align: center;
  }

  .value-cell.mono {
    font-size: 11px;
  }

  .value-cell.wide {
    max-width: 0;
    overflow-x: auto;
    white-space: nowrap;
  }

  .value-cell-with-copy {
    font-family: var(--font-mono);
    font-size: 12px;
    color: var(--text-primary);
    position: relative;
    padding-right: 28px;
  }

  .value-cell-with-copy.mono {
    font-size: 11px;
  }

  .value-cell-with-copy.wide {
    max-width: 0;
    overflow-x: auto;
    white-space: nowrap;
  }

  .value-text {
    display: inline-block;
  }

  /* Type Badges */
  .type-badge {
    display: inline-block;
    padding: 2px 6px;
    border-radius: var(--radius-sm);
    font-size: 10px;
    font-weight: 700;
    font-family: var(--font-mono);
    text-transform: uppercase;
    letter-spacing: 0.03em;
  }

  .type-badge.int-signed {
    background: rgba(37, 99, 235, 0.12);
    color: var(--color-int-signed);
  }

  .type-badge.int-unsigned {
    background: rgba(5, 150, 105, 0.12);
    color: var(--color-int-unsigned);
  }

  .type-badge.float {
    background: rgba(220, 38, 38, 0.12);
    color: var(--color-float);
  }

  .type-badge.binary {
    background: rgba(124, 58, 237, 0.12);
    color: var(--color-binary);
  }

  .type-badge.bytes {
    background: rgba(234, 88, 12, 0.12);
    color: var(--color-bytes);
  }

  /* Copy Button - Inline */
  .copy-btn-inline {
    background: transparent;
    border: none;
    padding: 2px;
    cursor: pointer;
    font-size: 12px;
    opacity: 0;
    transition: opacity 0.15s, transform 0.15s;
    position: absolute;
    right: 4px;
    top: 50%;
    transform: translateY(-50%);
    line-height: 1;
  }

  tr:hover .copy-btn-inline {
    opacity: 0.5;
  }

  .copy-btn-inline:hover {
    opacity: 1 !important;
    transform: translateY(-50%) scale(1.15);
  }

  .copy-btn-inline:active {
    transform: translateY(-50%) scale(0.9);
  }

  /* Toast */
  .toast {
    position: fixed;
    bottom: var(--spacing-lg);
    left: 50%;
    transform: translateX(-50%);
    background: var(--bg-tertiary);
    color: var(--text-primary);
    padding: var(--spacing-sm) var(--spacing-lg);
    border-radius: var(--radius-md);
    box-shadow: 0 2px 8px var(--shadow);
    animation: slideUp 0.2s ease-out;
    z-index: 1000;
    font-size: 12px;
  }

  /* Animations */
  @keyframes slideUp {
    from {
      opacity: 0;
      transform: translateX(-50%) translateY(10px);
    }
    to {
      opacity: 1;
      transform: translateX(-50%) translateY(0);
    }
  }

  /* Scrollbar Styling */
  .results::-webkit-scrollbar {
    width: 6px;
  }

  .results::-webkit-scrollbar-track {
    background: var(--bg-secondary);
  }

  .results::-webkit-scrollbar-thumb {
    background: var(--border-color);
    border-radius: 3px;
  }

  .results::-webkit-scrollbar-thumb:hover {
    background: var(--text-tertiary);
  }

  .value-cell.wide::-webkit-scrollbar {
    height: 4px;
  }

  .value-cell.wide::-webkit-scrollbar-track {
    background: transparent;
  }

  .value-cell.wide::-webkit-scrollbar-thumb {
    background: var(--border-color);
    border-radius: 2px;
  }

  /* Selection */
  ::selection {
    background: rgba(37, 99, 235, 0.25);
  }
</style>
