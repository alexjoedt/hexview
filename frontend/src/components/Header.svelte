<script>
  export let darkMode = false
  export let expertMode = false
  export let onThemeToggle
</script>

<header>
  <h1>⬡ Hexview</h1>
  
  <div class="header-controls">
    <label class="expert-mode-toggle">
      <input type="checkbox" bind:checked={expertMode} />
      <span class="checkbox-label">Expert Mode</span>
    </label>
    
    <button class="theme-toggle" on:click={onThemeToggle}>
      {darkMode ? '☀️' : '🌙'}
    </button>
  </div>
</header>

<style>
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

  .header-controls {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
  }

  .expert-mode-toggle {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    cursor: pointer;
    user-select: none;
  }

  .expert-mode-toggle input[type="checkbox"] {
    width: 16px;
    height: 16px;
    cursor: pointer;
    accent-color: var(--color-int-signed);
  }

  .checkbox-label {
    font-size: 12px;
    font-weight: 500;
    color: var(--text-secondary);
    transition: color 0.15s;
  }

  .expert-mode-toggle:hover .checkbox-label {
    color: var(--text-primary);
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
</style>
