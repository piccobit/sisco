// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/tag"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	mutation *ServiceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *ServiceCreate) SetName(s string) *ServiceCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *ServiceCreate) SetDescription(s string) *ServiceCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableDescription(s *string) *ServiceCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetProtocol sets the "protocol" field.
func (sc *ServiceCreate) SetProtocol(s string) *ServiceCreate {
	sc.mutation.SetProtocol(s)
	return sc
}

// SetHost sets the "host" field.
func (sc *ServiceCreate) SetHost(s string) *ServiceCreate {
	sc.mutation.SetHost(s)
	return sc
}

// SetPort sets the "port" field.
func (sc *ServiceCreate) SetPort(s string) *ServiceCreate {
	sc.mutation.SetPort(s)
	return sc
}

// SetAvailable sets the "available" field.
func (sc *ServiceCreate) SetAvailable(b bool) *ServiceCreate {
	sc.mutation.SetAvailable(b)
	return sc
}

// SetNillableAvailable sets the "available" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableAvailable(b *bool) *ServiceCreate {
	if b != nil {
		sc.SetAvailable(*b)
	}
	return sc
}

// SetHeartbeat sets the "heartbeat" field.
func (sc *ServiceCreate) SetHeartbeat(t time.Time) *ServiceCreate {
	sc.mutation.SetHeartbeat(t)
	return sc
}

// SetNillableHeartbeat sets the "heartbeat" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableHeartbeat(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetHeartbeat(*t)
	}
	return sc
}

// SetOwner sets the "owner" field.
func (sc *ServiceCreate) SetOwner(s string) *ServiceCreate {
	sc.mutation.SetOwner(s)
	return sc
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (sc *ServiceCreate) AddTagIDs(ids ...int) *ServiceCreate {
	sc.mutation.AddTagIDs(ids...)
	return sc
}

// AddTags adds the "tags" edges to the Tag entity.
func (sc *ServiceCreate) AddTags(t ...*Tag) *ServiceCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTagIDs(ids...)
}

// SetAreaID sets the "area" edge to the Area entity by ID.
func (sc *ServiceCreate) SetAreaID(id int) *ServiceCreate {
	sc.mutation.SetAreaID(id)
	return sc
}

// SetNillableAreaID sets the "area" edge to the Area entity by ID if the given value is not nil.
func (sc *ServiceCreate) SetNillableAreaID(id *int) *ServiceCreate {
	if id != nil {
		sc = sc.SetAreaID(*id)
	}
	return sc
}

// SetArea sets the "area" edge to the Area entity.
func (sc *ServiceCreate) SetArea(a *Area) *ServiceCreate {
	return sc.SetAreaID(a.ID)
}

// Mutation returns the ServiceMutation object of the builder.
func (sc *ServiceCreate) Mutation() *ServiceMutation {
	return sc.mutation
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	var (
		err  error
		node *Service
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Service)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ServiceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServiceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServiceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServiceCreate) defaults() {
	if _, ok := sc.mutation.Description(); !ok {
		v := service.DefaultDescription
		sc.mutation.SetDescription(v)
	}
	if _, ok := sc.mutation.Available(); !ok {
		v := service.DefaultAvailable
		sc.mutation.SetAvailable(v)
	}
	if _, ok := sc.mutation.Heartbeat(); !ok {
		v := service.DefaultHeartbeat
		sc.mutation.SetHeartbeat(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServiceCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Service.name"`)}
	}
	if _, ok := sc.mutation.Protocol(); !ok {
		return &ValidationError{Name: "protocol", err: errors.New(`ent: missing required field "Service.protocol"`)}
	}
	if _, ok := sc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "Service.host"`)}
	}
	if _, ok := sc.mutation.Port(); !ok {
		return &ValidationError{Name: "port", err: errors.New(`ent: missing required field "Service.port"`)}
	}
	if _, ok := sc.mutation.Available(); !ok {
		return &ValidationError{Name: "available", err: errors.New(`ent: missing required field "Service.available"`)}
	}
	if _, ok := sc.mutation.Heartbeat(); !ok {
		return &ValidationError{Name: "heartbeat", err: errors.New(`ent: missing required field "Service.heartbeat"`)}
	}
	if _, ok := sc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "Service.owner"`)}
	}
	return nil
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ServiceCreate) createSpec() (*Service, *sqlgraph.CreateSpec) {
	var (
		_node = &Service{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: service.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: service.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(service.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(service.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Protocol(); ok {
		_spec.SetField(service.FieldProtocol, field.TypeString, value)
		_node.Protocol = value
	}
	if value, ok := sc.mutation.Host(); ok {
		_spec.SetField(service.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := sc.mutation.Port(); ok {
		_spec.SetField(service.FieldPort, field.TypeString, value)
		_node.Port = value
	}
	if value, ok := sc.mutation.Available(); ok {
		_spec.SetField(service.FieldAvailable, field.TypeBool, value)
		_node.Available = value
	}
	if value, ok := sc.mutation.Heartbeat(); ok {
		_spec.SetField(service.FieldHeartbeat, field.TypeTime, value)
		_node.Heartbeat = value
	}
	if value, ok := sc.mutation.Owner(); ok {
		_spec.SetField(service.FieldOwner, field.TypeString, value)
		_node.Owner = value
	}
	if nodes := sc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   service.TagsTable,
			Columns: service.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AreaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.AreaTable,
			Columns: []string{service.AreaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: area.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.area_services = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceCreateBulk is the builder for creating many Service entities in bulk.
type ServiceCreateBulk struct {
	config
	builders []*ServiceCreate
}

// Save creates the Service entities in the database.
func (scb *ServiceCreateBulk) Save(ctx context.Context) ([]*Service, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Service, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServiceCreateBulk) SaveX(ctx context.Context) []*Service {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServiceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
