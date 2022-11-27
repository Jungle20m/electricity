package main

import (
	"context"
	"github.com/Jungle20m/electricity/component"
	"github.com/Jungle20m/electricity/config"
	"github.com/Jungle20m/electricity/internal/httpserver"
	mHttpServer "github.com/Jungle20m/electricity/sdk/httpserver"
	mLogger "github.com/Jungle20m/electricity/sdk/logger"
	mMysql "github.com/Jungle20m/electricity/sdk/mysql"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewConfig() (*config.Config, error) {
	conf, err := config.NewConfig("config/config.yaml")
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func NewLogger(conf *config.Config) (*mLogger.Logger, error) {
	logger, err := mLogger.New(
		mLogger.WithLevel(conf.Log.Level),
		mLogger.WithFormatterMode(conf.Log.Formatter),
		mLogger.WithOutputMode(conf.Log.Output),
		mLogger.WithDirectory(conf.Log.Folder),
		mLogger.WithFileName(conf.Log.FileName))
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func NewMysql(conf *config.Config) (*mMysql.Mysql, error) {
	msql, err := mMysql.New(conf.Mysql.Dns)
	if err != nil {
		return nil, err
	}
	return msql, nil
}

func NewAppContext(conf *config.Config, log *mLogger.Logger, msql *mMysql.Mysql) component.AppContext {
	appCtx := component.NewAppContext(msql, log, conf)
	return appCtx
}

func NewHttpServer(appCtx component.AppContext) (*mHttpServer.Server, error) {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	httpserver.NewRouter(handler, appCtx)
	server := mHttpServer.New(handler)
	return server, nil
}

func StartHTTPServer(lifecycle fx.Lifecycle, server *mHttpServer.Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				server.Start()
				return nil
			},
			OnStop: func(context.Context) error {
				server.Shutdown()
				return nil
			},
		},
	)
}

func StartGrpcServer() {}

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			NewMysql,
			NewAppContext,
			NewHttpServer,
		),
		fx.Invoke(
			StartHTTPServer,
		),
	).Run()
}
