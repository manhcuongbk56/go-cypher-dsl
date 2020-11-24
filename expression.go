package cypher_go_dsl

type IsExpression interface {
	Visitable
	IsExpression() bool
}

type Expression struct {
}

type Condition struct {
	expression Expression
}

//type ICondition interface {
//	IExpression
//}

func (lhs Expression) IsEqualTo(rhs Expression) Comparison {
	return Comparison{left: lhs, operator: EQUALITY, right: rhs}
}

func (lhs Expression) Accept(visitor Visitor) {
	panic("implement me")
}

func (lhs Condition) And(rhs Condition) {

}

type Operator string

const (
	EQUALITY = "equality"
)

type Comparison struct {
	Expression
	left     Expression
	operator Operator
	right    Expression
}

func (c Comparison) IsExpression() bool {
	return true
}

func (c Comparison) Accept(visitor Visitor) {
	panic("implement me")
}

func NameOrExpression(expression IsExpression) IsExpression  {
	named, isNamed := expression.(Named)
	if isNamed && named.getSymbolicName() !=  nil {
		return named.getSymbolicName()
	}
	return expression
}



