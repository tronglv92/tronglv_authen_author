package util

import (
	"context"
	"github/tronglv_authen_author/helper/define"
	"strings"
)

func SortOrder(s string) string {
	s = strings.ToLower(s)
	if s == "" || (s != string(define.OrderAsc) && s != string(define.OrderDesc)) {
		return string(define.OrderAsc)
	}
	return s
}

func ServiceName(ctx context.Context) string {
	if name, ok := ctx.Value(define.ServiceNameContextKey).(string); ok {
		return name
	}
	return define.EmptyString
}
