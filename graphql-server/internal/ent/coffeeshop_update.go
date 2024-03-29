// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent/coffeeshop"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent/predicate"
)

// CoffeeshopUpdate is the builder for updating Coffeeshop entities.
type CoffeeshopUpdate struct {
	config
	hooks    []Hook
	mutation *CoffeeshopMutation
}

// Where appends a list predicates to the CoffeeshopUpdate builder.
func (cu *CoffeeshopUpdate) Where(ps ...predicate.Coffeeshop) *CoffeeshopUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CoffeeshopUpdate) SetName(s string) *CoffeeshopUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetAddress sets the "address" field.
func (cu *CoffeeshopUpdate) SetAddress(s string) *CoffeeshopUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetLatitude sets the "latitude" field.
func (cu *CoffeeshopUpdate) SetLatitude(f float64) *CoffeeshopUpdate {
	cu.mutation.ResetLatitude()
	cu.mutation.SetLatitude(f)
	return cu
}

// AddLatitude adds f to the "latitude" field.
func (cu *CoffeeshopUpdate) AddLatitude(f float64) *CoffeeshopUpdate {
	cu.mutation.AddLatitude(f)
	return cu
}

// SetLongitude sets the "longitude" field.
func (cu *CoffeeshopUpdate) SetLongitude(f float64) *CoffeeshopUpdate {
	cu.mutation.ResetLongitude()
	cu.mutation.SetLongitude(f)
	return cu
}

// AddLongitude adds f to the "longitude" field.
func (cu *CoffeeshopUpdate) AddLongitude(f float64) *CoffeeshopUpdate {
	cu.mutation.AddLongitude(f)
	return cu
}

// Mutation returns the CoffeeshopMutation object of the builder.
func (cu *CoffeeshopUpdate) Mutation() *CoffeeshopMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CoffeeshopUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CoffeeshopMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CoffeeshopUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CoffeeshopUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CoffeeshopUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CoffeeshopUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(coffeeshop.Table, coffeeshop.Columns, sqlgraph.NewFieldSpec(coffeeshop.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(coffeeshop.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.SetField(coffeeshop.FieldAddress, field.TypeString, value)
	}
	if value, ok := cu.mutation.Latitude(); ok {
		_spec.SetField(coffeeshop.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedLatitude(); ok {
		_spec.AddField(coffeeshop.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.Longitude(); ok {
		_spec.SetField(coffeeshop.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedLongitude(); ok {
		_spec.AddField(coffeeshop.FieldLongitude, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coffeeshop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CoffeeshopUpdateOne is the builder for updating a single Coffeeshop entity.
type CoffeeshopUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoffeeshopMutation
}

// SetName sets the "name" field.
func (cuo *CoffeeshopUpdateOne) SetName(s string) *CoffeeshopUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *CoffeeshopUpdateOne) SetAddress(s string) *CoffeeshopUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetLatitude sets the "latitude" field.
func (cuo *CoffeeshopUpdateOne) SetLatitude(f float64) *CoffeeshopUpdateOne {
	cuo.mutation.ResetLatitude()
	cuo.mutation.SetLatitude(f)
	return cuo
}

// AddLatitude adds f to the "latitude" field.
func (cuo *CoffeeshopUpdateOne) AddLatitude(f float64) *CoffeeshopUpdateOne {
	cuo.mutation.AddLatitude(f)
	return cuo
}

// SetLongitude sets the "longitude" field.
func (cuo *CoffeeshopUpdateOne) SetLongitude(f float64) *CoffeeshopUpdateOne {
	cuo.mutation.ResetLongitude()
	cuo.mutation.SetLongitude(f)
	return cuo
}

// AddLongitude adds f to the "longitude" field.
func (cuo *CoffeeshopUpdateOne) AddLongitude(f float64) *CoffeeshopUpdateOne {
	cuo.mutation.AddLongitude(f)
	return cuo
}

// Mutation returns the CoffeeshopMutation object of the builder.
func (cuo *CoffeeshopUpdateOne) Mutation() *CoffeeshopMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CoffeeshopUpdate builder.
func (cuo *CoffeeshopUpdateOne) Where(ps ...predicate.Coffeeshop) *CoffeeshopUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CoffeeshopUpdateOne) Select(field string, fields ...string) *CoffeeshopUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Coffeeshop entity.
func (cuo *CoffeeshopUpdateOne) Save(ctx context.Context) (*Coffeeshop, error) {
	return withHooks[*Coffeeshop, CoffeeshopMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CoffeeshopUpdateOne) SaveX(ctx context.Context) *Coffeeshop {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CoffeeshopUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CoffeeshopUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CoffeeshopUpdateOne) sqlSave(ctx context.Context) (_node *Coffeeshop, err error) {
	_spec := sqlgraph.NewUpdateSpec(coffeeshop.Table, coffeeshop.Columns, sqlgraph.NewFieldSpec(coffeeshop.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Coffeeshop.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coffeeshop.FieldID)
		for _, f := range fields {
			if !coffeeshop.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coffeeshop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(coffeeshop.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.SetField(coffeeshop.FieldAddress, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Latitude(); ok {
		_spec.SetField(coffeeshop.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedLatitude(); ok {
		_spec.AddField(coffeeshop.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.Longitude(); ok {
		_spec.SetField(coffeeshop.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedLongitude(); ok {
		_spec.AddField(coffeeshop.FieldLongitude, field.TypeFloat64, value)
	}
	_node = &Coffeeshop{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coffeeshop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
