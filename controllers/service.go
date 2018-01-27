package controllers

import (
	"html/template"
	"strings"
	m "umami/models"
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

//ListEntityJSON  _
func (c *ServiceController) ListEntityJSON() {

	term := strings.TrimSpace(c.GetString("query"))
	entity := c.Ctx.Request.URL.Query().Get("entity")
	ret := m.NormalModel{}
	rowCount, lists, err := m.GetListEntity(entity, 15, term)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = rowCount
		ret.ListData = lists
		if rowCount == 0 {
			ret.RetOK = false
			ret.RetData = "ไม่พบข้อมูล"
		}
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}
