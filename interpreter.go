package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

var errRuntime = errors.New("RuntimeError")

type RunTimeError struct {
	message string
	token   Token
}

func (r RunTimeError) Error() string {
	return fmt.Sprintf("%s. %s", r.message, r.token)
}

type Interpreter struct {
	lox *Lox
}

func (i *Interpreter) Interpret(e Expr) (any, error) {
	val, err := e.Accept(i)
	if err != nil {
		loxErr := i.lox.runtimeError(err)
		if loxErr != nil {
			return nil, loxErr
		}

		return nil, fmt.Errorf("%w", err)
	}

	return val, nil
}

func (i *Interpreter) VisitLiteral(l *Literal) (any, error) {
	return l.Object, nil
}

func (i *Interpreter) VisitBinary(b *Binary) (any, error) {
	left, err := i.evaluate(b.Left)
	if err != nil {
		return nil, fmt.Errorf("%w. %w", errRuntime, err)
	}

	right, err := i.evaluate(b.Right)
	if err != nil {
		return nil, fmt.Errorf("%w. %w", errRuntime, err)
	}

	toFloat64 := func(num any) (*big.Float, error) {
		numFloat, ok := num.(float64)
		if ok {
			return big.NewFloat(numFloat), nil
		}

		numInt, ok := num.(int)
		if !ok {
			return &big.Float{}, fmt.Errorf("%w. can't cast %+v to float64", errRuntime, num)
		}

		return big.NewFloat(float64(numInt)), nil
	}

	leftAndRightToFloats64 := func() (*big.Float, *big.Float, error) {
		leftFloat, err := toFloat64(left)
		if err != nil {
			return &big.Float{}, &big.Float{}, err
		}

		rightFloat, err := toFloat64(right)
		if err != nil {
			return &big.Float{}, &big.Float{}, err
		}

		return leftFloat, rightFloat, err
	}

	toString := func() (string, string, error) {
		leftString, ok := left.(string)
		if !ok {
			return "", "", fmt.Errorf("%w. can't cast %+v to float64", errRuntime, left)
		}

		rightString, ok := right.(string)
		if !ok {
			return "", "", fmt.Errorf("%w. can't cast %+v to float64", errRuntime, right)
		}

		return leftString, rightString, nil
	}

	switch b.Operator.tType { //nolint:exhaustive
	case minus:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v - %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		res := leftNum.Sub(leftNum, rightNum)
		resFloat, _ := strconv.ParseFloat(res.String(), 64)

		return resFloat, nil
	case slash:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v / %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		res := leftNum.Quo(leftNum, rightNum)
		resFloat, _ := strconv.ParseFloat(res.String(), 64)

		return resFloat, nil
	case star:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v * %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		res := leftNum.Mul(leftNum, rightNum)
		resFloat, _ := strconv.ParseFloat(res.String(), 64)

		return resFloat, nil
	case plus:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err == nil {
			res := leftNum.Add(leftNum, rightNum)
			resFloat, _ := strconv.ParseFloat(res.String(), 64)

			return resFloat, nil
		}

		leftString, rightString, err := toString()
		if err != nil {
			return nil, fmt.Errorf("%w: addition error. `%+v + %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		return leftString + rightString, nil
	case greater:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v > %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		return leftNum.Cmp(rightNum) == 1, nil
	case greaterEqual:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v >= %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		cmpRes := leftNum.Cmp(rightNum)

		return (cmpRes == 0) || (cmpRes == 1), nil
	case less:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v < %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		return leftNum.Cmp(rightNum) == -1, nil
	case lessEqual:
		leftNum, rightNum, err := leftAndRightToFloats64()
		if err != nil {
			return nil, fmt.Errorf("%w: operands must be numbers. `%+v <= %+v` | %+v", errRuntime, left, right, b.Operator)
		}

		cmpRes := leftNum.Cmp(rightNum)

		return (cmpRes == 0) || (cmpRes == -1), nil
	case bangEqual:
		return left != right, nil
	case equal:
		return left == right, nil
	default:
		return nil, fmt.Errorf("%w: unknown operator type | %+v", errRuntime, b.Operator)
	}
}

func (i *Interpreter) VisitGrouping(g *Grouping) (any, error) {
	return i.evaluate(g.Expression)
}

func (i *Interpreter) VisitUnary(u *Unary) (any, error) {
	right, err := i.evaluate(u.Right)
	if err != nil {
		return nil, fmt.Errorf("%w. %w", errRuntime, err)
	}

	switch u.Operator.tType { //nolint:exhaustive
	case bang:
		return !i.isTruthy(right), nil

	case minus:
		rightNumFloat, ok := right.(float64)
		if ok {
			return -rightNumFloat, nil
		}

		rightNumInt, ok := right.(int)
		if ok {
			return -float64(rightNumInt), nil
		}

		return nil, fmt.Errorf("%w: operand must be a number `-%+v` | %+v", errRuntime, right, u.Operator)
	default:
		return nil, fmt.Errorf("%w: unknown operator type %+v", errRuntime, u.Operator)
	}
}

func (i *Interpreter) evaluate(e Expr) (any, error) {
	res, err := e.Accept(i)
	if err != nil {
		return nil, fmt.Errorf("%w. %w", errRuntime, err)
	}

	return res, nil
}

func (i *Interpreter) isTruthy(obj any) bool {
	if obj == nil {
		return false
	}

	b, ok := obj.(bool)
	if ok {
		return b
	}

	return true
}
