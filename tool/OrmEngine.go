package tool

import (
	"CloudRestaurant/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

/**
	orm 映射数据库的一个类
**/

// 全局
var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database //获取配置文件
	// 链接数据库
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSql) //是否显示sql语句

	// 结构体映射为数据库表
	err = engine.Sync2(new(model.SmsCode), new(model.Member))
	err = engine.Sync2(
		new(model.SmsCode),
		new(model.Member),
		new(model.FoodCategory),
		new(model.Shop),
		new(model.Service),
		new(model.ShopService),
		new(model.Goods),
	)
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	// 全局
	DbEngine = orm

	return orm, nil
}
