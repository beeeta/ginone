package user

import (
	"dbutils"
	"time"
)

type UserEntity struct {
	Id         int64
	Name       string    `xorm:"'name' unique" form:"name"`
	Password   string    `xorm:"'password' notnull" form:"password"`
	createTime time.Time `xorm:"created 'createTime'"`
	updateTime time.Time `xorm:"updated 'updateTime'"`
}

func (users *UserEntity) TableName() string {
	return "demo_user"
}

func (users *UserEntity) InsertUser(insertUser *UserEntity) bool {
	engine := dbutils.GetDb()
	//engine.Sync2(new(UserEntity))
	_, err := engine.Insert(&UserEntity{Name: "test", Password: "test2"})
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func (users *UserEntity) QueryUser() *[]UserEntity {
	engine := dbutils.GetDb()
	userss := make([]UserEntity, 0)
	err := engine.Find(&userss)
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &userss
}
