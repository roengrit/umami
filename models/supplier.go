package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//Supplier _
type Supplier struct {
	ID        int
	Lock      bool
	Name      string     `orm:"size(300)"`
	Address   string     `orm:"size(300)"`
	Province  *Provinces `orm:"rel(fk)"`
	PostCode  string     `orm:"size(10)"`
	Contact   string     `orm:"size(255)"`
	Tel       string     `orm:"size(100)"`
	Remark    string     `orm:"size(100)"`
	Creator   *User      `orm:"rel(fk)"`
	CreatedAt time.Time  `orm:"auto_now_add;type(datetime)"`
	Editor    *User      `orm:"null;rel(fk)"`
	EditedAt  time.Time  `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Supplier)) //Need to register model in init
}

//CreateSupplier _
func CreateSupplier(sup Supplier) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&sup)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateSupplier _
func UpdateSupplier(sup Supplier) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetSupplier(sup.ID)
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else {
		sup.CreatedAt = getUpdate.CreatedAt
		sup.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&sup); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//GetSupplier _
func GetSupplier(ID int) (sup *Supplier, errRet error) {
	supplier := &Supplier{}
	o := orm.NewOrm()
	o.QueryTable("supplier").Filter("ID", ID).RelatedSel().One(supplier)
	return supplier, errRet
}

//GetSupplierList _
func GetSupplierList(term string, limit int) (sup *[]Supplier, rowCount int, errRet error) {
	supplier := &[]Supplier{}
	o := orm.NewOrm()
	qs := o.QueryTable("supplier")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term).
		Or("Tel__icontains", term).
		Or("Contact__icontains", term).
		Or("Remark__icontains", term).
		Or("Address__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(supplier)
	return supplier, len(*supplier), errRet
}

//DeleteSupplier _
func DeleteSupplier(ID int) (errRet error) {
	o := orm.NewOrm()
	if num, errDelete := o.Delete(&Supplier{ID: ID}); errDelete != nil {
		errRet = errDelete
		_ = num
	}
	return errRet
}
