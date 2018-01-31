package controllers

import (
	"html/template"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	h "umami/helpers"
	m "umami/models"

	"github.com/go-playground/form"
	"github.com/google/uuid"
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

//CreateProduct _
func (c *ProductController) CreateProduct() {
	proID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if proID == 0 {
		c.Data["title"] = "สร้าง สินค้า"
	} else {
		c.Data["title"] = "แก้ไข สินค้า"
		pro, _ := m.GetProduct(int(proID))
		if len(pro.ImagePath1) > 0 {
			base64, _ := h.File64Encode(pro.ImagePath1)
			pro.ImageBase64 = base64
		}
		c.Data["m"] = pro
	}
	c.Data["ret"] = m.NormalModel{}
	c.Data["ProductCategory"] = m.GetAllProductCategory()
	c.Data["Unit"] = m.GetAllProductUnit()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "product/product.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product/product-script.html"
	c.Render()
}

//UpdateProduct _
func (c *ProductController) UpdateProduct() {

	var pro m.Product
	decoder := form.NewDecoder()
	err := decoder.Decode(&pro, c.Ctx.Request.Form)
	ret := m.NormalModel{}
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))
	file, header, _ := c.GetFile("ImgProduct")
	isNewImage := false
	if file != nil {
		fileName := header.Filename
		fileName = uuid.New().String() + filepath.Ext(fileName)
		filePathSave := "data/product/" + fileName
		err = c.SaveToFile("ImgProduct", filePathSave)
		if err == nil {
			isNewImage = true
			pro.ImagePath1 = filePathSave
			base64, errBase64 := h.File64Encode(filePathSave)
			err = errBase64
			pro.ImageBase64 = base64
		}
	}
	_ = file

	ret.RetOK = true
	if err != nil {
		ret.RetOK = false
		ret.RetData = err.Error()
	} else if c.GetString("Name") == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุชื่อ"
	}

	if ret.RetOK && pro.ID == 0 {
		pro.CreatedAt = time.Now()
		pro.Creator = &actionUser
		ID, err := m.CreateProduct(pro)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			pro.ID = int(ID)
			ret.RetData = "บันทึกสำเร็จ"
		}
	} else if ret.RetOK && pro.ID > 0 {
		pro.EditedAt = time.Now()
		pro.Editor = &actionUser
		err := m.UpdateProduct(pro, isNewImage)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		} else {
			ret.RetData = "บันทึกสำเร็จ"
		}
	}
	if pro.ID == 0 {
		c.Data["title"] = "สร้าง สินค้า"
		c.Data["m"] = pro
	} else {
		c.Data["title"] = "แก้ไข สินค้า"
		pro, _ := m.GetProduct(int(pro.ID))
		if len(pro.ImagePath1) > 0 {
			base64, _ := h.File64Encode(pro.ImagePath1)
			pro.ImageBase64 = base64
		}
		c.Data["m"] = pro
	}
	c.Data["ret"] = ret
	c.Data["ProductCategory"] = m.GetAllProductCategory()
	c.Data["Unit"] = m.GetAllProductUnit()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "product/product.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["scripts"] = "product/product-script.html"
	c.Render()
}

//DeleteProduct _
func (c *ProductController) DeleteProduct() {
	ID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 32)
	ret := m.NormalModel{}
	ret.RetOK = true
	err := m.DeleteMember(int(ID))
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
