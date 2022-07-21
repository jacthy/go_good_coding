package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"time"
)

type Customer struct {
	CustomerId int32
	Name       string
	UpdateTime time.Time `ignore:"-"` // 测试通过tag忽略
	lowCase    string    // 小写字段 会被忽略掉, 因为无法Interface()
	IntArr     []int
	StrArr     []string
	Origin
}

func TestOrigin_All(t *testing.T) {
	c := new(Customer)
	c.Name = "1"
	c.CustomerId = 1
	c.SetForUpdate(*c)
	c.Name = "2"
	c.CustomerId = 2
	c.UpdateTime = time.Now().Add(10 * time.Second)
	modifyMap, err := c.GetModifyColumnMap(c)
	if err != nil {
		t.Error(err)
	}
	if _, ok := modifyMap["CustomerId"]; !ok {
		t.Error("未能捕获到customerId的更新")
	}
	if _, ok := modifyMap["Name"]; !ok {
		t.Error("未能捕获到Name的更新")
	}

	list, err := c.GetModifyColumnList(c)
	if err != nil {
		t.Error(err)
	}
	if strings.Join(list, ",") != "CustomerId,Name" {
		t.Error("未能准确获取modify字段列表")
	}
}

func TestOrigin_AddIgnoreFieldName(t *testing.T) {
	c := new(Customer)
	c.Name = "1"
	c.CustomerId = 1
	c.IntArr = []int{1, 2}
	c.SetForUpdate(*c)
	c.AddIgnoreFieldName("Name")
	c.Name = "2"
	c.CustomerId = 2
	c.IntArr = append(c.IntArr, 3)
	c.StrArr = append(make([]string, 0), "1", "2", "3")
	modify, _ := c.GetModifyColumnMap(c)
	Convey("test", t, func() {
		modifyStr := fmt.Sprintf("%+v", modify)
		So(modifyStr, ShouldEqual, "map[CustomerId:2 IntArr:[1 2 3] StrArr:[1 2 3]]")
	})
	c.SetForUpdate(*c)
	c.IntArr = append(c.IntArr, 1)
	modify, _ = c.GetModifyColumnMap(c)
	Convey("test", t, func() {
		modifyStr := fmt.Sprintf("%+v", modify)
		So(modifyStr, ShouldEqual, "map[IntArr:[1 2 3 1]]")
	})
}
