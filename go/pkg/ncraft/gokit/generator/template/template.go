package template

import (
	"embed"
	"io"
	"io/fs"
)

//go:embed PKGNAME/*
var services embed.FS

//go:embed NAME-client/*
var clients embed.FS

//go:embed PKGNAME-sidecar/*
var sidecars embed.FS

type FileGetter = func(path string) ([]byte, error)

func Service(path string) ([]byte, error) {
	return fileContent(services, path)
}

func Client(path string) ([]byte, error) {
	return fileContent(clients, path)
}

func Sidecar(path string) ([]byte, error) {
	return fileContent(clients, path)
}

func ServiceNames() []string {
	return fileNames(services)
}

func ClientNames() []string {
	return fileNames(clients)
}

func SidecarNames() []string {
	return fileNames(sidecars)
}

func fileContent(fs embed.FS, path string) ([]byte, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func fileNames(efs embed.FS) []string {
	var names []string
	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		names = append(names, path)
		return nil
	})

	if err != nil {
		return []string{}
	}

	return names
}
