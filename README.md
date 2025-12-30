<div align="center">
  <img src="./logo.svg" width="240" alt="packrat (jspackr logo)" />
</div>

<h1 align="center">jspackr 0.3</h1>

<p align="center">
  A simple, fast, and modern JavaScript bundler powered by esbuild.<br />
  Lightweight, easy to use, and perfect for small to medium projects.
</p>

<p align="center">
  <a href="#-installation">Installation</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-configuration-file">Configuration</a> •
  <a href="#-examples">Examples</a>
</p>

---

## ✨ Features

- ⚡ **Lightning Fast** - Bundles JavaScript in milliseconds using esbuild
- 🔧 **Zero Config** - Works out of the box, no configuration needed
- ✂️ **Minification** - Optional built-in minification
- 📊 **Build Reports** - See exactly what's contributing to your bundle size
- 👀 **Watch Mode** - Auto-rebuild on file changes during development
- 🗺️ **Source Maps** - Debug with linked or inline source maps
- 🎨 **Beautiful Output** - Colored CLI with helpful messages
- 🤖 **Non-Interactive Mode** - Perfect for CI/CD pipelines
- 🔄 **Config Files** - Support for JSON configuration files

---

## 📋 Prerequisites

- **Node.js**: v18 or higher (required for npm installation and Go binary)
- **Go**: v1.21 or higher (only if building from source)
- **Operating System**: Windows, macOS, or Linux

---

## 🚀 Installation

### Option 1: Install via npm (Recommended)

```bash
# Install globally
npm install -g jspackr

# Verify installation
jspackr --version
```

### Option 2: Install via Go

```bash
# Install latest version
go install github.com/kalokaradia/jspackr@latest

# Verify installation
jspackr --version
```

### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/kalokaradia/jspackr.git
cd jspackr

# Build the binary
go build -o bin/jspackr src/main/main.go

# Make it executable (Linux/macOS)
chmod +x bin/jspackr

# Run directly
./bin/jspackr --version
```

### Option 4: Manual Installation

Download the appropriate binary for your platform from the [releases page](https://github.com/kalokaradia/jspackr/releases):

```bash
# Linux/macOS
sudo mv jspackr /usr/local/bin/jspackr

# Windows
# Move jspackr.exe to your PATH
```

---

## 🎯 Quick Start

### Basic Bundling

Bundle a JavaScript file with default settings:

```bash
jspackr src/index.js
```

This creates:

```
dist/bundle.js
```

### With Options

```bash
jspackr -i src/index.js -o dist/app.js -m -r -w
```

This will:

- 📥 Bundle `src/index.js`
- 📤 Write to `dist/app.js`
- ✂️ Minify the output
- 📊 Show build report
- 👀 Watch for changes

---

## 📖 Usage

### Command Syntax

```bash
jspackr [options]
```

### CLI Options

| Short | Long Form             | Description                                 | Default          |
| ----- | --------------------- | ------------------------------------------- | ---------------- |
| `-i`  | `--input <file>`      | Entry JavaScript file                       | Required         |
| `-o`  | `--out <file>`        | Output bundle file                          | `dist/bundle.js` |
| `-c`  | `--config <file>`     | Path to config file                         | Optional         |
| `-m`  | `--minify`            | Minify the output                           | `false`          |
| `-r`  | `--report`            | Generate build report                       | `false`          |
| `-s`  | `--source <mode>`     | Source map mode: `none`, `linked`, `inline` | `none`           |
| `-w`  | `--watch`             | Enable watch mode                           | `false`          |
|       | `--log-level <level>` | Log level: `debug`, `info`, `warn`, `error` | `info`           |
| `-f`  | `--force`             | Force overwrite without confirmation        | `false`          |
| `-y`  | `--yes`               | Auto-confirm all prompts                    | `false`          |
| `-n`  | `--no-confirm`        | Skip all confirmation prompts               | `false`          |
| `-v`  | `--version`           | Show version (standalone)                   | -                |
| `-h`  | `--help`              | Show help message                           | -                |

### Help Command

```bash
# Show help
jspackr --help

# Show version
jspackr --version
```

---

## ⚙️ Configuration File

Create a `jspackr.config.json` file for persistent configuration:

```json
{
	"input": "./src/index.js",
	"output": "./dist/bundle.js",
	"minify": true,
	"report": true,
	"sourceMap": "linked",
	"watch": false,
	"logLevel": "info"
}
```

### Config File Options

| Option      | Type    | Description                                     |
| ----------- | ------- | ----------------------------------------------- |
| `input`     | string  | Entry JavaScript file path                      |
| `output`    | string  | Output bundle file path                         |
| `minify`    | boolean | Minify the output bundle                        |
| `report`    | boolean | Generate build report                           |
| `sourceMap` | string  | Source map mode: `none`, `linked`, `inline`     |
| `watch`     | boolean | Enable watch mode                               |
| `logLevel`  | string  | Log verbosity: `debug`, `info`, `warn`, `error` |

### Using Config File

```bash
# Use default config (jspackr.config.json)
jspackr

# Specify custom config file
jspackr -c custom.config.json
```

---

## 💡 Examples

### Example 1: Basic Bundle

```bash
# Bundle a single file
jspackr -i src/index.js

