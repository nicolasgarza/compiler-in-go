import "fmt"

package object

type ObjectType string

type Object interface {
	Tpye() ObjectType
	Inspect() string
}

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value book
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%b", b.Value)}

type Null struct{}
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }