package controllers

import (
	"html/template"
)

//ServiceController _
type ServiceController struct {
	BaseController
}

//GetXSRF _
func (c *ServiceController) GetXSRF() {
	c.Data["json"] = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.ServeJSON()
}
