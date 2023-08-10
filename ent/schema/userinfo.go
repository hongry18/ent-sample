package schema

import "entgo.io/ent"

// UserInfo holds the schema definition for the UserInfo entity.
type UserInfo struct {
	ent.Schema
}

// Fields of the UserInfo.
func (UserInfo) Fields() []ent.Field {
	return nil
}

// Edges of the UserInfo.
func (UserInfo) Edges() []ent.Edge {
	return nil
}
