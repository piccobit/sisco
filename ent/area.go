// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sisco/ent/area"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Area is the model entity for the Area schema.
type Area struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AreaQuery when eager-loading is set.
	Edges AreaEdges `json:"edges"`
}

// AreaEdges holds the relations/edges for other nodes in the graph.
type AreaEdges struct {
	// Services holds the value of the services edge.
	Services []*Service `json:"services,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ServicesOrErr returns the Services value or an error if the edge
// was not loaded in eager-loading.
func (e AreaEdges) ServicesOrErr() ([]*Service, error) {
	if e.loadedTypes[0] {
		return e.Services, nil
	}
	return nil, &NotLoadedError{edge: "services"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Area) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case area.FieldID:
			values[i] = new(sql.NullInt64)
		case area.FieldName, area.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Area", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Area fields.
func (a *Area) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case area.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case area.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case area.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		}
	}
	return nil
}

// QueryServices queries the "services" edge of the Area entity.
func (a *Area) QueryServices() *ServiceQuery {
	return (&AreaClient{config: a.config}).QueryServices(a)
}

// Update returns a builder for updating this Area.
// Note that you need to call Area.Unwrap() before calling this method if this Area
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Area) Update() *AreaUpdateOne {
	return (&AreaClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Area entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Area) Unwrap() *Area {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Area is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Area) String() string {
	var builder strings.Builder
	builder.WriteString("Area(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Areas is a parsable slice of Area.
type Areas []*Area

func (a Areas) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
