package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct {

}

/**
 * shop模块的路由解析
 */
func (sc *ShopController)Router(engine *gin.Engine){
	engine.GET("/api/shops",sc.GetShopList)
	engine.GET("/api/search_shops", sc.SearchShop)
}
func (sc *ShopController)SearchShop(context *gin.Context){
	longitude:=context.Query("longitude")
	latitude:=context.Query("latitude")
	keyword:=context.Query("keyword")
	if keyword==""{
		tool.Failed(context,"请重新输入商铺名称")
		return
	}
	if longitude==""||longitude=="undefined"||latitude==""||latitude=="undefined"{
		longitude="116.34"
		latitude="40.34"
	}
	shopService:=service.ShopService{}
	shops:=shopService.SearchShops(longitude,latitude,keyword)
	if len(shops)!=0{
		tool.Success(context,shops)
		return
	}
	tool.Failed(context,"暂未获取到商户信息")
}

/**
 * 获取商铺列表
 */
func (sc *ShopController)GetShopList(context *gin.Context){
	longitude:=context.Query("longitude")
	latitude:=context.Query("latitude")
	if longitude==""||longitude=="undefined"||latitude==""||latitude=="undefined"{
		longitude="116.34"
		latitude="40.34"
	}
	shopService:=service.ShopService{}
	shops:=shopService.ShopList(longitude,latitude)
	if len(shops)==0{
		tool.Failed(context,"暂未获取到商户信息")
		return
	}
	for _,shop:=range shops{
		shopServices:=shopService.GetService(shop.Id)
		if len(shopServices)==0{
			shop.Supports=nil
		}else {
			shop.Supports=shopServices
		}
		tool.Success(context,shop)
	}
}