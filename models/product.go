package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Product _
type Product struct {
	ID              int
	Lock            bool
	Name            string `orm:"size(300)"`
	AVerageCostType int
	AVerageCost     float64          `orm:"digits(12);decimals(2)"`
	BalanceCost     float64          `orm:"digits(12);decimals(2)"`
	SalePrice       float64          `orm:"digits(12);decimals(2)"`
	Unit            *Unit            `orm:"rel(fk)"`
	ProductCategory *ProductCategory `orm:"rel(fk)"`
	ProductType     *ProductType     `orm:"rel(fk)"`
	ImagePath1      string           `orm:"size(300)"`
	ImagePath2      string           `orm:"size(300)"`
	ImagePath3      string           `orm:"size(300)"`
	Active          bool
	Creator         *User     `orm:"rel(fk)"`
	CreatedAt       time.Time `orm:"auto_now_add;type(datetime)"`
	Editor          *User     `orm:"null;rel(fk)"`
	EditedAt        time.Time `orm:"null;auto_now;type(datetime)"`
}

//ProductCategory _
type ProductCategory struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

//ProductType _
type ProductType struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

//Unit _
type Unit struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(300)"`
	Active    bool
	Creator   *User     `orm:"rel(fk)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	Editor    *User     `orm:"null;rel(fk)"`
	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Product), new(ProductCategory), new(ProductType), new(Unit)) // Need to register model in init
}

//GetProductList _
func GetProductList(top int, term string) (num int64, productList []Product, err error) {
	var sql = `SELECT T0.i_d,T0.name,T0.lock, T1.i_d as unit_id,T1.name as unit_name
			   FROM product T0	
			   JOIN unit T1 ON T0.unit_id = T1.i_d		    
			   WHERE lower(T0.name) like lower(?) order by T0.name limit {0}`
	if top == 0 {
		sql = strings.Replace(sql, "limit {0}", "", -1)
	} else {
		sql = strings.Replace(sql, "{0}", strconv.Itoa(top), -1)
	}
	o := orm.NewOrm()
	num, err = o.Raw(sql, "%"+term+"%").QueryRows(&productList)
	return num, productList, err
}

//GetProduct _
func GetProduct(ID int) (sup *Product, errRet error) {
	supplier := &Product{}
	o := orm.NewOrm()
	o.QueryTable("product").Filter("ID", ID).RelatedSel().One(supplier)
	supplier.Creator = nil
	supplier.Editor = nil
	return supplier, errRet
}
