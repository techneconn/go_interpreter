
package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_BOJ = "BOOLEAN"
	NULL_OBJ = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}
func (i *Boolean) Type() ObjectType { return BOOLEAN_BOJ }
func (i *Boolean) Inspect() string { return fmt.Sprintf("%t", i.Value) }

type Null struct {
}
func (i *Null) Type() ObjectType { return NULL_OBJ }
func (i *Null) Inspect() string { return "null" }

