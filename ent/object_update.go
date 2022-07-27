// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/object"
	"drive/ent/predicate"
	"drive/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ObjectUpdate is the builder for updating Object entities.
type ObjectUpdate struct {
	config
	hooks    []Hook
	mutation *ObjectMutation
}

// Where appends a list predicates to the ObjectUpdate builder.
func (ou *ObjectUpdate) Where(ps ...predicate.Object) *ObjectUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCreatedBy sets the "created_by" field.
func (ou *ObjectUpdate) SetCreatedBy(i int64) *ObjectUpdate {
	ou.mutation.ResetCreatedBy()
	ou.mutation.SetCreatedBy(i)
	return ou
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ou *ObjectUpdate) SetNillableCreatedBy(i *int64) *ObjectUpdate {
	if i != nil {
		ou.SetCreatedBy(*i)
	}
	return ou
}

// AddCreatedBy adds i to the "created_by" field.
func (ou *ObjectUpdate) AddCreatedBy(i int64) *ObjectUpdate {
	ou.mutation.AddCreatedBy(i)
	return ou
}

// SetUpdatedBy sets the "updated_by" field.
func (ou *ObjectUpdate) SetUpdatedBy(i int64) *ObjectUpdate {
	ou.mutation.ResetUpdatedBy()
	ou.mutation.SetUpdatedBy(i)
	return ou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ou *ObjectUpdate) SetNillableUpdatedBy(i *int64) *ObjectUpdate {
	if i != nil {
		ou.SetUpdatedBy(*i)
	}
	return ou
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ou *ObjectUpdate) AddUpdatedBy(i int64) *ObjectUpdate {
	ou.mutation.AddUpdatedBy(i)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *ObjectUpdate) SetUpdatedAt(t time.Time) *ObjectUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetDeletedAt sets the "deleted_at" field.
func (ou *ObjectUpdate) SetDeletedAt(t time.Time) *ObjectUpdate {
	ou.mutation.SetDeletedAt(t)
	return ou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ou *ObjectUpdate) SetNillableDeletedAt(t *time.Time) *ObjectUpdate {
	if t != nil {
		ou.SetDeletedAt(*t)
	}
	return ou
}

// SetURL sets the "url" field.
func (ou *ObjectUpdate) SetURL(s string) *ObjectUpdate {
	ou.mutation.SetURL(s)
	return ou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *ObjectUpdate) SetUserID(id int64) *ObjectUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ou *ObjectUpdate) SetNillableUserID(id *int64) *ObjectUpdate {
	if id != nil {
		ou = ou.SetUserID(*id)
	}
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *ObjectUpdate) SetUser(u *User) *ObjectUpdate {
	return ou.SetUserID(u.ID)
}

// Mutation returns the ObjectMutation object of the builder.
func (ou *ObjectUpdate) Mutation() *ObjectMutation {
	return ou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ou *ObjectUpdate) ClearUser() *ObjectUpdate {
	ou.mutation.ClearUser()
	return ou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *ObjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *ObjectUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *ObjectUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *ObjectUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *ObjectUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := object.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

func (ou *ObjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   object.Table,
			Columns: object.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: object.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldCreatedBy,
		})
	}
	if value, ok := ou.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldCreatedBy,
		})
	}
	if value, ok := ou.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldUpdatedBy,
		})
	}
	if value, ok := ou.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldUpdatedBy,
		})
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: object.FieldURL,
		})
	}
	if ou.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{object.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ObjectUpdateOne is the builder for updating a single Object entity.
type ObjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ObjectMutation
}

// SetCreatedBy sets the "created_by" field.
func (ouo *ObjectUpdateOne) SetCreatedBy(i int64) *ObjectUpdateOne {
	ouo.mutation.ResetCreatedBy()
	ouo.mutation.SetCreatedBy(i)
	return ouo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ouo *ObjectUpdateOne) SetNillableCreatedBy(i *int64) *ObjectUpdateOne {
	if i != nil {
		ouo.SetCreatedBy(*i)
	}
	return ouo
}

// AddCreatedBy adds i to the "created_by" field.
func (ouo *ObjectUpdateOne) AddCreatedBy(i int64) *ObjectUpdateOne {
	ouo.mutation.AddCreatedBy(i)
	return ouo
}

// SetUpdatedBy sets the "updated_by" field.
func (ouo *ObjectUpdateOne) SetUpdatedBy(i int64) *ObjectUpdateOne {
	ouo.mutation.ResetUpdatedBy()
	ouo.mutation.SetUpdatedBy(i)
	return ouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouo *ObjectUpdateOne) SetNillableUpdatedBy(i *int64) *ObjectUpdateOne {
	if i != nil {
		ouo.SetUpdatedBy(*i)
	}
	return ouo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ouo *ObjectUpdateOne) AddUpdatedBy(i int64) *ObjectUpdateOne {
	ouo.mutation.AddUpdatedBy(i)
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *ObjectUpdateOne) SetUpdatedAt(t time.Time) *ObjectUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetDeletedAt sets the "deleted_at" field.
func (ouo *ObjectUpdateOne) SetDeletedAt(t time.Time) *ObjectUpdateOne {
	ouo.mutation.SetDeletedAt(t)
	return ouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouo *ObjectUpdateOne) SetNillableDeletedAt(t *time.Time) *ObjectUpdateOne {
	if t != nil {
		ouo.SetDeletedAt(*t)
	}
	return ouo
}

// SetURL sets the "url" field.
func (ouo *ObjectUpdateOne) SetURL(s string) *ObjectUpdateOne {
	ouo.mutation.SetURL(s)
	return ouo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *ObjectUpdateOne) SetUserID(id int64) *ObjectUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ouo *ObjectUpdateOne) SetNillableUserID(id *int64) *ObjectUpdateOne {
	if id != nil {
		ouo = ouo.SetUserID(*id)
	}
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *ObjectUpdateOne) SetUser(u *User) *ObjectUpdateOne {
	return ouo.SetUserID(u.ID)
}

// Mutation returns the ObjectMutation object of the builder.
func (ouo *ObjectUpdateOne) Mutation() *ObjectMutation {
	return ouo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *ObjectUpdateOne) ClearUser() *ObjectUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *ObjectUpdateOne) Select(field string, fields ...string) *ObjectUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Object entity.
func (ouo *ObjectUpdateOne) Save(ctx context.Context) (*Object, error) {
	var (
		err  error
		node *Object
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ouo *ObjectUpdateOne) SaveX(ctx context.Context) *Object {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *ObjectUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *ObjectUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *ObjectUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := object.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

func (ouo *ObjectUpdateOne) sqlSave(ctx context.Context) (_node *Object, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   object.Table,
			Columns: object.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: object.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Object.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, object.FieldID)
		for _, f := range fields {
			if !object.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != object.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldCreatedBy,
		})
	}
	if value, ok := ouo.mutation.AddedCreatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldCreatedBy,
		})
	}
	if value, ok := ouo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldUpdatedBy,
		})
	}
	if value, ok := ouo.mutation.AddedUpdatedBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: object.FieldUpdatedBy,
		})
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: object.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: object.FieldURL,
		})
	}
	if ouo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Object{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{object.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}