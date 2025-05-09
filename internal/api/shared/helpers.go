package shared

import (
	"context"
	"database/sql"
	"strconv"
	"time"
)

func NullStringPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func NullTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func StringToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func TimeToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDContextKey).(string)
	return userID, ok
}

func GetOrgID(ctx context.Context) (string, bool) {
	orgID, ok := ctx.Value(OrgIDContextKey).(string)
	return orgID, ok
}

func ConvertOrgIDToInt(o string) (int64, error) {
	return strconv.ParseInt(o, 10, 64)
}
