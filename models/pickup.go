package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//PickUp _
type PickUp struct {
	ID             int
	Flag           int
	Active         bool
	DocNo          string    `orm:"size(30)"`
	DocDate        time.Time `form:"-"orm:"null"`
	DocTime        string    `orm:"size(6)"`
	DocRefNo       string    `orm:"size(30)"`
	TableNo        string    `orm:"size(300)"`
	SupplierID     int
	SupplierName   string `orm:"size(300)"`
	DiscountType   int
	DiscountWord   string  `orm:"size(300)"`
	TotalDiscount  float64 `orm:"digits(12);decimals(2)"`
	TotalAmount    float64 `orm:"digits(12);decimals(2)"`
	TotalNetAmount float64 `orm:"digits(12);decimals(2)"`
	CreditDay      int
	CreditDate     time.Time   `orm:"type(date)"`
	Remark         string      `orm:"size(300)"`
	CancelRemark   string      `orm:"size(300)"`
	Creator        *User       `orm:"rel(fk)"`
	CreatedAt      time.Time   `orm:"auto_now_add;type(datetime)"`
	Editor         *User       `orm:"null;rel(fk)"`
	EditedAt       time.Time   `orm:"null;auto_now;type(datetime)"`
	CancelUser     *User       `orm:"null;rel(fk)"`
	CancelAt       time.Time   `orm:"null;type(datetime)"`
	PickUpSub      []PickUpSub `orm:"-"`
}

//PickUpSub _
type PickUpSub struct {
	ID          int
	Flag        int
	Active      bool
	DocNo       string    `orm:"size(30)"`
	Product     *Product  `orm:"rel(fk)"`
	Unit        *Unit     `orm:"rel(fk)"`
	Qty         float64   `orm:"digits(12);decimals(2)"`
	RemainQty   float64   `orm:"digits(12);decimals(2)"`
	BalanceCost float64   `orm:"digits(12);decimals(2)"`
	Price       float64   `orm:"digits(12);decimals(2)"`
	TotalPrice  float64   `orm:"digits(12);decimals(2)"`
	Creator     *User     `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(PickUp), new(PickUpSub)) // Need to register model in init
}

//CreatePickUp _
func CreatePickUp(PickUp PickUp, user User) (retID int64, errRet error) {
	PickUp.DocNo = GetMaxDoc("pick_up", "PI")
	PickUp.Creator = &user
	PickUp.CreatedAt = time.Now()
	PickUp.CreditDay = 0
	PickUp.CreditDate = time.Now()
	PickUp.Active = true
	var fullDataSub []PickUpSub
	for _, val := range PickUp.PickUpSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = PickUp.DocNo
			val.Flag = PickUp.Flag
			val.Active = true
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&PickUp)
	id, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	o.Commit()
	if err == nil {
		retID = id
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//UpdatePickUp _
func UpdatePickUp(PickUp PickUp, user User) (retID int64, errRet error) {
	docCheck, _ := GetPickUp(PickUp.ID)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	PickUp.Creator = docCheck.Creator
	PickUp.CreatedAt = docCheck.CreatedAt
	PickUp.CreditDay = docCheck.CreditDay
	PickUp.CreditDate = docCheck.CreditDate
	PickUp.EditedAt = time.Now()
	PickUp.Editor = &user
	PickUp.Active = true
	var fullDataSub []PickUpSub
	for _, val := range PickUp.PickUpSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = PickUp.DocNo
			val.Flag = PickUp.Flag
			val.Active = true
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Update(&PickUp)
	if err == nil {
		_, err = o.QueryTable("pick_up_sub").Filter("doc_no", PickUp.DocNo).Delete()
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

//GetPickUp _
func GetPickUp(ID int) (doc *PickUp, errRet error) {
	PickUpDoc := &PickUp{}
	o := orm.NewOrm()
	o.QueryTable("pick_up").Filter("ID", ID).RelatedSel().One(PickUpDoc)
	o.QueryTable("pick_up_sub").Filter("doc_no", PickUpDoc.DocNo).RelatedSel().All(&PickUpDoc.PickUpSub)
	doc = PickUpDoc
	return doc, errRet
}

//GetPickUpList _
func GetPickUpList(term string, limit int, dateBegin, dateEnd string) (sup *[]PickUp, rowCount int, errRet error) {
	PickUp := &[]PickUp{}
	o := orm.NewOrm()
	qs := o.QueryTable("pick_up")
	condSub1 := orm.NewCondition()
	condSub2 := orm.NewCondition()
	cond1 := condSub1.And("doc_date__gte", dateBegin).And("doc_date__lte", dateEnd)
	qs = qs.SetCond(cond1)
	if dateBegin != "" && dateEnd != "" {
		cond2 := condSub2.Or("Supplier__Name__icontains", term).Or("DocNo__icontains", term).Or("Remark__icontains", term)
		cond1 = cond1.AndCond(cond2)
		qs = qs.SetCond(cond1)
	}
	qs.RelatedSel().Limit(limit).All(PickUp)
	return PickUp, len(*PickUp), errRet
}

//UpdateCancelPickUp _
func UpdateCancelPickUp(ID int, remark string, user User) (retID int64, errRet error) {
	docCheck := &PickUp{}
	o := orm.NewOrm()
	o.QueryTable("pick_up").Filter("ID", ID).RelatedSel().One(docCheck)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	docCheck.Active = false
	docCheck.CancelRemark = remark
	docCheck.CancelAt = time.Now()
	docCheck.CancelUser = &user
	o.Begin()
	_, err := o.Update(docCheck)
	o.Commit()
	if err != nil {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}
