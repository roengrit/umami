package controllers

import (
	"time"
)

//ReceiveController _
type ReceiveController struct {
	BaseController
}

//Get Home page
func (c *ReceiveController) Get() {
	c.Data["title"] = "รับสินค้า/วัตถุดิบ"
	c.Data["CurrentDate"] = time.Now()
	c.Data["RetCount"] = 0
	//c.Data["m"] = m.Order{}
	c.Layout = "layout.html"
	c.TplName = "receive/receive.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "receive/receive-style.html"
	c.LayoutSections["scripts"] = "receive/receive-script.html"
	c.Render()
}

func hello(in string) (out string) {
	out = in + "world"
	return
}
