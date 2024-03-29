// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/directory"
	"drive/ent/object"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DirectoryCreate is the builder for creating a Directory entity.
type DirectoryCreate struct {
	config
	mutation *DirectoryMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (dc *DirectoryCreate) SetCreatedBy(i int64) *DirectoryCreate {
	dc.mutation.SetCreatedBy(i)
	return dc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableCreatedBy(i *int64) *DirectoryCreate {
	if i != nil {
		dc.SetCreatedBy(*i)
	}
	return dc
}

// SetUpdatedBy sets the "updated_by" field.
func (dc *DirectoryCreate) SetUpdatedBy(i int64) *DirectoryCreate {
	dc.mutation.SetUpdatedBy(i)
	return dc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableUpdatedBy(i *int64) *DirectoryCreate {
	if i != nil {
		dc.SetUpdatedBy(*i)
	}
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DirectoryCreate) SetCreatedAt(t time.Time) *DirectoryCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableCreatedAt(t *time.Time) *DirectoryCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DirectoryCreate) SetUpdatedAt(t time.Time) *DirectoryCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableUpdatedAt(t *time.Time) *DirectoryCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DirectoryCreate) SetDeletedAt(t time.Time) *DirectoryCreate {
	dc.mutation.SetDeletedAt(t)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableDeletedAt(t *time.Time) *DirectoryCreate {
	if t != nil {
		dc.SetDeletedAt(*t)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DirectoryCreate) SetName(s string) *DirectoryCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetIsPublic sets the "is_public" field.
func (dc *DirectoryCreate) SetIsPublic(b bool) *DirectoryCreate {
	dc.mutation.SetIsPublic(b)
	return dc
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableIsPublic(b *bool) *DirectoryCreate {
	if b != nil {
		dc.SetIsPublic(*b)
	}
	return dc
}

// SetParentID sets the "parent_id" field.
func (dc *DirectoryCreate) SetParentID(i int64) *DirectoryCreate {
	dc.mutation.SetParentID(i)
	return dc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableParentID(i *int64) *DirectoryCreate {
	if i != nil {
		dc.SetParentID(*i)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DirectoryCreate) SetID(i int64) *DirectoryCreate {
	dc.mutation.SetID(i)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DirectoryCreate) SetNillableID(i *int64) *DirectoryCreate {
	if i != nil {
		dc.SetID(*i)
	}
	return dc
}

// AddObjectIDs adds the "objects" edge to the Object entity by IDs.
func (dc *DirectoryCreate) AddObjectIDs(ids ...int64) *DirectoryCreate {
	dc.mutation.AddObjectIDs(ids...)
	return dc
}

// AddObjects adds the "objects" edges to the Object entity.
func (dc *DirectoryCreate) AddObjects(o ...*Object) *DirectoryCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return dc.AddObjectIDs(ids...)
}

// SetParent sets the "parent" edge to the Directory entity.
func (dc *DirectoryCreate) SetParent(d *Directory) *DirectoryCreate {
	return dc.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Directory entity by IDs.
func (dc *DirectoryCreate) AddChildIDs(ids ...int64) *DirectoryCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Directory entity.
func (dc *DirectoryCreate) AddChildren(d ...*Directory) *DirectoryCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// Mutation returns the DirectoryMutation object of the builder.
func (dc *DirectoryCreate) Mutation() *DirectoryMutation {
	return dc.mutation
}

// Save creates the Directory in the database.
func (dc *DirectoryCreate) Save(ctx context.Context) (*Directory, error) {
	var (
		err  error
		node *Directory
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DirectoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Directory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DirectoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DirectoryCreate) SaveX(ctx context.Context) *Directory {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DirectoryCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DirectoryCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DirectoryCreate) defaults() {
	if _, ok := dc.mutation.CreatedBy(); !ok {
		v := directory.DefaultCreatedBy
		dc.mutation.SetCreatedBy(v)
	}
	if _, ok := dc.mutation.UpdatedBy(); !ok {
		v := directory.DefaultUpdatedBy
		dc.mutation.SetUpdatedBy(v)
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := directory.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := directory.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		v := directory.DefaultDeletedAt
		dc.mutation.SetDeletedAt(v)
	}
	if _, ok := dc.mutation.IsPublic(); !ok {
		v := directory.DefaultIsPublic
		dc.mutation.SetIsPublic(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := directory.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DirectoryCreate) check() error {
	if _, ok := dc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Directory.created_by"`)}
	}
	if _, ok := dc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Directory.updated_by"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Directory.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Directory.updated_at"`)}
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Directory.deleted_at"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Directory.name"`)}
	}
	if _, ok := dc.mutation.IsPublic(); !ok {
		return &ValidationError{Name: "is_public", err: errors.New(`ent: missing required field "Directory.is_public"`)}
	}
	return nil
}

func (dc *DirectoryCreate) sqlSave(ctx context.Context) (*Directory, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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

func (dc *DirectoryCreate) createSpec() (*Directory, *sqlgraph.CreateSpec) {
	var (
		_node = &Directory{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: directory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: directory.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedBy(); ok {
		_spec.SetField(directory.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := dc.mutation.UpdatedBy(); ok {
		_spec.SetField(directory.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(directory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(directory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.SetField(directory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(directory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.IsPublic(); ok {
		_spec.SetField(directory.FieldIsPublic, field.TypeBool, value)
		_node.IsPublic = value
	}
	if nodes := dc.mutation.ObjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   directory.ObjectsTable,
			Columns: []string{directory.ObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: object.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   directory.ParentTable,
			Columns: []string{directory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: directory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   directory.ChildrenTable,
			Columns: []string{directory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: directory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DirectoryCreateBulk is the builder for creating many Directory entities in bulk.
type DirectoryCreateBulk struct {
	config
	builders []*DirectoryCreate
}

// Save creates the Directory entities in the database.
func (dcb *DirectoryCreateBulk) Save(ctx context.Context) ([]*Directory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Directory, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DirectoryMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DirectoryCreateBulk) SaveX(ctx context.Context) []*Directory {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DirectoryCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DirectoryCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
