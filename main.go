package main

import (
	"fmt"
	"log"

	"github.com/A-ichi11/xorm-search/infra"
	"github.com/A-ichi11/xorm-search/model"
	"xorm.io/xorm"
)

var searchWords string = "佐"

func main() {
	// engineを作成します。
	engine := infra.DBInit()

	// user1 := model.Users{
	// 	Name:    "田中太郎",
	// 	Address: "大阪府",
	// }
	// user2 := model.Users{
	// 	Name:    "田中隆二",
	// 	Address: "京都府",
	// }
	// user3 := model.Users{
	// 	Name:    "佐々木太一",
	// 	Address: "東京都",
	// }
	// user4 := model.Users{
	// 	Name:    "佐々木次郎",
	// 	Address: "東京都",
	// }

	// users := []*model.Users{
	// 	// &user1,
	// 	// &user2,
	// 	// &user3,
	// 	// &user3,
	// 	&user4,
	// }

	// create(*engine, users)
	// find(*engine)

	// searchForwordMatch(*engine, searchWords)
	// searchBackwordMatch(*engine, searchWords)
	searchPartialMatch(*engine, searchWords)

}

// create ユーザー作成
func create(engine xorm.Engine, users []*model.Users) {
	_, err := engine.Table("users").Insert(users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

// find 複数取得(完全一致)
func find(engine xorm.Engine) {
	users := []model.Users{}
	// addressが大阪府のuserを全件取得します
	err := engine.Where("address = ?", "大阪府").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

// searchForwordMatch 前方一致
func searchForwordMatch(engine xorm.Engine, searchWord string) {
	users := []model.Users{}
	err := engine.Table("users").Where("name LIKE?", searchWord+"%").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

// searchBackwordMatch 後方一致
func searchBackwordMatch(engine xorm.Engine, searchWord string) {
	users := []model.Users{}
	err := engine.Table("users").Where("name LIKE?", "%"+searchWord).Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

// searchPartialMatch 部分一致
func searchPartialMatch(engine xorm.Engine, searchWord string) {
	users := []model.Users{}
	err := engine.Table("users").Where("name LIKE?", "%"+searchWord+"%").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}
