<script>
  import { INPUT_MODES, INT_TYPE_OPTIONS } from '../lib/constants.js'
  
  export let inputMode = 'hex'
  export let intType = 'int16'
  export let inputValue = ''
  export let error = null
  export let onClear
  
  function handleKeydown(event) {
    if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
      event.preventDefault()
      if (onClear) onClear()
    }
  }
  
  $: placeholder = inputMode === 'hex' 
    ? 'Enter hex (e.g., 0x48656c6c6f or 48 65 6c 6c 6f)' 
    : inputMode === 'binary'
    ? 'Enter binary (e.g., 01001000 01100101)'
    : 'Enter integer value'
</script>

<div class="input-section">
  <div class="mode-selector">
    {#each INPUT_MODES as mode}
      <button 
        class="mode-btn" 
        class:active={inputMode === mode.value}
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
    
    {#if inputValue}
      <button class="clear-btn" on:click={onClear}>
        ✕ Clear
      </button>
    {/if}
  </div>
  
  <input
    type="text"
    bind:value={inputValue}
    {placeholder}
    class:error={error}
    on:keydown={handleKeydown}
    autocomplete="off"
    spellcheck="false"
  />
  
  {#if error}
    <div class="error-message">
      {error}
    </div>
  {/if}
  
  <div class="help-text">
    Press <kbd>⌘K</kbd> to clear input
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

  input {
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

  input:focus {
    outline: none;
    border-color: var(--color-int-signed);
    background: var(--bg-primary);
  }

  input.error {
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
