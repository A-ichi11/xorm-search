package model

type Users struct {
	ID      int    `xorm:"id pk autoink"`
	Name    string `xorm:"name"`
	Address string `xorm:"address"`
}
