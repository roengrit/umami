package controllers

import (
	"github.com/astaxie/beego"
)

//BaseController Login Validate
type BaseController struct {
	beego.Controller
}

const active = "active menu-open"
const parentActive = "active"

//Prepare Login Validate
func (b *BaseController) Prepare() {

}
