package component

import "github.com/Jungle20m/electricity/sdk/mysql"

type AppContext interface {
	GetMysql() *mysql.Mysql
}

type appCtx struct {
	Mysql *mysql.Mysql
}

func NewAppContext(msql *mysql.Mysql) *appCtx {
	return &appCtx{Mysql: msql}
}

func (ctx *appCtx) GetMysql() *mysql.Mysql {
	return ctx.Mysql
}
