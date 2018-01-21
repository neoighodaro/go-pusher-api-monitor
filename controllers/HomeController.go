package controllers

import (
	"fmt"

	"github.com/kataras/iris/mvc"
)

//HomeController - our home controller
type HomeController struct {
	mvc.BaseController
}

//Show - Render the home page
func (h HomeController) Show() {
	fmt.Println("Welcome!")
}

//ShowDashboard - Render the API monitor dashboard
func (h HomeController) ShowDashboard() {

}
