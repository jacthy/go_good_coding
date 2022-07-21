package main

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func main() {
	RegisterRelation(new(repoCustomerOld), new(CustomerOld))
	repo := new(CustomerOldRepo)
	model := CustomerOld{
		CustomerId: 111,
		Name:       "custorm",
		ActivityId: 111,
	}
	model.SetIsForUpdate()
	model.CustomerId = 222
	model.ActivityId = 222
	err := repo.UpdateCustomer(&model)
	if err != nil {
		println("eer:", err.Error())
	}
}

type CustomerOldRepo struct {
	db *gorm.DB
}

type CustomerOld struct {
	CustomerId     int32  // 客户ID
	Name           string // 客户名称
	ActivityId     int32  // 选房活动ID
	OrganizationId int32  // 商户ID
	IntArr         []int
	StrArr         []string
	origin         Origin
}

func (c *CustomerOldRepo) UpdateCustomer(customer *CustomerOld) error {
	model := new(repoCustomerOld)
	return c.updateColumns(customer, model, func(db *gorm.DB) *gorm.DB {
		return db.Model(model).Where("choose_room_user_id=?", customer.CustomerId)
	})
}

type scope func(*gorm.DB) *gorm.DB

// updateColumns 更新entity变动的字段
func (c *CustomerOldRepo) updateColumns(entity interface{}, model interface{}, s scope) error {
	iModifiedField, ok := entity.(ModifiedField)
	if !ok {
		panic(fmt.Sprintf("entity not realize ModifiedField interface %+v", entity))
	}
	modifiedFields := iModifiedField.GetModifiedFields()
	updateColumns := ToDBColumn(entity, modifiedFields, model)
	if len(updateColumns) <= 0 {
		return nil
	}
	// 使用 Updates 方法,会触发调用model的BeforeUpdate方法,可用于处理字段加密
	fmt.Printf("%+v", updateColumns)
	return nil
}

func (c *CustomerOld) SetIsForUpdate() {
	c.origin.Origin = nil
	c.origin.Origin = *c
}

func (c *CustomerOld) GetModifiedFields() map[string]interface{} {
	return c.origin.RealGetModifiedFields(c)
}

type ModifiedField interface {
	GetModifiedFields() map[string]interface{}
}

// ToDBColumn entity 对应修改的 字段,转换成想要的 model 改动数据库字段(key=>value)
func ToDBColumn(entity interface{}, modifiedColumnValue map[string]interface{}, model interface{}) map[string]interface{} {
	return relationCache.toDBColumn(entity, modifiedColumnValue, model)
}

// toDBColumn entity 对应修改的 字段,转换成想要的 model 改动数据库字段(key=>value)
func (r *relation) toDBColumn(entity interface{}, modifiedColumnValue map[string]interface{}, model interface{}) map[string]interface{} {
	modelType := reflect.TypeOf(model).Elem()
	entityType := reflect.TypeOf(entity).Elem()

	if _, ok := r.entityToModel[entityType]; !ok {
		panic("转换DB列名时没有找到注册的 entity")
	}
	entityToModelRelation, ok := r.entityToModel[entityType][modelType]
	if !ok {
		panic("转换DB列名时没有找到注册的 model")
	}

	updateColumns := make(map[string]interface{})
	// 转db 字段
	for entityField, v := range modifiedColumnValue {
		modelField, ok := entityToModelRelation.Field[entityField]
		if !ok {
			continue
		}
		dbColumn, ok := entityToModelRelation.DBColumn[modelField]
		if !ok {
			continue
		}

		updateColumns[dbColumn] = v
	}

	return updateColumns
}

var relationCache = newRelation()

func newRelation() relation {
	return relation{
		entityToModel: make(map[reflect.Type]modelRelations),
	}
}

type relation struct {
	// modelRelation entity -> model 关系映射
	entityToModel map[reflect.Type]modelRelations
}

func RegisterRelation(model interface{}, entities ...interface{}) {
	relationCache.registerRelation(model, entities...)
}

func (r *relation) registerRelation(model interface{}, entities ...interface{}) {
	for _, entity := range entities {

		modelType := reflect.TypeOf(model).Elem()
		entityType, fieldMap := relationCache.getStructNameAndFieldIndex(entity)

		entityToModelRelationRes := relationCache.getModelRelation(modelType, entityType, fieldMap)

		if _, ok := relationCache.entityToModel[entityType]; !ok {
			relationCache.entityToModel[entityType] = make(modelRelations)
		}

		relationCache.entityToModel[entityType][modelType] = entityToModelRelationRes
	}
}

func (r *relation) getStructNameAndFieldIndex(i interface{}) (reflect.Type, map[string]bool) {
	structType := reflect.TypeOf(i).Elem()
	fieldMap := r.deepFlatField(structType, "")
	return structType, fieldMap
}

// getModelRelation 返回model->relation &  entity -> relation

