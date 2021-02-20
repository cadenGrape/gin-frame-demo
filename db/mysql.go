package db

import (
	"errors"
	"fmt"
	orm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sync"
)

type MySqlPool struct {}

var (
	instance *MySqlPool
	once sync.Once
	db *orm.DB
	err error
)

//单例模式
func GetInstance() *MySqlPool {
	once.Do(func() {
		instance = &MySqlPool{}
	})
	return instance
}

func (pool *MySqlPool) InitPool() (isSuc bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"), viper.GetString("db.charset"))
	db, err = orm.Open("mysql", dsn)
	if err != nil {
		panic(errors.New("mysql连接失败"))
		return false
	}
	//读取连接数配置
	db.DB().SetMaxIdleConns(viper.GetInt("db.MaxIdleConns"))
	db.DB().SetMaxOpenConns(viper.GetInt("db.MaxOpenConns"))
	return true
}