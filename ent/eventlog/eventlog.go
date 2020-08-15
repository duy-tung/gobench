// Code generated by entc, DO NOT EDIT.

package eventlog

import (
	"time"
)

const (
	// Label holds the string label denoting the eventlog type in the database.
	Label = "event_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID        = "id"      // FieldName holds the string denoting the name vertex property in the database.
	FieldName      = "name"    // FieldMessage holds the string denoting the message vertex property in the database.
	FieldMessage   = "message" // FieldLevel holds the string denoting the level vertex property in the database.
	FieldLevel     = "level"   // FieldSource holds the string denoting the source vertex property in the database.
	FieldSource    = "source"  // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"

	// EdgeApplications holds the string denoting the applications edge name in mutations.
	EdgeApplications = "applications"

	// Table holds the table name of the eventlog in the database.
	Table = "event_logs"
	// ApplicationsTable is the table the holds the applications relation/edge.
	ApplicationsTable = "event_logs"
	// ApplicationsInverseTable is the table name for the Application entity.
	// It exists in this package in order to avoid circular dependency with the "application" package.
	ApplicationsInverseTable = "applications"
	// ApplicationsColumn is the table column denoting the applications relation/edge.
	ApplicationsColumn = "application_event_logs"
)

// Columns holds all SQL columns for eventlog fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMessage,
	FieldLevel,
	FieldSource,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the EventLog type.
var ForeignKeys = []string{
	"application_event_logs",
}

var (
	// DefaultName holds the default value on creation for the name field.
	DefaultName string
	// DefaultLevel holds the default value on creation for the level field.
	DefaultLevel string
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
)
