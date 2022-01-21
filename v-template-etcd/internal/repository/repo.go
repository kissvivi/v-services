package repository

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"v-template-etcd/config"
	"v-template-etcd/internal/model"
)

//  数据访问层

var RepoSet = wire.NewSet(NewData, NewMysql, NewTestRepo, NewDemoRepo)

// Data 数据访问连接工具 以及外部数据
type Data struct {
	Mysql *gorm.DB
	//log *log.Helper
}

//mysql

func NewMysql(cfg *config.Config) *gorm.DB {

	user := cfg.Database.Mysql.Username
	pwd := cfg.Database.Mysql.Password
	url := cfg.Database.Mysql.Url
	dbname := cfg.Database.Mysql.Dbname

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, url, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//初始化表结构
	db.AutoMigrate(&model.Test{})

	return db
}

//func NewRedis()  {
//
//}
//
//func NewMongo()  {
//
//}

func NewData(mysql *gorm.DB) (*Data, func(), error) {

	return &Data{Mysql: mysql}, func() {

	}, nil
}
