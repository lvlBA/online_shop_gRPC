package pagination

import (
	"testing"

	"github.com/c2fo/testify/assert"
	"github.com/doug-martin/goqu/v9"
)

func TestPagination_Calc(t *testing.T) {
	type fields struct {
		page  uint
		limit uint
	}
	tests := []struct {
		name       string
		fields     fields
		wantOffset uint
		wantLimit  uint
	}{
		{
			name: "test1",
			fields: fields{
				page:  1,
				limit: 50,
			},
			wantOffset: 0,
			wantLimit:  50,
		},
		{
			name: "test2",
			fields: fields{
				page:  2,
				limit: 50,
			},
			wantOffset: 50,
			wantLimit:  50,
		},
		{
			name: "test3",
			fields: fields{
				page:  0,
				limit: 50,
			},
			wantOffset: 0,
			wantLimit:  50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pagination{
				page:  tt.fields.page,
				limit: tt.fields.limit,
			}
			gotOffset, gotLimit := p.Calc()
			if gotOffset != tt.wantOffset {
				t.Errorf("Calc() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("Calc() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
		})
	}
}

func TestPagination_DataSet(t *testing.T) {
	t.Run("page_1", func(t *testing.T) {
		ds := goqu.From("test_table").Select("*")
		pg := &Pagination{
			page:  1,
			limit: 50,
		}
		want := "SELECT * FROM \"test_table\" LIMIT 50"
		ds = pg.DataSet(ds)
		got, _, err := ds.ToSQL()
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
	t.Run("page_2", func(t *testing.T) {
		ds := goqu.From("test_table").Select("*")
		pg := &Pagination{
			page:  2,
			limit: 50,
		}
		want := "SELECT * FROM \"test_table\" LIMIT 50 OFFSET 50"
		ds = pg.DataSet(ds)
		got, _, err := ds.ToSQL()
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
	t.Run("page_0", func(t *testing.T) {
		ds := goqu.From("test_table").Select("*")
		pg := &Pagination{
			page:  0,
			limit: 50,
		}
		want := "SELECT * FROM \"test_table\" LIMIT 50"
		ds = pg.DataSet(ds)
		got, _, err := ds.ToSQL()
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}
