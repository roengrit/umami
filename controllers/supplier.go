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

//SupplierController _
type SupplierController struct {
	BaseController
}

//CreateSupplier _
func (c *SupplierController) CreateSupplier() {
	supID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if supID == 0 {
		c.Data["title"] = "สร้าง ร้านค้า/Supplier"
	} else {
		c.Data["title"] = "แก้ไข ร้านค้า/Supplier"
		sup, _ := m.GetSupplier(int(supID))
		c.Data["m"] = sup
	} 
	c.Data["Province"] = m.GetAllProvince()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "supplier/sup.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "supplier/sup-script.html"
	c.Render()
}

//UpdateSupplier _
func (c *SupplierController) UpdateSupplier() {

	var sub m.Supplier
	decoder := form.NewDecoder()
	err := decoder.Decode(&sub, c.Ctx.Request.Form)
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

	if ret.RetOK && sub.ID == 0 {
		sub.CreatedAt = time.Now()
		sub.Creator = &actionUser
		_, err := m.CreateSupplier(sub)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && sub.ID > 0 {
		sub.EditedAt = time.Now()
		sub.Editor = &actionUser
		err := m.UpdateSupplier(sub)
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

//DeleteSupplier _
func (c *SupplierController) DeleteSupplier() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteSupplier(int(ID))
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

//SupplierList _
func (c *SupplierController) SupplierList() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "ร้านค้า/Supplier"
	c.Layout = "layout.html"
	c.TplName = "supplier/sup-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "supplier/sup-list-script.html"
	c.Render()
}

//GetSupplierList _
func (c *SupplierController) GetSupplierList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	lists, rowCount, err := m.GetSupplierList(term, int(top))
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenSupHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLSupNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLSupError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}
