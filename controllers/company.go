package controllers

import (
	"html/template"
	"time"
	h "umami/helpers"
	m "umami/models"

	"github.com/go-playground/form"
)

//CompanyController _
type CompanyController struct {
	BaseController
}

//CreateCom _
func (c *CompanyController) CreateCom() {
	company, _ := m.GetComFirst()
	if company.ID == 0 {
		c.Data["title"] = "ข้อมูลร้าน/บริษัท"
	} else {
		c.Data["title"] = "แก้ไข ข้อมูลร้าน/บริษัท"
		mem, _ := m.GetCom(int(company.ID))
		c.Data["m"] = mem
	}
	c.Data["Province"] = m.GetAllProvince()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "company/com.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "company/com-script.html"
	c.Render()
}

//UpdateCom _
func (c *CompanyController) UpdateCom() {

	var sub m.Company
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
		_, err := m.CreateCom(sub)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && sub.ID > 0 {
		sub.EditedAt = time.Now()
		sub.Editor = &actionUser
		err := m.UpdateCom(sub)
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
