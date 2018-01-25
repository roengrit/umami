package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//OrderTable _
type OrderTable struct {
	ID          int
	Name        string
	Qty         int
	InUse       bool
	Remark      string `orm:"size(300)"`
	ReserveDate time.Time
	ReserveUser *User     `orm:"rel(fk)"`
	Creator     *User     `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	Editor      *User     `orm:"rel(fk)"`
	EditAt      time.Time `orm:"auto_now;type(datetime)"`
}

//OrderTableMerg _
type OrderTableMerg struct {
	ID         int
	Parent     *OrderTable `orm:"rel(fk)"`
	ChildTable *OrderTable `orm:"rel(fk)"`
	Remark     string      `orm:"size(300)"`
	Creator    *User       `orm:"rel(fk)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	Editor     *User       `orm:"rel(fk)"`
	EditAt     time.Time   `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(OrderTable), new(OrderTableMerg)) //Need to register model in init
}
