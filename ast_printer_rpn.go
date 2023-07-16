package main

import (
	"errors"
	"fmt"
	"strings"
)

var errAstPrinterRPN = errors.New("[AstPrinterRPN]")

type AstPrinterRPN struct{}

func (a *AstPrinterRPN) Print(e Expr) (string, error) {
	res, err := e.Accept(a)
	if err != nil {
		return "", fmt.Errorf("%w. %w", errAstPrinterRPN, err)
	}

	s, ok := res.(string)
	if !ok {
		return "", fmt.Errorf("%w. %v type assertion", errAstPrinterRPN, res)
	}

	return s, nil
}

func (a *AstPrinterRPN) VisitBinary(e *Binary) (any, error) {
	return a.rpn(string(e.Operator.lexeme), e.Left, e.Right)
}

func (a *AstPrinterRPN) VisitGrouping(e *Grouping) (any, error) {
	return a.rpn("group", e.Expression)
}

func (a *AstPrinterRPN) VisitLiteral(e *Literal) (any, error) {
	return fmt.Sprintf("%v", e.Object), nil
}

func (a *AstPrinterRPN) VisitUnary(e *Unary) (any, error) {
	return a.rpn(string(e.Operator.lexeme), e.Right)
}

func (a *AstPrinterRPN) rpn(name string, exprs ...Expr) (string, error) {
	sb := strings.Builder{}

	for _, e := range exprs {
		res, err := e.Accept(a)
		if err != nil {
			return "", fmt.Errorf("%w. %w", errAstPrinterRPN, err)
		}

		s, ok := res.(string)
		if !ok {
			return "", fmt.Errorf("%w. %v type assertion", errAstPrinter, res)
		}

		sb.WriteString(s)
		sb.WriteString(" ")
	}

	sb.WriteString(name)

	return sb.String(), nil
}
