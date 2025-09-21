package querybus

import (
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/querybus"
)

type QueryHandlersInformation struct {
	queryHandlersMap map[string]querybus.QueryHandler
}

func NewQueryHandlersInformation(queryHandlers []querybus.QueryHandler) *QueryHandlersInformation {
	qhi := &QueryHandlersInformation{
		queryHandlersMap: make(map[string]querybus.QueryHandler),
	}
	qhi.formatHandlers(queryHandlers)
	return qhi
}

func (qhi *QueryHandlersInformation) formatHandlers(queryHandlers []querybus.QueryHandler) {
	for _, handler := range queryHandlers {
		qhi.queryHandlersMap[handler.SubscribedTo().QueryName()] = handler
	}
}

func (qhi *QueryHandlersInformation) Search(query querybus.Query) (querybus.QueryHandler, error) {
	handler, exists := qhi.queryHandlersMap[query.QueryName()]
	if handler == nil || !exists {
		return nil, querybus.QueryNotRegisteredError{Query: query}
	}
	return handler, nil
}
