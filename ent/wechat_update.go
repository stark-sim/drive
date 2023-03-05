// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"drive/ent/predicate"
	"drive/ent/social"
	"drive/ent/wechat"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WechatUpdate is the builder for updating Wechat entities.
type WechatUpdate struct {
	config
	hooks    []Hook
	mutation *WechatMutation
}

// Where appends a list predicates to the WechatUpdate builder.
func (wu *WechatUpdate) Where(ps ...predicate.Wechat) *WechatUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetCreatedBy sets the "created_by" field.
func (wu *WechatUpdate) SetCreatedBy(i int64) *WechatUpdate {
	wu.mutation.ResetCreatedBy()
	wu.mutation.SetCreatedBy(i)
	return wu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (wu *WechatUpdate) SetNillableCreatedBy(i *int64) *WechatUpdate {
	if i != nil {
		wu.SetCreatedBy(*i)
	}
	return wu
}

// AddCreatedBy adds i to the "created_by" field.
func (wu *WechatUpdate) AddCreatedBy(i int64) *WechatUpdate {
	wu.mutation.AddCreatedBy(i)
	return wu
}

// SetUpdatedBy sets the "updated_by" field.
func (wu *WechatUpdate) SetUpdatedBy(i int64) *WechatUpdate {
	wu.mutation.ResetUpdatedBy()
	wu.mutation.SetUpdatedBy(i)
	return wu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (wu *WechatUpdate) SetNillableUpdatedBy(i *int64) *WechatUpdate {
	if i != nil {
		wu.SetUpdatedBy(*i)
	}
	return wu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (wu *WechatUpdate) AddUpdatedBy(i int64) *WechatUpdate {
	wu.mutation.AddUpdatedBy(i)
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WechatUpdate) SetUpdatedAt(t time.Time) *WechatUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetDeletedAt sets the "deleted_at" field.
func (wu *WechatUpdate) SetDeletedAt(t time.Time) *WechatUpdate {
	wu.mutation.SetDeletedAt(t)
	return wu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wu *WechatUpdate) SetNillableDeletedAt(t *time.Time) *WechatUpdate {
	if t != nil {
		wu.SetDeletedAt(*t)
	}
	return wu
}

// SetName sets the "name" field.
func (wu *WechatUpdate) SetName(s string) *WechatUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wu *WechatUpdate) SetNillableName(s *string) *WechatUpdate {
	if s != nil {
		wu.SetName(*s)
	}
	return wu
}

// AddSocialIDs adds the "socials" edge to the Social entity by IDs.
func (wu *WechatUpdate) AddSocialIDs(ids ...int64) *WechatUpdate {
	wu.mutation.AddSocialIDs(ids...)
	return wu
}

// AddSocials adds the "socials" edges to the Social entity.
func (wu *WechatUpdate) AddSocials(s ...*Social) *WechatUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wu.AddSocialIDs(ids...)
}

// Mutation returns the WechatMutation object of the builder.
func (wu *WechatUpdate) Mutation() *WechatMutation {
	return wu.mutation
}

// ClearSocials clears all "socials" edges to the Social entity.
func (wu *WechatUpdate) ClearSocials() *WechatUpdate {
	wu.mutation.ClearSocials()
	return wu
}

// RemoveSocialIDs removes the "socials" edge to Social entities by IDs.
func (wu *WechatUpdate) RemoveSocialIDs(ids ...int64) *WechatUpdate {
	wu.mutation.RemoveSocialIDs(ids...)
	return wu
}

// RemoveSocials removes "socials" edges to Social entities.
func (wu *WechatUpdate) RemoveSocials(s ...*Social) *WechatUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wu.RemoveSocialIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WechatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WechatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WechatUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WechatUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WechatUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WechatUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := wechat.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

func (wu *WechatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wechat.Table,
			Columns: wechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: wechat.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.CreatedBy(); ok {
		_spec.SetField(wechat.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(wechat.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.UpdatedBy(); ok {
		_spec.SetField(wechat.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(wechat.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.SetField(wechat.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wu.mutation.DeletedAt(); ok {
		_spec.SetField(wechat.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.SetField(wechat.FieldName, field.TypeString, value)
	}
	if wu.mutation.SocialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wechat.SocialsTable,
			Columns: []string{wechat.SocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: social.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedSocialsIDs(); len(nodes) > 0 && !wu.mutation.SocialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wechat.SocialsTable,
			Columns: []string{wechat.SocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: social.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.SocialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wechat.SocialsTable,
			Columns: []string{wechat.SocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: social.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wechat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WechatUpdateOne is the builder for updating a single Wechat entity.
type WechatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WechatMutation
}

// SetCreatedBy sets the "created_by" field.
func (wuo *WechatUpdateOne) SetCreatedBy(i int64) *WechatUpdateOne {
	wuo.mutation.ResetCreatedBy()
	wuo.mutation.SetCreatedBy(i)
	return wuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (wuo *WechatUpdateOne) SetNillableCreatedBy(i *int64) *WechatUpdateOne {
	if i != nil {
		wuo.SetCreatedBy(*i)
	}
	return wuo
}

// AddCreatedBy adds i to the "created_by" field.
func (wuo *WechatUpdateOne) AddCreatedBy(i int64) *WechatUpdateOne {
	wuo.mutation.AddCreatedBy(i)
	return wuo
}

// SetUpdatedBy sets the "updated_by" field.
func (wuo *WechatUpdateOne) SetUpdatedBy(i int64) *WechatUpdateOne {
	wuo.mutation.ResetUpdatedBy()
	wuo.mutation.SetUpdatedBy(i)
	return wuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (wuo *WechatUpdateOne) SetNillableUpdatedBy(i *int64) *WechatUpdateOne {
	if i != nil {
		wuo.SetUpdatedBy(*i)
	}
	return wuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (wuo *WechatUpdateOne) AddUpdatedBy(i int64) *WechatUpdateOne {
	wuo.mutation.AddUpdatedBy(i)
	return wuo
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WechatUpdateOne) SetUpdatedAt(t time.Time) *WechatUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetDeletedAt sets the "deleted_at" field.
func (wuo *WechatUpdateOne) SetDeletedAt(t time.Time) *WechatUpdateOne {
	wuo.mutation.SetDeletedAt(t)
	return wuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wuo *WechatUpdateOne) SetNillableDeletedAt(t *time.Time) *WechatUpdateOne {
	if t != nil {
		wuo.SetDeletedAt(*t)
	}
	return wuo
}

// SetName sets the "name" field.
func (wuo *WechatUpdateOne) SetName(s string) *WechatUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wuo *WechatUpdateOne) SetNillableName(s *string) *WechatUpdateOne {
	if s != nil {
		wuo.SetName(*s)
	}
	return wuo
}

// AddSocialIDs adds the "socials" edge to the Social entity by IDs.
func (wuo *WechatUpdateOne) AddSocialIDs(ids ...int64) *WechatUpdateOne {
	wuo.mutation.AddSocialIDs(ids...)
	return wuo
}

// AddSocials adds the "socials" edges to the Social entity.
func (wuo *WechatUpdateOne) AddSocials(s ...*Social) *WechatUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wuo.AddSocialIDs(ids...)
}

// Mutation returns the WechatMutation object of the builder.
func (wuo *WechatUpdateOne) Mutation() *WechatMutation {
	return wuo.mutation
}

// ClearSocials clears all "socials" edges to the Social entity.
func (wuo *WechatUpdateOne) ClearSocials() *WechatUpdateOne {
	wuo.mutation.ClearSocials()
	return wuo
}

// RemoveSocialIDs removes the "socials" edge to Social entities by IDs.
func (wuo *WechatUpdateOne) RemoveSocialIDs(ids ...int64) *WechatUpdateOne {
	wuo.mutation.RemoveSocialIDs(ids...)
	return wuo
}

// RemoveSocials removes "socials" edges to Social entities.
func (wuo *WechatUpdateOne) RemoveSocials(s ...*Social) *WechatUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return wuo.RemoveSocialIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WechatUpdateOne) Select(field string, fields ...string) *WechatUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wechat entity.
func (wuo *WechatUpdateOne) Save(ctx context.Context) (*Wechat, error) {
	var (
		err  error
		node *Wechat
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WechatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Wechat)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WechatMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WechatUpdateOne) SaveX(ctx context.Context) *Wechat {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WechatUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WechatUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WechatUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := wechat.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

func (wuo *WechatUpdateOne) sqlSave(ctx context.Context) (_node *Wechat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wechat.Table,
			Columns: wechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: wechat.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wechat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wechat.FieldID)
		for _, f := range fields {
			if !wechat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wechat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.CreatedBy(); ok {
		_spec.SetField(wechat.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(wechat.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.UpdatedBy(); ok {
		_spec.SetField(wechat.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(wechat.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.SetField(wechat.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := wuo.mutation.DeletedAt(); ok {
		_spec.SetField(wechat.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.SetField(wechat.FieldName, field.TypeString, value)
	}
	if wuo.mutation.SocialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wechat.SocialsTable,
			Columns: []string{wechat.SocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: social.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedSocialsIDs(); len(nodes) > 0 && !wuo.mutation.SocialsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wechat.SocialsTable,
			Columns: []string{wechat.SocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: social.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.SocialsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wechat.SocialsTable,
			Columns: []string{wechat.SocialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: social.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wechat{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wechat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}