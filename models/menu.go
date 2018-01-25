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
	Editor    *User     `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Menu)) //Need to register model in init
}
