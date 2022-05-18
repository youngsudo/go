package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	return
}

// 小写的 db变量 不能对外暴露, 只能在包内部使用
// 所以在外面 关闭 db 连接的时候, 就不能使用 defer db.Close()
// 可以封装一个函数,可以在外面调用, 关闭 db 连接
func Close() {
	_ = db.Close()
}
