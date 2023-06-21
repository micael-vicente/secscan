package sensors

import (
	"ACME/secscan/pkg/util"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	results    []string
	exclusions []string
)

// RunDirectory runs sensor on a directory, beeping when candidate files are found
func RunDirectory(path string, config Config) []string {

	info, err := os.Stat(path)
	exclusions = config.ExcludedFileExtensions

	if !(err == nil && info.IsDir()) {
		log.Fatalf("path '%s' does not exist or is not a directory", path)
	}

	_ = filepath.WalkDir(path, walk)

	return results
}

func walk(s string, d fs.DirEntry, err error) error {
	if err == nil && !d.IsDir() {
		ext := path.Ext(s)
		if !util.Contains(exclusions, ext) {
			results = append(results, s)
		}
		return nil
	} else {
		return err
	}
}
