package templates

import (
	"embed"
	"io"
	"io/fs"
	path2 "path"
	"strings"
)

//go:embed service-go/*
var services embed.FS

//go:embed NAME-client/*
var clients embed.FS

//go:embed sidecar/*
var sidecars embed.FS

type FileGetter = func(path string) ([]byte, error)

func Service(path string) ([]byte, error) {
	return fileContent(services, path, "service-go/")
}

func Client(path string) ([]byte, error) {
	return fileContent(clients, path, "")
}

func Sidecar(path string) ([]byte, error) {
	return fileContent(clients, path, "")
}

func ServiceNames() []string {
	return fileNames(services, "service-go/")
}

func ClientNames() []string {
	return fileNames(clients, "")
}

func SidecarNames() []string {
	return fileNames(sidecars, "")
}

func fileContent(fs embed.FS, path string, prefix string) ([]byte, error) {
	f, err := fs.Open(path2.Join(prefix, path))
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func fileNames(efs embed.FS, prefixTrim string) []string {
	var names []string
	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if len(prefixTrim) > 0 {
			path = strings.TrimPrefix(path, prefixTrim)
		}
		names = append(names, path)
		return nil
	})

	if err != nil {
		return []string{}
	}

	return names
}
