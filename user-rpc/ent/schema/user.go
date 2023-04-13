package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("姓名"),
		field.Enum("gender").Values("male", "female").Optional().Default("male").Comment("性别(male男,female女)"),
		field.String("avatar").Comment("头像"),
		field.String("username").Comment("用户名"),
		field.String("mobile").Comment("手机号"),
		field.String("id_card").Optional().Comment("身份证号"),
		field.String("openpid").Optional().Comment("openpid微信开放的用户标识"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("openpid", "id_card"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}
