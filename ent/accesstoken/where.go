// Code generated by entc, DO NOT EDIT.

package accesstoken

import (
	"time"
	"wynn-member-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LineToken applies equality check predicate on the "line_token" field. It's identical to LineTokenEQ.
func LineToken(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLineToken), v))
	})
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// TokenExpire applies equality check predicate on the "token_expire" field. It's identical to TokenExpireEQ.
func TokenExpire(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenExpire), vc))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), vc))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), vc))
	})
}

// LineTokenEQ applies the EQ predicate on the "line_token" field.
func LineTokenEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLineToken), v))
	})
}

// LineTokenNEQ applies the NEQ predicate on the "line_token" field.
func LineTokenNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLineToken), v))
	})
}

// LineTokenIn applies the In predicate on the "line_token" field.
func LineTokenIn(vs ...string) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLineToken), v...))
	})
}

// LineTokenNotIn applies the NotIn predicate on the "line_token" field.
func LineTokenNotIn(vs ...string) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLineToken), v...))
	})
}

// LineTokenGT applies the GT predicate on the "line_token" field.
func LineTokenGT(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLineToken), v))
	})
}

// LineTokenGTE applies the GTE predicate on the "line_token" field.
func LineTokenGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLineToken), v))
	})
}

// LineTokenLT applies the LT predicate on the "line_token" field.
func LineTokenLT(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLineToken), v))
	})
}

// LineTokenLTE applies the LTE predicate on the "line_token" field.
func LineTokenLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLineToken), v))
	})
}

// LineTokenContains applies the Contains predicate on the "line_token" field.
func LineTokenContains(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLineToken), v))
	})
}

// LineTokenHasPrefix applies the HasPrefix predicate on the "line_token" field.
func LineTokenHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLineToken), v))
	})
}

// LineTokenHasSuffix applies the HasSuffix predicate on the "line_token" field.
func LineTokenHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLineToken), v))
	})
}

// LineTokenIsNil applies the IsNil predicate on the "line_token" field.
func LineTokenIsNil() predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLineToken)))
	})
}

// LineTokenNotNil applies the NotNil predicate on the "line_token" field.
func LineTokenNotNil() predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLineToken)))
	})
}

// LineTokenEqualFold applies the EqualFold predicate on the "line_token" field.
func LineTokenEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLineToken), v))
	})
}

// LineTokenContainsFold applies the ContainsFold predicate on the "line_token" field.
func LineTokenContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLineToken), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// TokenExpireEQ applies the EQ predicate on the "token_expire" field.
func TokenExpireEQ(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenExpire), vc))
	})
}

// TokenExpireNEQ applies the NEQ predicate on the "token_expire" field.
func TokenExpireNEQ(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenExpire), vc))
	})
}

// TokenExpireIn applies the In predicate on the "token_expire" field.
func TokenExpireIn(vs ...time.Time) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTokenExpire), v...))
	})
}

// TokenExpireNotIn applies the NotIn predicate on the "token_expire" field.
func TokenExpireNotIn(vs ...time.Time) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTokenExpire), v...))
	})
}

// TokenExpireGT applies the GT predicate on the "token_expire" field.
func TokenExpireGT(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenExpire), vc))
	})
}

// TokenExpireGTE applies the GTE predicate on the "token_expire" field.
func TokenExpireGTE(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenExpire), vc))
	})
}

// TokenExpireLT applies the LT predicate on the "token_expire" field.
func TokenExpireLT(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenExpire), vc))
	})
}

// TokenExpireLTE applies the LTE predicate on the "token_expire" field.
func TokenExpireLTE(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenExpire), vc))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPIsNil applies the IsNil predicate on the "ip" field.
func IPIsNil() predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIP)))
	})
}

// IPNotNil applies the NotNil predicate on the "ip" field.
func IPNotNil() predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIP)))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), vc))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), vc))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), vc))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), vc))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), vc))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), vc))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), vc))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), vc))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccessToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.AccessToken(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), vc))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), vc))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), vc))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccessToken {
	vc := time.Time(v)
	return predicate.AccessToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), vc))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
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
func Not(p predicate.AccessToken) predicate.AccessToken {
	return predicate.AccessToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