# With custom output
jspackr -i src/index.js -o dist/app.js
```

### Example 2: Production Build

```bash
# Minified with source map and report
jspackr -i src/index.js -o dist/app.min.js -m -r -s inline
```

### Example 3: Development with Watch Mode

```bash
# Watch for changes and rebuild automatically
jspackr -i src/index.js -o dist/app.js -w
```

### Example 4: CI/CD / Non-Interactive

```bash
# Force overwrite, skip all confirmations
jspackr -i src/index.js -o dist/app.js -f

# Or with yes flag
jspackr -i src/index.js -o dist/app.js -y

# Or completely non-interactive
jspackr -i src/index.js -o dist/app.js -n
```

### Example 5: Using Configuration File

**Create `jspackr.config.json`:**

```json
{
	"input": "./src/app.js",
	"output": "./build/app.js",
	"minify": true,
	"report": true
}
```

**Run:**

```bash
jspackr
```

### Example 6: Verbose Logging

```bash
# Debug mode
jspackr -i src/index.js --log-level debug

# Quiet mode
jspackr -i src/index.js --log-level warn
```

### Example 7: Multiple Options Combined

```bash
# Production-ready bundle with all features
jspackr \
  --input src/index.js \
  --output dist/app.js \
  --minify \
  --report \
  --source linked \
  --log-level info
```

---

## 📊 Build Report

When using the `--report` flag, jspackr generates a detailed breakdown:

```
Bundle Size: 45.2 KB
⏱️  Build Time: 123ms

📦 Dependencies:
  src/utils.js      12.3 KB
  src/helpers.js    8.1 KB
  src/index.js      2.4 KB
  ...
```

---

## 🗺️ Source Maps

| Mode   | Flag Value | Description                           |
| ------ | ---------- | ------------------------------------- |
| None   | `none`     | No source map (default)               |
| Linked | `linked`   | Separate .map file linked via comment |
| Inline | `inline`   | Embedded as base64 in the bundle      |

```bash
# Linked source map
jspackr -i src/index.js -s linked

# Inline source map
jspackr -i src/index.js -s inline
```

---

## 📁 Project Structure

```
jspackr/
├── bin/                    # Compiled binaries
├── src/
│   ├── cli/               # CLI output and styling
│   │   ├── logger.go      # Logging functionality
│   │   ├── styles.go      # Colored output styles
│   │   └── ui.go          # UI components
│   ├── config/            # Configuration management
│   │   ├── config.go      # Config structures
│   │   ├── loader.go      # JSON config loading
│   │   ├── merger.go      # Config merging
│   │   └── validator.go   # Config validation
│   ├── core/
│   │   ├── builder/       # Bundling logic
│   │   │   ├── builder.go # Main builder
│   │   │   ├── report.go  # Build reporting
│   │   │   └── sourcemap.go # Source map handling
│   │   └── watcher/       # File watching
│   │       ├── debouncer.go
│   │       ├── hasher.go
│   │       └── watcher.go
│   ├── main/
│   │   └── main.go        # Entry point
│   └── utils/
│       ├── confirm.go     # Confirmation prompts
│       ├── file.go        # File utilities
│       └── flags.go       # CLI flags parsing
├── .gitignore
├── .npmignore
├── go.mod
├── go.sum
├── LICENSE
├── logo.svg
├── package.json
└── README.md
```

---

## 🐛 Troubleshooting

### Command Not Found

If `jspackr` is not found after installation:

```bash
# Add to PATH (npm global)
npm config get prefix
# Add the output path to your shell profile

# Or for Go installation
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Permission Denied

```bash
# Linux/macOS
sudo chown $(whoami) /usr/local/bin/jspackr

# Or use npm with sudo (not recommended)
sudo npm install -g jspackr
```

### Node Version Too Old

```bash
# Check version
node --version

# Update using nvm (Linux/macOS)
nvm install 18
nvm use 18

# Or download from nodejs.org
```

### Build Fails

```bash
# Clear cache and rebuild
go clean -cache
go build -o bin/jspackr src/main/main.go

# Check dependencies
go mod download
go mod verify
```

### Watch Mode Not Working

```bash
# Ensure you're not using network drives
# Check file permissions
ls -la src/

# Try with absolute paths
jspackr -i /absolute/path/to/index.js
```

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

```bash
# Fork the repository
git clone https://github.com/YOUR-USERNAME/jspackr.git
cd jspackr

# Create a feature branch
git checkout -b feature/amazing-feature

# Make changes
# Test your changes
go test ./...

# Commit your changes
git commit -m "Add amazing feature"

# Push to GitHub
git push origin feature/amazing-feature

# Open a Pull Request
```

### Coding Standards

- Follow Go formatting conventions (`gofmt`)
- Add comments for public functions
- Write tests for new features
- Update documentation as needed

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- [esbuild](https://esbuild.github.io/) for the amazing bundling engine
- [fatih/color](https://github.com/fatih/color) for colored terminal output
- All contributors and users!

---

## 📧 Contact

- **Author**: [Kaloka Radia Nanda](https://kalokaradiananda.my.id)
- **GitHub**: [@kalokaradia](https://github.com/kalokaradia)
- **Issues**: [Report a bug](https://github.com/kalokaradia/jspackr/issues)

---

<p align="center">
  Made with ❤️ by <b>Kaloka Radia Nanda</b>
</p>
