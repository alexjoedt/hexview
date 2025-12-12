<script>
  import { INPUT_MODES, INT_TYPE_OPTIONS } from '../lib/constants.js'
  
  export let inputMode = 'hex'
  export let intType = 'int16'
  export let inputValue = ''
  export let error = null
  export let onClear
  export let scaleFactor = 1
  
  let scaleInput = '1'
  
  function handleKeydown(event) {
    if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
      event.preventDefault()
      if (onClear) onClear()
    }
  }
  
  // Reactively update scaleFactor when scaleInput changes
  $: {
    const val = parseFloat(scaleInput)
    if (!isNaN(val) && val !== 0) {
      scaleFactor = val
    }
  }
  
  $: placeholder = inputMode === 'hex' 
    ? 'Enter hex (e.g., 0x48656c6c6f or 48 65 6c 6c 6f)' 
    : inputMode === 'binary'
    ? 'Enter binary (e.g., 01001000 01100101)'
    : inputMode === 'modbus'
    ? 'Enter registers: hex (1234 5678), decimal (d1000 d2000), or mixed'
    : 'Enter integer value'
    
  $: isModbus = inputMode === 'modbus'
</script>

<div class="input-section">
  <div class="mode-selector">
    {#each INPUT_MODES as mode}
      <button 
        class="mode-btn" 
        class:active={inputMode === mode.value}
        class:modbus={mode.value === 'modbus' && inputMode === mode.value}
        on:click={() => inputMode = mode.value}
      >
        {mode.label}
      </button>
    {/each}
    
    {#if inputMode === 'int'}
      <select class="type-selector" bind:value={intType}>
        {#each INT_TYPE_OPTIONS as option}
          <option value={option.value}>{option.label}</option>
        {/each}
      </select>
    {/if}
    
    {#if isModbus}
      <div class="scale-input">
        <label for="scale">Scale:</label>
        <input 
          id="scale"
          type="text" 
          bind:value={scaleInput}
          placeholder="1"
        />
      </div>
    {/if}
    
    {#if inputValue}
      <button class="clear-btn" on:click={onClear}>
        ✕ Clear
      </button>
    {/if}
  </div>
  
  {#if isModbus}
    <textarea
      bind:value={inputValue}
      {placeholder}
      class:error={error}
      on:keydown={handleKeydown}
      autocomplete="off"
      spellcheck="false"
      rows="3"
    ></textarea>
  {:else}
    <input
      type="text"
      bind:value={inputValue}
      {placeholder}
      class:error={error}
      on:keydown={handleKeydown}
      autocomplete="off"
      spellcheck="false"
    />
  {/if}
  
  {#if error}
    <div class="error-message">
      {error}
    </div>
  {/if}
  
  <div class="help-text">
    {#if isModbus}
      Enter multiple 16-bit registers: <kbd>hex</kbd> (0x1234, 5678) or <kbd>d</kbd> prefix for decimal (d1000)
    {:else}
      Press <kbd>⌘K</kbd> to clear input
    {/if}
  </div>
</div>

<style>
  .input-section {
    margin-bottom: var(--spacing-md);
    flex-shrink: 0;
  }

  .mode-selector {
    display: flex;
    gap: var(--spacing-sm);
    margin-bottom: var(--spacing-sm);
    align-items: center;
    flex-wrap: wrap;
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

  .mode-btn.modbus {
    background: var(--color-int-unsigned);
    border-color: var(--color-int-unsigned);
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

  .scale-input {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
  }

  .scale-input label {
    font-size: 11px;
    color: var(--text-secondary);
    font-weight: 500;
  }

  .scale-input input {
    width: 70px;
    padding: var(--spacing-xs) var(--spacing-sm);
    font-size: 11px;
    font-family: var(--font-mono);
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    color: var(--text-primary);
    transition: all 0.15s;
  }

  .scale-input input:focus {
    outline: none;
    border-color: var(--color-int-unsigned);
    background: var(--bg-primary);
  }

  .clear-btn {
    padding: var(--spacing-xs) var(--spacing-sm);
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    color: var(--text-secondary);
    font-size: 11px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.15s;
    margin-left: auto;
  }

  .clear-btn:hover {
    background: var(--bg-hover);
    color: var(--text-primary);
  }

  input, textarea {
    width: 100%;
    padding: var(--spacing-sm) var(--spacing-md);
    font-size: 13px;
    font-family: var(--font-mono);
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    color: var(--text-primary);
    transition: all 0.15s;
    box-sizing: border-box;
  }

  textarea {
    resize: vertical;
    min-height: 80px;
    line-height: 1.5;
  }

  input:focus, textarea:focus {
    outline: none;
    border-color: var(--color-int-signed);
    background: var(--bg-primary);
  }

  input.error, textarea.error {
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
</style>
