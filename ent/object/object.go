// Code generated by ent, DO NOT EDIT.

package object

const (
	// Label holds the string label denoting the object type in the database.
	Label = "object"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the object in the database.
	Table = "objects"
)

// Columns holds all SQL columns for object fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
