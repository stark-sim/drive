package ent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	// 事务如此简单
	client := FromContext(ctx)
	return client.User.Create().SetName(*input.Name).SetPassword(input.Password).SetPhone(input.Phone).Save(ctx)
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// IsFucked is the resolver for the isFucked field.
func (r *userWhereInputResolver) IsFucked(ctx context.Context, obj *UserWhereInput, data *bool) error {
	panic(fmt.Errorf("not implemented: IsFucked - isFucked"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
