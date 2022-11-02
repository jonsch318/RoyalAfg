package system

import "context"

type Repository[T IAggregate] interface {
	Load(context.Context, string) (T, error)
	Save(context.Context, T) error
	Delete(context.Context, string) error
}
