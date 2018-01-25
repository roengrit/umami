package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Menu _
type Menu struct {
	ID        int
	Name      string
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"rel(fk)"`
	EditAt    time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Menu)) //Need to register model in init
}
