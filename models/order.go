package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Order _
type Order struct {
	ID             int
	Flag           int
	DocNo          string  `orm:"size(30)"`
	DocTime        string  `orm:"size(6)"`
	DocRefNo       string  `orm:"size(30)"`
	TableNo        string  `orm:"size(300)"`
	Member         *Member `orm:"rel(fk)"`
	MemberName     string  `orm:"size(300)"`
	DiscountType   int
	DiscountWord   string  `orm:"size(300)"`
	TotalDiscount  float64 `orm:"digits(12);decimals(2)"`
	TotalAmount    float64 `orm:"digits(12);decimals(2)"`
	TotalNetAmount float64 `orm:"digits(12);decimals(2)"`
	CreditDay      int
	CreditDate     time.Time  `orm:"type(date)"`
	Creator        *User      `orm:"rel(fk)"`
	CreatedAt      time.Time  `orm:"auto_now_add;type(datetime)"`
	Editor         *User      `orm:"null;rel(fk)"`
	EditedAt       time.Time  `orm:"null;auto_now;type(datetime)"`
	OrderSub       []OrderSub `orm:"-"`
}

//OrderSub _
type OrderSub struct {
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
	orm.RegisterModel(new(Order), new(OrderSub)) // Need to register model in init
}
