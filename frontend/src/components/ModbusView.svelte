<script>
  import CopyButton from './CopyButton.svelte'
  
  export let result = null
  export let onCopy = () => {}
  export let scaleFactor = 1
  
  // Endianness display names for better UX
  const endianLabels = {
    BE: 'Big Endian',
    LE: 'Little Endian',
    BADC: 'Mid-Big (BADC)',
    CDAB: 'Mid-Little (CDAB)'
  }
  
  // Reactive scale function - use $: to make Svelte track scaleFactor changes
  $: applyScale = (value) => {
    if (scaleFactor === 1 || value === undefined || value === null) return value
    const num = typeof value === 'string' ? parseFloat(value) : value
    if (isNaN(num)) return value
    return (num * scaleFactor).toFixed(6).replace(/\.?0+$/, '')
  }
  
  function formatNumber(value) {
    if (value === undefined || value === null) return '—'
    return String(value)
  }
</script>

{#if result}
  <div class="modbus-view">
    <!-- Individual Registers -->
    <section class="section registers-section">
      <h3 class="section-title">
        <span class="icon">📊</span>
        Individual Registers ({result.registers?.length || 0})
      </h3>
      <div class="table-wrapper">
        <table class="data-table">
          <thead>
            <tr>
              <th class="col-index">Reg #</th>
              <th class="col-hex">Hex</th>
              <th class="col-value">Unsigned</th>
              <th class="col-value">Signed</th>
              <th class="col-binary">Binary</th>
              <th class="col-copy"></th>
            </tr>
          </thead>
          <tbody>
            {#each result.registers || [] as reg}
              <tr class="register-row">
                <td class="col-index">
                  <span class="reg-badge">{reg.index}</span>
                </td>
                <td class="col-hex mono">{reg.hex}</td>
                <td class="col-value mono">{applyScale(reg.unsigned)}</td>
                <td class="col-value mono">{applyScale(reg.signed)}</td>
                <td class="col-binary mono">{reg.binary}</td>
                <td class="col-copy">
                  <CopyButton value={reg.hex} {onCopy} />
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </section>

    <!-- 32-bit Combined Values -->
    {#if result.combined32?.length > 0}
      <section class="section combined-section">
        <h3 class="section-title">
          <span class="icon">🔗</span>
          32-bit Combined (Register Pairs)
        </h3>
        {#each result.combined32 as combo}
          <div class="combo-card">
            <div class="combo-header">
              <span class="combo-regs">
                Reg {combo.registerStart} + {combo.registerStart + 1}
              </span>
              <span class="combo-hex mono">{combo.hex}</span>
              <CopyButton value={combo.hex} {onCopy} />
            </div>
            <div class="combo-grid">
              <div class="combo-group">
                <h4>Unsigned INT32</h4>
                <div class="value-row">
                  <span class="endian-label">BE</span>
                  <span class="value mono">{applyScale(combo.uint32BE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">LE</span>
                  <span class="value mono">{applyScale(combo.uint32LE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">BADC</span>
                  <span class="value mono">{applyScale(combo.uint32BADC)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">CDAB</span>
                  <span class="value mono">{applyScale(combo.uint32CDAB)}</span>
                </div>
              </div>
              <div class="combo-group">
                <h4>Signed INT32</h4>
                <div class="value-row">
                  <span class="endian-label">BE</span>
                  <span class="value mono">{applyScale(combo.int32BE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">LE</span>
                  <span class="value mono">{applyScale(combo.int32LE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">BADC</span>
                  <span class="value mono">{applyScale(combo.int32BADC)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">CDAB</span>
                  <span class="value mono">{applyScale(combo.int32CDAB)}</span>
                </div>
              </div>
              <div class="combo-group">
                <h4>FLOAT32</h4>
                <div class="value-row">
                  <span class="endian-label">BE</span>
                  <span class="value mono float-value">{applyScale(combo.float32BE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">LE</span>
                  <span class="value mono float-value">{applyScale(combo.float32LE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">BADC</span>
                  <span class="value mono float-value">{applyScale(combo.float32BADC)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">CDAB</span>
                  <span class="value mono float-value">{applyScale(combo.float32CDAB)}</span>
                </div>
              </div>
            </div>
          </div>
        {/each}
      </section>
    {/if}

    <!-- 64-bit Combined Values -->
    {#if result.combined64?.length > 0}
      <section class="section combined-section">
        <h3 class="section-title">
          <span class="icon">🔗</span>
          64-bit Combined (4 Registers)
        </h3>
        {#each result.combined64 as combo}
          <div class="combo-card">
            <div class="combo-header">
              <span class="combo-regs">
                Reg {combo.registerStart}–{combo.registerStart + 3}
              </span>
              <span class="combo-hex mono">{combo.hex}</span>
              <CopyButton value={combo.hex} {onCopy} />
            </div>
            <div class="combo-grid">
              <div class="combo-group">
                <h4>Unsigned INT64</h4>
                <div class="value-row">
                  <span class="endian-label">BE</span>
                  <span class="value mono">{applyScale(combo.uint64BE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">LE</span>
                  <span class="value mono">{applyScale(combo.uint64LE)}</span>
                </div>
              </div>
              <div class="combo-group">
                <h4>Signed INT64</h4>
                <div class="value-row">
                  <span class="endian-label">BE</span>
                  <span class="value mono">{applyScale(combo.int64BE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">LE</span>
                  <span class="value mono">{applyScale(combo.int64LE)}</span>
                </div>
              </div>
              <div class="combo-group">
                <h4>FLOAT64</h4>
                <div class="value-row">
                  <span class="endian-label">BE</span>
                  <span class="value mono float-value">{applyScale(combo.float64BE)}</span>
                </div>
                <div class="value-row">
                  <span class="endian-label">LE</span>
                  <span class="value mono float-value">{applyScale(combo.float64LE)}</span>
                </div>
              </div>
            </div>
          </div>
        {/each}
      </section>
    {/if}

    <!-- Raw Data -->
    <section class="section raw-section">
      <h3 class="section-title">
        <span class="icon">📝</span>
        Raw Data
      </h3>
      <div class="raw-grid">
        <div class="raw-item">
          <span class="raw-label">Hex</span>
          <span class="raw-value mono">{result.rawHex}</span>
          <CopyButton value={result.rawHex} {onCopy} />
        </div>
        <div class="raw-item">
          <span class="raw-label">ASCII</span>
          <span class="raw-value mono ascii">{result.ascii}</span>
          <CopyButton value={result.ascii} {onCopy} />
        </div>
      </div>
    </section>
  </div>
{/if}

<style>
  .modbus-view {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
  }

  .section {
    background: var(--bg-secondary);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    overflow: hidden;
  }

  .section-title {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    padding: var(--spacing-sm) var(--spacing-md);
    margin: 0;
    font-size: 12px;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    color: var(--text-secondary);
    background: var(--bg-tertiary);
    border-bottom: 1px solid var(--border-color);
  }

  .section-title .icon {
    font-size: 14px;
  }

  /* Table Styles */
  .table-wrapper {
    overflow-x: auto;
  }

  .data-table {
    width: 100%;
    border-collapse: collapse;
    font-size: 12px;
  }

  .data-table th,
  .data-table td {
    padding: var(--spacing-xs) var(--spacing-sm);
    text-align: left;
    border-bottom: 1px solid var(--border-color);
  }

  .data-table th {
    font-weight: 600;
    color: var(--text-tertiary);
    font-size: 10px;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    background: var(--bg-tertiary);
  }

  .data-table tbody tr:last-child td {
    border-bottom: none;
  }

  .register-row:hover {
    background: var(--bg-hover);
  }

  .col-index { width: 60px; text-align: center; }
  .col-hex { width: 80px; }
  .col-value { width: 100px; }
  .col-binary { min-width: 150px; }
  .col-copy { width: 40px; text-align: center; }

  .reg-badge {
    display: inline-block;
    padding: 2px 8px;
    background: var(--color-int-unsigned);
    color: white;
    border-radius: var(--radius-sm);
    font-size: 10px;
    font-weight: 600;
  }

  .mono {
    font-family: var(--font-mono);
  }

  /* Combined Cards */
  .combo-card {
    border-bottom: 1px solid var(--border-color);
    padding: var(--spacing-md);
  }

  .combo-card:last-child {
    border-bottom: none;
  }

  .combo-header {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    margin-bottom: var(--spacing-sm);
    padding-bottom: var(--spacing-sm);
    border-bottom: 1px dashed var(--border-color);
  }

  .combo-regs {
    font-weight: 600;
    color: var(--color-int-signed);
    font-size: 13px;
  }

  .combo-hex {
    color: var(--text-secondary);
    font-size: 12px;
    background: var(--bg-tertiary);
    padding: 2px 8px;
    border-radius: var(--radius-sm);
  }

  .combo-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
    gap: var(--spacing-md);
  }

  .combo-group h4 {
    margin: 0 0 var(--spacing-xs) 0;
    font-size: 11px;
    font-weight: 600;
    color: var(--text-tertiary);
    text-transform: uppercase;
    letter-spacing: 0.3px;
  }

  .value-row {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    padding: 2px 0;
  }

  .endian-label {
    font-size: 10px;
    font-weight: 500;
    color: var(--text-tertiary);
    min-width: 40px;
  }

  .value {
    font-size: 12px;
    color: var(--text-primary);
  }

  .float-value {
    color: var(--color-float);
  }

  /* Raw Section */
  .raw-grid {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xs);
    padding: var(--spacing-md);
  }

  .raw-item {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    padding: var(--spacing-xs);
    background: var(--bg-tertiary);
    border-radius: var(--radius-sm);
  }

  .raw-label {
    font-size: 10px;
    font-weight: 600;
    color: var(--text-tertiary);
    text-transform: uppercase;
    min-width: 50px;
  }

  .raw-value {
    flex: 1;
    font-size: 12px;
    color: var(--text-primary);
    word-break: break-all;
  }

  .ascii {
    color: var(--color-binary);
    letter-spacing: 1px;
  }
</style>
