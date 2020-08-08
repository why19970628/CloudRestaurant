package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"
)

type GoodDao struct {
	*tool.Orm
}

func NewGoodDao()*GoodDao{
	return &GoodDao{tool.DbEngine}
}

//根据商家的id查询商户下所拥有的所有的食品数据
func (gd *GoodDao)QueryFoods(shop_id int64)[]model.Goods{
	var goods []model.Goods
	err:=gd.Where("shop_id = ?",shop_id).Find(&goods)
	if err!=nil{
		fmt.Println(err.Error())
		return nil
	}
	return  goods
}
