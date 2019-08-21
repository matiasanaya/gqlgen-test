package embedding

import (
	"context"
)

//go:generate go run github.com/99designs/gqlgen

// Resolvers model
type Resolvers struct {
	*queryResolver
}

// Query resolver
func (r *Resolvers) Query() QueryResolver {
	return r.queryResolver
}

// NewResolvers constructor
func NewResolvers() *Resolvers {
	return &Resolvers{&queryResolver{}}
}

type queryResolver struct{}

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return []*Todo{&Todo{}}, nil
}

// Todo model
type Todo struct {
	*ider
	*namer
}

// // ID resolver
// func (t *Todo) ID(ctx context.Context) (string, error) {
// 	return t.ider.ID(ctx)
// }
//
// // Name resolver
// func (t *Todo) Name(ctx context.Context) (string, error) {
// 	return t.namer.Name(ctx)
// }

type ider struct{}

func (*ider) ID(ctx context.Context) (string, error) {
	return "1", nil
}

type namer struct{}

func (*namer) Name(ctx context.Context) (string, error) {
	return "name", nil
}
