package db

type UniTable struct {
	Name string
}

//	// Examples: 'age >= 18', 'age >= 18 and verified == true'
//	// Comparison operators: '==', '!=', '<', '>', '<=', '>='
//	// Logical operator: 'and'
//	// Dot access is supported, eg: 'user.age == 11'
//	// Accessing list elements is not supported yet.
// ["or", ["and, ["==", b, c], ["!=", e, f]], "a.b.c"]
// ((b == c) && (e != f)) || a.b.c
// type Filter struct {
//	Value string
//	Expr *lang.Expression
// }
//
// func (d *UniTable) Create(tx context.Context, record *UniRecord) (string, error) {
// 	return "", nil
// }
//
// func (d *UniTable) Update(tx context.Context, record *UniRecord) error {
// 	return nil
// }
//
// func (d *UniTable) Get(ctx context.Context, id string) (*UniRecord, error) {
// 	return nil, nil
// }
//
// func (d *UniTable) List(ctx context.Context, filter string, pagination Pagination, orderBy core.Ordering, options core.Options) ([]*UniRecord, error) {
// 	return nil, nil
// }
//
// func (d *UniTable) Delete(tx context.Context, id string) error {
// 	return nil
// }
//
// func (d *UniTable) Truncate(tx context.Context) error {
// 	return nil
// }
//
// func (d *UniTable) Count(tx context.Context) (int32, error) {
// 	return 0, nil
// }
