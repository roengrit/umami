package controllers

import (
	"html/template"
	"time"
	m "umami/models"

	"github.com/go-playground/form"
)

//ReceiveController _
type ReceiveController struct {
	BaseController
}

//Get _
func (c *ReceiveController) Get() {
	c.Data["title"] = "รับสินค้า/วัตถุดิบ"
	c.Data["CurrentDate"] = time.Now()
	c.Data["RetCount"] = 0
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "receive/receive.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "receive/receive-style.html"
	c.LayoutSections["scripts"] = "receive/receive-script.html"
	c.Render()
}

//Post _
func (c *ReceiveController) Post() {
	doc := m.Receive{}
	decoder := form.NewDecoder()
	parsFormErr := decoder.Decode(&doc, c.Ctx.Request.Form)
	if parsFormErr != nil {

	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = doc
	c.ServeJSON()
}
