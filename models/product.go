package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Product _
type Product struct {
	ID              int
	Lock            bool
	Name            string           `orm:"size(300)"`
	BalanceCost     float64          `orm:"digits(12);decimals(2)"`
	SalePrice       float64          `orm:"digits(12);decimals(2)"`
	Unit            *Unit            `orm:"rel(fk)"`
	ProductCategory *ProductCategory `orm:"rel(fk)"`
	ProductType     *ProductType     `orm:"rel(fk)"`
	ImagePath       string           `orm:"size(300)"`
	Active          bool
	Creator         *User     `orm:"rel(fk)"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	Editor          *User     `orm:"rel(fk)"`
	EditAt          time.Time `orm:"auto_now;type(datetime)"`
}

//ProductCategory _
type ProductCategory struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"rel(fk)"`
	EditAt    time.Time `orm:"auto_now;type(datetime)"`
}

//ProductType _
type ProductType struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"rel(fk)"`
	EditAt    time.Time `orm:"auto_now;type(datetime)"`
}

//Unit _
type Unit struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"rel(fk)"`
	EditAt    time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Product), new(ProductCategory), new(ProductType), new(Unit)) // Need to register model in init
}
