package utils

import (
	"context"
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"log"
)

type M map[string]interface{}

func ToJson(v interface{}) string {
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Cannot marshal %v", v)
	}
	return string(marshalled)
}

func ToRawJson(v interface{}) []byte {
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Cannot marshal %v", v)
	}
	return marshalled
}

func StrCmp(v1 interface{}, v2 interface{}) bool {
	return v1.(string) == v2.(string)
}

func GetRoleFromContext(ctx context.Context) roles.Role {
	return ctx.Value(contextKeys.UserRoleKey).(roles.Role)
}

func GetUserIdFromContext(ctx context.Context) users.UserId {
	return ctx.Value(contextKeys.UserIdKey).(users.UserId)
}
