// Code generated by ent, DO NOT EDIT.

package ent

import (
	"drive/ent/email"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Email is the model entity for the Email schema.
type Email struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmailQuery when eager-loading is set.
	Edges EmailEdges `json:"edges"`
}

// EmailEdges holds the relations/edges for other nodes in the graph.
type EmailEdges struct {
	// Socials holds the value of the socials edge.
	Socials []*Social `json:"socials,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedSocials map[string][]*Social
}

// SocialsOrErr returns the Socials value or an error if the edge
// was not loaded in eager-loading.
func (e EmailEdges) SocialsOrErr() ([]*Social, error) {
	if e.loadedTypes[0] {
		return e.Socials, nil
	}
	return nil, &NotLoadedError{edge: "socials"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Email) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case email.FieldID, email.FieldCreatedBy, email.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case email.FieldName:
			values[i] = new(sql.NullString)
		case email.FieldCreatedAt, email.FieldUpdatedAt, email.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Email", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Email fields.
func (e *Email) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case email.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int64(value.Int64)
		case email.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				e.CreatedBy = value.Int64
			}
		case email.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				e.UpdatedBy = value.Int64
			}
		case email.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case email.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case email.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				e.DeletedAt = value.Time
			}
		case email.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		}
	}
	return nil
}

// QuerySocials queries the "socials" edge of the Email entity.
func (e *Email) QuerySocials() *SocialQuery {
	return (&EmailClient{config: e.config}).QuerySocials(e)
}

// Update returns a builder for updating this Email.
// Note that you need to call Email.Unwrap() before calling this method if this Email
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Email) Update() *EmailUpdateOne {
	return (&EmailClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Email entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Email) Unwrap() *Email {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Email is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Email) String() string {
	var builder strings.Builder
	builder.WriteString("Email(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", e.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", e.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(e.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSocials returns the Socials named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Email) NamedSocials(name string) ([]*Social, error) {
	if e.Edges.namedSocials == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedSocials[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Email) appendNamedSocials(name string, edges ...*Social) {
	if e.Edges.namedSocials == nil {
		e.Edges.namedSocials = make(map[string][]*Social)
	}
	if len(edges) == 0 {
		e.Edges.namedSocials[name] = []*Social{}
	} else {
		e.Edges.namedSocials[name] = append(e.Edges.namedSocials[name], edges...)
	}
}

// Emails is a parsable slice of Email.
type Emails []*Email

func (e Emails) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}