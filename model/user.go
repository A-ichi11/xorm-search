package model

type Users struct {
	UserID int `xorm:"user_id pk autoink"`
	Name   string
	Age    int
}
