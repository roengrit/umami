package controllers

import (
	"html/template"
	"strconv"
	"strings"
	"time"
	h "umami/helpers"
	m "umami/models"

	"github.com/go-playground/form"
)

//TableController _
type TableController struct {
	BaseController
}

//CreateTable _
func (c *TableController) CreateTable() {
	unitID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if unitID == 0 {
		c.Data["title"] = "สร้างโต๊ะ"
	} else {
		c.Data["title"] = "แก้ไขโต๊ะ"
		unit, _ := m.GetOrderTable(int(unitID))
		c.Data["m"] = unit
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "order-table/table.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "order-table/table-script.html"
	c.Render()
}

//UpdateTable _
func (c *TableController) UpdateTable() {

	var table m.OrderTable
	decoder := form.NewDecoder()
	err := decoder.Decode(&table, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))

	ret.RetOK = true
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if c.GetString("Name") == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && table.ID == 0 {
		table.CreatedAt = time.Now()
		table.Creator = &actionUser
		_, err := m.CreateOrderTable(table)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && table.ID > 0 {
		table.EditedAt = time.Now()
		table.Editor = &actionUser
		err := m.UpdateOrderTable(table)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//DeleteTable _
func (c *TableController) DeleteTable() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteOrderTable(int(ID))
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else {
		ret.RetData = "ลบข้อมูลสำเร็จ"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//TableList _
func (c *TableController) TableList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "โต๊ะ"
	c.Layout = "layout.html"
	c.TplName = "order-table/table-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "order-table/table-list-script.html"
	c.Render()
}

//GetTableList _
func (c *TableController) GetTableList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetOrderTableList(term, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenTableHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLTableNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLTableError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
