package ent

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
)

func (r *directoryResolver) ID(ctx context.Context, obj *Directory) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *directoryResolver) ParentID(ctx context.Context, obj *Directory) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *objectResolver) ID(ctx context.Context, obj *Object) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *objectResolver) UserID(ctx context.Context, obj *Object) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *createUserInputResolver) ObjectIDs(ctx context.Context, obj *CreateUserInput, data []string) error {
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

type directoryResolver struct{ *Resolver }
type objectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
