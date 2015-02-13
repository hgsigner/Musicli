package main

import "reflect"

func SliceContains(arg interface{}, match interface{}) bool {

	slice, ok := CheckInterfaceType(arg, reflect.Slice)
	if !ok {
		return false
	}

	length := slice.Len()
	for i := 0; i < length; i++ {
		if slice.Index(i).Interface() == match {
			return true
		}
	}

	return false
}

func CheckInterfaceType(arg interface{}, kind reflect.Kind) (reflect.Value, bool) {
	val := reflect.ValueOf(arg)
	if val.Kind() == kind {
		return val, true
	}

	return val, false
}
