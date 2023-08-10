// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent_sample/ent/predicate"
	"ent_sample/ent/userinfo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserInfoDelete is the builder for deleting a UserInfo entity.
type UserInfoDelete struct {
	config
	hooks    []Hook
	mutation *UserInfoMutation
}

// Where appends a list predicates to the UserInfoDelete builder.
func (uid *UserInfoDelete) Where(ps ...predicate.UserInfo) *UserInfoDelete {
	uid.mutation.Where(ps...)
	return uid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uid *UserInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uid.sqlExec, uid.mutation, uid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uid *UserInfoDelete) ExecX(ctx context.Context) int {
	n, err := uid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uid *UserInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userinfo.Table, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	if ps := uid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uid.mutation.done = true
	return affected, err
}

// UserInfoDeleteOne is the builder for deleting a single UserInfo entity.
type UserInfoDeleteOne struct {
	uid *UserInfoDelete
}

// Where appends a list predicates to the UserInfoDelete builder.
func (uido *UserInfoDeleteOne) Where(ps ...predicate.UserInfo) *UserInfoDeleteOne {
	uido.uid.mutation.Where(ps...)
	return uido
}

// Exec executes the deletion query.
func (uido *UserInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := uido.uid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uido *UserInfoDeleteOne) ExecX(ctx context.Context) {
	if err := uido.Exec(ctx); err != nil {
		panic(err)
	}
}