package http

const VersionTypeName = "Version"
const VersionTypeFullName = "mojo.http.Version"

func DefaultVersion() *Version {
    return &Version{
        Major: 1,
        Minor: 1,
    }
}

func (x *Version) IsEmpty() bool {
    return x == nil || x.Major == 0
}
