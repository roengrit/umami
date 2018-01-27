package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//Receive _
type Receive struct {
	ID             int
	Flag           int
	DocNo          string    `orm:"size(30)"`
	DocDate        time.Time `form:"-"orm:"null"`
	DocTime        string    `orm:"size(6)"`
	DocRefNo       string    `orm:"size(30)"`
	TableNo        string    `orm:"size(300)"`
	Supplier       *Supplier `orm:"rel(fk)"`
	SupplierName   string    `orm:"size(300)"`
	DiscountType   int
	DiscountWord   string  `orm:"size(300)"`
	TotalDiscount  float64 `orm:"digits(12);decimals(2)"`
	TotalAmount    float64 `orm:"digits(12);decimals(2)"`
	TotalNetAmount float64 `orm:"digits(12);decimals(2)"`
	CreditDay      int
	CreditDate     time.Time    `orm:"type(date)"`
	Remark         string       `orm:"size(300)"`
	Creator        *User        `orm:"rel(fk)"`
	CreatedAt      time.Time    `orm:"auto_now_add;type(datetime)"`
	Editor         *User        `orm:"null;rel(fk)"`
	EditedAt       time.Time    `orm:"null;auto_now;type(datetime)"`
	ReceiveSub     []ReceiveSub `orm:"-"`
}

//ReceiveSub _
type ReceiveSub struct {
	ID         int
	Flag       int
	DocNo      string    `orm:"size(30)"`
	Product    *Product  `orm:"rel(fk)"`
	Unit       *Unit     `orm:"rel(fk)"`
	Qty        float64   `orm:"digits(12);decimals(2)"`
	RemainQty  float64   `orm:"digits(12);decimals(2)"`
	Price      float64   `orm:"digits(12);decimals(2)"`
	TotalPrice float64   `orm:"digits(12);decimals(2)"`
	Creator    *User     `orm:"rel(fk)"`
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	Editor     *User     `orm:"null;rel(fk)"`
	EditedAt   time.Time `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Receive), new(ReceiveSub)) // Need to register model in init
}

//CreateReceive _
func CreateReceive(receive Receive, user User) (retID int64, errRet error) {
	receive.DocNo = GetMaxDoc("receive", "REC")
	receive.Creator = &user
	receive.CreatedAt = time.Now()
	receive.CreditDay = 0
	receive.CreditDate = time.Now()
	var fullDataSub []ReceiveSub
	for _, val := range receive.ReceiveSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = receive.DocNo
			val.Flag = receive.Flag
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&receive)
	if err == nil {
		_, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	}
	o.Commit()
	if err == nil {
		retID = id
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//UpdateReceive _
func UpdateReceive(receive Receive, user User) (retID int64, errRet error) {
	docCheck, _ := GetReceive(receive.ID)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	receive.Creator = docCheck.Creator
	receive.CreatedAt = docCheck.CreatedAt
	receive.CreditDay = docCheck.CreditDay
	receive.CreditDate = docCheck.CreditDate
	receive.EditedAt = time.Now()
	receive.Editor = &user
	var fullDataSub []ReceiveSub
	for _, val := range receive.ReceiveSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.EditedAt = time.Now()
			val.Editor = &user
			val.DocNo = receive.DocNo
			val.Flag = receive.Flag
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Update(&receive)
	if err == nil {
		_, err = o.QueryTable("receive_sub").Filter("doc_no", receive.DocNo).Delete()
	}
	if err == nil {
		_, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	}
	o.Commit()
	if err == nil {
		retID = id
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//GetReceive _
func GetReceive(ID int) (doc *Receive, errRet error) {
	receiveDoc := &Receive{}
	o := orm.NewOrm()
	o.QueryTable("receive").Filter("ID", ID).RelatedSel().One(receiveDoc)
	o.QueryTable("receive_sub").Filter("doc_no", receiveDoc.DocNo).RelatedSel().All(&receiveDoc.ReceiveSub)
	doc = receiveDoc
	return doc, errRet
}

//GetReceiveList _
func GetReceiveList(term string, limit int, dateBegin, dateEnd string) (sup *[]Receive, rowCount int, errRet error) {
	receive := &[]Receive{}
	orm.Debug = true
	o := orm.NewOrm()
	qs := o.QueryTable("receive")
	condSub1 := orm.NewCondition()
	condSub2 := orm.NewCondition()
	cond1 := condSub1.Or("Supplier__Name__icontains", term).
		Or("DocNo__icontains", term).
		Or("Remark__icontains", term)
	qs = qs.SetCond(cond1)
	if dateBegin != "" && dateEnd != "" {
		cond2 := condSub2.And("doc_date__gte", dateBegin).And("doc_date__lte", dateEnd)
		qs = qs.SetCond(cond2)
	}
	qs.RelatedSel().Limit(limit).All(receive)
	return receive, len(*receive), errRet
}
