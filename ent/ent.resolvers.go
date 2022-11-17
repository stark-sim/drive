package ent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"drive/common"
	"fmt"
	"strconv"
)

// ID is the resolver for the id field.
func (r *directoryResolver) ID(ctx context.Context, obj *Directory) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ParentID is the resolver for the parentID field.
func (r *directoryResolver) ParentID(ctx context.Context, obj *Directory) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *objectResolver) ID(ctx context.Context, obj *Object) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// UserID is the resolver for the userID field.
func (r *objectResolver) UserID(ctx context.Context, obj *Object) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (Noder, error) {
	return r.client.Noder(ctx, common.StringToInt64(id))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]Noder, error) {
	intIDs := make([]int64, len(ids))
	for i, v := range ids {
		intIDs[i] = common.StringToInt64(v)
	}
	return r.client.Noders(ctx, intIDs)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder) (*UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, WithUserOrder(orderBy))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *User) (string, error) {
	return strconv.FormatInt(obj.ID, 10), nil
}

// ObjectIDs is the resolver for the objectIDs field.
func (r *createUserInputResolver) ObjectIDs(ctx context.Context, obj *CreateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// AddObjectIDs is the resolver for the addObjectIDs field.
func (r *updateUserInputResolver) AddObjectIDs(ctx context.Context, obj *UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// RemoveObjectIDs is the resolver for the removeObjectIDs field.
func (r *updateUserInputResolver) RemoveObjectIDs(ctx context.Context, obj *UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// Directory returns DirectoryResolver implementation.
func (r *Resolver) Directory() DirectoryResolver { return &directoryResolver{r} }

// Object returns ObjectResolver implementation.
func (r *Resolver) Object() ObjectResolver { return &objectResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// CreateUserInput returns CreateUserInputResolver implementation.
func (r *Resolver) CreateUserInput() CreateUserInputResolver { return &createUserInputResolver{r} }

// UpdateUserInput returns UpdateUserInputResolver implementation.
func (r *Resolver) UpdateUserInput() UpdateUserInputResolver { return &updateUserInputResolver{r} }

type directoryResolver struct{ *Resolver }
type objectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
type updateUserInputResolver struct{ *Resolver }
