package util

import (
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
