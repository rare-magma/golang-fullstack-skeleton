package querybus

import "context"

type QueryHandler interface {
	Handle(ctx context.Context, query Query) (Response, error)
	SubscribedTo() Query
}
