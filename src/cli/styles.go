package cli

import "github.com/fatih/color"

// Icons provides consistent icon constants for CLI output
type Icons struct {
	Success    string // ✓
	Error      string // ✗
	Warn       string // ⚠
	Info       string // ℹ
	Debug      string // 🔍
	Watch      string // 👀
	Rebuild    string // ↻
	Build      string // ⚙️
	Dir        string // 📁
	Tip        string // 💡
	ArrowRight string // →
	Space      string //  
}

// Common icons used throughout the CLI
var IconsDefault = Icons{
	Success:    "✓",
	Error:      "✗",
	Warn:       "⚠",
	Info:       "ℹ",
	Debug:      "🔍",
	Watch:      "👀",
	Rebuild:    "↻",
	Build:      "⚙️",
	Dir:        "📁",
	Tip:        "💡",
	ArrowRight: "→",
	Space:      " ",
}

// Styles defines color schemes for different UI elements
type Styles struct {
	Title      *color.Color
	Subtitle   *color.Color
	Section    *color.Color
	Key        *color.Color
	Value      *color.Color
	Path       *color.Color
	Highlight  *color.Color
	Dim        *color.Color
	Stats      *color.Color
	Badge      *color.Color
	Warn       *color.Color
	Error      *color.Color
}

// DefaultStyles provides the standard color scheme
var DefaultStyles = Styles{
	Title:     color.New(color.FgCyan, color.Bold),
	Subtitle:  color.New(color.FgWhite),
	Section:   color.New(color.FgMagenta, color.Bold),
	Key:       color.New(color.FgYellow),
	Value:     color.New(color.FgGreen),
	Path:      color.New(color.FgCyan),
	Highlight: color.New(color.FgWhite, color.Bold),
	Dim:       color.New(color.FgWhite),
	Stats:     color.New(color.FgBlue),
	Badge:     color.New(color.FgBlack, color.BgWhite),
	Warn:      color.New(color.FgYellow),
	Error:     color.New(color.FgRed),
}

// PrintSection prints a section header
func PrintSection(title string) {
	width := 60
	padding := (width - len(title) - 2) / 2
	divider := ""
	for i := 0; i < width; i++ {
		divider += "─"
	}

	DefaultStyles.Section.Println("\n" + divider)
	DefaultStyles.Section.Printf(" %*s %s %*s \n", padding, "", title, padding, "")
	DefaultStyles.Section.Println(divider)
}

// PrintKeyValue prints a key-value pair with consistent formatting
func PrintKeyValue(key, value string, indent int) {
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}
	DefaultStyles.Key.Printf("%s%s%s:", indentStr, IconsDefault.Space, key)
	DefaultStyles.Value.Printf(" %s\n", value)
}

// PrintDivider prints a visual divider
func PrintDivider() {
	divider := ""
	for i := 0; i < 60; i++ {
		divider += "─"
	}
	DefaultStyles.Dim.Println(divider)
}

// PrintStat prints a statistic with label
func PrintStat(label string, value string) {
	DefaultStyles.Key.Printf("  %s %s ", IconsDefault.Space, label)
	DefaultStyles.Stats.Printf("%s\n", value)
}

// PrintPath prints a file/directory path with styling
func PrintPath(path string) {
	DefaultStyles.Path.Println(path)
}

// PrintHighlight prints highlighted text
func PrintHighlight(text string) {
	DefaultStyles.Highlight.Println(text)
}

// PrintBadge prints text in a badge style
func PrintBadge(text string) {
	DefaultStyles.Badge.Printf(" %s ", text)
}

// PrintBox prints text inside a box
func PrintBox(lines []string) {
	if len(lines) == 0 {
		return
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	border := "┌"
	for i := 0; i < maxLen+2; i++ {
		border += "─"
	}
	border += "┐"

	DefaultStyles.Section.Println(border)
	for _, line := range lines {
		DefaultStyles.Section.Printf("│ %-*s │\n", maxLen, line)
	}

	border = "└"
	for i := 0; i < maxLen+2; i++ {
		border += "─"
	}
	border += "┘"

	DefaultStyles.Section.Println(border)
}

