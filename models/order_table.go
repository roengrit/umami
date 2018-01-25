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
	ReserveUser *User `orm:"rel(fk)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderTableMerg struct {
	ID         int
	Parent     *OrderTable `orm:"rel(fk)"`
	ChildTable *OrderTable `orm:"rel(fk)"`
	Remark     string      `orm:"size(300)"`
	MergUser   *User       `orm:"rel(fk)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func init() {
	orm.RegisterModel(new(OrderTable), new(OrderTableMerg)) //Need to register model in init
}
