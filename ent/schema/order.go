package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.Float("order_total").SchemaType(map[string]string{
			dialect.MySQL: "decimal(6,2)",
		}),
		field.Time("order_placed").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("order_paid").Default(false),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
