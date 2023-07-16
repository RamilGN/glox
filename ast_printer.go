package main

import (
	"errors"
	"fmt"
	"strings"
)

var errAstPrinter = errors.New("[AstPrinter]")

type AstPrinter []struct{}

func (a *AstPrinter) Print(e Expr) (string, error) {
	res, err := e.Accept(a)
	if err != nil {
		return "", fmt.Errorf("%w. %w", errAstPrinter, err)
	}

	s, ok := res.(string)
	if !ok {
		return "", fmt.Errorf("%w. %v type assertion", errAstPrinter, res)
	}

	return s, nil
}

func (a *AstPrinter) VisitBinary(e *Binary) (any, error) {
	return a.parenthesize(string(e.Operator.lexeme), e.Left, e.Right)
}

func (a *AstPrinter) VisitGrouping(e *Grouping) (any, error) {
	return a.parenthesize("group", e.Expression)
}

func (a *AstPrinter) VisitLiteral(e *Literal) (any, error) {
	return fmt.Sprintf("%v", e.Object), nil
}

func (a *AstPrinter) VisitUnary(e *Unary) (any, error) {
	return a.parenthesize(string(e.Operator.lexeme), e.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) (string, error) {
	sb := strings.Builder{}

	sb.WriteString("(")
	sb.WriteString(name)

	for _, e := range exprs {
		sb.WriteString(" ")

		res, err := e.Accept(a)
		if err != nil {
			return "", fmt.Errorf("%w. %w", errAstPrinter, err)
		}

		s, ok := res.(string)
		if !ok {
			return "", fmt.Errorf("%w. %v type assertion", errAstPrinter, res)
		}

		sb.WriteString(s)
	}

	sb.WriteString(")")

	return sb.String(), nil
}
