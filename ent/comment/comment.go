// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCommentID holds the string denoting the comment_id field in the database.
	FieldCommentID = "comment_id"
	// FieldUpdateDatetime holds the string denoting the update_datetime field in the database.
	FieldUpdateDatetime = "update_datetime"
	// FieldCreateDatetime holds the string denoting the create_datetime field in the database.
	FieldCreateDatetime = "create_datetime"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "comments"
	// OwnerInverseTable is the table name for the Blog entity.
	// It exists in this package in order to avoid circular dependency with the "blog" package.
	OwnerInverseTable = "blogs"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "blog_comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPassword,
	FieldContent,
	FieldCommentID,
	FieldUpdateDatetime,
	FieldCreateDatetime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"blog_comments",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCommentID holds the default value on creation for the "comment_id" field.
	DefaultCommentID int64
	// DefaultUpdateDatetime holds the default value on creation for the "update_datetime" field.
	DefaultUpdateDatetime func() time.Time
	// UpdateDefaultUpdateDatetime holds the default value on update for the "update_datetime" field.
	UpdateDefaultUpdateDatetime func() time.Time
	// DefaultCreateDatetime holds the default value on creation for the "create_datetime" field.
	DefaultCreateDatetime func() time.Time
)
