// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/mnsv1511/kasikorn-line-assignment/ent/accountflags"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountflagsFields := schema.AccountFlags{}.Fields()
	_ = accountflagsFields
	// accountflagsDescCreatedAt is the schema descriptor for created_at field.
	accountflagsDescCreatedAt := accountflagsFields[4].Descriptor()
	// accountflags.DefaultCreatedAt holds the default value on creation for the created_at field.
	accountflags.DefaultCreatedAt = accountflagsDescCreatedAt.Default.(func() time.Time)
	// accountflagsDescUpdatedAt is the schema descriptor for updated_at field.
	accountflagsDescUpdatedAt := accountflagsFields[5].Descriptor()
	// accountflags.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	accountflags.DefaultUpdatedAt = accountflagsDescUpdatedAt.Default.(func() time.Time)
	// accountflags.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	accountflags.UpdateDefaultUpdatedAt = accountflagsDescUpdatedAt.UpdateDefault.(func() time.Time)
}
