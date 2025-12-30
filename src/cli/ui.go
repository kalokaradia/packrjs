package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kalokaradia/jspackr/src/config"
)

// PromptOptions defines options for styled prompts
type PromptOptions struct {
	Icon       string
	IconColor  *color.Color
	TextColor  *color.Color
	ArrowColor *color.Color
	Tip        string
	TipColor   *color.Color
}

// DefaultPromptOptions returns default prompt styling
func DefaultPromptOptions() PromptOptions {
	return PromptOptions{
		Icon:       IconsDefault.Space,
		IconColor:  color.New(color.FgCyan),
		TextColor:  color.New(color.FgWhite),
		ArrowColor: color.New(color.FgYellow),
		Tip:        "",
		TipColor:   color.New(color.FgBlue),
	}
}

// Confirm prompts user for yes/no confirmation with styling
func Confirm(message string, defaultYes bool) bool {
	opts := DefaultPromptOptions()
	return ConfirmWithOptions(message, opts, defaultYes)
}

// ConfirmWithOptions prompts user with custom styling
func ConfirmWithOptions(message string, opts PromptOptions, defaultYes bool) bool {
	// Create the prompt line
	prompt := ""
	
	if opts.Icon != "" {
		prompt += opts.Icon + " "
	}
	prompt += message + " "
	
	// Default indicator
	defaultStr := "[y/N]"
	if defaultYes {
		defaultStr = "[Y/n]"
	}
	
	opts.ArrowColor.Print(defaultStr + ": ")
	
	// Get input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	// Handle empty input with default
	if input == "" {
		return defaultYes
	}
	
	// Normalize input
	lower := strings.ToLower(input)
	return lower == "y" || lower == "yes"
}

// ConfirmCreateDir prompts user to create a directory
func ConfirmCreateDir(path string) bool {
	message := fmt.Sprintf("Directory '%s' does not exist. Create it?", path)
	return Confirm(message, false)
}

// Spinner represents an animated spinner
type Spinner struct {
	message  string
	interval time.Duration
	stopChan chan struct{}
	done     chan struct{}
	idx      int
}

// Spinner frames for animation
var spinnerFrames = []string{
	"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
}

// NewSpinner creates a new spinner with message
func NewSpinner(message string) *Spinner {
	return &Spinner{
		message:  message,
		interval: 80 * time.Millisecond,
		stopChan: make(chan struct{}),
		done:     make(chan struct{}),
		idx:      0,
	}
}

// Start begins the spinner animation
func (s *Spinner) Start() {
	go func() {
		for {
			select {
			case <-s.stopChan:
				// Clear the line and show final state
				fmt.Print("\r")
				for i := 0; i < 80; i++ {
					fmt.Print(" ")
				}
				fmt.Print("\r")
				close(s.done)
				return
			case <-time.After(s.interval):
				frame := spinnerFrames[s.idx]
				color.New(color.FgCyan).Printf("\r%s %s", frame, s.message)
				s.idx = (s.idx + 1) % len(spinnerFrames)
			}
		}
	}()
}

// Stop stops the spinner and shows completion
func (s *Spinner) Stop(success bool) {
	close(s.stopChan)
	<-s.done
	
	if success {
		color.New(color.FgGreen).Printf("\r✓ %s\n", s.message)
	} else {
		color.New(color.FgRed).Printf("\r✗ %s\n", s.message)
	}
}

// StopWithMessage stops spinner with custom message
func (s *Spinner) StopWithMessage(message string, success bool) {
	close(s.stopChan)
	<-s.done
	
	if success {
		color.New(color.FgGreen).Printf("\r✓ %s\n", message)
	} else {
		color.New(color.FgRed).Printf("\r✗ %s\n", message)
	}
}

// ProgressBar represents a simple progress bar
type ProgressBar struct {
	total     int
	width     int
	prefix    string
	fillChar  string
	emptyChar string
	fillColor *color.Color
}

// NewProgressBar creates a new progress bar
func NewProgressBar(total int, prefix string) *ProgressBar {
	return &ProgressBar{
		total:     total,
		width:     40,
		prefix:    prefix,
		fillChar:  "█",
		emptyChar: "░",
		fillColor: color.New(color.FgCyan),
	}
}

// Render draws the progress bar at current progress
func (p *ProgressBar) Render(current int) {
	if current > p.total {
		current = p.total
	}
	
	percent := float64(current) / float64(p.total)
	filled := int(float64(p.width) * percent)
	empty := p.width - filled
	
	bar := strings.Repeat(p.fillChar, filled) + strings.Repeat(p.emptyChar, empty)
	
	p.fillColor.Printf("\r%s [%s] %3d%%", p.prefix, bar, int(percent*100))
}

// Finish completes the progress bar
func (p *ProgressBar) Finish(message string) {
	p.Render(p.total)
	fmt.Println()
	if message != "" {
		color.New(color.FgGreen).Printf("✓ %s\n", message)
	}
}

// PrintStatus prints a status message with consistent formatting
func PrintStatus(icon, message string, iconColor *color.Color, messageColor *color.Color) {
	if iconColor == nil {
		iconColor = color.New(color.FgCyan)
	}
	if messageColor == nil {
		messageColor = color.New(color.FgWhite)
	}
	
	iconColor.Printf("%s ", icon)
	messageColor.Println(message)
}

// PrintWelcome prints a title for the application
func PrintTitle() {
	fmt.Println()
	color.New(color.FgCyan).Println("jspackr 0.3")
	fmt.Println()
}

// PrintBuildSummary prints a summary of the build configuration
func PrintBuildSummary(cfg *config.Config) {
	DefaultStyles.Section.Println("Build Configuration")

	PrintKeyValue("Input", cfg.Input, 0)
	PrintKeyValue("Output", cfg.Output, 0)
	PrintKeyValue("Minify", fmt.Sprintf("%t", cfg.Minify), 0)
	PrintKeyValue("Source Map", cfg.SourceMap, 0)
	PrintKeyValue("Report", fmt.Sprintf("%t", cfg.Report), 0)
	PrintKeyValue("Watch Mode", fmt.Sprintf("%t", cfg.Watch), 0)
	PrintKeyValue("Log Level", cfg.LogLevel, 0)

	PrintDivider()
}

// PrintBuildResult prints the result of a build operation
func PrintBuildResult(success bool, message string) {
	if success {
		if IconsDefault.Success != "" {
			DefaultStyles.Value.Printf("  %s %s\n", IconsDefault.Success, message)
		} else {
			DefaultStyles.Value.Printf("  %s\n", message)
		}
	} else {
		if IconsDefault.Error != "" {
			DefaultStyles.Key.Printf("  %s %s\n", IconsDefault.Error, message)
		} else {
			DefaultStyles.Key.Printf("  %s\n", message)
		}
	}
}

// PrintHelpInfo prints help/tip information
func PrintHelpInfo(tip string) {
	if tip == "" {
		return
	}
	if IconsDefault.Tip != "" {
		DefaultStyles.Dim.Printf("  %s %s\n", IconsDefault.Tip, tip)
	} else {
		DefaultStyles.Dim.Printf("  %s\n", tip)
	}
}

// NewLine prints a new line
func NewLine() {
	fmt.Println()
}

