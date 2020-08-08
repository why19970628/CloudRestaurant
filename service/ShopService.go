package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"strconv"
)

type ShopService struct {

}

func (ss *ShopService)GetService(shopId int64)[]model.Service{
	shopDao:=dao.NewShopDao()
	return shopDao.QueryServiceByShopId(shopId)
}
func (ss *ShopService)SearchShops(long,lat,keyword string)[]model.Shop{
	shopDao:=dao.NewShopDao()
	longitude,err:=strconv.ParseFloat(long,10)
	if err!=nil{
		return nil
	}
	latitude,err:=strconv.ParseFloat(lat,10)
	if err!=nil{
		return nil
	}
	return shopDao.QueryShops(longitude,latitude,keyword)
}

/**
 * 查询商铺列表数据
 */
func (ss *ShopService)ShopList(long,lat string)[]model.Shop{
	shopDao:=dao.NewShopDao()
	longitude,err:=strconv.ParseFloat(long,10)
	if err!=nil{
		return nil
	}
	latitude,err:=strconv.ParseFloat(lat,10)
	if err!=nil{
		return nil
	}
	return shopDao.QueryShops(longitude,latitude,"")
}
