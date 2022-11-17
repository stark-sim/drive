package ent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	return r.client.User.Create().SetName(*input.Name).SetPassword(input.Password).SetPhone(input.Phone).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
