package model

import (
	"fmt"
	"go_yangzhou/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DbHandler *gorm.DB

func init()  {
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn,
		config.InitConfig.Database.DbUsername,
		config.InitConfig.Database.DbPassword,
		config.InitConfig.Database.DbHost,
		config.InitConfig.Database.DbPort,
		config.InitConfig.Database.DbDatabase)
	var err error
		DbHandler ,err = gorm.Open(mysql.Open(dsn),&gorm.Config{
			//默认关闭事务
			SkipDefaultTransaction: true,
		//DarRun 模式会生成但不执行 SQL，可以用于准备或测试生成的 SQL
		DryRun: false,
		//执行任何 SQL 以提高后续的调用时都创建并缓存预编译语句，可速度
		PrepareStmt: false,
		//禁止外键
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	//使用数据库连接错误
	if err != nil {
		panic(fmt.Sprintf("connect database failed %s \n",err))
	}
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	handler,_ := DbHandler.DB()
	//设置连接池的最大连接数，空闲连接数
	handler.SetMaxOpenConns(config.InitConfig.MaxConnection)
	handler.SetMaxIdleConns(config.InitConfig.MaxIdle)
	//最大空闲时常，最大复用时间
	//handler.SetConnMaxIdleTime(time.Second * 30)
	handler.SetConnMaxLifetime(time.Hour)
	// Ping
	handler.Ping()
	// Close
	//sqlDB.Close()
	// 返回数据库统计信息
	handler.Stats()
}