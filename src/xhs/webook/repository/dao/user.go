package dao

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

var (
	ComfilctErr     = errors.New("已注册")
	UserNotFoundErr = gorm.ErrRecordNotFound
)

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (dao *UserDao) FindEmail(c context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao *UserDao) FindPhone(c context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("phone = ?", phone).First(&u).Error
	return u, err
}

func (dao *UserDao) FindId(c context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("`id` = ?", id).First(&u).Error
	return u, err
}

func (dao *UserDao) Insert(c context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.Create(&u).Error
	if mysqlerr, ok := err.(*mysql.MySQLError); ok {
		if mysqlerr.Number == 1062 {
			return ComfilctErr
		}
	}
	return err
}

type User struct {
	Id       int64          `gorm:"primaryKey,autoIncrement"`
	Email    sql.NullString `gorm:"unique"`
	Password string

	Phone sql.NullString `gorm:"unique"`

	Ctime int64
	Utime int64
}
