package main

import (
	"fmt"
	"goggles/controllers"
	"goggles/models"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	tmpl := iris.Handlebars("./templates", ".html")
	tmpl.Reload(true) //dev mode

	tmpl.AddFunc("status_class", func(status int) string {
		if status >= 200 && status < 300 {
			return "success"
		} else if status >= 300 && status < 400 {
			return "warning"
		} else if status >= 400 {
			return "danger"
		}

		return "success"
	})

	app.RegisterView(tmpl)

	//initialise ORM
	db, err := gorm.Open("sqlite3", "./db/gorm.db")

	if err != nil {
		fmt.Println("Error", err)
	}

	//AutoMigrate models
	db.AutoMigrate(&models.Movies{}, &models.EndPoints{}, &models.EndPointCalls{})

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

	app.Get("/admin/endpoints", func(ctx iris.Context) {
		dashBoard := controllers.DashBoardController{Cntx: ctx}
		dashBoard.ShowEndpoints()
	})

	app.Get("/", func(ctx iris.Context) {
		(new(controllers.HomeController)).Show()
	})

	app.Run(iris.Addr("127.0.0.1:1234"))
}
