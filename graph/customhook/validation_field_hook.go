package customhook

import (
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/vektah/gqlparser/v2/ast"
)

func ValidationFieldHook(td *ast.Definition, fd *ast.FieldDefinition, f *modelgen.Field) (*modelgen.Field, error) {
	c := fd.Directives.ForName("constraint")
	if c != nil {
		maxConstraint := c.Arguments.ForName("max")
		if maxConstraint != nil {
			f.Tag += " validate:\"max=" + maxConstraint.Value.String() + "\""
		}
	}
	return f, nil
}
