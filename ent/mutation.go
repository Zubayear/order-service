// Code generated by entc, DO NOT EDIT.

package ent

import (
	"OrderService/ent/order"
	"OrderService/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOrder = "Order"
)

// OrderMutation represents an operation that mutates the Order nodes in the graph.
type OrderMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	user_id        *uuid.UUID
	order_total    *float64
	addorder_total *float64
	order_placed   *time.Time
	order_paid     *bool
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Order, error)
	predicates     []predicate.Order
}

var _ ent.Mutation = (*OrderMutation)(nil)

// orderOption allows management of the mutation configuration using functional options.
type orderOption func(*OrderMutation)

// newOrderMutation creates new mutation for the Order entity.
func newOrderMutation(c config, op Op, opts ...orderOption) *OrderMutation {
	m := &OrderMutation{
		config:        c,
		op:            op,
		typ:           TypeOrder,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOrderID sets the ID field of the mutation.
func withOrderID(id uuid.UUID) orderOption {
	return func(m *OrderMutation) {
		var (
			err   error
			once  sync.Once
			value *Order
		)
		m.oldValue = func(ctx context.Context) (*Order, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Order.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOrder sets the old Order of the mutation.
func withOrder(node *Order) orderOption {
	return func(m *OrderMutation) {
		m.oldValue = func(context.Context) (*Order, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OrderMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OrderMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Order entities.
func (m *OrderMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OrderMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OrderMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Order.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUserID sets the "user_id" field.
func (m *OrderMutation) SetUserID(u uuid.UUID) {
	m.user_id = &u
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *OrderMutation) UserID() (r uuid.UUID, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldUserID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *OrderMutation) ResetUserID() {
	m.user_id = nil
}

// SetOrderTotal sets the "order_total" field.
func (m *OrderMutation) SetOrderTotal(f float64) {
	m.order_total = &f
	m.addorder_total = nil
}

// OrderTotal returns the value of the "order_total" field in the mutation.
func (m *OrderMutation) OrderTotal() (r float64, exists bool) {
	v := m.order_total
	if v == nil {
		return
	}
	return *v, true
}

// OldOrderTotal returns the old "order_total" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldOrderTotal(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOrderTotal is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOrderTotal requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOrderTotal: %w", err)
	}
	return oldValue.OrderTotal, nil
}

// AddOrderTotal adds f to the "order_total" field.
func (m *OrderMutation) AddOrderTotal(f float64) {
	if m.addorder_total != nil {
		*m.addorder_total += f
	} else {
		m.addorder_total = &f
	}
}

// AddedOrderTotal returns the value that was added to the "order_total" field in this mutation.
func (m *OrderMutation) AddedOrderTotal() (r float64, exists bool) {
	v := m.addorder_total
	if v == nil {
		return
	}
	return *v, true
}

// ResetOrderTotal resets all changes to the "order_total" field.
func (m *OrderMutation) ResetOrderTotal() {
	m.order_total = nil
	m.addorder_total = nil
}

// SetOrderPlaced sets the "order_placed" field.
func (m *OrderMutation) SetOrderPlaced(t time.Time) {
	m.order_placed = &t
}

// OrderPlaced returns the value of the "order_placed" field in the mutation.
func (m *OrderMutation) OrderPlaced() (r time.Time, exists bool) {
	v := m.order_placed
	if v == nil {
		return
	}
	return *v, true
}

// OldOrderPlaced returns the old "order_placed" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldOrderPlaced(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOrderPlaced is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOrderPlaced requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOrderPlaced: %w", err)
	}
	return oldValue.OrderPlaced, nil
}

// ResetOrderPlaced resets all changes to the "order_placed" field.
func (m *OrderMutation) ResetOrderPlaced() {
	m.order_placed = nil
}

// SetOrderPaid sets the "order_paid" field.
func (m *OrderMutation) SetOrderPaid(b bool) {
	m.order_paid = &b
}

// OrderPaid returns the value of the "order_paid" field in the mutation.
func (m *OrderMutation) OrderPaid() (r bool, exists bool) {
	v := m.order_paid
	if v == nil {
		return
	}
	return *v, true
}

// OldOrderPaid returns the old "order_paid" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldOrderPaid(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOrderPaid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOrderPaid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOrderPaid: %w", err)
	}
	return oldValue.OrderPaid, nil
}

// ResetOrderPaid resets all changes to the "order_paid" field.
func (m *OrderMutation) ResetOrderPaid() {
	m.order_paid = nil
}

// Where appends a list predicates to the OrderMutation builder.
func (m *OrderMutation) Where(ps ...predicate.Order) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *OrderMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Order).
func (m *OrderMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OrderMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.user_id != nil {
		fields = append(fields, order.FieldUserID)
	}
	if m.order_total != nil {
		fields = append(fields, order.FieldOrderTotal)
	}
	if m.order_placed != nil {
		fields = append(fields, order.FieldOrderPlaced)
	}
	if m.order_paid != nil {
		fields = append(fields, order.FieldOrderPaid)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OrderMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case order.FieldUserID:
		return m.UserID()
	case order.FieldOrderTotal:
		return m.OrderTotal()
	case order.FieldOrderPlaced:
		return m.OrderPlaced()
	case order.FieldOrderPaid:
		return m.OrderPaid()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OrderMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case order.FieldUserID:
		return m.OldUserID(ctx)
	case order.FieldOrderTotal:
		return m.OldOrderTotal(ctx)
	case order.FieldOrderPlaced:
		return m.OldOrderPlaced(ctx)
	case order.FieldOrderPaid:
		return m.OldOrderPaid(ctx)
	}
	return nil, fmt.Errorf("unknown Order field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) SetField(name string, value ent.Value) error {
	switch name {
	case order.FieldUserID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case order.FieldOrderTotal:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOrderTotal(v)
		return nil
	case order.FieldOrderPlaced:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOrderPlaced(v)
		return nil
	case order.FieldOrderPaid:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOrderPaid(v)
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OrderMutation) AddedFields() []string {
	var fields []string
	if m.addorder_total != nil {
		fields = append(fields, order.FieldOrderTotal)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OrderMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case order.FieldOrderTotal:
		return m.AddedOrderTotal()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) AddField(name string, value ent.Value) error {
	switch name {
	case order.FieldOrderTotal:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddOrderTotal(v)
		return nil
	}
	return fmt.Errorf("unknown Order numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OrderMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OrderMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OrderMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Order nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OrderMutation) ResetField(name string) error {
	switch name {
	case order.FieldUserID:
		m.ResetUserID()
		return nil
	case order.FieldOrderTotal:
		m.ResetOrderTotal()
		return nil
	case order.FieldOrderPlaced:
		m.ResetOrderPlaced()
		return nil
	case order.FieldOrderPaid:
		m.ResetOrderPaid()
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OrderMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OrderMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OrderMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OrderMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OrderMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OrderMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OrderMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Order unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OrderMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Order edge %s", name)
}