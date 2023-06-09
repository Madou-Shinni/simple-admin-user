// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple-admin-user-rpc/ent/predicate"
	"simple-admin-user-rpc/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(u user.Gender) *UserUpdate {
	uu.mutation.SetGender(u)
	return uu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGender(u *user.Gender) *UserUpdate {
	if u != nil {
		uu.SetGender(*u)
	}
	return uu
}

// ClearGender clears the value of the "gender" field.
func (uu *UserUpdate) ClearGender() *UserUpdate {
	uu.mutation.ClearGender()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetIDCard sets the "id_card" field.
func (uu *UserUpdate) SetIDCard(s string) *UserUpdate {
	uu.mutation.SetIDCard(s)
	return uu
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIDCard(s *string) *UserUpdate {
	if s != nil {
		uu.SetIDCard(*s)
	}
	return uu
}

// ClearIDCard clears the value of the "id_card" field.
func (uu *UserUpdate) ClearIDCard() *UserUpdate {
	uu.mutation.ClearIDCard()
	return uu
}

// SetOpenpid sets the "openpid" field.
func (uu *UserUpdate) SetOpenpid(s string) *UserUpdate {
	uu.mutation.SetOpenpid(s)
	return uu
}

// SetNillableOpenpid sets the "openpid" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOpenpid(s *string) *UserUpdate {
	if s != nil {
		uu.SetOpenpid(*s)
	}
	return uu
}

// ClearOpenpid clears the value of the "openpid" field.
func (uu *UserUpdate) ClearOpenpid() *UserUpdate {
	uu.mutation.ClearOpenpid()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeEnum, value)
	}
	if uu.mutation.GenderCleared() {
		_spec.ClearField(user.FieldGender, field.TypeEnum)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if value, ok := uu.mutation.IDCard(); ok {
		_spec.SetField(user.FieldIDCard, field.TypeString, value)
	}
	if uu.mutation.IDCardCleared() {
		_spec.ClearField(user.FieldIDCard, field.TypeString)
	}
	if value, ok := uu.mutation.Openpid(); ok {
		_spec.SetField(user.FieldOpenpid, field.TypeString, value)
	}
	if uu.mutation.OpenpidCleared() {
		_spec.ClearField(user.FieldOpenpid, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(u user.Gender) *UserUpdateOne {
	uuo.mutation.SetGender(u)
	return uuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGender(u *user.Gender) *UserUpdateOne {
	if u != nil {
		uuo.SetGender(*u)
	}
	return uuo
}

// ClearGender clears the value of the "gender" field.
func (uuo *UserUpdateOne) ClearGender() *UserUpdateOne {
	uuo.mutation.ClearGender()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetIDCard sets the "id_card" field.
func (uuo *UserUpdateOne) SetIDCard(s string) *UserUpdateOne {
	uuo.mutation.SetIDCard(s)
	return uuo
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIDCard(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIDCard(*s)
	}
	return uuo
}

// ClearIDCard clears the value of the "id_card" field.
func (uuo *UserUpdateOne) ClearIDCard() *UserUpdateOne {
	uuo.mutation.ClearIDCard()
	return uuo
}

// SetOpenpid sets the "openpid" field.
func (uuo *UserUpdateOne) SetOpenpid(s string) *UserUpdateOne {
	uuo.mutation.SetOpenpid(s)
	return uuo
}

// SetNillableOpenpid sets the "openpid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOpenpid(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOpenpid(*s)
	}
	return uuo
}

// ClearOpenpid clears the value of the "openpid" field.
func (uuo *UserUpdateOne) ClearOpenpid() *UserUpdateOne {
	uuo.mutation.ClearOpenpid()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "User.gender": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeEnum, value)
	}
	if uuo.mutation.GenderCleared() {
		_spec.ClearField(user.FieldGender, field.TypeEnum)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IDCard(); ok {
		_spec.SetField(user.FieldIDCard, field.TypeString, value)
	}
	if uuo.mutation.IDCardCleared() {
		_spec.ClearField(user.FieldIDCard, field.TypeString)
	}
	if value, ok := uuo.mutation.Openpid(); ok {
		_spec.SetField(user.FieldOpenpid, field.TypeString, value)
	}
	if uuo.mutation.OpenpidCleared() {
		_spec.ClearField(user.FieldOpenpid, field.TypeString)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
