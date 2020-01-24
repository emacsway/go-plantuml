package astParser

import (
	"github.com/bykof/go-plantuml/domain"
	"github.com/magiconair/properties/assert"
	"go/ast"
	"testing"
)

func Test_starExprToField(t *testing.T) {
	assert.Equal(
		t,
		starExprToField("test", &ast.StarExpr{
			X: &ast.Ident{
				Name: "string",
			},
		}),
		domain.Field{
			Name:     "test",
			Type:     "*string",
			Nullable: true,
		},
	)
}

func Test_identToField(t *testing.T) {
	assert.Equal(
		t,
		identToField("test", &ast.Ident{Name: "string"}),
		domain.Field{
			Name: "test",
			Type: "string",
		},
	)
}

func Test_selectorExprToField(t *testing.T) {
	assert.Equal(
		t,
		selectorExprToField(
			"test",
			&ast.SelectorExpr{
				Sel: &ast.Ident{Name: "testSelect"},
				X:   &ast.Ident{Name: "testPackage"},
			},
		),
		domain.Field{
			Name:    "test",
			Type:    "testSelect",
			Package: "testPackage",
		},
	)
}
