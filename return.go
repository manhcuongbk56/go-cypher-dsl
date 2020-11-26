package cypher_go_dsl

import "fmt"

type Return struct {
	distinct *Distinct
	body     *ReturnBody
	key      string
}

func (r Return) getKey() string {
	return r.key
}

func (r Return) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
	(*visitor).Enter(r)
	VisitIfNotNull(r.distinct, visitor)
	r.body.accept(visitor)
	(*visitor).Leave(r)
}

func ReturnByMultiVariable(distinct bool, returnItems ExpressionList, order *Order, skip *Skip, limit *Limit) *Return {
	var distinctInstance *Distinct
	if distinct {
		distinctInstance = &Distinct{}
	}
	body := ReturnBody{
		returnItems: returnItems,
		order:       order,
		skip:        skip,
		limit:       limit,
	}
	return &Return{
		distinct: distinctInstance,
		body:     &body,
	}
}

func (r Return) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("RETURN  ")
}

func (r Return) leave(renderer *CypherRenderer) {
}
