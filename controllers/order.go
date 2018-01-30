package controllers

import (
	"bytes"
	"html/template"
	"strconv"
	"strings"
	"time"
	h "umami/helpers"
	m "umami/models"

	"github.com/go-playground/form"
)

//OrderController _
type OrderController struct {
	BaseController
}

//Get _
func (c *OrderController) Get() {
	docID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	if strings.Contains(c.Ctx.Request.URL.RequestURI(), "read") {
		c.Data["r"] = "readonly"
	}
	if docID == 0 {
		c.Data["title"] = "ขาย"
	} else {
		doc, _ := m.GetOrder(int(docID))
		c.Data["m"] = doc
		if !doc.Active {
			c.Data["r"] = "readonly"
		}
		c.Data["RetCount"] = len(doc.OrderSub)
		c.Data["title"] = "แก้ไข ขาย : " + doc.DocNo
	}
	c.Data["CurrentDate"] = time.Now()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layout.html"
	c.TplName = "order/order.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "order/order-style.html"
	c.LayoutSections["scripts"] = "order/order-script.html"
	c.Render()
}

//Post _
func (c *OrderController) Post() {
	doc := m.Order{}
	doc.Flag = 3 // ขาย
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
			_, parsFormErr = m.CreateOrder(doc, actionUser)
			if parsFormErr == nil {
				retJSON.RetOK = true
				retJSON.RetData = "บันทึกสำเร็จ"
			} else {
				retJSON.RetOK = false
				retJSON.RetData = parsFormErr.Error()
			}
		}
		if retJSON.RetOK && doc.ID != 0 {
			_, parsFormErr = m.UpdateOrder(doc, actionUser)
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

//OrderList _
func (c *OrderController) OrderList() {
	c.Data["beginDate"] = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["endDate"] = time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "ขาย"
	c.Layout = "layout.html"
	c.TplName = "order/order-list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["html_head"] = "order/order-style.html"
	c.LayoutSections["scripts"] = "order/order-list-script.html"
	c.Render()
}

//GetOrderList _
func (c *OrderController) GetOrderList() {
	ret := m.NormalModel{}
	ret.RetOK = true
	top, _ := strconv.ParseInt(c.GetString("top"), 10, 32)
	term := c.GetString("txt-search")
	dateBegin := c.GetString("txt-date-begin")
	dateEnd := c.GetString("txt-date-end")
	if dateBegin != "" {
		sp := strings.Split(dateBegin, "-")
		dateBegin = sp[2] + "-" + sp[1] + "-" + sp[0]
	}
	if dateEnd != "" {
		sp := strings.Split(dateEnd, "-")
		dateEnd = sp[2] + "-" + sp[1] + "-" + sp[0]
	}
	lists, rowCount, err := m.GetOrderList(term, int(top), dateBegin, dateEnd)
	if err == nil {
		ret.RetOK = true
		ret.RetCount = int64(rowCount)
		ret.RetData = h.GenOrderHTML(*lists)
		if rowCount == 0 {
			ret.RetData = h.HTMLOrderNotFoundRows
		}
	} else {
		ret.RetOK = false
		ret.RetData = strings.Replace(h.HTMLOrderError, "{err}", err.Error(), -1)
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//CancelOrder _
func (c *OrderController) CancelOrder() {
	ID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	ret := m.NormalModel{}
	dataTemplate := m.NormalModel{}
	dataTemplate.ID = ID
	dataTemplate.Title = "กรุณาระบุ หมายเหตุ การยกเลิก"
	dataTemplate.XSRF = c.XSRFToken()
	t, err := template.ParseFiles("views/order/order-cancel.html")
	var tpl bytes.Buffer
	if err = t.Execute(&tpl, dataTemplate); err != nil {
		ret.RetOK = err != nil
		ret.RetData = err.Error()
	} else {
		ret.RetOK = true
		ret.RetData = tpl.String()
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//UpdateCancelOrder _
func (c *OrderController) UpdateCancelOrder() {
	actionUser, _ := m.GetUser(h.GetUser(c.Ctx.Request))
	ret := m.NormalModel{}
	ID, _ := c.GetInt("ID")
	remark := c.GetString("Remark")
	ret.RetOK = true
	if remark == "" {
		ret.RetOK = false
		ret.RetData = "กรุณาระบุหมายเหตุ"
	}
	if ret.RetOK {
		_, err := m.UpdateCancelOrder(ID, remark, actionUser)
		if err != nil {
			ret.RetOK = false
			ret.RetData = err.Error()
		}
	}
	ret.XSRF = c.XSRFToken()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["json"] = ret
	c.ServeJSON()
}

//Print _
func (c *OrderController) Print() {
	docID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("id"), 10, 32)
	doc, _ := m.GetOrder(int(docID))
	c.Data["m"] = doc
	if !doc.Active {
		c.Data["r"] = "readonly"
	}
	c.Data["RetCount"] = len(doc.OrderSub)
	c.Data["title"] = "พิมพ์ : " + doc.DocNo
	c.Data["CurrentDate"] = time.Now()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "order/invoice.html"
	c.Render()
}
