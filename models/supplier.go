package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Supplier _
type Supplier struct {
	ID        int
	Name      string    `orm:"size(300)"`
	Address   string    `orm:"size(1000)"`
	Tel       string    `orm:"size(100)"`
	Email     string    `orm:"size(100)"`
	Line      string    `orm:"size(100)"`
	Facebook  string    `orm:"size(100)"`
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Supplier)) //Need to register model in init
}
