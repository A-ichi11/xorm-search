package model

type Users struct {
	ID      int    `xorm:"id pk autoincr"`
	Name    string `xorm:"name"`
	Address string `xorm:"address"`
}
