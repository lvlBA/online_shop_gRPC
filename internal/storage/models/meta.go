package models

import "time"

type Meta struct {
	ID      string    `db:"id"         goqu:"skipinsert,skipupdate"`
	Created time.Time `db:"created_at" goqu:"skipupdate"`
	Changed time.Time `db:"changed_at" goqu:"skipinsert"`
}

func (m *Meta) UpdateMeta() {
	if m.Created.IsZero() {
		m.Created = time.Now()
	}

	m.Changed = time.Now()
}