func (r *relation) getModelRelation(modelType reflect.Type, entityType reflect.Type, fieldMap map[string]bool) modelRelation {
	entityToModelRelation := newEntityToModel()

	entityName := entityType.Name()

	for i := 0; i < modelType.NumField(); i++ {
		modelFieldName := modelType.Field(i).Name

		fieldTag := modelType.Field(i).Tag

		structField := r.flatTag(fieldTag, relationTagName, ",")
		for entityStruct, entityField := range structField {
			if (entityStruct != "" && entityField == "") || (entityStruct == "" && entityField != "") {
				panic("注册model与entity关系,tag 标签错误 " + entityStruct + " " + entityField)
			}

			if entityStruct != "" && entityField != "" && modelFieldName != "" {
				if entityStruct == entityName {
					if !fieldMap[entityField] {
						panic("注册model与entity关系,entity字段不存在 " + entityName + "." + entityField)
					}

					dbTagMap := r.flatTag(fieldTag, gormTagName, ";")
					if dbColumnName, ok := dbTagMap[gormColumnTagName]; ok {
						entityToModelRelation.DBColumn[modelFieldName] = dbColumnName
					} else {
						entityToModelRelation.DBColumn[modelFieldName] = ToColumnName(modelFieldName)
					}
					entityToModelRelation.Field[entityField] = modelFieldName
				}
			}
		}
	}

	return entityToModelRelation
}

func ToColumnName(string2 string) string {
	return string2
}

var (
	relationTagName   = "relation"
	gormTagName       = "gorm"
	gormColumnTagName = "column"
)

// flatTag 打平对应的字段的tag, key 为对应需要解析的tag 名称
func (r *relation) flatTag(tag reflect.StructTag, key, keyValSeparate string) map[string]string {
	strTag := tag.Get(key)
	names := strings.Split(strTag, keyValSeparate)
	structField := make(map[string]string)
	for _, v := range names {
		realTag := strings.Split(v, ":")
		for i := 0; i < len(realTag)-1; i += 2 {
			structField[realTag[i]] = realTag[i+1]
		}
	}
	if key == relationTagName && strTag != "" && len(structField) <= 0 {
		panic("注册model与entity关系,tag 标签错误 " + strTag)
	}

	return structField
}

func (r *relation) deepFlatField(structType reflect.Type, parentFieldName string) map[string]bool {
	addToMap := func(a map[string]bool, b map[string]bool) {
		for k, v := range b {
			a[k] = v
		}
	}
	appendFieldName := func(parentFieldName, fieldName string) string {
		if parentFieldName == "" {
			return fieldName
		}
		return parentFieldName + "." + fieldName
	}
	fieldMap := make(map[string]bool)
	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fk := appendFieldName(parentFieldName, fieldName)
		if structType.Field(i).Type.Kind() == reflect.Ptr {
			st := structType.Field(i).Type.Elem()
			if st.Kind() == reflect.Struct {
				addToMap(fieldMap, r.deepFlatField(st, fk))
			}
		} else if structType.Field(i).Type.Kind() == reflect.Struct {
			addToMap(fieldMap, r.deepFlatField(structType.Field(i).Type, fk))
		}
		fieldMap[fk] = true
	}
	return fieldMap
}

type modelRelations map[reflect.Type]modelRelation

// modelRelation model 的映射关系
type modelRelation struct {
	// Field entity 对应的结构体名称 -> model 对应的结构体名称
	Field map[string]string
	// DBColumn model 结构体字段 -> 数据库表字段名称
	DBColumn map[string]string
}

// newEntityToModel 创建entityToModel
func newEntityToModel() modelRelation {
	return modelRelation{
		Field:    make(map[string]string),
		DBColumn: make(map[string]string),
	}
}

type Origin struct {
	Origin interface{}
}

// RealGetModifiedFields 获取修改结构体和原始结构体中差异部分
func (origin *Origin) RealGetModifiedFields(modified interface{}) map[string]interface{} {
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

func (origin *Origin) modifiedFields(parentFieldName string, originElem, modifiedElem reflect.Value) map[string]interface{} {
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

func (origin *Origin) appendFieldName(parentFieldName, fieldName string) string {
	if parentFieldName == "" {
		return fieldName
	}
	return parentFieldName + "." + fieldName
}

func (origin *Origin) addToMap(a map[string]interface{}, b map[string]interface{}) {
	for k, v := range b {
		a[k] = v
	}
}

type repoCustomerOld struct {
	ChooseRoomUserId     int32    `gorm:"primary_key;" relation:"CustomerOld:CustomerId"` // 资格用户ID
	ChooseRoomActivityId int32    `relation:"CustomerOld:ActivityId"`                     // 活动ID
	UserName             string   `relation:"CustomerOld:Name"`                           // 用户名称
	IntArr               []int    `relation:"CustomerOld:IntArr"`
	StrArr               []string `relation:"CustomerOld:IntArr"`
}
