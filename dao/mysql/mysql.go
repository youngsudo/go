package mysql

import (
	"app/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// 小写的 db变量 不能对外暴露, 只能在包内部使用
// 所以在外面 关闭 db 连接的时候, 就不能使用 defer db.Close()
// 可以封装一个函数,可以在外面调用, 关闭 db 连接
func Close() {
	_ = db.Close()
}
