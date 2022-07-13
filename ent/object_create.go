// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/object"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ObjectCreate is the builder for creating a Object entity.
type ObjectCreate struct {
	config
	mutation *ObjectMutation
	hooks    []Hook
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

// check runs all checks and user-defined validators on the builder.
func (oc *ObjectCreate) check() error {
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
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *ObjectCreate) createSpec() (*Object, *sqlgraph.CreateSpec) {
	var (
		_node = &Object{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: object.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: object.FieldID,
			},
		}
	)
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
