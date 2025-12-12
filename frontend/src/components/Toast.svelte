<script>
  export let message = ''
  export let show = false
</script>

{#if show}
  <div class="toast">
    {message}
  </div>
{/if}

<style>
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
</style>
