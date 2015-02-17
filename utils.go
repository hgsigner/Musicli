package main

import (
	"fmt"
	"reflect"
)

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

func FormatStructToText(str interface{}) string {

	st := reflect.ValueOf(str)
	fields_count := st.NumField()
	retrieved_urls := []string{}

	for i := 0; i < fields_count; i++ {
		if field_value := st.Field(i).String(); field_value != "" {
			retrieved_urls = append(retrieved_urls, field_value)
		}
	}

	if len(retrieved_urls) == 0 {
		return ""
	}

	formated_urls := ""
	for _, url := range retrieved_urls {
		formated_urls = fmt.Sprintf("%s%s\n", formated_urls, url)
	}

	return formated_urls

}
