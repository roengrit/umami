package controllers

import (
	"html/template"
	"strconv"
	"strings"
	"time"
	m "umami/models"
)

//PickUpController _
type PickUpController struct {
	BaseController
}

//Get _
func (c *PickUpController) Get() {
	docID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if docID == 0 {
		c.Data["title"] = "เบิกสินค้า/วัตถุดิบ"
	} else {
		doc, _ := m.GetPickUp(int(docID))
		c.Data["m"] = doc
		if !doc.Active {
			c.Data["r"] = "readonly"
		}
		c.Data["RetCount"] = len(doc.PickUpSub)
		c.Data["title"] = "แก้ไข เบิกสินค้า/วัตถุดิบ : " + doc.DocNo
	}
	c.Data["CurrentDate"] = time.Now()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "pickup/pickup.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "pickup/pickup-style.html"
	c.LayoutSections["scripts"] = "pickup/pickup-script.html"
	c.Render()
}
