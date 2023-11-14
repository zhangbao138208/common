// Code generated by ent, DO NOT EDIT.

package winconfig

import (
	"entgo.io/ent/dialect/sql"
	"gitlab.skig.tech/zero-core/common/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldTitle, v))
}

// TitleZh applies equality check predicate on the "title_zh" field. It's identical to TitleZhEQ.
func TitleZh(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldTitleZh, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldValue, v))
}

// ShowApp applies equality check predicate on the "show_app" field. It's identical to ShowAppEQ.
func ShowApp(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldShowApp, v))
}

// CanModify applies equality check predicate on the "can_modify" field. It's identical to CanModifyEQ.
func CanModify(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldCanModify, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldContainsFold(FieldTitle, v))
}

// TitleZhEQ applies the EQ predicate on the "title_zh" field.
func TitleZhEQ(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldTitleZh, v))
}

// TitleZhNEQ applies the NEQ predicate on the "title_zh" field.
func TitleZhNEQ(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldTitleZh, v))
}

// TitleZhIn applies the In predicate on the "title_zh" field.
func TitleZhIn(vs ...string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldTitleZh, vs...))
}

// TitleZhNotIn applies the NotIn predicate on the "title_zh" field.
func TitleZhNotIn(vs ...string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldTitleZh, vs...))
}

// TitleZhGT applies the GT predicate on the "title_zh" field.
func TitleZhGT(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldTitleZh, v))
}

// TitleZhGTE applies the GTE predicate on the "title_zh" field.
func TitleZhGTE(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldTitleZh, v))
}

// TitleZhLT applies the LT predicate on the "title_zh" field.
func TitleZhLT(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldTitleZh, v))
}

// TitleZhLTE applies the LTE predicate on the "title_zh" field.
func TitleZhLTE(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldTitleZh, v))
}

// TitleZhContains applies the Contains predicate on the "title_zh" field.
func TitleZhContains(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldContains(FieldTitleZh, v))
}

// TitleZhHasPrefix applies the HasPrefix predicate on the "title_zh" field.
func TitleZhHasPrefix(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldHasPrefix(FieldTitleZh, v))
}

// TitleZhHasSuffix applies the HasSuffix predicate on the "title_zh" field.
func TitleZhHasSuffix(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldHasSuffix(FieldTitleZh, v))
}

// TitleZhEqualFold applies the EqualFold predicate on the "title_zh" field.
func TitleZhEqualFold(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEqualFold(FieldTitleZh, v))
}

// TitleZhContainsFold applies the ContainsFold predicate on the "title_zh" field.
func TitleZhContainsFold(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldContainsFold(FieldTitleZh, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldContainsFold(FieldValue, v))
}

// ShowAppEQ applies the EQ predicate on the "show_app" field.
func ShowAppEQ(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldShowApp, v))
}

// ShowAppNEQ applies the NEQ predicate on the "show_app" field.
func ShowAppNEQ(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldShowApp, v))
}

// ShowAppIn applies the In predicate on the "show_app" field.
func ShowAppIn(vs ...int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldShowApp, vs...))
}

// ShowAppNotIn applies the NotIn predicate on the "show_app" field.
func ShowAppNotIn(vs ...int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldShowApp, vs...))
}

// ShowAppGT applies the GT predicate on the "show_app" field.
func ShowAppGT(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldShowApp, v))
}

// ShowAppGTE applies the GTE predicate on the "show_app" field.
func ShowAppGTE(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldShowApp, v))
}

// ShowAppLT applies the LT predicate on the "show_app" field.
func ShowAppLT(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldShowApp, v))
}

// ShowAppLTE applies the LTE predicate on the "show_app" field.
func ShowAppLTE(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldShowApp, v))
}

// CanModifyEQ applies the EQ predicate on the "can_modify" field.
func CanModifyEQ(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldCanModify, v))
}

// CanModifyNEQ applies the NEQ predicate on the "can_modify" field.
func CanModifyNEQ(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldCanModify, v))
}

// CanModifyIn applies the In predicate on the "can_modify" field.
func CanModifyIn(vs ...int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldCanModify, vs...))
}

// CanModifyNotIn applies the NotIn predicate on the "can_modify" field.
func CanModifyNotIn(vs ...int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldCanModify, vs...))
}

// CanModifyGT applies the GT predicate on the "can_modify" field.
func CanModifyGT(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldCanModify, v))
}

// CanModifyGTE applies the GTE predicate on the "can_modify" field.
func CanModifyGTE(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldCanModify, v))
}

// CanModifyLT applies the LT predicate on the "can_modify" field.
func CanModifyLT(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldCanModify, v))
}

// CanModifyLTE applies the LTE predicate on the "can_modify" field.
func CanModifyLTE(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldCanModify, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int32) predicate.WinConfig {
	return predicate.WinConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WinConfig) predicate.WinConfig {
	return predicate.WinConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WinConfig) predicate.WinConfig {
	return predicate.WinConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WinConfig) predicate.WinConfig {
	return predicate.WinConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}