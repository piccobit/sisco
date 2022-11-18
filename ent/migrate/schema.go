// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AreasColumns holds the columns for the "areas" table.
	AreasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
	}
	// AreasTable holds the schema information for the "areas" table.
	AreasTable = &schema.Table{
		Name:       "areas",
		Columns:    AreasColumns,
		PrimaryKey: []*schema.Column{AreasColumns[0]},
	}
	// ServicesColumns holds the columns for the "services" table.
	ServicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "protocol", Type: field.TypeString},
		{Name: "host", Type: field.TypeString},
		{Name: "port", Type: field.TypeString},
		{Name: "available", Type: field.TypeBool, Default: false},
		{Name: "heartbeat", Type: field.TypeTime},
		{Name: "area_services", Type: field.TypeInt, Nullable: true},
	}
	// ServicesTable holds the schema information for the "services" table.
	ServicesTable = &schema.Table{
		Name:       "services",
		Columns:    ServicesColumns,
		PrimaryKey: []*schema.Column{ServicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "services_areas_services",
				Columns:    []*schema.Column{ServicesColumns[8]},
				RefColumns: []*schema.Column{AreasColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeString, Unique: true},
		{Name: "token", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// ServiceTagsColumns holds the columns for the "service_tags" table.
	ServiceTagsColumns = []*schema.Column{
		{Name: "service_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// ServiceTagsTable holds the schema information for the "service_tags" table.
	ServiceTagsTable = &schema.Table{
		Name:       "service_tags",
		Columns:    ServiceTagsColumns,
		PrimaryKey: []*schema.Column{ServiceTagsColumns[0], ServiceTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "service_tags_service_id",
				Columns:    []*schema.Column{ServiceTagsColumns[0]},
				RefColumns: []*schema.Column{ServicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "service_tags_tag_id",
				Columns:    []*schema.Column{ServiceTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AreasTable,
		ServicesTable,
		TagsTable,
		TokensTable,
		ServiceTagsTable,
	}
)

func init() {
	ServicesTable.ForeignKeys[0].RefTable = AreasTable
	ServiceTagsTable.ForeignKeys[0].RefTable = ServicesTable
	ServiceTagsTable.ForeignKeys[1].RefTable = TagsTable
}
