# Hexview

A lightweight desktop tool for to quickly convert hexadecimal values and perform related conversions. Built with Wails v2, Hexview provides a fast, intuitive interface for rapid conversions.

## Features

- **Comprehensive hex conversions**: Convert between hex, decimal, binary, and ASCII
- **Multiple data types**: Support for int8-64, uint8-64, float32/64
- **Endianness support**: Big-endian, little-endian, and mid-endian byte orders
- **Flexible input parsing**: Accepts various hex formats (0x prefix, spaces, colons, commas)
- **Real-time conversion**: Instant feedback as you type
- **Native desktop app**: Standalone application without browser overhead
- **Cross-platform**: Available for macOS, Windows, and Linux

## Technology Stack

- **Backend**: Go with Wails v2 framework
- **Frontend**: Svelte with Vite
- **Conversion library**: Standalone `convert` package

## Installation

### Build from Source

**Prerequisites:**
- Go 1.25 or later
- Node.js and npm
- Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

**Clone and build:**
```bash
git clone https://github.com/alexjoedt/hexview.git
cd hexview
wails build
```

The binary will be created in `build/bin/`.

## Usage

Launch the application and enter hexadecimal values in any of these formats:

```
0x48656c6c6f
48 65 6c 6c 6f
48:65:6C:6C:6F
48,65,6c,6c,6f
```

The app will automatically display conversions for:
- Signed/unsigned integers (8, 16, 32, 64-bit)
- Floating-point numbers (32, 64-bit)
- Binary representation
- ASCII text (when applicable)
- Multiple endianness formats

## Development

### Running in Development Mode

```bash
# Start Wails dev server with hot reload
wails dev

# Or use Taskfile
task dev
```

### Project Structure

```
hexview/
├── app.go              # Main application logic and conversion handlers
├── main.go             # Wails application entry point
├── convert/            # Standalone conversion package
│   ├── convert.go      # Conversion functions
│   ├── convert_test.go # Comprehensive test suite (92% coverage)
│   └── README.md       # Package documentation
├── frontend/           # Svelte UI
│   ├── src/
│   │   ├── App.svelte  # Main application component
│   │   └── components/ # Reusable UI components
│   └── wailsjs/        # Auto-generated Go bindings (do not edit)
└── Taskfile.yml        # Task automation

```

### Testing

```bash
# Run Go tests with coverage
go test --cover ./...

# Or use Taskfile
task test
```

### Building

```bash
# Build for current platform
wails build

# Build optimized production binary
task build:prod

# Build for specific platforms
task build:mac          # Universal macOS binary
task build:mac:arm      # Apple Silicon only
task build:windows      # Windows executable
task build:linux        # Linux binary
```

See `Taskfile.yml` for all available build tasks.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is open source. See the LICENSE file for details.

## Acknowledgments

Built with [Wails](https://wails.io/) - a framework for building desktop applications using Go and web technologies.
