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
  
  // Toast state
  let toastMessage = ''
  let showToast = false
  
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
      toastMessage = 'Copied to clipboard!'
      showToast = true
      setTimeout(() => showToast = false, 2000)
    } catch (err) {
      toastMessage = 'Failed to copy'
      showToast = true
      setTimeout(() => showToast = false, 2000)
    }
  }
  
  function clearInput() {
    hexInput = ''
    result = null
    error = null
  }
  
  function handleKeydown(event) {
    if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
      event.preventDefault()
      clearInput()
    }
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
        Accepts: 0x prefix, spaces, colons, commas • Case insensitive • <kbd>⌘K</kbd> to clear
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
        
        <!-- Signed Integers - Big Endian -->
        {#if result.int8BE !== null || result.int16BE !== null || result.int32BE !== null || result.int64BE !== null}
          <div class="result-group signed-integers">
            <h3 class="group-title">Signed Integers (Big Endian)</h3>
            <div class="result-grid">
              {#if result.int8BE !== null && result.int8BE !== undefined}
                <div class="result-item">
                  <span class="type-label int-signed">INT8</span>
                  <span class="value">{result.int8BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.int8BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.int16BE !== null && result.int16BE !== undefined}
                <div class="result-item">
                  <span class="type-label int-signed">INT16</span>
                  <span class="value">{result.int16BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.int16BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.int32BE !== null && result.int32BE !== undefined}
                <div class="result-item">
                  <span class="type-label int-signed">INT32</span>
                  <span class="value">{result.int32BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.int32BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.int64BE !== null && result.int64BE !== undefined}
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
              {#if result.int16LE !== null && result.int16LE !== undefined}
                <div class="result-item">
                  <span class="type-label int-signed">INT16 LE</span>
                  <span class="value">{result.int16LE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.int16LE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.int32LE !== null && result.int32LE !== undefined}
                <div class="result-item">
                  <span class="type-label int-signed">INT32 LE</span>
                  <span class="value">{result.int32LE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.int32LE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.int64LE !== null && result.int64LE !== undefined}
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

        <!-- Unsigned Integers - Big Endian -->
        {#if result.uint8BE !== null || result.uint16BE !== null || result.uint32BE !== null || result.uint64BE !== null}
          <div class="result-group unsigned-integers">
            <h3 class="group-title">Unsigned Integers (Big Endian)</h3>
            <div class="result-grid">
              {#if result.uint8BE !== null && result.uint8BE !== undefined}
                <div class="result-item">
                  <span class="type-label int-unsigned">UINT8</span>
                  <span class="value">{result.uint8BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.uint8BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.uint16BE !== null && result.uint16BE !== undefined}
                <div class="result-item">
                  <span class="type-label int-unsigned">UINT16</span>
                  <span class="value">{result.uint16BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.uint16BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.uint32BE !== null && result.uint32BE !== undefined}
                <div class="result-item">
                  <span class="type-label int-unsigned">UINT32</span>
                  <span class="value">{result.uint32BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.uint32BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.uint64BE !== null && result.uint64BE !== undefined}
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
              {#if result.uint16LE !== null && result.uint16LE !== undefined}
                <div class="result-item">
                  <span class="type-label int-unsigned">UINT16 LE</span>
                  <span class="value">{result.uint16LE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.uint16LE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.uint32LE !== null && result.uint32LE !== undefined}
                <div class="result-item">
                  <span class="type-label int-unsigned">UINT32 LE</span>
                  <span class="value">{result.uint32LE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.uint32LE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.uint64LE !== null && result.uint64LE !== undefined}
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

        <!-- Floating Point Numbers -->
        {#if result.float32BE !== null || result.float64BE !== null || result.float32LE !== null || result.float64LE !== null}
          <div class="result-group float-numbers">
            <h3 class="group-title">Floating Point</h3>
            <div class="result-grid">
              {#if result.float32BE !== null && result.float32BE !== undefined}
                <div class="result-item">
                  <span class="type-label float">FLOAT32 BE</span>
                  <span class="value">{result.float32BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.float32BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.float64BE !== null && result.float64BE !== undefined}
                <div class="result-item">
                  <span class="type-label float">FLOAT64 BE</span>
                  <span class="value">{result.float64BE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.float64BE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.float32LE !== null && result.float32LE !== undefined}
                <div class="result-item">
                  <span class="type-label float">FLOAT32 LE</span>
                  <span class="value">{result.float32LE}</span>
                  <button class="copy-btn" on:click={() => copyToClipboard(result.float32LE)} title="Copy">
                    📋
                  </button>
                </div>
              {/if}
              
              {#if result.float64LE !== null && result.float64LE !== undefined}
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
  
  .input-hint kbd {
    background: var(--bg-tertiary);
    padding: 0.1rem 0.3rem;
    border-radius: var(--radius-sm);
    font-family: var(--font-sans);
    font-size: 0.7rem;
    font-style: normal;
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
    animation: fadeIn 0.3s ease-out;
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

  /* Toast */
  .toast {
    position: fixed;
    bottom: var(--spacing-xl);
    left: 50%;
    transform: translateX(-50%);
    background: var(--bg-tertiary);
    color: var(--text-primary);
    padding: var(--spacing-sm) var(--spacing-lg);
    border-radius: var(--radius-md);
    box-shadow: 0 4px 12px var(--shadow);
    animation: slideUp 0.3s ease-out;
    z-index: 1000;
  }

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

  @keyframes slideUp {
    from {
      opacity: 0;
      transform: translateX(-50%) translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateX(-50%) translateY(0);
    }
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
</style>
