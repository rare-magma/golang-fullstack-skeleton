package querybus

import "context"

type QueryBus interface {
	Ask(ctx context.Context, query Query) (Response, error)
}
