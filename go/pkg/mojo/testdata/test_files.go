package testdata

import "embed"

//go:embed mojo-alias/*
var AliasCaseFiles embed.FS

//go:embed mojo-entity/*
var EntityCaseFiles embed.FS

//go:embed mojo-enum/*
var EnumCaseFiles embed.FS

//go:embed mojo-inherits/*
var InheritsCaseFiles embed.FS

//go:embed mojo-package/*
var PackageCaseFiles embed.FS
