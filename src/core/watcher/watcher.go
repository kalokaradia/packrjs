package watcher

import (
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/kalokaradia/jspackr/src/cli"
	"github.com/kalokaradia/jspackr/src/core/builder"
)

var (
	fileHashes = make(map[string][32]byte)
)

// WatchFiles watch file changes and trigger rebuilds
func WatchFiles(entry string, opts builder.Options, logger *cli.Logger) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Use default logger if nil
	if logger == nil {
		logger = cli.New("info")
	}

	// consistent absolute path
	entryPath, err := filepath.Abs(entry)
	if err != nil {
		return err
	}

	// baseline hash
	if hash, err := HashFile(entryPath); err == nil {
		fileHashes[entryPath] = hash
	}

	logger.PrintWatch(entryPath)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&(fsnotify.Write|fsnotify.Create) == 0 {
					continue
				}

				eventPath, err := filepath.Abs(event.Name)
				if err != nil || eventPath != entryPath {
					continue
				}

				// debounce: wait 300ms before rebuild
				StartDebounce(300*time.Millisecond, func() {
					newHash, err := HashFile(entryPath)
					if err != nil {
						return
					}
					if newHash == fileHashes[entryPath] {
						return
					}

					fileHashes[entryPath] = newHash
					logger.PrintRebuild()

					if err := builder.Run(opts); err != nil {
						logger.Error("Build failed: %v", err)
					} else {
						logger.PrintSuccess()
					}
				})

			case err, ok := <-watcher.Errors:
				if ok {
					logger.Error("Watcher error: %v", err)
				}
			}
		}
	}()

	// watch entry file
	if err := watcher.Add(entryPath); err != nil {
		return err
	}

	// block forever
	select {}
}

