package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/julienschmidt/httprouter"
	"github.com/kataras/iris/sessions"
	"wemall-go/config"
	"wemall-go/model"
	"wemall-go/route"
)

func init() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	if config.DBConfig.SQLLog {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns);
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)

	model.DB = db;
}

func main() {
	app := iris.New(iris.Configuration{
        Gzip    : true, 
        Charset : "UTF-8",
	})

	if config.ServerConfig.Debug {
		app.Adapt(iris.DevLogger())
	}

	app.Adapt(sessions.New(sessions.Config{
		Cookie: config.ServerConfig.SessionID,
		Expires: time.Minute * 20,
	}))

	app.Adapt(httprouter.New())

	route.Route(app)

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{
			"errNo" : model.ErrorCode.NotFound,
			"msg"   : "Not Found",
			"data"  : iris.Map{},
		})

	})

	app.OnError(500, func(ctx *iris.Context) {
		ctx.JSON(iris.StatusInternalServerError, iris.Map{
			"errNo" : model.ErrorCode.ERROR,
			"msg"   : "error",
			"data"  : iris.Map{},
		})
	})

	app.Listen(":" + strconv.Itoa(config.ServerConfig.Port))
}



