package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

}

type Origin struct {
	ignoreTag       ignoreTag // 需要忽略的tag及标记值，默认使用"ignore" != ""
	ignoreFieldName []string
	Origin          interface{}
}

// SetIgnoreTag 设置自定义的忽略标签
func (o *Origin) SetIgnoreTag(tag, value string) {
	o.ignoreTag.tag = tag
	o.ignoreTag.val = value
}

// AddIgnoreFieldName 设置自定义的字段，Origin是默认忽略字段
func (o *Origin) AddIgnoreFieldName(fieldNames ...string) {
	o.ignoreFieldName = append(o.ignoreFieldName, fieldNames...)
}

func (o *Origin) SetForUpdate(origin interface{}) {
	o.Origin = origin
	o.SetIgnoreTag("ignore", "-")
	o.AddIgnoreFieldName("Origin")
}

// GetModifyColumnMap 获取原数据与修改后数据字段值不同的值的键值对：map[fieldName]value
func (o *Origin) GetModifyColumnMap(modify interface{}) (map[string]interface{}, error) {
	originElem, modifiedElem, err := o.getOriginAndModifyElem(modify)
	if err != nil {
		return nil, err
	}
	return o.getModifyFieldsMap(originElem, modifiedElem), nil
}

// 获取原数据reflect.Value 和 修改后数据reflect.Value 若两个数据源类型不同，会导致报错
func (o *Origin) getOriginAndModifyElem(modify interface{}) (originElem, modifiedElem reflect.Value, err error) {
	if o.Origin == nil {
		err = errors.New("have no origin model value")
		return
	}
	originStructElem := reflect.ValueOf(o).Elem()
	originElem = originStructElem.FieldByName("Origin").Elem()
	modifiedElem = reflect.ValueOf(modify).Elem()
	if originElem.Type() != modifiedElem.Type() {
		err = errors.New(fmt.Sprintf("原始类型[%v]和修改类型[%v]不一致", originElem.Type(), modifiedElem.Type()))
	}
	return
}

// GetModifyColumnList 获取原数据与修改后数据字段值不同的字段
func (o *Origin) GetModifyColumnList(modify interface{}) ([]string, error) {
	originElem, modifiedElem, err := o.getOriginAndModifyElem(modify)
	if err != nil {
		return nil, err
	}
	return o.getModifyFieldsList(originElem, modifiedElem), nil
}

func (o *Origin) isIgnoreField(fileName string) bool {
	for _, s := range o.ignoreFieldName {
		if s == fileName {
			return true
		}
	}
	return false
}

// 获取原数据与修改后数据字段值不同的值的键值对：map[fieldName]value
func (o *Origin) getModifyFieldsMap(originElem, modifyElem reflect.Value) map[string]interface{} {
	modifyMap := make(map[string]interface{})
	handler := func(fieldName string, modifyValue interface{}) {
		modifyMap[fieldName] = modifyValue
	}
	o.diffIterator(originElem, modifyElem, handler)
	return modifyMap
}

// 遍历器,遇到字段不同时，会调用handler
func (o *Origin) diffIterator(originElem, modifyElem reflect.Value, handler func(fieldName string, modifyValue interface{})) {
	for i := 0; i < originElem.NumField(); i++ {
		if !originElem.Field(i).CanInterface() {
			continue
		}
		if originElem.Type().Field(i).Tag.Get(o.ignoreTag.tag) == o.ignoreTag.val {
			continue
		}
		fieldName := originElem.Type().Field(i).Name
		if o.isIgnoreField(fieldName) {
			continue
		}
		modifyValue := modifyElem.FieldByName(fieldName).Interface()
		originValue := originElem.Field(i).Interface()
		if !reflect.DeepEqual(originValue, modifyValue) {
			handler(fieldName, modifyValue)
		}
	}
}

// 获取原数据与修改后数据字段值不同的字段
func (o *Origin) getModifyFieldsList(originElem, modifyElem reflect.Value) []string {
	var resultList []string
	handler := func(fieldName string, modifyValue interface{}) {
		resultList = append(resultList, fieldName)
	}
	o.diffIterator(originElem, modifyElem, handler)
	return resultList
}

type ignoreTag struct {
	tag string
	val string
}
