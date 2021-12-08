package test

import "embed"

//go:embed case-alias/*
var AliasCaseFiles embed.FS

//go:embed case-enum/*
var EnumCaseFiles embed.FS

//go:embed case-inherits/*
var InheritsCaseFiles embed.FS

//go:embed case-package/*
var PackageCaseFiles embed.FS
