package generator

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Snapshot reader inputs so scanning siblings cannot consume the input later
// used to preserve a middleware/hook file, or a subsequent generation run.
func (o *Options) snapshotPreviousFiles() (map[string][]byte, error) {
	files := map[string][]byte{}
	for name, reader := range o.PreviousFiles {
		if reader == nil || !strings.HasSuffix(name, ".go") {
			continue
		}
		source, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("read previous file %s: %w", name, err)
		}
		o.PreviousFiles[name] = bytes.NewReader(source)
		files[path.Clean(filepath.ToSlash(name))] = source
	}
	return files, nil
}

func (o *Options) handlerSources(target string, previous map[string][]byte) (map[string][]byte, error) {
	files := map[string][]byte{}
	dir := path.Dir(target)
	if o.Output != "" {
		entries, err := os.ReadDir(filepath.Join(o.Output, filepath.FromSlash(dir)))
		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") || strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
				continue
			}
			relative := path.Join(dir, name)
			// Explicit in-memory sources override the output directory.
			if _, ok := previous[relative]; ok {
				continue
			}
			source, err := os.ReadFile(filepath.Join(o.Output, filepath.FromSlash(relative)))
			if err != nil {
				return nil, err
			}
			files[relative] = source
		}
	}
	for name, source := range previous {
		if path.Dir(name) == dir {
			files[name] = source
		}
	}
	return files, nil
}
