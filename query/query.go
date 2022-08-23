package query

import "context"

type Querier[T any] interface {
	Query(ctx context.Context) (*T, error)
}

type QuerierFactory[T any] interface {
	Create(ctx context.Context) Querier[T]
}
