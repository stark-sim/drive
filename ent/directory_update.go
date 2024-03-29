// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/directory"
	"drive/ent/object"
	"drive/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DirectoryUpdate is the builder for updating Directory entities.
type DirectoryUpdate struct {
	config
	hooks    []Hook
	mutation *DirectoryMutation
}

// Where appends a list predicates to the DirectoryUpdate builder.
func (du *DirectoryUpdate) Where(ps ...predicate.Directory) *DirectoryUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedBy sets the "created_by" field.
func (du *DirectoryUpdate) SetCreatedBy(i int64) *DirectoryUpdate {
	du.mutation.ResetCreatedBy()
	du.mutation.SetCreatedBy(i)
	return du
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (du *DirectoryUpdate) SetNillableCreatedBy(i *int64) *DirectoryUpdate {
	if i != nil {
		du.SetCreatedBy(*i)
	}
	return du
}

// AddCreatedBy adds i to the "created_by" field.
func (du *DirectoryUpdate) AddCreatedBy(i int64) *DirectoryUpdate {
	du.mutation.AddCreatedBy(i)
	return du
}

// SetUpdatedBy sets the "updated_by" field.
func (du *DirectoryUpdate) SetUpdatedBy(i int64) *DirectoryUpdate {
	du.mutation.ResetUpdatedBy()
	du.mutation.SetUpdatedBy(i)
	return du
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (du *DirectoryUpdate) SetNillableUpdatedBy(i *int64) *DirectoryUpdate {
	if i != nil {
		du.SetUpdatedBy(*i)
	}
	return du
}

// AddUpdatedBy adds i to the "updated_by" field.
func (du *DirectoryUpdate) AddUpdatedBy(i int64) *DirectoryUpdate {
	du.mutation.AddUpdatedBy(i)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DirectoryUpdate) SetUpdatedAt(t time.Time) *DirectoryUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DirectoryUpdate) SetDeletedAt(t time.Time) *DirectoryUpdate {
	du.mutation.SetDeletedAt(t)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DirectoryUpdate) SetNillableDeletedAt(t *time.Time) *DirectoryUpdate {
	if t != nil {
		du.SetDeletedAt(*t)
	}
	return du
}

// SetName sets the "name" field.
func (du *DirectoryUpdate) SetName(s string) *DirectoryUpdate {
	du.mutation.SetName(s)
	return du
}

// SetIsPublic sets the "is_public" field.
func (du *DirectoryUpdate) SetIsPublic(b bool) *DirectoryUpdate {
	du.mutation.SetIsPublic(b)
	return du
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (du *DirectoryUpdate) SetNillableIsPublic(b *bool) *DirectoryUpdate {
	if b != nil {
		du.SetIsPublic(*b)
	}
	return du
}

// SetParentID sets the "parent_id" field.
func (du *DirectoryUpdate) SetParentID(i int64) *DirectoryUpdate {
	du.mutation.SetParentID(i)
	return du
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (du *DirectoryUpdate) SetNillableParentID(i *int64) *DirectoryUpdate {
	if i != nil {
		du.SetParentID(*i)
	}
	return du
}

// ClearParentID clears the value of the "parent_id" field.
func (du *DirectoryUpdate) ClearParentID() *DirectoryUpdate {
	du.mutation.ClearParentID()
	return du
}

// AddObjectIDs adds the "objects" edge to the Object entity by IDs.
func (du *DirectoryUpdate) AddObjectIDs(ids ...int64) *DirectoryUpdate {
	du.mutation.AddObjectIDs(ids...)
	return du
}

// AddObjects adds the "objects" edges to the Object entity.
func (du *DirectoryUpdate) AddObjects(o ...*Object) *DirectoryUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return du.AddObjectIDs(ids...)
}

// SetParent sets the "parent" edge to the Directory entity.
func (du *DirectoryUpdate) SetParent(d *Directory) *DirectoryUpdate {
	return du.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Directory entity by IDs.
func (du *DirectoryUpdate) AddChildIDs(ids ...int64) *DirectoryUpdate {
	du.mutation.AddChildIDs(ids...)
	return du
}

// AddChildren adds the "children" edges to the Directory entity.
func (du *DirectoryUpdate) AddChildren(d ...*Directory) *DirectoryUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddChildIDs(ids...)
}

// Mutation returns the DirectoryMutation object of the builder.
func (du *DirectoryUpdate) Mutation() *DirectoryMutation {
	return du.mutation
}

// ClearObjects clears all "objects" edges to the Object entity.
func (du *DirectoryUpdate) ClearObjects() *DirectoryUpdate {
	du.mutation.ClearObjects()
	return du
}

// RemoveObjectIDs removes the "objects" edge to Object entities by IDs.
func (du *DirectoryUpdate) RemoveObjectIDs(ids ...int64) *DirectoryUpdate {
	du.mutation.RemoveObjectIDs(ids...)
	return du
}

// RemoveObjects removes "objects" edges to Object entities.
func (du *DirectoryUpdate) RemoveObjects(o ...*Object) *DirectoryUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return du.RemoveObjectIDs(ids...)
}

// ClearParent clears the "parent" edge to the Directory entity.
func (du *DirectoryUpdate) ClearParent() *DirectoryUpdate {
	du.mutation.ClearParent()
	return du
}

// ClearChildren clears all "children" edges to the Directory entity.
func (du *DirectoryUpdate) ClearChildren() *DirectoryUpdate {
	du.mutation.ClearChildren()
	return du
}

// RemoveChildIDs removes the "children" edge to Directory entities by IDs.
func (du *DirectoryUpdate) RemoveChildIDs(ids ...int64) *DirectoryUpdate {
	du.mutation.RemoveChildIDs(ids...)
	return du
}

// RemoveChildren removes "children" edges to Directory entities.
func (du *DirectoryUpdate) RemoveChildren(d ...*Directory) *DirectoryUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DirectoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DirectoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DirectoryUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DirectoryUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DirectoryUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DirectoryUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := directory.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

func (du *DirectoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   directory.Table,
			Columns: directory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: directory.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedBy(); ok {
		_spec.SetField(directory.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := du.mutation.AddedCreatedBy(); ok {
		_spec.AddField(directory.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := du.mutation.UpdatedBy(); ok {
		_spec.SetField(directory.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := du.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(directory.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(directory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.SetField(directory.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(directory.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.IsPublic(); ok {
		_spec.SetField(directory.FieldIsPublic, field.TypeBool, value)
	}
	if du.mutation.ObjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedObjectsIDs(); len(nodes) > 0 && !du.mutation.ObjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ObjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !du.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{directory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DirectoryUpdateOne is the builder for updating a single Directory entity.
type DirectoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DirectoryMutation
}

// SetCreatedBy sets the "created_by" field.
func (duo *DirectoryUpdateOne) SetCreatedBy(i int64) *DirectoryUpdateOne {
	duo.mutation.ResetCreatedBy()
	duo.mutation.SetCreatedBy(i)
	return duo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (duo *DirectoryUpdateOne) SetNillableCreatedBy(i *int64) *DirectoryUpdateOne {
	if i != nil {
		duo.SetCreatedBy(*i)
	}
	return duo
}

// AddCreatedBy adds i to the "created_by" field.
func (duo *DirectoryUpdateOne) AddCreatedBy(i int64) *DirectoryUpdateOne {
	duo.mutation.AddCreatedBy(i)
	return duo
}

// SetUpdatedBy sets the "updated_by" field.
func (duo *DirectoryUpdateOne) SetUpdatedBy(i int64) *DirectoryUpdateOne {
	duo.mutation.ResetUpdatedBy()
	duo.mutation.SetUpdatedBy(i)
	return duo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (duo *DirectoryUpdateOne) SetNillableUpdatedBy(i *int64) *DirectoryUpdateOne {
	if i != nil {
		duo.SetUpdatedBy(*i)
	}
	return duo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (duo *DirectoryUpdateOne) AddUpdatedBy(i int64) *DirectoryUpdateOne {
	duo.mutation.AddUpdatedBy(i)
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DirectoryUpdateOne) SetUpdatedAt(t time.Time) *DirectoryUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DirectoryUpdateOne) SetDeletedAt(t time.Time) *DirectoryUpdateOne {
	duo.mutation.SetDeletedAt(t)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DirectoryUpdateOne) SetNillableDeletedAt(t *time.Time) *DirectoryUpdateOne {
	if t != nil {
		duo.SetDeletedAt(*t)
	}
	return duo
}

// SetName sets the "name" field.
func (duo *DirectoryUpdateOne) SetName(s string) *DirectoryUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetIsPublic sets the "is_public" field.
func (duo *DirectoryUpdateOne) SetIsPublic(b bool) *DirectoryUpdateOne {
	duo.mutation.SetIsPublic(b)
	return duo
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (duo *DirectoryUpdateOne) SetNillableIsPublic(b *bool) *DirectoryUpdateOne {
	if b != nil {
		duo.SetIsPublic(*b)
	}
	return duo
}

// SetParentID sets the "parent_id" field.
func (duo *DirectoryUpdateOne) SetParentID(i int64) *DirectoryUpdateOne {
	duo.mutation.SetParentID(i)
	return duo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (duo *DirectoryUpdateOne) SetNillableParentID(i *int64) *DirectoryUpdateOne {
	if i != nil {
		duo.SetParentID(*i)
	}
	return duo
}

// ClearParentID clears the value of the "parent_id" field.
func (duo *DirectoryUpdateOne) ClearParentID() *DirectoryUpdateOne {
	duo.mutation.ClearParentID()
	return duo
}

// AddObjectIDs adds the "objects" edge to the Object entity by IDs.
func (duo *DirectoryUpdateOne) AddObjectIDs(ids ...int64) *DirectoryUpdateOne {
	duo.mutation.AddObjectIDs(ids...)
	return duo
}

// AddObjects adds the "objects" edges to the Object entity.
func (duo *DirectoryUpdateOne) AddObjects(o ...*Object) *DirectoryUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return duo.AddObjectIDs(ids...)
}

// SetParent sets the "parent" edge to the Directory entity.
func (duo *DirectoryUpdateOne) SetParent(d *Directory) *DirectoryUpdateOne {
	return duo.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Directory entity by IDs.
func (duo *DirectoryUpdateOne) AddChildIDs(ids ...int64) *DirectoryUpdateOne {
	duo.mutation.AddChildIDs(ids...)
	return duo
}

// AddChildren adds the "children" edges to the Directory entity.
func (duo *DirectoryUpdateOne) AddChildren(d ...*Directory) *DirectoryUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddChildIDs(ids...)
}

// Mutation returns the DirectoryMutation object of the builder.
func (duo *DirectoryUpdateOne) Mutation() *DirectoryMutation {
	return duo.mutation
}

// ClearObjects clears all "objects" edges to the Object entity.
func (duo *DirectoryUpdateOne) ClearObjects() *DirectoryUpdateOne {
	duo.mutation.ClearObjects()
	return duo
}

// RemoveObjectIDs removes the "objects" edge to Object entities by IDs.
func (duo *DirectoryUpdateOne) RemoveObjectIDs(ids ...int64) *DirectoryUpdateOne {
	duo.mutation.RemoveObjectIDs(ids...)
	return duo
}

// RemoveObjects removes "objects" edges to Object entities.
func (duo *DirectoryUpdateOne) RemoveObjects(o ...*Object) *DirectoryUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return duo.RemoveObjectIDs(ids...)
}

// ClearParent clears the "parent" edge to the Directory entity.
func (duo *DirectoryUpdateOne) ClearParent() *DirectoryUpdateOne {
	duo.mutation.ClearParent()
	return duo
}

// ClearChildren clears all "children" edges to the Directory entity.
func (duo *DirectoryUpdateOne) ClearChildren() *DirectoryUpdateOne {
	duo.mutation.ClearChildren()
	return duo
}

// RemoveChildIDs removes the "children" edge to Directory entities by IDs.
func (duo *DirectoryUpdateOne) RemoveChildIDs(ids ...int64) *DirectoryUpdateOne {
	duo.mutation.RemoveChildIDs(ids...)
	return duo
}

// RemoveChildren removes "children" edges to Directory entities.
func (duo *DirectoryUpdateOne) RemoveChildren(d ...*Directory) *DirectoryUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DirectoryUpdateOne) Select(field string, fields ...string) *DirectoryUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Directory entity.
func (duo *DirectoryUpdateOne) Save(ctx context.Context) (*Directory, error) {
	var (
		err  error
		node *Directory
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DirectoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (duo *DirectoryUpdateOne) SaveX(ctx context.Context) *Directory {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DirectoryUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DirectoryUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DirectoryUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := directory.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

func (duo *DirectoryUpdateOne) sqlSave(ctx context.Context) (_node *Directory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   directory.Table,
			Columns: directory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: directory.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Directory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, directory.FieldID)
		for _, f := range fields {
			if !directory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != directory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedBy(); ok {
		_spec.SetField(directory.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := duo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(directory.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := duo.mutation.UpdatedBy(); ok {
		_spec.SetField(directory.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := duo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(directory.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(directory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.SetField(directory.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(directory.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.IsPublic(); ok {
		_spec.SetField(directory.FieldIsPublic, field.TypeBool, value)
	}
	if duo.mutation.ObjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedObjectsIDs(); len(nodes) > 0 && !duo.mutation.ObjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ObjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !duo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Directory{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{directory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
