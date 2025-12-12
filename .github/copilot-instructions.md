# Hexview - Copilot Instructions

## Project Overview

Hexview is a **lightweight desktop tool for network engineers** to quickly convert hexadecimal values and perform related conversions. Built with Wails v2, it prioritizes:
- **Simplicity**: Small, focused tool that does one thing well
- **Excellent UX**: Fast, intuitive interface for rapid conversions during network troubleshooting
- **Desktop native**: Standalone app without browser overhead

**Tech stack:**
- **Backend**: Go (Wails v2 framework) with a comprehensive `convert` package for hex/binary conversions
- **Frontend**: Svelte with Vite for the UI
- **Bridge**: Wails automatically generates TypeScript bindings from Go methods (see `frontend/wailsjs/go/`)

## Architecture

### Backend Structure
- `main.go`: Wails application entry point, embeds frontend assets
- `app.go`: Application struct with context and exported methods (e.g., `Greet`)
- `convert/`: Standalone package for hex/binary conversions (can be used independently)

### Frontend Integration
- **Auto-generated bindings**: `frontend/wailsjs/go/main/App.js` provides typed JS functions for Go methods
- **DO NOT EDIT** generated files in `frontend/wailsjs/` - they regenerate on build
- Import Go functions: `import {Greet} from '../wailsjs/go/main/App.js'`
- Call Go from Svelte: `Greet(name).then(result => ...)`

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
### Updating Frontend UI
1. Edit `frontend/src/App.svelte` for UI changes
2. Call Go functions via imported bindings
3. Use `wails dev` for live reload
4. Generated bindings auto-update on save

**UX Principles for Network Engineers:**
- Minimize clicks: Auto-focus inputs, keyboard shortcuts for common actions
- Instant feedback: Show conversions in real-time as user types
- Forgiving input: Accept various formats (0x, spaces, colons) without validation errors
- Clear output: Display results in multiple formats simultaneously
- Copy-friendly: Easy to copy results for pasting into other tools
### Updating Frontend UI
1. Edit `frontend/src/App.svelte` for UI changes
2. Call Go functions via imported bindings
3. Use `wails dev` for live reload
4. Generated bindings auto-update on save

### Building for Release
```bash
wails build              # Creates platform-specific binary
wails build -platform    # Cross-compile (darwin, windows, linux)
```
