package data

type PaginationResult BoxedArray

func (p *PaginationResult) GetPackageName() string {
    return p.PackageName
}

func (p *PaginationResult) GetFullName() string {
    return p.FullName
}
