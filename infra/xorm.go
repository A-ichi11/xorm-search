package infra

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func DBInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// xormで使ったSQLをログに吐きます
	engine.ShowSQL(true)

	// ユーザーテーブルが存在するかチェック
	exist, err := engine.IsTableExist("users")
	if err != nil {
		log.Fatal(err)
	}

	// 存在しなければテーブルを作成します
	if !exist {
		engine.CreateTables(&model.Users{})
	}

	return engine
}
