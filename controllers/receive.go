package controllers

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"
	h "umami/helpers"
	m "umami/models"

	"github.com/go-playground/form"
)

//ReceiveController _
type ReceiveController struct {
	BaseController
}

//Get _
func (c *ReceiveController) Get() {

	docID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if docID == 0 {
		c.Data["title"] = "รับสินค้า/วัตถุดิบ"
	} else {
		doc, _ := m.GetReceive(int(docID))
		fmt.Println(doc.ReceiveSub)
		fmt.Println(len(doc.ReceiveSub))
		c.Data["m"] = doc
		c.Data["RetCount"] = len(doc.ReceiveSub)
		c.Data["title"] = "แก้ไข รับสินค้า/วัตถุดิบ : " + doc.DocNo
	}
	c.Data["CurrentDate"] = time.Now()
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
	doc.Flag = 1 // รับ
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))
	retJSON := m.NormalModel{RetOK: true}
	decoder := form.NewDecoder()
	parsFormErr := decoder.Decode(&doc, c.Ctx.Request.Form)
	if parsFormErr == nil {
		if docDate, err := h.ValidateDate(c.GetString("DocDate")); err == nil {
			doc.DocDate = docDate
		} else {
			retJSON.RetOK = false
			retJSON.RetData = "มีข้อมูลบางอย่างไม่ครบถ้วน"
		}
		if retJSON.RetOK && doc.ID == 0 {
			_, parsFormErr = m.CreateReceive(doc, actionUser)
			if parsFormErr == nil {
				retJSON.RetOK = true
				retJSON.RetData = "บันทึกสำเร็จ"
			} else {
				retJSON.RetOK = false
				retJSON.RetData = parsFormErr.Error()
			}
		}
		if retJSON.RetOK && doc.ID != 0 {
			_, parsFormErr = m.UpdateReceive(doc, actionUser)
			if parsFormErr == nil {
				retJSON.RetOK = true
				retJSON.RetData = "บันทึกสำเร็จ"
			} else {
				retJSON.RetOK = false
				retJSON.RetData = parsFormErr.Error()
			}
		}
	} else {
		retJSON.RetOK = false
		retJSON.RetData = "มีข้อมูลบางอย่างไม่ครบถ้วน"
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = retJSON
	c.ServeJSON()
}
