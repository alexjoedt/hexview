# Hexview - Copilot Instructions

## Project Overview

Hexview is a **lightweight desktop tool for network engineers** to quickly convert hexadecimal values and perform related conversions. Built with Wails v2, it prioritizes:
- **Simplicity**: Small, focused tool that does one thing well
- **Excellent UX**: Fast, intuitive interface for rapid conversions during network troubleshooting
- **Desktop native**: Standalone app without browser overhead

**Tech stack:**
- **Backend**: Go (Wails v2 framework) with layered architecture
- **Frontend**: Svelte with Vite for the UI
- **Bridge**: Wails automatically generates TypeScript bindings from Go methods (see `frontend/wailsjs/go/`)

## Architecture (3-Layer Design)

### 1. Wails App Layer (`app.go`)
Thin glue layer exposing Go methods to frontend. All exported methods become JS functions via Wails bindings.
```go
func (a *App) ConvertHex(hexInput string) (*models.ConversionResult, error)
```

### 2. Service Layer (`service/converter.go`)
Business logic that orchestrates conversions and builds comprehensive result structures. Calls `convert` package functions and assembles `models.ConversionResult` with all endianness variants (BE, LE, BADC, CDAB).

### 3. Convert Package (`convert/`)
Pure conversion algorithms. **Can be used standalone** (no Wails dependency). Provides:
- Flexible hex parsing (accepts `0x`, spaces, colons, commas)
- Generic functions with type constraints: `hexToInt[T integer]()`
- Endianness naming: default BE, `*LE()` suffix for little-endian, `*BADC()` and `*CDAB()` for mid-endian

### Data Flow
Frontend → `frontend/wailsjs/go/main/App.js` (auto-generated) → `app.go` → `service/converter.go` → `convert/` → `models.ConversionResult` → Frontend

### Frontend Integration
- **Auto-generated bindings**: `frontend/wailsjs/go/main/App.js` provides typed JS functions
- **DO NOT EDIT** generated files in `frontend/wailsjs/` - regenerate on build
- Wrapper layer: `frontend/src/lib/api.js` provides app-specific abstractions over raw bindings
- Debounced conversion: App.svelte uses 300ms debounce for real-time conversion

## Development Workflow

### Running the App
```bash
# Live development with hot reload
wails dev

# Build production binary
wails build

# Alternative: Use Taskfile
task test    # Run Go tests with coverage
task build   # Build binary to bin/
task tidy    # Format and tidy Go code
```

### Testing
- **Go tests**: Located in `convert/convert_test.go` (1061 lines, 92% coverage)
- **Test pattern**: Table-driven tests with descriptive names
- **Run tests**: `go test --cover ./...` or `task test`

### Frontend Development
```bash
cd frontend
npm install          # Install dependencies
npm run dev         # Vite dev server (standalone)
npm run build       # Production build
```

## Convert Package Patterns

### Key Design Principles
1. **Flexible input parsing**: `ParseHex()` accepts multiple formats (0x, spaces, commas, colons)
   ```go
   // All valid: "0x48656c6c6f", "48 65 6c 6c 6f", "48:65:6C:6C:6F"
   bytes, err := convert.HexToBytes("0x48 65 6c 6c 6f")
   ```

2. **Generic helpers**: Use type constraints for DRY code
   ```go
   // Generic conversion functions (internal)
   func hexToInt[T integer](hexStr string, byteSize int, endian binary.ByteOrder) (T, error)
   ```

3. **Endianness**: Default big-endian, explicit `LE` suffix for little-endian
   ```go
   val := convert.HexToInt32("7fffffff")      // Big-endian
   val := convert.HexToInt32LE("ffffff7f")    // Little-endian
   ```

4. **Comprehensive error handling**: Specific error types exported
   - `ErrInvalidHexChar`, `ErrInvalidLength`, `ErrOverflow`, `ErrEmptyInput`, `ErrInvalidBinaryChar`

### Test Conventions
- Use table-driven tests with `name`, `input`, `want`, `wantErr` fields
- Group tests by function with clear separators: `// ============================================================================`
- Test edge cases: max/min values, negative numbers, wrong lengths, invalid chars
- Helper functions: `bytesEqual()`, `floatsEqual()` for comparisons

## Wails-Specific Patterns

### Exporting Go Methods to Frontend
1. **Add method to App struct** in `app.go`:
   ```go
   func (a *App) MethodName(param string) string {
       return result
   }
   ```
2. **Rebuild** to regenerate TypeScript bindings
3. **Import in Svelte**: Bindings appear automatically in `wailsjs/go/main/App.js`

### Context Usage
- `app.ctx` is set during `startup()` lifecycle hook
- Use for runtime methods: `runtime.LogInfo(a.ctx, "message")`

## Code Style (Go)

- Follow **gofmt** formatting (enforced by `task tidy`)
- Package documentation at top of files with examples
- Exported functions have full doc comments
- Error messages use `fmt.Errorf` with wrapped errors: `fmt.Errorf("%w: %v", ErrType, details)`

## Build Configuration

- **Taskfile.yml**: Primary task runner (test, build, tidy, clean)
- **wails.json**: Wails-specific config (frontend commands, output name)
- **go.mod**: Module name is `hexview` (not a full path)

## Common Tasks

### Adding New Conversion Function
1. **Add logic to `service/converter.go`**: Implement conversion logic, call `convert` package functions
2. **Extend `models/result.go`**: Add new fields to `ConversionResult` (use pointers for optional values)
3. **Update `app.go`** (if needed): Add new exported method for frontend
4. **Rebuild**: `wails dev` auto-regenerates TypeScript bindings
5. **Frontend**: Import from `wailsjs/go/main/App.js` and update Svelte components

**Example**: Adding a new Modbus register conversion required:
- New `ModbusResult` struct in `models/result.go` with fields for registers, combined 32/64-bit values
- `ConvertModbusRegisters()` method in both `service/converter.go` and `app.go`
- Frontend component `ModbusView.svelte` consuming the new API

### Updating Frontend UI
1. Edit components in `frontend/src/components/` (tables, input sections)
2. Main app logic in `frontend/src/App.svelte`
3. Use `wails dev` for live reload (backend + frontend hot reload)
4. Theme handling: Uses localStorage + `prefers-color-scheme` (see App.svelte lines 22-30)

**UX Principles for Network Engineers:**
- Debounced conversion (300ms): Real-time feedback without excessive backend calls
- Forgiving input: Parse multiple hex formats in `convert.ParseHex()`
- Copy-friendly: CopyButton component with toast notifications
- Multi-format display: Show BE, LE, BADC, CDAB endianness simultaneously

### Building for Release
```bash
wails build                          # Current platform
wails build -platform darwin/universal  # macOS universal binary
task build:prod                      # Stripped debug symbols (-w -s)
```
See `Taskfile.yml` for cross-platform build targets (requires native platform or CI/CD).
