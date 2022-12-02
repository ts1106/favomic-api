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
	"github.com/google/uuid"
	"github.com/ts1106/favomic-api/ent/comic"
	"github.com/ts1106/favomic-api/ent/episode"
	"github.com/ts1106/favomic-api/ent/predicate"
)

// EpisodeUpdate is the builder for updating Episode entities.
type EpisodeUpdate struct {
	config
	hooks    []Hook
	mutation *EpisodeMutation
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (eu *EpisodeUpdate) Where(ps ...predicate.Episode) *EpisodeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetTitle sets the "title" field.
func (eu *EpisodeUpdate) SetTitle(s string) *EpisodeUpdate {
	eu.mutation.SetTitle(s)
	return eu
}

// SetURL sets the "url" field.
func (eu *EpisodeUpdate) SetURL(s string) *EpisodeUpdate {
	eu.mutation.SetURL(s)
	return eu
}

// SetThumbnail sets the "Thumbnail" field.
func (eu *EpisodeUpdate) SetThumbnail(s string) *EpisodeUpdate {
	eu.mutation.SetThumbnail(s)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EpisodeUpdate) SetUpdatedAt(t time.Time) *EpisodeUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableUpdatedAt(t *time.Time) *EpisodeUpdate {
	if t != nil {
		eu.SetUpdatedAt(*t)
	}
	return eu
}

// SetComicID sets the "comic_id" field.
func (eu *EpisodeUpdate) SetComicID(u uuid.UUID) *EpisodeUpdate {
	eu.mutation.SetComicID(u)
	return eu
}

// SetComic sets the "comic" edge to the Comic entity.
func (eu *EpisodeUpdate) SetComic(c *Comic) *EpisodeUpdate {
	return eu.SetComicID(c.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (eu *EpisodeUpdate) Mutation() *EpisodeMutation {
	return eu.mutation
}

// ClearComic clears the "comic" edge to the Comic entity.
func (eu *EpisodeUpdate) ClearComic() *EpisodeUpdate {
	eu.mutation.ClearComic()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EpisodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EpisodeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EpisodeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EpisodeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EpisodeUpdate) check() error {
	if v, ok := eu.mutation.Title(); ok {
		if err := episode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Episode.title": %w`, err)}
		}
	}
	if v, ok := eu.mutation.URL(); ok {
		if err := episode.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Episode.url": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Thumbnail(); ok {
		if err := episode.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "Thumbnail", err: fmt.Errorf(`ent: validator failed for field "Episode.Thumbnail": %w`, err)}
		}
	}
	if _, ok := eu.mutation.ComicID(); eu.mutation.ComicCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Episode.comic"`)
	}
	return nil
}

func (eu *EpisodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   episode.Table,
			Columns: episode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: episode.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Title(); ok {
		_spec.SetField(episode.FieldTitle, field.TypeString, value)
	}
	if value, ok := eu.mutation.URL(); ok {
		_spec.SetField(episode.FieldURL, field.TypeString, value)
	}
	if value, ok := eu.mutation.Thumbnail(); ok {
		_spec.SetField(episode.FieldThumbnail, field.TypeString, value)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(episode.FieldUpdatedAt, field.TypeTime, value)
	}
	if eu.mutation.ComicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ComicTable,
			Columns: []string{episode.ComicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: comic.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ComicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ComicTable,
			Columns: []string{episode.ComicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: comic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EpisodeUpdateOne is the builder for updating a single Episode entity.
type EpisodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EpisodeMutation
}

// SetTitle sets the "title" field.
func (euo *EpisodeUpdateOne) SetTitle(s string) *EpisodeUpdateOne {
	euo.mutation.SetTitle(s)
	return euo
}

// SetURL sets the "url" field.
func (euo *EpisodeUpdateOne) SetURL(s string) *EpisodeUpdateOne {
	euo.mutation.SetURL(s)
	return euo
}

// SetThumbnail sets the "Thumbnail" field.
func (euo *EpisodeUpdateOne) SetThumbnail(s string) *EpisodeUpdateOne {
	euo.mutation.SetThumbnail(s)
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EpisodeUpdateOne) SetUpdatedAt(t time.Time) *EpisodeUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableUpdatedAt(t *time.Time) *EpisodeUpdateOne {
	if t != nil {
		euo.SetUpdatedAt(*t)
	}
	return euo
}

// SetComicID sets the "comic_id" field.
func (euo *EpisodeUpdateOne) SetComicID(u uuid.UUID) *EpisodeUpdateOne {
	euo.mutation.SetComicID(u)
	return euo
}

// SetComic sets the "comic" edge to the Comic entity.
func (euo *EpisodeUpdateOne) SetComic(c *Comic) *EpisodeUpdateOne {
	return euo.SetComicID(c.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (euo *EpisodeUpdateOne) Mutation() *EpisodeMutation {
	return euo.mutation
}

// ClearComic clears the "comic" edge to the Comic entity.
func (euo *EpisodeUpdateOne) ClearComic() *EpisodeUpdateOne {
	euo.mutation.ClearComic()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EpisodeUpdateOne) Select(field string, fields ...string) *EpisodeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Episode entity.
func (euo *EpisodeUpdateOne) Save(ctx context.Context) (*Episode, error) {
	var (
		err  error
		node *Episode
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Episode)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EpisodeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EpisodeUpdateOne) SaveX(ctx context.Context) *Episode {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EpisodeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EpisodeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EpisodeUpdateOne) check() error {
	if v, ok := euo.mutation.Title(); ok {
		if err := episode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Episode.title": %w`, err)}
		}
	}
	if v, ok := euo.mutation.URL(); ok {
		if err := episode.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Episode.url": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Thumbnail(); ok {
		if err := episode.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "Thumbnail", err: fmt.Errorf(`ent: validator failed for field "Episode.Thumbnail": %w`, err)}
		}
	}
	if _, ok := euo.mutation.ComicID(); euo.mutation.ComicCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Episode.comic"`)
	}
	return nil
}

func (euo *EpisodeUpdateOne) sqlSave(ctx context.Context) (_node *Episode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   episode.Table,
			Columns: episode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: episode.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Episode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, episode.FieldID)
		for _, f := range fields {
			if !episode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != episode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Title(); ok {
		_spec.SetField(episode.FieldTitle, field.TypeString, value)
	}
	if value, ok := euo.mutation.URL(); ok {
		_spec.SetField(episode.FieldURL, field.TypeString, value)
	}
	if value, ok := euo.mutation.Thumbnail(); ok {
		_spec.SetField(episode.FieldThumbnail, field.TypeString, value)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(episode.FieldUpdatedAt, field.TypeTime, value)
	}
	if euo.mutation.ComicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ComicTable,
			Columns: []string{episode.ComicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: comic.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ComicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ComicTable,
			Columns: []string{episode.ComicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: comic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Episode{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}