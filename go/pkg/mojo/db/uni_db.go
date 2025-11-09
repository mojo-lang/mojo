package db

type UniDb struct {
	Tables map[string]*UniTable
}

func (d *UniDb) Table(name string) *UniTable {
	return nil
}

func (d *UniDb) RenameTable() {
}

func (d *UniDb) ListTables() {
}

func (d *UniDb) DropTable() {
}
