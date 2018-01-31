package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Product _
type Product struct {
	ID              int
	Lock            bool
	Name            string           `orm:"size(300)"`
	Barcode         string           `orm:"size(13)"`
	AverageCost     float64          `orm:"digits(12);decimals(2)"`
	BalanceQty      float64          `orm:"digits(12);decimals(2)"`
	SalePrice       float64          `orm:"digits(12);decimals(2)"`
	Unit            *Unit            `orm:"rel(fk)"`
	ProductCategory *ProductCategory `orm:"rel(fk)"`
	ProductType     int
	ImagePath1      string `orm:"size(300)"`
	ImageBase64     string `orm:"-"`
	Remark          string `orm:"size(100)"`
	FixCost          bool
	Active           bool
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

// //ProductType _
// type ProductType struct {
// 	ID        int
// 	Lock      bool
// 	Name      string `orm:"size(300)"`
// 	Active    bool
// 	Creator   *User     `orm:"rel(fk)"`
// 	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
// 	Editor    *User     `orm:"null;rel(fk)"`
// 	EditedAt  time.Time `orm:"null;auto_now;type(datetime)"`
// }

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
	orm.RegisterModel(new(Product), new(ProductCategory), new(Unit)) // Need to register model in init
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
func GetProduct(ID int) (pro *Product, errRet error) {
	Product := &Product{}
	o := orm.NewOrm()
	o.QueryTable("product").Filter("ID", ID).RelatedSel().One(Product)
	return Product, errRet
}

//GetAllProductCategory _
func GetAllProductCategory() (pro *[]ProductCategory) {
	ProductCategory := &[]ProductCategory{}
	o := orm.NewOrm()
	o.QueryTable("product_category").RelatedSel().All(ProductCategory)
	return ProductCategory
}

//GetAllProductUnit _
func GetAllProductUnit() (pro *[]Unit) {
	Unit := &[]Unit{}
	o := orm.NewOrm()
	o.QueryTable("unit").RelatedSel().All(Unit)
	return Unit
}

//CreateProduct _
func CreateProduct(pro Product) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&pro)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateProduct _
func UpdateProduct(pro Product, isNewImage bool) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProduct(pro.ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else if errRet == nil {
		if !isNewImage {
			pro.ImagePath1 = getUpdate.ImagePath1
		}
		pro.BalanceQty = getUpdate.BalanceQty
		pro.AverageCost = getUpdate.AverageCost
		pro.CreatedAt = getUpdate.CreatedAt
		pro.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&pro); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteProduct _
func DeleteProduct(ID int) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetProduct(ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if num, errDelete := o.Delete(&Product{ID: ID}); errDelete != nil && errRet == nil {
		errRet = errDelete
		_ = num
	}
	return errRet
}

//GetManagmentProductList _
func GetManagmentProductList(term string, limit int) (pro *[]Product, rowCount int, errRet error) {
	productlist := &[]Product{}
	o := orm.NewOrm()
	qs := o.QueryTable("product")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term).
		Or("Remark__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(productlist)
	return productlist, len(*productlist), errRet
}
