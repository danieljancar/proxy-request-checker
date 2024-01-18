package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Walks through all the files and directories located under the "./reports" directory, make sure to run from project root.
// If a file is found with the extension ".json" and the filename contains "_report", it will be deleted.
// Any encountered error is returned by the function, and if not nil, it panics.
func main() {
	err := filepath.Walk("./reports", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" && strings.Contains(filepath.Base(path), "_report") {
			err := os.Remove(path)
			if err != nil {
				return nil
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
