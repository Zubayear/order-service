// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "order_total", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(6,2)"}},
		{Name: "order_placed", Type: field.TypeTime},
		{Name: "order_paid", Type: field.TypeBool, Default: false},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
	}
)

func init() {
}
