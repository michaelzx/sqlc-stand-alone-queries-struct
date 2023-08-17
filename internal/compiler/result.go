package compiler

import (
	"github.com/michaelzx/sqlc-stand-alone-queries-struct/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
