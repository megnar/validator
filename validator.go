package validator

import (
	"errors"
	"reflect"
)

func Validator(x any) error {
	to := reflect.TypeOf(x)
	vo := reflect.ValueOf(x)
	numField := to.NumField()
	for x := 0; x < numField; x++ {
		val := vo.Field(x).String()
		name := to.Field(x).Name
		if name == "Title" {
			if val == "" || len(val) > 100 {
				return errors.New("Title incorrect")
			}
		}
		if name == "Text" {
			if val == "" || len(val) > 500 {
				return errors.New("Text incorrect")
			}
		}
	}
	return nil
}
