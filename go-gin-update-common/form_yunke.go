package main

import (
	"fmt"
	"reflect"
	"time"
)

type StepOneYunke struct {
	Id           int32       `json:"id" gorm:"column:customer_id" relation:"StepOne:CustomerId"`
	Name         string      `json:"name" gorm:"column:name" relation:"StepOne:Name"`
	UpdTime      time.Time   `json:"updTime" gorm:"column:update_time" relation:"StepOne:UpdTime"`
	BooleanValue bool        `json:"booleanValue" gorm:"column:boolean_value" relation:"StepOne:BooleanValue"`
	StrList      StringSlice `json:"strList" gorm:"column:str_list" relation:"StepOne:StrList"`
	Ignore       StringSlice `json:"strListIgnore" gorm:"-"`
	NoGormTag    StringSlice `json:"noGormTag"`
	origin       OriginYunke
}

type OriginYunke struct {
	Origin interface{}
}

// RealGetModifiedFields 获取修改结构体和原始结构体中差异部分
func (origin *OriginYunke) RealGetModifiedFields(modified interface{}) map[string]interface{} {
	if origin.Origin == nil {
		return nil
	}
	originStructElem := reflect.ValueOf(origin).Elem()
	originElem := originStructElem.FieldByName("Origin").Elem()
	modifiedElem := reflect.ValueOf(modified).Elem()

	if originElem.Type() != modifiedElem.Type() {
		panic(fmt.Sprintln("原始类型和修改类型不一致", originElem.Type(), modifiedElem.Type()))
	}

	modifiedColumnValue := origin.modifiedFields("", originElem, modifiedElem)

	return modifiedColumnValue
}

func (origin *OriginYunke) modifiedFields(parentFieldName string, originElem, modifiedElem reflect.Value) map[string]interface{} {
	numField := originElem.NumField()
	modifiedColumnValue := make(map[string]interface{})
	for i := 0; i < numField; i++ {
		if originElem.Type().Field(i).Tag.Get("origin") == "-" {
			continue
		}
		if !originElem.Field(i).CanInterface() {
			continue
		}
		fieldName := originElem.Type().Field(i).Name
		// 结构体是否需要深入处理
		if originElem.Type().Field(i).Tag.Get("origin") == "deep" {
			ne := originElem.Field(i)
			me := modifiedElem.FieldByName(fieldName)
			if ne.Kind() == reflect.Ptr {
				ne = ne.Elem()
				me = me.Elem()
			}
			if ne.Kind() == reflect.Struct {
				origin.addToMap(modifiedColumnValue, origin.modifiedFields(origin.appendFieldName(parentFieldName, fieldName), ne, me))
			}
		}

		modifiedFieldValue := modifiedElem.FieldByName(fieldName).Interface()
		originFieldValue := originElem.Field(i).Interface()
		if !reflect.DeepEqual(modifiedFieldValue, originFieldValue) {
			modifiedColumnValue[origin.appendFieldName(parentFieldName, fieldName)] = modifiedFieldValue
		}
	}
	return modifiedColumnValue
}

func (origin *OriginYunke) appendFieldName(parentFieldName, fieldName string) string {
	if parentFieldName == "" {
		return fieldName
	}
	return parentFieldName + "." + fieldName
}

func (origin *OriginYunke) addToMap(a map[string]interface{}, b map[string]interface{}) {
	for k, v := range b {
		a[k] = v
	}
}
