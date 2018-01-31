package controllers

import (
	"html/template"
	"path/filepath"
	"time"
	h "umami/helpers"
	m "umami/models"

	"github.com/go-playground/form"
	"github.com/google/uuid"
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
		if len(company.ImageLogo) > 0 {
			base64, _ := h.File64Encode(company.ImageLogo)
			company.ImageBase64 = base64
		}
		c.Data["m"] = company
	}
	c.Data["ret"] = m.NormalModel{}
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
	isNewImage := false
	file, header, _ := c.GetFile("ImgLogo")
	if file != nil {
		fileName := header.Filename
		fileName = uuid.New().String() + filepath.Ext(fileName)
		filePathSave := "data/company/" + fileName
		err = c.SaveToFile("ImgLogo", filePathSave)
		if err == nil {
			isNewImage = true
			sub.ImageLogo = filePathSave
			h.RemoveContentsExcludeFile("data/company", fileName)
			base64, errBase64 := h.File64Encode(filePathSave)
			err = errBase64
			sub.ImageBase64 = base64
		}
	}
	_ = file

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
		err := m.UpdateCom(sub, isNewImage)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	company, _ := m.GetComFirst()
	if company.ID == 0 {
		c.Data["title"] = "ข้อมูลร้าน/บริษัท"
	} else {
		c.Data["title"] = "แก้ไข ข้อมูลร้าน/บริษัท"
		if len(company.ImageLogo) > 0 {
			base64, _ := h.File64Encode(company.ImageLogo)
			company.ImageBase64 = base64
		}
		c.Data["m"] = company
	}
	c.Data["ret"] = ret
	c.Data["Province"] = m.GetAllProvince()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "company/com.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "company/com-script.html"
	c.Render()
}
