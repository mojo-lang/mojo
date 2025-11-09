package data

type ArrayResponse BoxedArray

func (r *ArrayResponse) GetPackageName() string {
	return r.PackageName
}

func (r *ArrayResponse) GetFullName() string {
	return r.FullName
}
