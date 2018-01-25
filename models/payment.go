package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Payment _
type Payment struct {
	ID          int
	Flag        int
	DocNo       string       `orm:"size(30)"`
	Remark      string       `orm:"size(1000)"`
	PaymentType *PaymentType `orm:"rel(fk)"`
	Amount      float64      `orm:"digits(12);decimals(2)"`
	Creator     *User        `orm:"rel(fk)"`
	CreatedAt   time.Time    `orm:"auto_now_add;type(datetime)"`
	Editor      *User        `orm:"null;rel(fk)"`
	EditedAt    time.Time    `orm:"null;auto_now;type(datetime)"`
}

//PaymentType _
type PaymentType struct {
	ID        int
	Name      string    `orm:"size(300)"`
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Payment), new(PaymentType)) //Need to register model in init
}
