package main

import (
	"fmt"
	"github.com/Jungle20m/electricity/component"
	"github.com/Jungle20m/electricity/config"
	"github.com/Jungle20m/electricity/internal/httpserver/app"
	"github.com/Jungle20m/electricity/sdk/httpserver"
	"github.com/Jungle20m/electricity/sdk/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Mysql
	msql, err := mysql.New("anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// App context
	appCtx := component.NewAppContext(msql)

	// Http server
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	app.NewRouter(handler, appCtx)
	httpserver := httpserver.New(handler)

	// Listen signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		fmt.Println("app - Run - signal: " + s.String())
	case err := <-httpserver.Notify():
		fmt.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpserver.Shutdown()
	if err != nil {
		fmt.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
