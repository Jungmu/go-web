// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jungmu/go-web/ent/bloglog"
)

// BlogLogCreate is the builder for creating a BlogLog entity.
type BlogLogCreate struct {
	config
	mutation *BlogLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBlogID sets the "blog_id" field.
func (blc *BlogLogCreate) SetBlogID(i int64) *BlogLogCreate {
	blc.mutation.SetBlogID(i)
	return blc
}

// SetURL sets the "url" field.
func (blc *BlogLogCreate) SetURL(s string) *BlogLogCreate {
	blc.mutation.SetURL(s)
	return blc
}

// SetReason sets the "reason" field.
func (blc *BlogLogCreate) SetReason(s string) *BlogLogCreate {
	blc.mutation.SetReason(s)
	return blc
}

// SetDetail sets the "detail" field.
func (blc *BlogLogCreate) SetDetail(s string) *BlogLogCreate {
	blc.mutation.SetDetail(s)
	return blc
}

// SetClientIP sets the "client_ip" field.
func (blc *BlogLogCreate) SetClientIP(s string) *BlogLogCreate {
	blc.mutation.SetClientIP(s)
	return blc
}

// SetCreateDatetime sets the "create_datetime" field.
func (blc *BlogLogCreate) SetCreateDatetime(t time.Time) *BlogLogCreate {
	blc.mutation.SetCreateDatetime(t)
	return blc
}

// SetNillableCreateDatetime sets the "create_datetime" field if the given value is not nil.
func (blc *BlogLogCreate) SetNillableCreateDatetime(t *time.Time) *BlogLogCreate {
	if t != nil {
		blc.SetCreateDatetime(*t)
	}
	return blc
}

// SetID sets the "id" field.
func (blc *BlogLogCreate) SetID(i int64) *BlogLogCreate {
	blc.mutation.SetID(i)
	return blc
}

// Mutation returns the BlogLogMutation object of the builder.
func (blc *BlogLogCreate) Mutation() *BlogLogMutation {
	return blc.mutation
}

// Save creates the BlogLog in the database.
func (blc *BlogLogCreate) Save(ctx context.Context) (*BlogLog, error) {
	var (
		err  error
		node *BlogLog
	)
	blc.defaults()
	if len(blc.hooks) == 0 {
		if err = blc.check(); err != nil {
			return nil, err
		}
		node, err = blc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlogLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = blc.check(); err != nil {
				return nil, err
			}
			blc.mutation = mutation
			if node, err = blc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(blc.hooks) - 1; i >= 0; i-- {
			if blc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = blc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, blc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*BlogLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlogLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (blc *BlogLogCreate) SaveX(ctx context.Context) *BlogLog {
	v, err := blc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (blc *BlogLogCreate) Exec(ctx context.Context) error {
	_, err := blc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blc *BlogLogCreate) ExecX(ctx context.Context) {
	if err := blc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (blc *BlogLogCreate) defaults() {
	if _, ok := blc.mutation.CreateDatetime(); !ok {
		v := bloglog.DefaultCreateDatetime()
		blc.mutation.SetCreateDatetime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (blc *BlogLogCreate) check() error {
	if _, ok := blc.mutation.BlogID(); !ok {
		return &ValidationError{Name: "blog_id", err: errors.New(`ent: missing required field "BlogLog.blog_id"`)}
	}
	if _, ok := blc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "BlogLog.url"`)}
	}
	if _, ok := blc.mutation.Reason(); !ok {
		return &ValidationError{Name: "reason", err: errors.New(`ent: missing required field "BlogLog.reason"`)}
	}
	if _, ok := blc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New(`ent: missing required field "BlogLog.detail"`)}
	}
	if _, ok := blc.mutation.ClientIP(); !ok {
		return &ValidationError{Name: "client_ip", err: errors.New(`ent: missing required field "BlogLog.client_ip"`)}
	}
	if _, ok := blc.mutation.CreateDatetime(); !ok {
		return &ValidationError{Name: "create_datetime", err: errors.New(`ent: missing required field "BlogLog.create_datetime"`)}
	}
	return nil
}

func (blc *BlogLogCreate) sqlSave(ctx context.Context) (*BlogLog, error) {
	_node, _spec := blc.createSpec()
	if err := sqlgraph.CreateNode(ctx, blc.driver, _spec); err != nil {
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

func (blc *BlogLogCreate) createSpec() (*BlogLog, *sqlgraph.CreateSpec) {
	var (
		_node = &BlogLog{config: blc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bloglog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: bloglog.FieldID,
			},
		}
	)
	_spec.OnConflict = blc.conflict
	if id, ok := blc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := blc.mutation.BlogID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: bloglog.FieldBlogID,
		})
		_node.BlogID = value
	}
	if value, ok := blc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bloglog.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := blc.mutation.Reason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bloglog.FieldReason,
		})
		_node.Reason = value
	}
	if value, ok := blc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bloglog.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := blc.mutation.ClientIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bloglog.FieldClientIP,
		})
		_node.ClientIP = value
	}
	if value, ok := blc.mutation.CreateDatetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bloglog.FieldCreateDatetime,
		})
		_node.CreateDatetime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BlogLog.Create().
//		SetBlogID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlogLogUpsert) {
//			SetBlogID(v+v).
//		}).
//		Exec(ctx)
func (blc *BlogLogCreate) OnConflict(opts ...sql.ConflictOption) *BlogLogUpsertOne {
	blc.conflict = opts
	return &BlogLogUpsertOne{
		create: blc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BlogLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (blc *BlogLogCreate) OnConflictColumns(columns ...string) *BlogLogUpsertOne {
	blc.conflict = append(blc.conflict, sql.ConflictColumns(columns...))
	return &BlogLogUpsertOne{
		create: blc,
	}
}

type (
	// BlogLogUpsertOne is the builder for "upsert"-ing
	//  one BlogLog node.
	BlogLogUpsertOne struct {
		create *BlogLogCreate
	}

	// BlogLogUpsert is the "OnConflict" setter.
	BlogLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetBlogID sets the "blog_id" field.
func (u *BlogLogUpsert) SetBlogID(v int64) *BlogLogUpsert {
	u.Set(bloglog.FieldBlogID, v)
	return u
}

// UpdateBlogID sets the "blog_id" field to the value that was provided on create.
func (u *BlogLogUpsert) UpdateBlogID() *BlogLogUpsert {
	u.SetExcluded(bloglog.FieldBlogID)
	return u
}

// AddBlogID adds v to the "blog_id" field.
func (u *BlogLogUpsert) AddBlogID(v int64) *BlogLogUpsert {
	u.Add(bloglog.FieldBlogID, v)
	return u
}

// SetURL sets the "url" field.
func (u *BlogLogUpsert) SetURL(v string) *BlogLogUpsert {
	u.Set(bloglog.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *BlogLogUpsert) UpdateURL() *BlogLogUpsert {
	u.SetExcluded(bloglog.FieldURL)
	return u
}

// SetReason sets the "reason" field.
func (u *BlogLogUpsert) SetReason(v string) *BlogLogUpsert {
	u.Set(bloglog.FieldReason, v)
	return u
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *BlogLogUpsert) UpdateReason() *BlogLogUpsert {
	u.SetExcluded(bloglog.FieldReason)
	return u
}

// SetDetail sets the "detail" field.
func (u *BlogLogUpsert) SetDetail(v string) *BlogLogUpsert {
	u.Set(bloglog.FieldDetail, v)
	return u
}

// UpdateDetail sets the "detail" field to the value that was provided on create.
func (u *BlogLogUpsert) UpdateDetail() *BlogLogUpsert {
	u.SetExcluded(bloglog.FieldDetail)
	return u
}

// SetClientIP sets the "client_ip" field.
func (u *BlogLogUpsert) SetClientIP(v string) *BlogLogUpsert {
	u.Set(bloglog.FieldClientIP, v)
	return u
}

// UpdateClientIP sets the "client_ip" field to the value that was provided on create.
func (u *BlogLogUpsert) UpdateClientIP() *BlogLogUpsert {
	u.SetExcluded(bloglog.FieldClientIP)
	return u
}

// SetCreateDatetime sets the "create_datetime" field.
func (u *BlogLogUpsert) SetCreateDatetime(v time.Time) *BlogLogUpsert {
	u.Set(bloglog.FieldCreateDatetime, v)
	return u
}

// UpdateCreateDatetime sets the "create_datetime" field to the value that was provided on create.
func (u *BlogLogUpsert) UpdateCreateDatetime() *BlogLogUpsert {
	u.SetExcluded(bloglog.FieldCreateDatetime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BlogLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(bloglog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BlogLogUpsertOne) UpdateNewValues() *BlogLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(bloglog.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BlogLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BlogLogUpsertOne) Ignore() *BlogLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlogLogUpsertOne) DoNothing() *BlogLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlogLogCreate.OnConflict
// documentation for more info.
func (u *BlogLogUpsertOne) Update(set func(*BlogLogUpsert)) *BlogLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlogLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetBlogID sets the "blog_id" field.
func (u *BlogLogUpsertOne) SetBlogID(v int64) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetBlogID(v)
	})
}

// AddBlogID adds v to the "blog_id" field.
func (u *BlogLogUpsertOne) AddBlogID(v int64) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.AddBlogID(v)
	})
}

// UpdateBlogID sets the "blog_id" field to the value that was provided on create.
func (u *BlogLogUpsertOne) UpdateBlogID() *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateBlogID()
	})
}

// SetURL sets the "url" field.
func (u *BlogLogUpsertOne) SetURL(v string) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *BlogLogUpsertOne) UpdateURL() *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateURL()
	})
}

// SetReason sets the "reason" field.
func (u *BlogLogUpsertOne) SetReason(v string) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *BlogLogUpsertOne) UpdateReason() *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateReason()
	})
}

// SetDetail sets the "detail" field.
func (u *BlogLogUpsertOne) SetDetail(v string) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetDetail(v)
	})
}

// UpdateDetail sets the "detail" field to the value that was provided on create.
func (u *BlogLogUpsertOne) UpdateDetail() *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateDetail()
	})
}

// SetClientIP sets the "client_ip" field.
func (u *BlogLogUpsertOne) SetClientIP(v string) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetClientIP(v)
	})
}

// UpdateClientIP sets the "client_ip" field to the value that was provided on create.
func (u *BlogLogUpsertOne) UpdateClientIP() *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateClientIP()
	})
}

// SetCreateDatetime sets the "create_datetime" field.
func (u *BlogLogUpsertOne) SetCreateDatetime(v time.Time) *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetCreateDatetime(v)
	})
}

// UpdateCreateDatetime sets the "create_datetime" field to the value that was provided on create.
func (u *BlogLogUpsertOne) UpdateCreateDatetime() *BlogLogUpsertOne {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateCreateDatetime()
	})
}

// Exec executes the query.
func (u *BlogLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlogLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlogLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BlogLogUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BlogLogUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BlogLogCreateBulk is the builder for creating many BlogLog entities in bulk.
type BlogLogCreateBulk struct {
	config
	builders []*BlogLogCreate
	conflict []sql.ConflictOption
}

// Save creates the BlogLog entities in the database.
func (blcb *BlogLogCreateBulk) Save(ctx context.Context) ([]*BlogLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(blcb.builders))
	nodes := make([]*BlogLog, len(blcb.builders))
	mutators := make([]Mutator, len(blcb.builders))
	for i := range blcb.builders {
		func(i int, root context.Context) {
			builder := blcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlogLogMutation)
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
					_, err = mutators[i+1].Mutate(root, blcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = blcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, blcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, blcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (blcb *BlogLogCreateBulk) SaveX(ctx context.Context) []*BlogLog {
	v, err := blcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (blcb *BlogLogCreateBulk) Exec(ctx context.Context) error {
	_, err := blcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blcb *BlogLogCreateBulk) ExecX(ctx context.Context) {
	if err := blcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BlogLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlogLogUpsert) {
//			SetBlogID(v+v).
//		}).
//		Exec(ctx)
func (blcb *BlogLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *BlogLogUpsertBulk {
	blcb.conflict = opts
	return &BlogLogUpsertBulk{
		create: blcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BlogLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (blcb *BlogLogCreateBulk) OnConflictColumns(columns ...string) *BlogLogUpsertBulk {
	blcb.conflict = append(blcb.conflict, sql.ConflictColumns(columns...))
	return &BlogLogUpsertBulk{
		create: blcb,
	}
}

// BlogLogUpsertBulk is the builder for "upsert"-ing
// a bulk of BlogLog nodes.
type BlogLogUpsertBulk struct {
	create *BlogLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BlogLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(bloglog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BlogLogUpsertBulk) UpdateNewValues() *BlogLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(bloglog.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BlogLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BlogLogUpsertBulk) Ignore() *BlogLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlogLogUpsertBulk) DoNothing() *BlogLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlogLogCreateBulk.OnConflict
// documentation for more info.
func (u *BlogLogUpsertBulk) Update(set func(*BlogLogUpsert)) *BlogLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlogLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetBlogID sets the "blog_id" field.
func (u *BlogLogUpsertBulk) SetBlogID(v int64) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetBlogID(v)
	})
}

// AddBlogID adds v to the "blog_id" field.
func (u *BlogLogUpsertBulk) AddBlogID(v int64) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.AddBlogID(v)
	})
}

// UpdateBlogID sets the "blog_id" field to the value that was provided on create.
func (u *BlogLogUpsertBulk) UpdateBlogID() *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateBlogID()
	})
}

// SetURL sets the "url" field.
func (u *BlogLogUpsertBulk) SetURL(v string) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *BlogLogUpsertBulk) UpdateURL() *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateURL()
	})
}

// SetReason sets the "reason" field.
func (u *BlogLogUpsertBulk) SetReason(v string) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetReason(v)
	})
}

// UpdateReason sets the "reason" field to the value that was provided on create.
func (u *BlogLogUpsertBulk) UpdateReason() *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateReason()
	})
}

// SetDetail sets the "detail" field.
func (u *BlogLogUpsertBulk) SetDetail(v string) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetDetail(v)
	})
}

// UpdateDetail sets the "detail" field to the value that was provided on create.
func (u *BlogLogUpsertBulk) UpdateDetail() *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateDetail()
	})
}

// SetClientIP sets the "client_ip" field.
func (u *BlogLogUpsertBulk) SetClientIP(v string) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetClientIP(v)
	})
}

// UpdateClientIP sets the "client_ip" field to the value that was provided on create.
func (u *BlogLogUpsertBulk) UpdateClientIP() *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateClientIP()
	})
}

// SetCreateDatetime sets the "create_datetime" field.
func (u *BlogLogUpsertBulk) SetCreateDatetime(v time.Time) *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.SetCreateDatetime(v)
	})
}

// UpdateCreateDatetime sets the "create_datetime" field to the value that was provided on create.
func (u *BlogLogUpsertBulk) UpdateCreateDatetime() *BlogLogUpsertBulk {
	return u.Update(func(s *BlogLogUpsert) {
		s.UpdateCreateDatetime()
	})
}

// Exec executes the query.
func (u *BlogLogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BlogLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlogLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlogLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}