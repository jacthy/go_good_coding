package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	val1 := StepOne{
		CustomerId:   1111,
		Name:         "name",
		UpdTime:      time.Now(),
		BooleanValue: true,
		StrList:      StringSlice{"1", "2", "3"},
		Ignore:       StringSlice{"1", "2", "3"},
		NoGormTag:    StringSlice{"1", "2", "3"},
	}
	val1.SetForUpdate()
	//err := StructClone(&val1, &val1.Origin.Origin)
	err := StructClone(&val1, new(StepOne))
	if err != nil {
		println(err.Error())
	}
}

type Customer struct {
	CustomerId int32
	Name       string
	UpdTime    time.Time
	TagList    []string
	SheetList  []Sheet
	Origin     interface{}
	Layer      Layer
}

type Sheet struct {
	SheetId int
	Name    string
}

type Layer struct {
	CodeList []int
}

// StepOne 扁平结构下的各种类型，以及gorm不同标签，测试效果
type StepOne struct {
	CustomerId   int32       `json:"customerId" gorm:"column:customer_id"`
	Name         string      `json:"name" gorm:"column:name"`
	UpdTime      time.Time   `json:"updTime" gorm:"column:update_time"`
	BooleanValue bool        `json:"booleanValue" gorm:"column:boolean_value"`
	StrList      StringSlice `json:"strList" gorm:"column:str_list"`
	Ignore       StringSlice `json:"strListIgnore" gorm:"-"`
	NoGormTag    StringSlice `json:"noGormTag"`
	Origin       interface{}
}

type Origin struct {
	Origin interface{}
}

func (s *StepOne) SetForUpdate() {
	s.Origin = nil
	s.Origin = *s
}

type StringSlice []string

func (s *StringSlice) Scan(value interface{}) error {
	jsonData, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("assertion failed, type: %T, value: %+v", value, value)
	}
	if string(jsonData) == "" {
		return nil
	}
	err := json.Unmarshal(jsonData, s)
	if err != nil {
		return err
	}
	return nil
}

func (s StringSlice) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}

	jsonData, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return string(jsonData), nil
}

func StructClone(srcPtr interface{}, dstPtr interface{}) (err error) {
	srcv, dstv := reflect.ValueOf(srcPtr), reflect.ValueOf(dstPtr)
	srct, dstt := reflect.TypeOf(srcPtr), reflect.TypeOf(dstPtr)

	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr || srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		return fmt.Errorf("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		return fmt.Errorf("Fatal error:value of parameters should not be nil")
	}

	srcValue := srcv.Elem()
	dstValue := dstv.Elem()
	srcfields := deepFields(srcValue.Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstValue.FieldByName(v.Name)
		src := srcValue.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		// 支持 string-->[]string 的合理变化【英文,分割】
		if src.Kind() == reflect.String && dst.Kind() == reflect.Slice && dst.CanSet() {
			srcStr := src.String()
			switch dst.Interface().(type) {
			case []string:
				dstStrArr := []string{}
				if srcStr != "" {
					dstStrArr = strings.Split(srcStr, ",")
				}
				dst.Set(reflect.ValueOf(dstStrArr))
				continue
			}
		}

		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

func deepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}
