package cypher_go_dsl

import (
	"fmt"
	"strconv"
)

type NumberLiteral struct {
	content int
	key     string
}

func (n NumberLiteral) getKey() string {
	return n.key
}

func (n NumberLiteral) IsExpression() bool {
	return true
}

func (n NumberLiteral) GetContent() interface{} {
	return n.content
}

func (n NumberLiteral) AsString() string {
	return strconv.Itoa(n.content)
}

func (n NumberLiteral) accept(visitor *CypherRenderer) {
	n.key = fmt.Sprint(&n)
	(*visitor).enter(n)
	(*visitor).Leave(n)
}

func (n NumberLiteral) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(n.AsString())
}

func (n NumberLiteral) leave(renderer *CypherRenderer) {
}