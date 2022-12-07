package fun

import (
	"fmt"
	"reflect"
	"strings"
)

func Struct2Map(obj interface{}, tag string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { //识别指针
		t = t.Elem()
		v = v.Elem()
	}
	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { //结构体
		for i := 0; i < t.NumField(); i++ {
			if tag != "" {
				tagArr := strings.Split(t.Field(i).Tag.Get(tag), ",")
				if len(tagArr) > 1 && tagArr[1] == "omitempty" {
					if v.Field(i).IsZero() {
						continue
					} else if v.Field(i).Kind() == reflect.String && v.Field(i).String() == "" {
						continue
					} else if v.Field(i).Kind() == reflect.Slice && v.Field(i).Len() == 0 {
						continue
					} else if v.Field(i).Kind() == reflect.Map && v.Field(i).Len() == 0 {
						continue
					} else if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
						continue
					} else if v.Field(i).Kind() == reflect.Interface && v.Field(i).IsNil() {
						continue
					}
				}
				data[tagArr[0]] = v.Field(i).Interface()
			} else {
				data[t.Field(i).Name] = v.Field(i).Interface()
			}
		}
	}
	return data
}

func Struct2Str(obj interface{}, tag string) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { //识别指针
		t = t.Elem()
		v = v.Elem()

	}
	var data string
	if t.Kind() == reflect.Struct { //结构体
		for i := 0; i < t.NumField(); i++ {
			if tag != "" {
				tagArr := strings.Split(t.Field(i).Tag.Get(tag), ",")
				if len(tagArr) > 1 && tagArr[1] == "omitempty" {
					if v.Field(i).IsZero() {
						continue
					} else if v.Field(i).Kind() == reflect.String && v.Field(i).String() == "" {
						continue
					} else if v.Field(i).Kind() == reflect.Slice && v.Field(i).Len() == 0 {
						continue
					} else if v.Field(i).Kind() == reflect.Map && v.Field(i).Len() == 0 {
						continue
					} else if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
						continue
					} else if v.Field(i).Kind() == reflect.Interface && v.Field(i).IsNil() {
						continue
					}
				}
				data += tagArr[0] + " : " + fmt.Sprintf("%v", v.Field(i).Interface()) + "\n"
			} else {
				data += t.Field(i).Name + " : " + fmt.Sprintf("%v", v.Field(i).Interface()) + "\n"
			}
		}
	}
	return data
}

func Struct2Url(obj interface{}, tag string, isEnc bool) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { //识别指针
		t = t.Elem()
		v = v.Elem()
	}
	var data string
	if t.Kind() == reflect.Struct { //结构体
		for i := 0; i < t.NumField(); i++ {
			//if !v.Field(i).IsZero() {
			if tag != "" {
				tagArr := strings.Split(t.Field(i).Tag.Get(tag), ",")
				if len(tagArr) > 1 && tagArr[1] == "omitempty" {
					if v.Field(i).IsZero() {
						continue
					} else if v.Field(i).Kind() == reflect.String && v.Field(i).String() == "" {
						continue
					} else if v.Field(i).Kind() == reflect.Slice && v.Field(i).Len() == 0 {
						continue
					} else if v.Field(i).Kind() == reflect.Map && v.Field(i).Len() == 0 {
						continue
					} else if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
						continue
					} else if v.Field(i).Kind() == reflect.Interface && v.Field(i).IsNil() {
						continue
					}
				}
				if isEnc {
					data += "&" + tagArr[0] + "=" + UrlEncode(fmt.Sprintf("%v", v.Field(i).Interface()))
				} else {
					data += "&" + tagArr[0] + "=" + fmt.Sprintf("%v", v.Field(i).Interface())
				}

			} else {
				if isEnc {
					data += "&" + t.Field(i).Name + "=" + UrlEncode(fmt.Sprintf("%v", v.Field(i).Interface()))
				} else {
					data += "&" + t.Field(i).Name + "=" + fmt.Sprintf("%v", v.Field(i).Interface())
				}
			}
			//}
		}
	}
	if len(data) > 0 {
		return data[1:]
	} else {
		return ""
	}
}

func ArrayKeys(m map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
