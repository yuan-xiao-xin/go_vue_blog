package model

import (
	"fmt"
	"go_vue_blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

// InitDB 初始化数据库连接
func InitDB() {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(
		mysql.Open(fmt.Sprintf(
			dsn,
			utils.DbUser,
			utils.DbPassword,
			utils.DbHost,
			utils.DbPort,
			utils.DbName,
		)),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	//获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取底层sql.DB失败，请检查参数：", err)
	}
	//设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//设置连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	
	//测试连接是否有效
	if err := sqlDB.Ping(); err != nil {
		fmt.Println("连接失效，请检查参数：", err)
		sqlDB.Close() // 仅在Ping失败时关闭
	}
	
	//进行模型迁移
	err = AutoMigrateModels(db)
	if err != nil {
		fmt.Println("自动迁移模型失败：", err)
	}
	
}

func AutoMigrateModels(db *gorm.DB) error {
	model := []any{
		&User{},
		&Article{},
		&Category{},
	}
	if err := db.AutoMigrate(model...); err != nil {
		return err
	}
	return nil
}
