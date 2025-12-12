<script>
  import { copyToClipboard } from '../lib/utils.js'
  
  export let value
  export let onCopy = null
  
  async function handleCopy() {
    try {
      await copyToClipboard(value)
      if (onCopy) onCopy(true)
    } catch (err) {
      if (onCopy) onCopy(false)
    }
  }
</script>

<button class="copy-btn-inline" on:click={handleCopy}>
  📋
</button>

<style>
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

  :global(tr:hover) .copy-btn-inline {
    opacity: 0.5;
  }

  .copy-btn-inline:hover {
    opacity: 1 !important;
    transform: translateY(-50%) scale(1.15);
  }

  .copy-btn-inline:active {
    transform: translateY(-50%) scale(0.9);
  }
</style>
