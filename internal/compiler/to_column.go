package compiler

import (
	"strings"

	"github.com/michaelzx/sqlc-stand-alone-queries-struct/internal/sql/ast"
	"github.com/michaelzx/sqlc-stand-alone-queries-struct/internal/sql/astutils"
)

func arrayDims(n *ast.TypeName) int {
	if n == nil || n.ArrayBounds == nil {
		return 0
	}
	return len(n.ArrayBounds.Items)
}

func toColumn(n *ast.TypeName) *Column {
	if n == nil {
		panic("can't build column for nil type name")
	}
	typ, err := ParseTypeName(n)
	if err != nil {
		panic("toColumn: " + err.Error())
	}
	arrayDims := arrayDims(n)
	return &Column{
		Type:      typ,
		DataType:  strings.TrimPrefix(astutils.Join(n.Names, "."), "."),
		NotNull:   true, // XXX: How do we know if this should be null?
		IsArray:   arrayDims > 0,
		ArrayDims: arrayDims,
	}
}
