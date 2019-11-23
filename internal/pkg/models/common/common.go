package common

import (
	"github.com/jinzhu/gorm"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/db"
)

// First
func First(where interface{},out interface{})(notFound bool,err error){
	err= db.DB.Where(where).First(out).Error
	if err!=nil {
		notFound=gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find
func Find(where interface{},out interface{},orders ...string)error{
	db:=db.DB.Where(where)
	if len(orders)>0 {
		for _,order:=range orders {
			db=db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Updates
func Updates(where interface{},value interface{})error{
	return db.DB.Model(where).Updates(value).Error
}

// Create
func Create(value interface{})error{
	return db.DB.Create(value).Error
}

// GetPage
func GetPage(model, where interface{}, out interface{}, pageIndex, pageSize uint64, totalCount *int, whereOrder ...PageWhereOrder) error {
	db := db.DB.Model(model).Where(where)
	if len(whereOrder)>0 {
		for _,wo:=range whereOrder {
			if wo.Order !="" {
				db=db.Order(wo.Order)
			}
			if wo.Where !="" {
				db=db.Where(wo.Where,wo.Value...)
			}
		}
	}

	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}

	if *totalCount == 0 {
		return nil
	}

	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// GetAllInfos
func GetInfos(model, where interface{}, out interface{}, totalCount *int ) error{
	db:=db.DB.Model(model).Where(where)
	err:=db.Count(totalCount).Error
	if err!=nil{
		return err
	}
	if *totalCount==0{
		return nil
	}
	return db.Find(out).Error
}


