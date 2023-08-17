package postgresql

import (
	"github.com/michaelzx/sqlc-stand-alone-queries-struct/internal/sql/catalog"
)

func pgTemp() *catalog.Schema {
	return &catalog.Schema{Name: "pg_temp"}
}
