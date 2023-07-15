package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
)

type Expr struct {
	TypeName string
	Fields   []ExprField
}

type ExprField struct {
	Name     string
	TypeName string
}

var errAst = errors.New("[ast]")

func main() {
	tmpl, err := template.New("ast.go.tmpl").ParseFiles("tools/ast/ast.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	exprTypes := []Expr{
		{
			TypeName: "Binary",
			Fields: []ExprField{
				{Name: "Left", TypeName: "Expr"},
				{Name: "Operator", TypeName: "Token"},
				{Name: "Right", TypeName: "Expr"},
			},
		},
		{
			TypeName: "Grouping",
			Fields: []ExprField{
				{Name: "Expression", TypeName: "Expr"},
			},
		},
		{
			TypeName: "Literal",
			Fields: []ExprField{
				{Name: "Object", TypeName: "any"},
			},
		},
		{
			TypeName: "Unary",
			Fields: []ExprField{
				{Name: "Operator", TypeName: "Token"},
				{Name: "Right", TypeName: "Expr"},
			},
		},
	}

	if err := func() error {
		fileName := "ast.gen.go"

		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("%w. %w", errAst, err)
		}
		defer file.Close()

		err = tmpl.Execute(file, exprTypes)
		if err != nil {
			return fmt.Errorf("%w. %w", errAst, err)
		}

		command := exec.Command("go", "fmt", fileName)

		err = command.Run()
		if err != nil {
			return fmt.Errorf("%w. %w", errAst, err)
		}

		return nil
	}(); err != nil {
		log.Fatal(err)
	}
}
