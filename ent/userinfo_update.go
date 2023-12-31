// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent_sample/ent/predicate"
	"ent_sample/ent/user"
	"ent_sample/ent/userinfo"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserInfoUpdate is the builder for updating UserInfo entities.
type UserInfoUpdate struct {
	config
	hooks    []Hook
	mutation *UserInfoMutation
}

// Where appends a list predicates to the UserInfoUpdate builder.
func (uiu *UserInfoUpdate) Where(ps ...predicate.UserInfo) *UserInfoUpdate {
	uiu.mutation.Where(ps...)
	return uiu
}

// SetEtc sets the "etc" field.
func (uiu *UserInfoUpdate) SetEtc(s string) *UserInfoUpdate {
	uiu.mutation.SetEtc(s)
	return uiu
}

// SetNillableEtc sets the "etc" field if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableEtc(s *string) *UserInfoUpdate {
	if s != nil {
		uiu.SetEtc(*s)
	}
	return uiu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (uiu *UserInfoUpdate) SetUsersID(id int) *UserInfoUpdate {
	uiu.mutation.SetUsersID(id)
	return uiu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (uiu *UserInfoUpdate) SetNillableUsersID(id *int) *UserInfoUpdate {
	if id != nil {
		uiu = uiu.SetUsersID(*id)
	}
	return uiu
}

// SetUsers sets the "users" edge to the User entity.
func (uiu *UserInfoUpdate) SetUsers(u *User) *UserInfoUpdate {
	return uiu.SetUsersID(u.ID)
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiu *UserInfoUpdate) Mutation() *UserInfoMutation {
	return uiu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (uiu *UserInfoUpdate) ClearUsers() *UserInfoUpdate {
	uiu.mutation.ClearUsers()
	return uiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uiu *UserInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uiu.sqlSave, uiu.mutation, uiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uiu *UserInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := uiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uiu *UserInfoUpdate) Exec(ctx context.Context) error {
	_, err := uiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiu *UserInfoUpdate) ExecX(ctx context.Context) {
	if err := uiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uiu *UserInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userinfo.Table, userinfo.Columns, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	if ps := uiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiu.mutation.Etc(); ok {
		_spec.SetField(userinfo.FieldEtc, field.TypeString, value)
	}
	if uiu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userinfo.UsersTable,
			Columns: []string{userinfo.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uiu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userinfo.UsersTable,
			Columns: []string{userinfo.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uiu.mutation.done = true
	return n, nil
}

// UserInfoUpdateOne is the builder for updating a single UserInfo entity.
type UserInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserInfoMutation
}

// SetEtc sets the "etc" field.
func (uiuo *UserInfoUpdateOne) SetEtc(s string) *UserInfoUpdateOne {
	uiuo.mutation.SetEtc(s)
	return uiuo
}

// SetNillableEtc sets the "etc" field if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableEtc(s *string) *UserInfoUpdateOne {
	if s != nil {
		uiuo.SetEtc(*s)
	}
	return uiuo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (uiuo *UserInfoUpdateOne) SetUsersID(id int) *UserInfoUpdateOne {
	uiuo.mutation.SetUsersID(id)
	return uiuo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (uiuo *UserInfoUpdateOne) SetNillableUsersID(id *int) *UserInfoUpdateOne {
	if id != nil {
		uiuo = uiuo.SetUsersID(*id)
	}
	return uiuo
}

// SetUsers sets the "users" edge to the User entity.
func (uiuo *UserInfoUpdateOne) SetUsers(u *User) *UserInfoUpdateOne {
	return uiuo.SetUsersID(u.ID)
}

// Mutation returns the UserInfoMutation object of the builder.
func (uiuo *UserInfoUpdateOne) Mutation() *UserInfoMutation {
	return uiuo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (uiuo *UserInfoUpdateOne) ClearUsers() *UserInfoUpdateOne {
	uiuo.mutation.ClearUsers()
	return uiuo
}

// Where appends a list predicates to the UserInfoUpdate builder.
func (uiuo *UserInfoUpdateOne) Where(ps ...predicate.UserInfo) *UserInfoUpdateOne {
	uiuo.mutation.Where(ps...)
	return uiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uiuo *UserInfoUpdateOne) Select(field string, fields ...string) *UserInfoUpdateOne {
	uiuo.fields = append([]string{field}, fields...)
	return uiuo
}

// Save executes the query and returns the updated UserInfo entity.
func (uiuo *UserInfoUpdateOne) Save(ctx context.Context) (*UserInfo, error) {
	return withHooks(ctx, uiuo.sqlSave, uiuo.mutation, uiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) SaveX(ctx context.Context) *UserInfo {
	node, err := uiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uiuo *UserInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := uiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uiuo *UserInfoUpdateOne) ExecX(ctx context.Context) {
	if err := uiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uiuo *UserInfoUpdateOne) sqlSave(ctx context.Context) (_node *UserInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(userinfo.Table, userinfo.Columns, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	id, ok := uiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userinfo.FieldID)
		for _, f := range fields {
			if !userinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uiuo.mutation.Etc(); ok {
		_spec.SetField(userinfo.FieldEtc, field.TypeString, value)
	}
	if uiuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userinfo.UsersTable,
			Columns: []string{userinfo.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uiuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userinfo.UsersTable,
			Columns: []string{userinfo.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserInfo{config: uiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uiuo.mutation.done = true
	return _node, nil
}
