package controllers

import (
	"strconv"
	"strings"
	m "umami/models"
)

//ProductController _
type ProductController struct {
	BaseController
}

//ListProductJSON  _
func (c *ProductController) ListProductJSON() {
	term := strings.TrimSpace(c.GetString("query"))
	ret := m.NormalModel{}
	rowCount, lists, err := m.GetProductList(15, term)
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

//GetProductJSON  _
func (c *ProductController) GetProductJSON() {
	ID, _ := strconv.ParseInt(c.GetString("id"), 10, 32)
	ret := m.NormalModel{}
	product, err := m.GetProduct(int(ID))
	if err == nil {
		ret.RetOK = true
		ret.Data1 = product
	} else {
		ret.RetOK = false
		ret.RetData = "ไม่พบข้อมูล"
	}
	c.Data["json"] = ret
	c.ServeJSON()
}
