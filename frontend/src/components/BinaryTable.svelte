<script>
  import CopyButton from "./CopyButton.svelte";
  import { formatValue, hasValue } from "../lib/utils.js";

  export let result;
  export let onCopy;
</script>

<div class="table-wrapper">
  <div class="table-title">Binary Representations</div>
  <table>
    <thead>
      <tr>
        <th>Type</th>
        <th>Value</th>
      </tr>
    </thead>
    <tbody>
      <tr class:unavailable={!hasValue(result?.bytes)}>
        <td class="type-cell">
          <span class="type-badge bytes">HEX</span>
        </td>
        <td class="value-cell-with-copy mono wide">
          <span class="value-text">{formatValue(result?.bytes)}</span>
          {#if hasValue(result?.bytes)}
            <CopyButton value={result.bytes} {onCopy} />
          {/if}
        </td>
      </tr>
      <tr class:unavailable={!hasValue(result?.binary)}>
        <td class="type-cell">
          <span class="type-badge binary">BINARY</span>
        </td>
        <td class="value-cell-with-copy mono wide">
          <span class="value-text">{formatValue(result?.binary)}</span>
          {#if hasValue(result?.binary)}
            <CopyButton value={result.binary} {onCopy} />
          {/if}
        </td>
      </tr>
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

  tbody tr:not(:last-child) td {
    border-bottom: 1px solid var(--border-color);
  }

  td {
    padding: var(--spacing-sm) var(--spacing-md);
  }

  .type-cell {
    width: 90px;
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

  .value-cell-with-copy.wide::-webkit-scrollbar {
    height: 4px;
  }

  .value-cell-with-copy.wide::-webkit-scrollbar-track {
    background: transparent;
  }

  .value-cell-with-copy.wide::-webkit-scrollbar-thumb {
    background: var(--border-color);
    border-radius: 2px;
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

  .type-badge.binary {
    background: rgba(124, 58, 237, 0.12);
    color: var(--color-binary);
  }

  .type-badge.bytes {
    background: rgba(234, 88, 12, 0.12);
    color: var(--color-bytes);
  }
</style>
