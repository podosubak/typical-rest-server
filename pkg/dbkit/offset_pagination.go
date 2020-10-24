package dbkit

import (
	sq "github.com/Masterminds/squirrel"
)

// OffsetPagination param
type OffsetPagination struct {
	Offset uint64
	Limit  uint64
}

var _ SelectOption = (*OffsetPagination)(nil)

// CompileSelect to compile select query for pagination
func (p *OffsetPagination) CompileSelect(base sq.SelectBuilder) (sq.SelectBuilder, error) {
	base = base.Offset(p.Offset)
	if p.Limit > 0 {
		base = base.Limit(p.Limit)
	}
	return base, nil
}
