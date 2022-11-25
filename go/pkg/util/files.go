package util

import (
	"bufio"
	"bytes"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

func ClearGeneratedFiles(path string, suffixes ...string) error {
	return clearGeneratedFiles(path, false, suffixes...)
}

func DeepClearGeneratedFiles(path string, suffixes ...string) error {
	return clearGeneratedFiles(path, true, suffixes...)
}

func clearGeneratedFiles(path string, recursive bool, suffixes ...string) error {
	return clearFiles(path, false, func(file string) bool {
		matchedSuffix := false
		for _, suffix := range suffixes {
			if strings.HasSuffix(file, suffix) {
				matchedSuffix = true
			}
		}
		if len(suffixes) == 0 || matchedSuffix {
			return IsAllGeneratedFile(file)
		}
		return false
	})
}

func ClearFiles(path string, suffixes ...string) error {
	return clearFiles(path, false, func(file string) bool {
		return hasSuffix(file, suffixes...)
	})
}

func DeepClearFiles(path string, suffixes ...string) error {
	return clearFiles(path, true, func(file string) bool {
		return hasSuffix(file, suffixes...)
	})
}

func hasSuffix(file string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(file, suffix) {
			return true
		}
	}
	return false
}

func IsAllGeneratedFile(path string) bool {
	if core.IsExist(path) {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return false
		}

		firstLine, _ := bufio.NewReader(bytes.NewReader(content)).ReadString('\n')
		firstLine = strings.TrimSpace(firstLine)

		return strings.HasPrefix(firstLine, "// Code generated") &&
			(strings.HasSuffix(firstLine, "DO NOT EDIT.") ||
				strings.HasSuffix(firstLine, "DO NOT EDIT!"))
	}
	return false
}

func clearFiles(path string, recursive bool, filter func(file string) bool) error {
	var files []string
	var err error
	if recursive {
		err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if f != nil && !f.IsDir() && filter(path) {
				files = append(files, path)
			}
			return nil
		})
	} else {
		var fileInfos []fs.FileInfo
		fileInfos, err = ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, f := range fileInfos {
			name := filepath.Join(path, f.Name())
			if !f.IsDir() && filter(name) {
				files = append(files, name)
			}
		}
	}

	if err != nil {
		logs.Errorw("filepath.Walk() failed", "error", err.Error())
		return err
	}

	for _, file := range files {
		os.Remove(file)
	}

	return nil
}

func GetAbsolutePath(pwd string, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(pwd, path)
}
