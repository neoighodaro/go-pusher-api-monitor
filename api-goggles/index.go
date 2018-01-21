package main

import (
	"goggles/controllers"
	"goggles/models"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	//initialise ORM
	db, _ := gorm.Open("sqlite3", "./db/gorm.db")

	//AutoMigrate models
	db.AutoMigrate(&models.Movies{})

	//api endpoints
	app.Get("/api/movies", func(ctx iris.Context) {
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Get()
	})

	app.Post("/api/movies", func(ctx iris.Context) {
		mv := controllers.MoviesController{Cntx: ctx}
		mv.Add()
	})

	app.Get("/api/movies/{id}", func(ctx iris.Context) {
		ID, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
		mv := controllers.MoviesController{Cntx: ctx}
		mv.GetByID(ID)
	})

	app.Put("/api/movies/{id}", func(ctx iris.Context) {
		ID, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

		mv := controllers.MoviesController{Cntx: ctx}
		mv.Edit(ID)
	})

	app.Delete("/api/movies/{id}", func(ctx iris.Context) {
		ID, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

		mv := controllers.MoviesController{Cntx: ctx}
		mv.Delete(ID)
	})

	app.Get("/", func(ctx iris.Context) {

	})

	app.Run(iris.Addr("127.0.0.1:1234"))
}
