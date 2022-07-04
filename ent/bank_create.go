// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wynn-member-api/ent/bank"
	"wynn-member-api/ent/bankaccount"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BankCreate is the builder for creating a Bank entity.
type BankCreate struct {
	config
	mutation *BankMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (bc *BankCreate) SetUUID(u uuid.UUID) *BankCreate {
	bc.mutation.SetUUID(u)
	return bc
}

// SetName sets the "name" field.
func (bc *BankCreate) SetName(s string) *BankCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetShortName sets the "short_name" field.
func (bc *BankCreate) SetShortName(s string) *BankCreate {
	bc.mutation.SetShortName(s)
	return bc
}

// SetNillableShortName sets the "short_name" field if the given value is not nil.
func (bc *BankCreate) SetNillableShortName(s *string) *BankCreate {
	if s != nil {
		bc.SetShortName(*s)
	}
	return bc
}

// SetNameTh sets the "name_th" field.
func (bc *BankCreate) SetNameTh(s string) *BankCreate {
	bc.mutation.SetNameTh(s)
	return bc
}

// SetNillableNameTh sets the "name_th" field if the given value is not nil.
func (bc *BankCreate) SetNillableNameTh(s *string) *BankCreate {
	if s != nil {
		bc.SetNameTh(*s)
	}
	return bc
}

// SetShortNameTh sets the "short_name_th" field.
func (bc *BankCreate) SetShortNameTh(s string) *BankCreate {
	bc.mutation.SetShortNameTh(s)
	return bc
}

// SetNillableShortNameTh sets the "short_name_th" field if the given value is not nil.
func (bc *BankCreate) SetNillableShortNameTh(s *string) *BankCreate {
	if s != nil {
		bc.SetShortNameTh(*s)
	}
	return bc
}

// SetBankAccountName sets the "bank_account_name" field.
func (bc *BankCreate) SetBankAccountName(s string) *BankCreate {
	bc.mutation.SetBankAccountName(s)
	return bc
}

// SetNillableBankAccountName sets the "bank_account_name" field if the given value is not nil.
func (bc *BankCreate) SetNillableBankAccountName(s *string) *BankCreate {
	if s != nil {
		bc.SetBankAccountName(*s)
	}
	return bc
}

// SetLogo sets the "logo" field.
func (bc *BankCreate) SetLogo(s string) *BankCreate {
	bc.mutation.SetLogo(s)
	return bc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (bc *BankCreate) SetNillableLogo(s *string) *BankCreate {
	if s != nil {
		bc.SetLogo(*s)
	}
	return bc
}

// SetBankID sets the "bank_id" field.
func (bc *BankCreate) SetBankID(s string) *BankCreate {
	bc.mutation.SetBankID(s)
	return bc
}

// SetNillableBankID sets the "bank_id" field if the given value is not nil.
func (bc *BankCreate) SetNillableBankID(s *string) *BankCreate {
	if s != nil {
		bc.SetBankID(*s)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BankCreate) SetStatus(b bank.Status) *BankCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bc *BankCreate) SetNillableStatus(b *bank.Status) *BankCreate {
	if b != nil {
		bc.SetStatus(*b)
	}
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BankCreate) SetCreatedAt(t time.Time) *BankCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BankCreate) SetNillableCreatedAt(t *time.Time) *BankCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BankCreate) SetUpdatedAt(t time.Time) *BankCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BankCreate) SetNillableUpdatedAt(t *time.Time) *BankCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// AddAccountIDs adds the "accounts" edge to the BankAccount entity by IDs.
func (bc *BankCreate) AddAccountIDs(ids ...int) *BankCreate {
	bc.mutation.AddAccountIDs(ids...)
	return bc
}

// AddAccounts adds the "accounts" edges to the BankAccount entity.
func (bc *BankCreate) AddAccounts(b ...*BankAccount) *BankCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddAccountIDs(ids...)
}

// Mutation returns the BankMutation object of the builder.
func (bc *BankCreate) Mutation() *BankMutation {
	return bc.mutation
}

// Save creates the Bank in the database.
func (bc *BankCreate) Save(ctx context.Context) (*Bank, error) {
	var (
		err  error
		node *Bank
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BankCreate) SaveX(ctx context.Context) *Bank {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BankCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BankCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BankCreate) defaults() {
	if _, ok := bc.mutation.UUID(); !ok {
		v := bank.DefaultUUID()
		bc.mutation.SetUUID(v)
	}
	if _, ok := bc.mutation.Status(); !ok {
		v := bank.DefaultStatus
		bc.mutation.SetStatus(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := bank.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := bank.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BankCreate) check() error {
	if _, ok := bc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "uuid"`)}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := bank.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (bc *BankCreate) sqlSave(ctx context.Context) (*Bank, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BankCreate) createSpec() (*Bank, *sqlgraph.CreateSpec) {
	var (
		_node = &Bank{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bank.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bank.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: bank.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.ShortName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldShortName,
		})
		_node.ShortName = value
	}
	if value, ok := bc.mutation.NameTh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldNameTh,
		})
		_node.NameTh = value
	}
	if value, ok := bc.mutation.ShortNameTh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldShortNameTh,
		})
		_node.ShortNameTh = value
	}
	if value, ok := bc.mutation.BankAccountName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldBankAccountName,
		})
		_node.BankAccountName = value
	}
	if value, ok := bc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := bc.mutation.BankID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bank.FieldBankID,
		})
		_node.BankID = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: bank.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bank.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bank.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   bank.AccountsTable,
			Columns: []string{bank.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bankaccount.FieldID,
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

// BankCreateBulk is the builder for creating many Bank entities in bulk.
type BankCreateBulk struct {
	config
	builders []*BankCreate
}

// Save creates the Bank entities in the database.
func (bcb *BankCreateBulk) Save(ctx context.Context) ([]*Bank, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bank, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BankMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BankCreateBulk) SaveX(ctx context.Context) []*Bank {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BankCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BankCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
