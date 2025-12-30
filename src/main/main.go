package main

import (
	"os"

	"github.com/kalokaradia/jspackr/src/cli"
	"github.com/kalokaradia/jspackr/src/config"
	"github.com/kalokaradia/jspackr/src/core/builder"
	"github.com/kalokaradia/jspackr/src/core/watcher"
	"github.com/kalokaradia/jspackr/src/utils"
)

func main() {
	const version = "0.3.0"
	flagCfg, configPath, showVersion, help := utils.ParseFlags()

	// Handle version flag
	if err := utils.ValidateVersionFlag(showVersion); err != nil {
		cli.DefaultStyles.Key.Printf("\n✗ %v\n", err)
		os.Exit(1)
	}

	if showVersion {
		utils.ShowVersion()
		return
	}

	if help {
		utils.ShowUsage()
		return
	}

	// Load configuration
	finalCfg := config.Default()

	if configPath == "" {
		defaultConfig, _ := utils.FindConfigFile()
		if defaultConfig != "" {
			configPath = defaultConfig
		}
	}

	if configPath != "" {
		fileCfg, err := config.Load(configPath)
		if err != nil {
			cli.DefaultStyles.Key.Printf("\n✗ Failed to load config: %v\n", err)
			os.Exit(2)
		}
		finalCfg = fileCfg
	}

	config.Merge(finalCfg, flagCfg)

	if err := config.Validate(finalCfg); err != nil {
		cli.DefaultStyles.Key.Printf("\n✗ %v\n", err)
		os.Exit(2)
	}

	logger := cli.New(finalCfg.LogLevel)

	// Print welcome banner
	cli.PrintTitle()

	// Print full build configuration summary
	cli.PrintBuildSummary(finalCfg)

	// Validate input path exists
	if err := config.ValidateInputPath(finalCfg.Input); err != nil {
		logger.FatalErr(err, "Invalid input path")
	}

	// Validate output path and handle directory creation
	outDir := utils.GetOutputParent(finalCfg.Output)
	if outDir != "." {
		if _, err := config.ValidateOutputPath(finalCfg.Output); err != nil {
			// Output directory doesn't exist, ask user to create it
			// Skip confirmation if noConfirm flag is set
			if !finalCfg.NoConfirm {
				if !cli.ConfirmCreateDir(outDir) {
					cli.DefaultStyles.Warn.Println("\n⚠ Build cancelled")
					return
				}
			}
			if err := utils.CreateDir(outDir); err != nil {
				logger.FatalErr(err, "Failed to create directory")
			}
			logger.PrintDirCreated(outDir)
		}
	}

	if err := utils.ValidateOutputFile(finalCfg.Output); err != nil {
		logger.FatalErr(err, "Invalid output path")
	}

	// Check if we should overwrite existing file
	// Skip confirmation if force, yes, or noConfirm flags are set
	if !utils.ConfirmOverwrite(finalCfg.Output, finalCfg.Force, finalCfg.Yes, finalCfg.NoConfirm) {
		cli.DefaultStyles.Warn.Println("\n⚠ Build cancelled")
		return
	}

	if outDir != "." {
		if notEmpty, _ := utils.DirNotEmpty(outDir); notEmpty {
			logger.WarnWithTip(
				"Output directory not empty: "+outDir,
				"Existing files may be overwritten",
			)
		}
	}

	opts := builder.Options{
		Input:     finalCfg.Input,
		Output:    finalCfg.Output,
		Minify:    finalCfg.Minify,
		Report:    finalCfg.Report,
		SourceMap: finalCfg.SourceMap,
	}

	if finalCfg.Watch {
		logger.Info("Watch mode enabled")
		watcher.WatchFiles(finalCfg.Input, opts, logger)
		return
	}

	// Start build
	logger.PrintBuildStart()
	spinner := cli.NewSpinner("Bundling...")
	spinner.Start()

	if err := builder.Run(opts); err != nil {
		spinner.Stop(false)
		logger.FatalErr(err, "Build failed")
	}

	spinner.Stop(true)

	logger.PrintSuccess()
}

