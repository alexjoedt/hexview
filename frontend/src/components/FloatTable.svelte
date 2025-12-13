<script>
  import CopyButton from './CopyButton.svelte'
  import { formatValue, hasValue } from '../lib/utils.js'
  import { FLOAT_TYPES, getResultValue, getResultHex } from '../lib/constants.js'
  
  export let result
  export let expertMode = false
  export let onCopy
</script>

<div class="table-wrapper">
  <div class="table-title">Floating Point Conversions</div>
  <table>
    <thead>
      <tr>
        <th>Type</th>
        <th>Big-Endian (BE)</th>
        {#if expertMode}<th>BE Hex</th>{/if}
        <th>Little-Endian (LE)</th>
        {#if expertMode}<th>LE Hex</th>{/if}
        <th>Mid-Big (BADC)</th>
        {#if expertMode}<th>BADC Hex</th>{/if}
        <th>Mid-Little (CDAB)</th>
        {#if expertMode}<th>CDAB Hex</th>{/if}
      </tr>
    </thead>
    <tbody>
      {#each FLOAT_TYPES as type}
        {@const values = type.endianness.map(endian => ({
          endian,
          value: getResultValue(result, type.key, endian),
          hex: getResultHex(result, type.key, endian)
        }))}
        {@const hasAnyValue = values.some(v => hasValue(v.value))}
        
        <tr class:unavailable={!hasAnyValue} class:highlighted={hasAnyValue}>
          <td class="type-cell">
            <span class="type-badge float">{type.name}</span>
          </td>
          
          {#each values as {value, hex}}
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
      {/each}
    </tbody>
  </table>
</div>

<style>
  .table-wrapper {
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    overflow-x: hidden;
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
    font-size: 10px;
    text-transform: uppercase;
    letter-spacing: 0.03em;
    color: var(--text-secondary);
    padding: var(--spacing-xs) var(--spacing-sm);
    border-bottom: 1px solid var(--border-color);
  }

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
    background: rgba(220, 38, 38, 0.08) !important;
    border-left: 3px solid var(--color-float);
  }

  :global(.dark) tr.highlighted {
    background: rgba(248, 113, 113, 0.12) !important;
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

  .type-badge.float {
    background: rgba(220, 38, 38, 0.12);
    color: var(--color-float);
  }
</style>
