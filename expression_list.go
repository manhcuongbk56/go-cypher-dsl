package cypher_go_dsl

import "fmt"

type ExpressionList struct {
	expressions []Expression
	key         string
}

func (e ExpressionList) getKey() string {
	return e.key
}

func (e ExpressionList) PrepareVisit(child Visitable) Visitable {
	expression, isExpression := child.(Expression)
	if !isExpression {
		panic("Can not prepare un expression type in expression list")
	}
	return NameOrExpression(&expression)
}

func (e ExpressionList) accept(visitor *CypherRenderer) {
	e.key = fmt.Sprint(&e)
	(*visitor).enter(e)
	for _, expression := range e.expressions {
		e.PrepareVisit(expression).accept(visitor)
	}
	(*visitor).Leave(e)
}

func (e ExpressionList) enter(renderer *CypherRenderer) {
}

func (e ExpressionList) leave(renderer *CypherRenderer) {
}

func NewExpressionList(expression ...Expression) ExpressionList {
	expressions := make([]Expression, len(expression))
	expressions = append(expressions, expression...)
	return ExpressionList{expressions: expressions}
}