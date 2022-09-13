// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/jungmu/go-web/ent/blog"
	"github.com/jungmu/go-web/ent/bloglog"
	"github.com/jungmu/go-web/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blogFields := schema.Blog{}.Fields()
	_ = blogFields
	// blogDescSubTitle is the schema descriptor for sub_title field.
	blogDescSubTitle := blogFields[2].Descriptor()
	// blog.DefaultSubTitle holds the default value on creation for the sub_title field.
	blog.DefaultSubTitle = blogDescSubTitle.Default.(string)
	// blogDescTags is the schema descriptor for tags field.
	blogDescTags := blogFields[3].Descriptor()
	// blog.DefaultTags holds the default value on creation for the tags field.
	blog.DefaultTags = blogDescTags.Default.(string)
	// blogDescUpdateDatetime is the schema descriptor for update_datetime field.
	blogDescUpdateDatetime := blogFields[5].Descriptor()
	// blog.DefaultUpdateDatetime holds the default value on creation for the update_datetime field.
	blog.DefaultUpdateDatetime = blogDescUpdateDatetime.Default.(func() time.Time)
	// blog.UpdateDefaultUpdateDatetime holds the default value on update for the update_datetime field.
	blog.UpdateDefaultUpdateDatetime = blogDescUpdateDatetime.UpdateDefault.(func() time.Time)
	// blogDescCreateDatetime is the schema descriptor for create_datetime field.
	blogDescCreateDatetime := blogFields[6].Descriptor()
	// blog.DefaultCreateDatetime holds the default value on creation for the create_datetime field.
	blog.DefaultCreateDatetime = blogDescCreateDatetime.Default.(func() time.Time)
	bloglogFields := schema.BlogLog{}.Fields()
	_ = bloglogFields
	// bloglogDescCreateDatetime is the schema descriptor for create_datetime field.
	bloglogDescCreateDatetime := bloglogFields[6].Descriptor()
	// bloglog.DefaultCreateDatetime holds the default value on creation for the create_datetime field.
	bloglog.DefaultCreateDatetime = bloglogDescCreateDatetime.Default.(func() time.Time)
}
