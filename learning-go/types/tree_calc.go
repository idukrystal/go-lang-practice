package main

import (
	"errors"
	"fmt"
)

var endOfExpr error = errors.New("End of Expression")

type ExprNode struct{
	exprTerm
	Next *ExprNode
	Prev *ExprNode
}

func (node *ExprNode) Connect(next  *ExprNode) {
	node.Next = next
	next.Prev = node
}

type exprTerm interface {
	asString() string
}

type NumTerm struct {
	Number float64
}

func (numTerm NumTerm) asString() string {
	return "number"
}

type OprTerm struct {
	Opr
}

func (oprTerm OprTerm) asString() string {
	return "operation"
}

type Opr interface {
	Operate (float64, float64) (float64, error)
	fmt.Stringer
}

type Plus struct {}

func (p Plus) Operate(x float64, y float64) (float64, error) {
	return x + y, nil
}

func (p Plus) String() string {
	return "+"
}

type Minus struct {}

func (m Minus) Operate(x, y float64) (float64, error) {
	return x - y, nil
}

func(m  Minus) String() string {
	return "-"
}

func Calculate(node *ExprNode) (float64, error) {
	if node == nil {
		return 0, endOfExpr
	}

	switch term := node.exprTerm.(type) {
	case NumTerm:
		ans, err := Calculate(node.Next)
		if err != nil {
			if err == endOfExpr {
				return term.Number, nil
			}
			return 0, err
		}
		return ans, nil
	case OprTerm:
		ans, err := Calculate(node.Next)
		if err != nil {
			if err == endOfExpr {
				return 0, fmt.Errorf("No Right operand for: %s", term)
			}
			return 0, err
		}
		if node.Prev == nil {
			return 0, fmt.Errorf("No left operand for: %s", term)
		}
		prev, ok := node.Prev.exprTerm.(NumTerm)
		if !ok {
			return 0, fmt.Errorf("Invalid left operand to: %s", term)
		}
		return term.Operate(prev.Number, ans)
	default:
		return 0, errors.New("Unknown expr type in node")
	}
}

func ttt(ss fmt.Stringer) {}

func main() {
	a := ExprNode{NumTerm{3}, nil, nil}
	b := ExprNode{NumTerm{5}, nil, nil}
	c := ExprNode{NumTerm{4}, nil, nil}

	d := ExprNode{OprTerm{Plus{}}, nil, nil}
	e := ExprNode{OprTerm{Minus{}}, nil, nil}

	a.Connect(&d)
	d.Connect(&b)
	b.Connect(&e)
	e.Connect(&c)
	fmt.Println(Calculate(&d))
}

