package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//Order _
type Order struct {
	ID                       int
	Flag                     int
	Active                   bool
	DocType                  int
	DocNo                    string    `orm:"size(30)"`
	DocDate                  time.Time `form:"-" orm:"null"`
	DocTime                  string    `orm:"size(6)"`
	DocRefNo                 string    `orm:"size(30)"`
	TableNo                  string    `orm:"size(300)"`
	Member                   *Member   `orm:"rel(fk)"`
	MemberName               string    `orm:"size(300)"`
	DiscountType             int
	DiscountWord             string `orm:"size(300)"`
	VatType                  int
	VatWord                  string  `orm:"size(300)"`
	TotalDiscount            float64 `orm:"digits(12);decimals(2)"`
	TotalVatValue            float64 `orm:"digits(12);decimals(2)"`
	TotalAmount              float64 `orm:"digits(12);decimals(2)"`
	TotalNetAmount           float64 `orm:"digits(12);decimals(2)"`
	TotalInCludeVatNetAmount float64 `orm:"digits(12);decimals(2)"`
	CreditDay                int
	CreditDate               time.Time  `orm:"type(date)"`
	Remark                   string     `orm:"size(300)"`
	CancelRemark             string     `orm:"size(300)"`
	Creator                  *User      `orm:"rel(fk)"`
	CreatedAt                time.Time  `orm:"auto_now_add;type(datetime)"`
	Editor                   *User      `orm:"null;rel(fk)"`
	EditedAt                 time.Time  `orm:"null;auto_now;type(datetime)"`
	CancelUser               *User      `orm:"null;rel(fk)"`
	CancelAt                 time.Time  `orm:"null;type(datetime)"`
	OrderSub                 []OrderSub `orm:"-"`
}

//OrderSub _
type OrderSub struct {
	ID          int
	Flag        int
	Active      bool
	DocNo       string    `orm:"size(30)"`
	DocDate     time.Time `form:"-" orm:"null"`
	Product     *Product  `orm:"rel(fk)"`
	Unit        *Unit     `orm:"rel(fk)"`
	Qty         float64   `orm:"digits(12);decimals(2)"`
	RemainQty   float64   `orm:"digits(12);decimals(2)"`
	AverageCost float64   `orm:"digits(12);decimals(2)"`
	Price       float64   `orm:"digits(12);decimals(2)"`
	TotalPrice  float64   `orm:"digits(12);decimals(2)"`
	Creator     *User     `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Order), new(OrderSub)) // Need to register model in init
}

//CreateOrder _
func CreateOrder(order Order, user User) (retID int64, errRet error) {
	order.DocNo = GetMaxDoc("receive", "IV")
	order.Creator = &user
	order.CreatedAt = time.Now()
	order.CreditDay = 0
	order.CreditDate = time.Now()
	order.Active = true
	var fullDataSub []OrderSub
	for _, val := range order.OrderSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = order.DocNo
			val.Flag = order.Flag
			val.Active = true
			val.DocDate = order.DocDate
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&order)
	id, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	if err == nil {
		retID = id
		o.Commit()
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//UpdateOrder _
func UpdateOrder(order Order, user User) (retID int64, errRet error) {
	docCheck, _ := GetReceive(order.ID)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	order.Creator = docCheck.Creator
	order.CreatedAt = docCheck.CreatedAt
	order.CreditDay = docCheck.CreditDay
	order.CreditDate = docCheck.CreditDate
	order.EditedAt = time.Now()
	order.Editor = &user
	order.Active = docCheck.Active
	var fullDataSub []OrderSub
	for _, val := range order.OrderSub {
		if val.Product.ID != 0 {
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = order.DocNo
			val.Flag = order.Flag
			val.Active = order.Active
			val.DocDate = order.DocDate
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Update(&order)
	if err == nil {
		_, err = o.QueryTable("order_sub").Filter("doc_no", order.DocNo).Delete()
	}
	if err == nil {
		_, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	}
	if err == nil {
		retID = id
		o.Commit()
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//GetOrder _
func GetOrder(ID int) (doc *Order, errRet error) {
	order := &Order{}
	o := orm.NewOrm()
	o.QueryTable("order").Filter("ID", ID).RelatedSel().One(order)
	o.QueryTable("order_sub").Filter("doc_no", order.DocNo).RelatedSel().All(&order.OrderSub)
	doc = order
	return doc, errRet
}

//GetOrderList _
func GetOrderList(term string, limit int, dateBegin, dateEnd string) (doc *[]Order, rowCount int, errRet error) {
	order := &[]Order{}
	o := orm.NewOrm()
	qs := o.QueryTable("order")
	condSub1 := orm.NewCondition()
	condSub2 := orm.NewCondition()
	cond1 := condSub1.And("doc_date__gte", dateBegin).And("doc_date__lte", dateEnd)
	qs = qs.SetCond(cond1)
	if dateBegin != "" && dateEnd != "" {
		cond2 := condSub2.Or("Member__Name__icontains", term).Or("DocNo__icontains", term).Or("Remark__icontains", term)
		cond1 = cond1.AndCond(cond2)
		qs = qs.SetCond(cond1)
	}
	qs.RelatedSel().Limit(limit).All(order)
	return order, len(*order), errRet
}

//UpdateCancelOrder _
func UpdateCancelOrder(ID int, remark string, user User) (retID int64, errRet error) {
	docCheck := &Receive{}
	o := orm.NewOrm()
	o.QueryTable("order").Filter("ID", ID).RelatedSel().One(docCheck)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	docCheck.Active = false
	docCheck.CancelRemark = remark
	docCheck.CancelAt = time.Now()
	docCheck.CancelUser = &user
	o.Begin()
	_, err := o.Update(docCheck)
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	errRet = err
	return retID, errRet
}
