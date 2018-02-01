package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//OrderTable _
type OrderTable struct {
	ID          int
	Lock        bool
	Name        string
	Qty         int
	InUse       bool
	Remark      string    `orm:"size(300)"`
	ReserveDate time.Time `orm:"null"`
	ReserveTime string    `orm:"size(5)"`
	ReserveUser string    `orm:"size(300)"`
	Creator     *User     `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	Editor      *User     `orm:"null;rel(fk)"`
	EditedAt    time.Time `orm:"null;auto_now;type(datetime)"`
}

//OrderTableMerg _
type OrderTableMerg struct {
	ID         int
	Lock       bool
	Parent     *OrderTable `orm:"rel(fk)"`
	ChildTable *OrderTable `orm:"rel(fk)"`
	Remark     string      `orm:"size(300)"`
	Creator    *User       `orm:"rel(fk)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	Editor     *User       `orm:"null;rel(fk)"`
	EditedAt   time.Time   `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(OrderTable), new(OrderTableMerg)) //Need to register model in init
}

//GetOrderTable _
func GetOrderTable(ID int) (table *OrderTable, errRet error) {
	Unit := &OrderTable{}
	o := orm.NewOrm()
	o.QueryTable("order_table").Filter("ID", ID).RelatedSel().One(Unit)
	return Unit, errRet
}

//CreateOrderTable _
func CreateOrderTable(table OrderTable) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&table)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateOrderTable _
func UpdateOrderTable(table OrderTable) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetOrderTable(table.ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else if errRet == nil {
		table.Remark = getUpdate.Remark
		table.ReserveDate = getUpdate.ReserveDate
		table.ReserveUser = getUpdate.ReserveUser
		table.ReserveTime = getUpdate.ReserveTime
		table.InUse = getUpdate.InUse
		table.CreatedAt = getUpdate.CreatedAt
		table.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&table); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//DeleteOrderTable _
func DeleteOrderTable(ID int) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetOrderTable(ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if num, errDelete := o.Delete(&OrderTable{ID: ID}); errDelete != nil && errRet == nil {
		errRet = errDelete
		_ = num
	}
	return errRet
}

//GetOrderTableList _
func GetOrderTableList(term string, limit int) (tableOut *[]OrderTable, rowCount int, errRet error) {
	table := &[]OrderTable{}
	o := orm.NewOrm()
	qs := o.QueryTable("order_table")
	cond := orm.NewCondition()
	cond1 := cond.Or("Name__icontains", term)
	qs.SetCond(cond1).RelatedSel().Limit(limit).All(table)
	return table, len(*table), errRet
}
