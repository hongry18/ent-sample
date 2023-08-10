// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent_sample/ent/userinfo"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserInfoCreate is the builder for creating a UserInfo entity.
type UserInfoCreate struct {
	config
	mutation *UserInfoMutation
	hooks    []Hook
}

// Mutation returns the UserInfoMutation object of the builder.
func (uic *UserInfoCreate) Mutation() *UserInfoMutation {
	return uic.mutation
}

// Save creates the UserInfo in the database.
func (uic *UserInfoCreate) Save(ctx context.Context) (*UserInfo, error) {
	return withHooks(ctx, uic.sqlSave, uic.mutation, uic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uic *UserInfoCreate) SaveX(ctx context.Context) *UserInfo {
	v, err := uic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uic *UserInfoCreate) Exec(ctx context.Context) error {
	_, err := uic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uic *UserInfoCreate) ExecX(ctx context.Context) {
	if err := uic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserInfoCreate) check() error {
	return nil
}

func (uic *UserInfoCreate) sqlSave(ctx context.Context) (*UserInfo, error) {
	if err := uic.check(); err != nil {
		return nil, err
	}
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uic.mutation.id = &_node.ID
	uic.mutation.done = true
	return _node, nil
}

func (uic *UserInfoCreate) createSpec() (*UserInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &UserInfo{config: uic.config}
		_spec = sqlgraph.NewCreateSpec(userinfo.Table, sqlgraph.NewFieldSpec(userinfo.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// UserInfoCreateBulk is the builder for creating many UserInfo entities in bulk.
type UserInfoCreateBulk struct {
	config
	builders []*UserInfoCreate
}

// Save creates the UserInfo entities in the database.
func (uicb *UserInfoCreateBulk) Save(ctx context.Context) ([]*UserInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uicb.builders))
	nodes := make([]*UserInfo, len(uicb.builders))
	mutators := make([]Mutator, len(uicb.builders))
	for i := range uicb.builders {
		func(i int, root context.Context) {
			builder := uicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) SaveX(ctx context.Context) []*UserInfo {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uicb *UserInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := uicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicb *UserInfoCreateBulk) ExecX(ctx context.Context) {
	if err := uicb.Exec(ctx); err != nil {
		panic(err)
	}
}
