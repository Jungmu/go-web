// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/jungmu/go-web/ent/bloglog"
)

// BlogLog is the model entity for the BlogLog schema.
type BlogLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// blog_id
	BlogID int64 `json:"blog_id,omitempty"`
	// request url
	URL string `json:"url,omitempty"`
	// reason
	Reason string `json:"reason,omitempty"`
	// detail
	Detail string `json:"detail,omitempty"`
	// client_ip
	ClientIP string `json:"client_ip,omitempty"`
	// update_datetime
	CreateDatetime time.Time `json:"create_datetime,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlogLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case bloglog.FieldID, bloglog.FieldBlogID:
			values[i] = new(sql.NullInt64)
		case bloglog.FieldURL, bloglog.FieldReason, bloglog.FieldDetail, bloglog.FieldClientIP:
			values[i] = new(sql.NullString)
		case bloglog.FieldCreateDatetime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BlogLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlogLog fields.
func (bl *BlogLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bloglog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bl.ID = int64(value.Int64)
		case bloglog.FieldBlogID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field blog_id", values[i])
			} else if value.Valid {
				bl.BlogID = value.Int64
			}
		case bloglog.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				bl.URL = value.String
			}
		case bloglog.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				bl.Reason = value.String
			}
		case bloglog.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				bl.Detail = value.String
			}
		case bloglog.FieldClientIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_ip", values[i])
			} else if value.Valid {
				bl.ClientIP = value.String
			}
		case bloglog.FieldCreateDatetime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_datetime", values[i])
			} else if value.Valid {
				bl.CreateDatetime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this BlogLog.
// Note that you need to call BlogLog.Unwrap() before calling this method if this BlogLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (bl *BlogLog) Update() *BlogLogUpdateOne {
	return (&BlogLogClient{config: bl.config}).UpdateOne(bl)
}

// Unwrap unwraps the BlogLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bl *BlogLog) Unwrap() *BlogLog {
	_tx, ok := bl.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlogLog is not a transactional entity")
	}
	bl.config.driver = _tx.drv
	return bl
}

// String implements the fmt.Stringer.
func (bl *BlogLog) String() string {
	var builder strings.Builder
	builder.WriteString("BlogLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bl.ID))
	builder.WriteString("blog_id=")
	builder.WriteString(fmt.Sprintf("%v", bl.BlogID))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(bl.URL)
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(bl.Reason)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(bl.Detail)
	builder.WriteString(", ")
	builder.WriteString("client_ip=")
	builder.WriteString(bl.ClientIP)
	builder.WriteString(", ")
	builder.WriteString("create_datetime=")
	builder.WriteString(bl.CreateDatetime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BlogLogs is a parsable slice of BlogLog.
type BlogLogs []*BlogLog

func (bl BlogLogs) config(cfg config) {
	for _i := range bl {
		bl[_i].config = cfg
	}
}