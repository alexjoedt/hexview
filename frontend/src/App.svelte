<script>
  import Header from './components/Header.svelte'
  import InputSection from './components/InputSection.svelte'
  import IntegerTable from './components/IntegerTable.svelte'
  import FloatTable from './components/FloatTable.svelte'
  import BinaryTable from './components/BinaryTable.svelte'
  import ModbusView from './components/ModbusView.svelte'
  import Toast from './components/Toast.svelte'
  import { convert } from './lib/api.js'
  
  // State
  let inputMode = 'hex'
  let intType = 'int16'
  let inputValue = ''
  let result = null
  let modbusResult = null
  let error = null
  let isLoading = false
  let debounceTimer = null
  let scaleFactor = 1
  
  // Theme state - load from localStorage or fall back to system preference
  let darkMode = (() => {
    const saved = localStorage.getItem('hexview-theme')
    if (saved !== null) {
      return saved === 'dark'
    }
    return window.matchMedia('(prefers-color-scheme: dark)').matches
  })()
  
  // Expert mode state
  let expertMode = false
  
  // Toast state
  let toastMessage = ''
  let showToast = false
  
  // Reactive conversion with debounce
  $: {
    if (inputValue.trim() === '') {
      result = null
      modbusResult = null
      error = null
    } else {
      debouncedConvert(inputValue, inputMode, intType)
    }
  }
  
  function debouncedConvert(input, mode, type) {
    clearTimeout(debounceTimer)
    isLoading = true
    
    debounceTimer = setTimeout(async () => {
      try {
        const convertResult = await convert(input, mode, type)
        if (mode === 'modbus') {
          modbusResult = convertResult
          result = null
        } else {
          result = convertResult
          modbusResult = null
        }
        error = null
      } catch (err) {
        error = err.message || String(err)
        result = null
        modbusResult = null
      } finally {
        isLoading = false
      }
    }, 300)
  }
  
  function toggleTheme() {
    darkMode = !darkMode
    // Save theme preference to localStorage
    localStorage.setItem('hexview-theme', darkMode ? 'dark' : 'light')
  }
  
  function handleCopy(success) {
    toastMessage = success ? 'Copied!' : 'Failed to copy'
    showToast = true
    setTimeout(() => showToast = false, 1500)
  }
  
  function clearInput() {
    inputValue = ''
    result = null
    modbusResult = null
    error = null
  }
</script>

<main class:dark={darkMode}>
  <div class="container">
    <Header 
      {darkMode} 
      bind:expertMode
      onThemeToggle={toggleTheme}
    />
    
    <InputSection 
      bind:inputMode
      bind:intType
      bind:inputValue
      bind:scaleFactor
      {error}
      onClear={clearInput}
    />
    
    {#if inputMode === 'modbus' && modbusResult && !error}
      <section class="results">
        <ModbusView result={modbusResult} {scaleFactor} onCopy={handleCopy} />
      </section>
    {:else if result && !error}
      <section class="results">
        <IntegerTable {result} {expertMode} onCopy={handleCopy} />
        <FloatTable {result} {expertMode} onCopy={handleCopy} />
        <BinaryTable {result} onCopy={handleCopy} />
      </section>
    {/if}
    
    <Toast message={toastMessage} show={showToast} />
  </div>
</main>

<style>
  main {
    min-height: 100vh;
    background: var(--bg-primary);
    color: var(--text-primary);
    transition: background 0.3s, color 0.3s;
  }
  
  .container {
    height: 100vh;
    display: flex;
    flex-direction: column;
    padding: var(--spacing-md);
  }

  .results {
    flex: 1;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: var(--spacing-sm);
  }

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
</style>
