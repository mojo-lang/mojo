package core

import (
	"runtime"
	"strings"
)

const PlatformTypeName = "Platform"
const PlatformTypeFullName = "mojo.core.Platform"

func NewPlatform(os OS, arch Architecture, variant string) *Platform {
	return &Platform{
		Os:           os,
		Architecture: arch,
		Variant:      variant,
	}
}

func (x *Platform) IsEmpty() bool {
	return x == nil || x.Os == 0
}

func (x *Platform) SetOsName(name string) *Platform {
	if x != nil {
		x.OsName = name
	}
	return x
}

func (x *Platform) SetOsVersion(version string) *Platform {
	if x != nil {
		x.OsVersion = version
	}
	return x
}

func (x *Platform) IsLinux() bool {
	return x != nil && x.Os == OS_OS_LINUX
}

func (x *Platform) IsArmArch() bool {
	return x != nil && (x.Architecture == Architecture_ARCHITECTURE_ARM || x.Architecture == Architecture_ARCHITECTURE_ARM64)
}

func (x *Platform) Is64Arch() bool {
	return x != nil && (x.Architecture == Architecture_ARCHITECTURE_AMD64 || x.Architecture == Architecture_ARCHITECTURE_ARM64)
}

// normalizeOS https://github.com/containerd/containerd/blob/main/platforms/platforms.go
func normalizeOS(os string) string {
	if os == "" {
		return runtime.GOOS
	}
	os = strings.ToLower(os)

	switch os {
	case "macos":
		os = "darwin"
	}
	return os
}

// normalizeArch normalizes the architecture.
func normalizeArch(arch, variant string) (string, string) {
	arch, variant = strings.ToLower(arch), strings.ToLower(variant)
	switch arch {
	case "i386":
		arch = "386"
		variant = ""
	case "x86_64", "x86-64", "amd64":
		arch = "amd64"
		if variant == "v1" {
			variant = ""
		}
	case "aarch64", "arm64":
		arch = "arm64"
		switch variant {
		case "8", "v8":
			variant = ""
		}
	case "armhf":
		arch = "arm"
		variant = "v7"
	case "armel":
		arch = "arm"
		variant = "v6"
	case "arm":
		switch variant {
		case "", "7":
			variant = "v7"
		case "5", "6", "8":
			variant = "v" + variant
		}
	}

	return arch, variant
}
