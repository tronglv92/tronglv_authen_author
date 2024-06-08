package util

import (
	"encoding/json"
	"fmt"
)

func Marshal(val any) string {
	v, _ := json.Marshal(val)
	return string(v)
}

func MapToStruct[T any](val any) (*T, error) {
	j, err := json.Marshal(val)
	if err != nil {
		return nil, fmt.Errorf("[err] marshal any to JSON: %v", err)
	}

	var result = new(T)
	if err = json.Unmarshal(j, result); err != nil {
		return nil, fmt.Errorf("[err] unmarshalling JSON to struct: %v", err)
	}
	return result, nil
}
