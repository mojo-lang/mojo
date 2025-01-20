package util

import (
	"errors"
	"io"
	"os"
	"path"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

type GeneratedFile struct {
	// The file name, relative to the output directory.  The name must not
	// contain "." or ".." components and must be relative, not be absolute (so,
	// the file cannot lie outside the output directory).  "/" must be used as
	// the path separator, not "\".
	//
	// If the name is omitted, the content will be appended to the previous
	// file.  This allows the generator to break large files into small chunks,
	// and allows the generated text to be streamed back to protoc so that large
	// files need not reside completely in memory at one time.  Note that as of
	// this writing protoc does not optimize for this -- it will read the entire
	// CodeGeneratorResponse before writing files to disk.
	Name string

	// The file contents.
	Content string

	Reader io.Reader

	SkipIfExist         bool
	SkipIfUserCodeMixed bool
}

type GeneratedFiles []*GeneratedFile

func (c *GeneratedFile) WriteTo(output string, guard *PathGuard) error {
	if len(c.Name) == 0 {
		return errors.New("not valid file: has no file name")
	}
	if len(c.Content) == 0 && c.Reader == nil {
		logs.Warnw("skip an empty file", "name", c.Name)
		return nil
	}

	name := path.Join(output, c.Name)
	dir := path.Dir(name)

	if guard != nil {
		if err := guard.Check(dir); err != nil {
			return err
		}
	}

	if core.IsExist(name) {
		if c.SkipIfExist {
			return nil
		}

		if c.SkipIfUserCodeMixed && !IsAllGeneratedFile(name) {
			return nil
		}
	}

	if len(c.Content) == 0 {
		if bytes, err := io.ReadAll(c.Reader); err != nil {
			return err
		} else {
			c.Content = string(bytes)
		}
	}

	perm := os.FileMode(0o666)
	if path.Ext(name) == ".sh" {
		perm = os.ModePerm
	}

	return os.WriteFile(name, []byte(c.Content), perm)
}
