package querybus

import "fmt"

type QueryNotRegisteredError struct {
	Query Query
}

func (e QueryNotRegisteredError) Error() string {
	return fmt.Sprintf("no handler registered for query: %s", e.Query.QueryName())
}
