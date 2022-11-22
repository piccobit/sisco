// Code generated by ent, DO NOT EDIT.

package service

import (
	"time"
)

const (
	// Label holds the string label denoting the service type in the database.
	Label = "service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldProtocol holds the string denoting the protocol field in the database.
	FieldProtocol = "protocol"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldAvailable holds the string denoting the available field in the database.
	FieldAvailable = "available"
	// FieldHeartbeat holds the string denoting the heartbeat field in the database.
	FieldHeartbeat = "heartbeat"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeArea holds the string denoting the area edge name in mutations.
	EdgeArea = "area"
	// Table holds the table name of the service in the database.
	Table = "services"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "service_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// AreaTable is the table that holds the area relation/edge.
	AreaTable = "services"
	// AreaInverseTable is the table name for the Area entity.
	// It exists in this package in order to avoid circular dependency with the "area" package.
	AreaInverseTable = "areas"
	// AreaColumn is the table column denoting the area relation/edge.
	AreaColumn = "area_services"
)

// Columns holds all SQL columns for service fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldProtocol,
	FieldHost,
	FieldPort,
	FieldAvailable,
	FieldHeartbeat,
	FieldOwner,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "services"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"area_services",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"service_id", "tag_id"}
)

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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultAvailable holds the default value on creation for the "available" field.
	DefaultAvailable bool
	// DefaultHeartbeat holds the default value on creation for the "heartbeat" field.
	DefaultHeartbeat time.Time
)
