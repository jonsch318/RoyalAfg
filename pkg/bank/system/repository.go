package system

import "context"

type Repository[T IAggregate] interface {
	Load(context.Context, id string) (T, error)
	Save(context.Context, T) error
}
