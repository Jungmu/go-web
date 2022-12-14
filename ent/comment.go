// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/jungmu/go-web/ent/blog"
	"github.com/jungmu/go-web/ent/comment"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// name
	Name string `json:"name,omitempty"`
	// password
	Password string `json:"password,omitempty"`
	// content
	Content string `json:"content,omitempty"`
	// comment_id
	CommentID int64 `json:"comment_id,omitempty"`
	// update_datetime
	UpdateDatetime time.Time `json:"update_datetime,omitempty"`
	// update_datetime
	CreateDatetime time.Time `json:"create_datetime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges         CommentEdges `json:"edges"`
	blog_comments *int64
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Blog `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) OwnerOrErr() (*Blog, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: blog.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldCommentID:
			values[i] = new(sql.NullInt64)
		case comment.FieldName, comment.FieldPassword, comment.FieldContent:
			values[i] = new(sql.NullString)
		case comment.FieldUpdateDatetime, comment.FieldCreateDatetime:
			values[i] = new(sql.NullTime)
		case comment.ForeignKeys[0]: // blog_comments
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case comment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case comment.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				c.Password = value.String
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.FieldCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_id", values[i])
			} else if value.Valid {
				c.CommentID = value.Int64
			}
		case comment.FieldUpdateDatetime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_datetime", values[i])
			} else if value.Valid {
				c.UpdateDatetime = value.Time
			}
		case comment.FieldCreateDatetime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_datetime", values[i])
			} else if value.Valid {
				c.CreateDatetime = value.Time
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field blog_comments", value)
			} else if value.Valid {
				c.blog_comments = new(int64)
				*c.blog_comments = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Comment entity.
func (c *Comment) QueryOwner() *BlogQuery {
	return (&CommentClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(c.Password)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(c.Content)
	builder.WriteString(", ")
	builder.WriteString("comment_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CommentID))
	builder.WriteString(", ")
	builder.WriteString("update_datetime=")
	builder.WriteString(c.UpdateDatetime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("create_datetime=")
	builder.WriteString(c.CreateDatetime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
