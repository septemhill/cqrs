package execute

import "context"

type Executor[T any] interface {
	Execute(ctx context.Context) (*T, error)
}

type ExecutorFactory[T any] interface {
	Create(ctx context.Context) Executor[T]
}
