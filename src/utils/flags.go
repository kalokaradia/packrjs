package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/kalokaradia/jspackr/src/config"
)

// ParseFlags parses command line flags and returns configuration
func ParseFlags() (*config.Config, string, bool, bool) {
	cfg := &config.Config{}
	var configPath string
	var showVersion bool
	var help bool

	flag.StringVar(&configPath, "c", "", "Path to config file")
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.StringVar(&cfg.Input, "i", "", "Entry file")
	flag.StringVar(&cfg.Input, "input", "", "Entry file")
	flag.StringVar(&cfg.Output, "o", "", "Output file")
	flag.StringVar(&cfg.Output, "out", "", "Output file")
	flag.BoolVar(&cfg.Minify, "m", false, "Minify")
	flag.BoolVar(&cfg.Minify, "minify", false, "Minify")
	flag.BoolVar(&cfg.Report, "r", false, "Build report")
	flag.BoolVar(&cfg.Report, "report", false, "Build report")
	flag.StringVar(&cfg.SourceMap, "s", "", "Source map")
	flag.StringVar(&cfg.SourceMap, "source", "", "Source map")
	flag.BoolVar(&cfg.Watch, "w", false, "Watch mode")
	flag.BoolVar(&cfg.Watch, "watch", false, "Watch mode")
	flag.StringVar(&cfg.LogLevel, "log-level", "", "Log level")
	// Force flags for non-interactive mode
	flag.BoolVar(&cfg.Force, "f", false, "Force overwrite (skip confirmation)")
	flag.BoolVar(&cfg.Force, "force", false, "Force overwrite (skip confirmation)")
	flag.BoolVar(&cfg.Yes, "y", false, "Yes to overwrite (auto-confirm)")
	flag.BoolVar(&cfg.Yes, "yes", false, "Yes to overwrite (auto-confirm)")
	flag.BoolVar(&cfg.NoConfirm, "n", false, "No confirmations (skip all prompts)")
	flag.BoolVar(&cfg.NoConfirm, "no-confirm", false, "No confirmations (skip all prompts)")
	flag.BoolVar(&showVersion, "v", false, "Version")
	flag.BoolVar(&showVersion, "version", false, "Version")
	flag.BoolVar(&help, "h", false, "Help")
	flag.BoolVar(&help, "help", false, "Help")
	flag.Parse()

	return cfg, configPath, showVersion, help
}

// FindConfigFile looks for default config file in current directory
func FindConfigFile() (string, error) {
	if _, err := os.Stat("jspackr.config.json"); err == nil {
		return "jspackr.config.json", nil
	}
	return "", nil
}

// ValidateVersionFlag checks if version flag is used correctly
func ValidateVersionFlag(showVersion bool) error {
	if showVersion {
		if len(os.Args) > 2 {
			return fmt.Errorf("--version cannot be combined with other flags")
		}
	}
	return nil
}

// ShowVersion prints the version information
func ShowVersion() {
	color.New(color.FgCyan).Println("jspackr 0.3")
}

// ShowUsage displays colored and formatted help message
func ShowUsage() {
	// Colors
	titleColor := color.New(color.FgCyan, color.Bold)
	flagColor := color.New(color.FgGreen)
	descColor := color.New(color.FgWhite)
	dimColor := color.New(color.FgHiBlack)
	sectionColor := color.New(color.FgYellow, color.Bold)

	fmt.Println()
	titleColor.Println("╔══════════════════════════════════════════════════════════════╗")
	titleColor.Println("║                        JSPACKR 0.3                           ║")
	titleColor.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Usage
	sectionColor.Println("📖 USAGE")
	descColor.Print("  ")
	flagColor.Println("jspackr [options]")
	fmt.Println()

	// Description
	sectionColor.Println("📝 DESCRIPTION")
	descColor.Println("  A fast JavaScript bundler that packs your code efficiently")
	fmt.Println()

	// Options section
	sectionColor.Println("⚡ OPTIONS")
	fmt.Println()

	// Core options
	dimColor.Println("  ┌─────────────────────────────────────────────────────────────┐")
	dimColor.Println("  │                      CORE OPTIONS                           │")
	dimColor.Println("  └─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	flagColor.Println("  -i, --input <file>     ")
	descColor.Println("    Entry JavaScript file to bundle")
	fmt.Println()

	flagColor.Println("  -o, --out <file>       ")
	descColor.Println("    Output bundle file")
	fmt.Println()

	flagColor.Println("  -c, --config <file>    ")
	descColor.Println("    Path to configuration file")
	fmt.Println()

	// Build options
	dimColor.Println("  ┌─────────────────────────────────────────────────────────────┐")
	dimColor.Println("  │                      BUILD OPTIONS                          │")
	dimColor.Println("  └─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	flagColor.Println("  -m, --minify           ")
	descColor.Println("    Minify the output bundle")
	fmt.Println()

	flagColor.Println("  -s, --source <type>    ")
	descColor.Println("    Generate source map (inline, external, none)")
	fmt.Println()

	flagColor.Println("  -r, --report           ")
	descColor.Println("    Generate build report")
	fmt.Println()

	flagColor.Println("  -w, --watch            ")
	descColor.Println("    Watch for file changes")
	fmt.Println()

	flagColor.Println("  --log-level <level>    ")
	descColor.Println("    Set log level (debug, info, warn, error)")
	fmt.Println()

	// Non-interactive options
	dimColor.Println("  ┌─────────────────────────────────────────────────────────────┐")
	dimColor.Println("  │                 NON-INTERACTIVE OPTIONS                     │")
	dimColor.Println("  └─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	flagColor.Println("  -f, --force            ")
	descColor.Println("    Skip file overwrite confirmation")
	fmt.Println()

	flagColor.Println("  -y, --yes              ")
	descColor.Println("    Auto-confirm overwrite prompts")
	fmt.Println()

	flagColor.Println("  -n, --no-confirm       ")
	descColor.Println("    Skip all confirmation prompts")
	fmt.Println()

	// Info options
	dimColor.Println("  ┌─────────────────────────────────────────────────────────────┐")
	dimColor.Println("  │                        INFO OPTIONS                         │")
	dimColor.Println("  └─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	flagColor.Println("  -v, --version          ")
	descColor.Println("    Show version information")
	fmt.Println()

	flagColor.Println("  -h, --help             ")
	descColor.Println("    Show this help message")
	fmt.Println()

	// Examples section
	sectionColor.Println("💡 EXAMPLES")
	fmt.Println()
	dimColor.Println("  # Basic usage")
	descColor.Println("    jspackr -i src/index.js -o dist/bundle.js")
	fmt.Println()
	dimColor.Println("  # With minification and watch mode")
	descColor.Println("    jspackr -i src/index.js -o dist/bundle.js --minify --watch")
	fmt.Println()
	dimColor.Println("  # Non-interactive mode")
	descColor.Println("    jspackr -i src/index.js -o dist/bundle.js --force")
	fmt.Println()
	dimColor.Println("  # Using config file")
	descColor.Println("    jspackr -c jspackr.config.json")
	fmt.Println()

	// Footer
	titleColor.Println("╔══════════════════════════════════════════════════════════════╗")
	titleColor.Println("║     For more info, visit: github.com/kalokaradia/jspackr     ║")
	titleColor.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

// ValidateOutputFile validates the output file path
func ValidateOutputFile(path string) error {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return fmt.Errorf("output path is a directory: %s", path)
	}
	return nil
}

// GetAbsolutePath returns the absolute path for a given path
func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}
