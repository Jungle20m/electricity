package component

import (
	"github.com/Jungle20m/electricity/config"
	mLogger "github.com/Jungle20m/electricity/sdk/logger"
	mMysql "github.com/Jungle20m/electricity/sdk/mysql"
)

type AppContext interface {
	GetMysql() *mMysql.Mysql
	GetLogger() *mLogger.Logger
	GetConfig() *config.Config
}

type appCtx struct {
	Mysql  *mMysql.Mysql
	Logger *mLogger.Logger
	Config *config.Config
}

func NewAppContext(msql *mMysql.Mysql, logger *mLogger.Logger, conf *config.Config) *appCtx {
	return &appCtx{
		Mysql:  msql,
		Logger: logger,
		Config: conf,
	}
}

func (ctx *appCtx) GetMysql() *mMysql.Mysql {
	return ctx.Mysql
}

func (ctx *appCtx) GetLogger() *mLogger.Logger {
	return ctx.Logger
}

func (ctx *appCtx) GetConfig() *config.Config {
	return ctx.Config
}
