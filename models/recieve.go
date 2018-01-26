package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Receive _
type Receive struct {
	ID            int
	Flag          int
	DocNo         string    `orm:"size(30)"`
	DocRefNo      string    `orm:"size(30)"`
	Supplier      *Supplier `orm:"rel(fk)"`
	SupplierName  string    `orm:"size(300)"`
	DiscountType  int
	DiscountWord  string  `orm:"size(300)"`
	TotalDiscount float64 `orm:"digits(12);decimals(2)"`
	TotalAmount   float64 `orm:"digits(12);decimals(2)"`
	CreditDay     int
	CreditDate    time.Time `orm:"type(date)"`
	Creator       *User     `orm:"rel(fk)"`
	CreatedAt     time.Time `orm:"auto_now_add;type(datetime)"`
	Editor        *User     `orm:"null;rel(fk)"`
	EditedAt      time.Time `orm:"null;auto_now;type(datetime)"`
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
