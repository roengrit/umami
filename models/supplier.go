package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Supplier _
type Supplier struct {
	ID        int
	Name      string `orm:"size(300)"`
	Address   string `orm:"size(1000)"`
	Tel       string `orm:"size(100)"`
	Email     string `orm:"size(100)"`
	Line      string `orm:"size(100)"`
	Facebook  string `orm:"size(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	orm.RegisterModel(new(Supplier)) //Need to register model in init
}
