package controllers

import (
	"goggles/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//DashBoardController - Controller object for Endpoints dashboard
type DashBoardController struct {
	mvc.BaseController
	Cntx iris.Context
}

//ShowEndpoints - show list of endpoints
func (d DashBoardController) ShowEndpoints() {
	endpoints := (models.EndPoints{}).GetWithCallSummary()

	d.Cntx.ViewData("endpoints", endpoints)
	d.Cntx.View("endpoints.html")
}
