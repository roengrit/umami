package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Stock _
type Stock struct {
	ID      int
	LotDate time.Time `orm:"type(datetime)"`
	Product *Product  `orm:"rel(fk)"`
	Qty     int
}

func init() {
	orm.RegisterModel(new(Stock)) //Need to register model in init
}
