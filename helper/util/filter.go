package util

import "github/tronglv_authen_author/helper/define"

func Filter[T any](i map[string]interface{}, key string) (T, bool) {
	var result T
	if val, ok := i[key]; ok {
		if i, ok := val.(T); ok {
			return i, true
		}
	}
	return result, false
}

func FilterNotEmpty(i map[string]interface{}, key string) (string, bool) {
	if val, ok := Filter[string](i, key); ok {
		if len(val) > 0 {
			return val, ok
		}
	}
	return define.EmptyString, false
}

func FilterNotZero(i map[string]interface{}, key string) (int, bool) {
	if val, ok := Filter[int](i, key); ok {
		if val > 0 {
			return val, ok
		}
	}
	return 0, false
}
