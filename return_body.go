package cypher_go_dsl

type ReturnBody struct {
	returnItems ExpressionList
	order *Order
	skip *Skip
	limit *Limit
}

func (r ReturnBody) Accept(visitor Visitor) {
	r.returnItems.Accept(visitor)
	VisitIfNotNull(r.order, visitor)
	VisitIfNotNull(r.skip, visitor)
	VisitIfNotNull(r.limit, visitor)
}

