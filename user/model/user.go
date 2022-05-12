package models

import (
	"gorm.io/gorm"
	"qiyaowu-go-zero/basic/pgsql"
)

const STATUS_ACTIVE = 1

type (
	UserModel interface {
		GetItem(userId int64) *User
	}

	defaultUserModel struct {
		conn *gorm.DB
	}
	User struct {
		Id                int64  `gorm:"primary_key" json:"id"`
		Uuid              string `gorm:"column:uuid;type:varchar(255)" json:"uuid"`
		UserName          string `gorm:"column:username;type:varchar(255)" json:"username"`
		PasswordHash      string `gorm:"column:password_hash;type:varchar(255)" json:"password_hash"`
		Nickname          string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
		Phone             string `gorm:"column:phone;type:varchar(255)" json:"phone"`
		Email             string `gorm:"column:email;type:varchar(255)" json:"email"`
		Role              string `gorm:"column:role;type:varchar(255)" json:"role"`
		ItemType          string `gorm:"column:item_type;type:varchar(255)" json:"item_type"`
		Status            int16  `gorm:"column:status;type:int4" json:"status"`
		Locked            bool
		LockedInfo        lockedInfo
		LoginsError       bool
		LoginErrorInfo    loginErrorInfo
		DeletedAt         string `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
		CreatedAt         string `gorm:"column:created_at;type:timestamp" json:"created_at" description:"创建时间"`
		UpdatedAt         string `gorm:"column:updated_at;type:timestamp" json:"updated_at" description:"修改时间"`
		UpdatedPasswordAt string `gorm:"column:updated_password_at;type:timestamp" json:"updated_password_at" description:"密码修改时间"`
	}
	lockedInfo struct {
		UserId     int64 `json:"user_id"`
		Number     int16 `json:"number"`
		LockedTime int64 `json:"locked_time"`
	}
	loginErrorInfo struct {
		UserId    int64 `json:"user_id"`
		Number    int16 `json:"number"`
		ErrorTime int64 `json:"error_time"`
	}
)

func NewUserModel() UserModel {
	return &defaultUserModel{
		conn: pgsql.GetDB(),
	}
}

func (u User) TableName() string {
	return "user"
}
func (this defaultUserModel) GetItem(userId int64) *User {
	users := &User{}
	if err := this.conn.Where(User{Id: userId}).First(&users).Error; err == nil {
		return users
	}
	return nil
}
