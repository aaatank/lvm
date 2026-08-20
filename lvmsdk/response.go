package lvmsdk

import (
	"fmt"
	"reflect"
)

type Response[D any] struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
	Data   D      `json:"data"`
}

func (r *Response[D]) check() error {
	v := reflect.ValueOf(r.Data)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	kind := v.Kind()
	if kind != reflect.Struct && kind != reflect.Map && kind != reflect.Slice && kind != reflect.Array {
		return fmt.Errorf("data must be struct, map, slice or array, got %T", r.Data)
	}
	return nil
}
