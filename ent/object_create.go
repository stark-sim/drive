// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/object"
	"drive/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ObjectCreate is the builder for creating a Object entity.
type ObjectCreate struct {
	config
	mutation *ObjectMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (oc *ObjectCreate) SetCreatedBy(i int64) *ObjectCreate {
	oc.mutation.SetCreatedBy(i)
	return oc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableCreatedBy(i *int64) *ObjectCreate {
	if i != nil {
		oc.SetCreatedBy(*i)
	}
	return oc
}

// SetUpdatedBy sets the "updated_by" field.
func (oc *ObjectCreate) SetUpdatedBy(i int64) *ObjectCreate {
	oc.mutation.SetUpdatedBy(i)
	return oc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableUpdatedBy(i *int64) *ObjectCreate {
	if i != nil {
		oc.SetUpdatedBy(*i)
	}
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *ObjectCreate) SetCreatedAt(t time.Time) *ObjectCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableCreatedAt(t *time.Time) *ObjectCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *ObjectCreate) SetUpdatedAt(t time.Time) *ObjectCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableUpdatedAt(t *time.Time) *ObjectCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetDeletedAt sets the "deleted_at" field.
func (oc *ObjectCreate) SetDeletedAt(t time.Time) *ObjectCreate {
	oc.mutation.SetDeletedAt(t)
	return oc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableDeletedAt(t *time.Time) *ObjectCreate {
	if t != nil {
		oc.SetDeletedAt(*t)
	}
	return oc
}

// SetURL sets the "url" field.
func (oc *ObjectCreate) SetURL(s string) *ObjectCreate {
	oc.mutation.SetURL(s)
	return oc
}

// SetID sets the "id" field.
func (oc *ObjectCreate) SetID(i int64) *ObjectCreate {
	oc.mutation.SetID(i)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *ObjectCreate) SetNillableID(i *int64) *ObjectCreate {
	if i != nil {
		oc.SetID(*i)
	}
	return oc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (oc *ObjectCreate) SetUserID(id int64) *ObjectCreate {
	oc.mutation.SetUserID(id)
	return oc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (oc *ObjectCreate) SetNillableUserID(id *int64) *ObjectCreate {
	if id != nil {
		oc = oc.SetUserID(*id)
	}
	return oc
}

// SetUser sets the "user" edge to the User entity.
func (oc *ObjectCreate) SetUser(u *User) *ObjectCreate {
	return oc.SetUserID(u.ID)
}

// Mutation returns the ObjectMutation object of the builder.
func (oc *ObjectCreate) Mutation() *ObjectMutation {
	return oc.mutation
}

// Save creates the Object in the database.
func (oc *ObjectCreate) Save(ctx context.Context) (*Object, error) {
	var (
		err  error
		node *Object
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Object)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ObjectMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *ObjectCreate) SaveX(ctx context.Context) *Object {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *ObjectCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *ObjectCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *ObjectCreate) defaults() {
	if _, ok := oc.mutation.CreatedBy(); !ok {
		v := object.DefaultCreatedBy
		oc.mutation.SetCreatedBy(v)
	}
	if _, ok := oc.mutation.UpdatedBy(); !ok {
		v := object.DefaultUpdatedBy
		oc.mutation.SetUpdatedBy(v)
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := object.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := object.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.DeletedAt(); !ok {
		v := object.DefaultDeletedAt
		oc.mutation.SetDeletedAt(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		v := object.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *ObjectCreate) check() error {
	if _, ok := oc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Object.created_by"`)}
	}
	if _, ok := oc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Object.updated_by"`)}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Object.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Object.updated_at"`)}
	}
	if _, ok := oc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Object.deleted_at"`)}
	}
	if _, ok := oc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Object.url"`)}
	}
	return nil
}

func (oc *ObjectCreate) sqlSave(ctx context.Context) (*Object, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (oc *ObjectCreate) createSpec() (*Object, *sqlgraph.CreateSpec) {
	var (
		_node = &Object{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: object.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: object.FieldID,
			},
		}
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := oc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := oc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: object.FieldURL,
		})
		_node.URL = value
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   object.UserTable,
			Columns: []string{object.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_objects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ObjectCreateBulk is the builder for creating many Object entities in bulk.
type ObjectCreateBulk struct {
	config
	builders []*ObjectCreate
}

// Save creates the Object entities in the database.
func (ocb *ObjectCreateBulk) Save(ctx context.Context) ([]*Object, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Object, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ObjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *ObjectCreateBulk) SaveX(ctx context.Context) []*Object {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *ObjectCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *ObjectCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
