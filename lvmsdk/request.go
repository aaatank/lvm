package lvmsdk

import (
	"fmt"
	"reflect"
)

type DoRequest struct {
	Fn      string `json:"fn"`
	Content string `json:"content"`
}

type CallRequest[P any] struct {
	Fn       string `json:"fn"`
	Content  string `json:"content"`
	Function string `json:"function"`
	Params   P      `json:"params"`
}

func (cr *CallRequest[P]) check() error {
	v := reflect.ValueOf(cr.Params)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	kind := v.Kind()
	if kind != reflect.Struct && kind != reflect.Map {
		return fmt.Errorf("params must be struct or map, got %T", cr.Params)
	}
	return nil
}
