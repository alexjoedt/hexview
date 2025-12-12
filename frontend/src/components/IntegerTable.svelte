<script>
  import TableRow from './TableRow.svelte'
  import { INT_TYPES } from '../lib/constants.js'
  
  export let result
  export let expertMode = false
  export let onCopy
</script>

<div class="table-wrapper">
  <div class="table-title">Integer Conversions</div>
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
      {#each INT_TYPES as type}
        <TableRow {type} {result} {expertMode} {onCopy} />
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
    font-size: 11px;
    text-transform: uppercase;
    letter-spacing: 0.03em;
    color: var(--text-secondary);
    padding: var(--spacing-sm) var(--spacing-md);
    border-bottom: 1px solid var(--border-color);
  }
</style>
