package util

import "encoding/json"

func AnyToStruct[T any](val any) (*T, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	t := new(T)
	if err = json.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	return t, nil
}
