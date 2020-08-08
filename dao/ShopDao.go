package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao()*ShopDao{
	return &ShopDao{tool.DbEngine}
}

func (sd *ShopDao)QueryServiceByShopId(shopId int64)[]model.Service{
	var services []model.Service
	err:=sd.Table("service").Join("INNER","shop_service",
		"service.id=shop_service.service_id and shop_id=?",shopId).Find(&services)
	if err!=nil{
		return nil
	}
	return services
}

const DEFAULT_RANGE = 5

/**
 * 操作数据库查询商铺数据列表
 */
func (sd *ShopDao)QueryShops(Longitude,latitude float64,keyword string)[]model.Shop{
	var shops []model.Shop
	if keyword==""{
		err:=sd.Engine.Where("Longitude<? and Longitude>? and latitude<? and latitude>? and status=1",
			Longitude+DEFAULT_RANGE,Longitude-DEFAULT_RANGE,latitude+DEFAULT_RANGE,latitude-DEFAULT_RANGE).Find(&shops)
		if err!=nil{
			fmt.Println(err.Error())
			return nil
		}
	}else{
		err:=sd.Engine.Where("Longitude<? and Longitude>? and latitude<? and latitude>? and status=1 and name like ?",
			Longitude+DEFAULT_RANGE,Longitude-DEFAULT_RANGE,latitude+DEFAULT_RANGE,latitude-DEFAULT_RANGE,keyword).Find(&shops)
		if err!=nil{
			fmt.Println(err.Error())
			return nil
		}
	}

	return shops
}