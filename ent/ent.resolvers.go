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
func (r *queryResolver) Users(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder, where *UserWhereInput) (*UserConnection, error) {
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

// ID is the resolver for the id field.
func (r *directoryWhereInputResolver) ID(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *directoryWhereInputResolver) IDNeq(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *directoryWhereInputResolver) IDIn(ctx context.Context, obj *DirectoryWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *directoryWhereInputResolver) IDNotIn(ctx context.Context, obj *DirectoryWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *directoryWhereInputResolver) IDGt(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *directoryWhereInputResolver) IDGte(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *directoryWhereInputResolver) IDLt(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *directoryWhereInputResolver) IDLte(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ParentID is the resolver for the parentID field.
func (r *directoryWhereInputResolver) ParentID(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// ParentIDNeq is the resolver for the parentIDNEQ field.
func (r *directoryWhereInputResolver) ParentIDNeq(ctx context.Context, obj *DirectoryWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentIDNeq - parentIDNEQ"))
}

// ParentIDIn is the resolver for the parentIDIn field.
func (r *directoryWhereInputResolver) ParentIDIn(ctx context.Context, obj *DirectoryWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDIn - parentIDIn"))
}

// ParentIDNotIn is the resolver for the parentIDNotIn field.
func (r *directoryWhereInputResolver) ParentIDNotIn(ctx context.Context, obj *DirectoryWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDNotIn - parentIDNotIn"))
}

// ID is the resolver for the id field.
func (r *objectWhereInputResolver) ID(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *objectWhereInputResolver) IDNeq(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *objectWhereInputResolver) IDIn(ctx context.Context, obj *ObjectWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *objectWhereInputResolver) IDNotIn(ctx context.Context, obj *ObjectWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *objectWhereInputResolver) IDGt(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *objectWhereInputResolver) IDGte(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *objectWhereInputResolver) IDLt(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *objectWhereInputResolver) IDLte(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// UserID is the resolver for the userID field.
func (r *objectWhereInputResolver) UserID(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserID - userID"))
}

// UserIDNeq is the resolver for the userIDNEQ field.
func (r *objectWhereInputResolver) UserIDNeq(ctx context.Context, obj *ObjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: UserIDNeq - userIDNEQ"))
}

// UserIDIn is the resolver for the userIDIn field.
func (r *objectWhereInputResolver) UserIDIn(ctx context.Context, obj *ObjectWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: UserIDIn - userIDIn"))
}

// UserIDNotIn is the resolver for the userIDNotIn field.
func (r *objectWhereInputResolver) UserIDNotIn(ctx context.Context, obj *ObjectWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: UserIDNotIn - userIDNotIn"))
}

// AddObjectIDs is the resolver for the addObjectIDs field.
func (r *updateUserInputResolver) AddObjectIDs(ctx context.Context, obj *UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// RemoveObjectIDs is the resolver for the removeObjectIDs field.
func (r *updateUserInputResolver) RemoveObjectIDs(ctx context.Context, obj *UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

// ID is the resolver for the id field.
func (r *userWhereInputResolver) ID(ctx context.Context, obj *UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *userWhereInputResolver) IDNeq(ctx context.Context, obj *UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *userWhereInputResolver) IDIn(ctx context.Context, obj *UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userWhereInputResolver) IDNotIn(ctx context.Context, obj *UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *userWhereInputResolver) IDGt(ctx context.Context, obj *UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *userWhereInputResolver) IDGte(ctx context.Context, obj *UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *userWhereInputResolver) IDLt(ctx context.Context, obj *UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *userWhereInputResolver) IDLte(ctx context.Context, obj *UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
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

// DirectoryWhereInput returns DirectoryWhereInputResolver implementation.
func (r *Resolver) DirectoryWhereInput() DirectoryWhereInputResolver {
	return &directoryWhereInputResolver{r}
}

// ObjectWhereInput returns ObjectWhereInputResolver implementation.
func (r *Resolver) ObjectWhereInput() ObjectWhereInputResolver { return &objectWhereInputResolver{r} }

// UpdateUserInput returns UpdateUserInputResolver implementation.
func (r *Resolver) UpdateUserInput() UpdateUserInputResolver { return &updateUserInputResolver{r} }

// UserWhereInput returns UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() UserWhereInputResolver { return &userWhereInputResolver{r} }

type directoryResolver struct{ *Resolver }
type objectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
type directoryWhereInputResolver struct{ *Resolver }
type objectWhereInputResolver struct{ *Resolver }
type updateUserInputResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
