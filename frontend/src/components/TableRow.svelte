<script>
  import CopyButton from './CopyButton.svelte'
  import { formatValue, hasValue } from '../lib/utils.js'
  import { getResultValue, getResultHex } from '../lib/constants.js'
  
  export let type
  export let result
  export let expertMode = false
  export let onCopy
  
  $: values = type.endianness.map(endian => ({
    endian,
    value: getResultValue(result, type.key, endian),
    hex: getResultHex(result, type.key, endian)
  }))
  
  $: hasAnyValue = values.some(v => hasValue(v.value))
  $: highlighted = hasAnyValue
</script>

<tr class:unavailable={!hasAnyValue} class:highlighted={highlighted}>
  <td class="type-cell">
    <span class="type-badge {type.signed ? 'int-signed' : 'int-unsigned'}">
      {type.name}
    </span>
  </td>
  
  {#each values as {endian, value, hex}}
    <td class="value-cell-with-copy">
      <span class="value-text">{formatValue(value)}</span>
      {#if hasValue(value)}
        <CopyButton {value} {onCopy} />
      {/if}
    </td>
    
    {#if expertMode}
      <td class="hex-cell">
        <span class="mono-hex">{formatValue(hex)}</span>
      </td>
    {/if}
  {/each}
</tr>

<style>
  tr {
    background: var(--bg-primary);
    transition: background 0.1s;
  }

  tr:hover {
    background: var(--bg-hover);
  }

  tr.unavailable {
    opacity: 0.4;
  }

  tr.highlighted {
    background: rgba(37, 99, 235, 0.08) !important;
    border-left: 3px solid var(--color-int-signed);
  }

  :global(.dark) tr.highlighted {
    background: rgba(96, 165, 250, 0.12) !important;
  }

  tr:not(:last-child) td {
    border-bottom: 1px solid var(--border-color);
  }

  td {
    padding: var(--spacing-xs) var(--spacing-sm);
  }

  .type-cell {
    width: 90px;
  }

  .hex-cell {
    background: var(--bg-secondary);
  }

  .mono-hex {
    font-family: var(--font-mono);
    font-size: 11px;
    color: var(--text-secondary);
  }

  .value-cell-with-copy {
    font-family: var(--font-mono);
    font-size: 12px;
    color: var(--text-primary);
    position: relative;
    padding-right: 28px;
  }

  .value-text {
    display: inline-block;
  }

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
</style>
