// Code generated by ent, DO NOT EDIT.

package ent

import (
	"drive/ent/email"
	"drive/ent/social"
	"drive/ent/wechat"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Social is the model entity for the Social schema.
type Social struct {
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
	// RelationID holds the value of the "relation_id" field.
	RelationID int64 `json:"relation_id,omitempty"`
	// Type holds the value of the "type" field.
	Type int32 `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SocialQuery when eager-loading is set.
	Edges SocialEdges `json:"edges"`
}

// SocialEdges holds the relations/edges for other nodes in the graph.
type SocialEdges struct {
	// Email holds the value of the email edge.
	Email *Email `json:"email,omitempty"`
	// Wechat holds the value of the wechat edge.
	Wechat *Wechat `json:"wechat,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// EmailOrErr returns the Email value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SocialEdges) EmailOrErr() (*Email, error) {
	if e.loadedTypes[0] {
		if e.Email == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: email.Label}
		}
		return e.Email, nil
	}
	return nil, &NotLoadedError{edge: "email"}
}

// WechatOrErr returns the Wechat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SocialEdges) WechatOrErr() (*Wechat, error) {
	if e.loadedTypes[1] {
		if e.Wechat == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: wechat.Label}
		}
		return e.Wechat, nil
	}
	return nil, &NotLoadedError{edge: "wechat"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Social) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case social.FieldID, social.FieldCreatedBy, social.FieldUpdatedBy, social.FieldRelationID, social.FieldType:
			values[i] = new(sql.NullInt64)
		case social.FieldName:
			values[i] = new(sql.NullString)
		case social.FieldCreatedAt, social.FieldUpdatedAt, social.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Social", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Social fields.
func (s *Social) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case social.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case social.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				s.CreatedBy = value.Int64
			}
		case social.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				s.UpdatedBy = value.Int64
			}
		case social.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case social.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case social.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case social.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case social.FieldRelationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relation_id", values[i])
			} else if value.Valid {
				s.RelationID = value.Int64
			}
		case social.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryEmail queries the "email" edge of the Social entity.
func (s *Social) QueryEmail() *EmailQuery {
	return (&SocialClient{config: s.config}).QueryEmail(s)
}

// QueryWechat queries the "wechat" edge of the Social entity.
func (s *Social) QueryWechat() *WechatQuery {
	return (&SocialClient{config: s.config}).QueryWechat(s)
}

// Update returns a builder for updating this Social.
// Note that you need to call Social.Unwrap() before calling this method if this Social
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Social) Update() *SocialUpdateOne {
	return (&SocialClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Social entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Social) Unwrap() *Social {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Social is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Social) String() string {
	var builder strings.Builder
	builder.WriteString("Social(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("relation_id=")
	builder.WriteString(fmt.Sprintf("%v", s.RelationID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Socials is a parsable slice of Social.
type Socials []*Social

func (s Socials) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}