package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//Company _
type Company struct {
	ID          int
	Lock        bool
	Name        string     `orm:"size(300)"`
	Address     string     `orm:"size(300)"`
	Province    *Provinces `orm:"rel(fk)"`
	PostCode    string     `orm:"size(10)"`
	Contact     string     `orm:"size(255)"`
	Tel         string     `orm:"size(100)"`
	ImageLogo   string     `orm:"size(300)"`
	ImageBase64 string     `orm:"-"`
	Remark      string     `orm:"size(100)"`
	Creator     *User      `orm:"rel(fk)"`
	CreatedAt   time.Time  `orm:"auto_now_add;type(datetime)"`
	Editor      *User      `orm:"null;rel(fk)"`
	EditedAt    time.Time  `orm:"null;auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Company)) //Need to register model in init
}

//CreateCom _
func CreateCom(company Company) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&company)
	o.Commit()
	if err == nil {
		retID = id
	}
	return retID, err
}

//UpdateCom _
func UpdateCom(company Company, isNewImage bool) (errRet error) {
	o := orm.NewOrm()
	getUpdate, _ := GetCom(company.ID)
	if getUpdate.Lock {
		errRet = errors.New("ข้อมูลถูก Lock ไม่สามารถแก้ไขได้")
	}
	if getUpdate == nil {
		errRet = errors.New("ไม่พบข้อมูล")
	} else if errRet == nil {
		if !isNewImage {
			company.ImageLogo = getUpdate.ImageLogo
		}
		company.CreatedAt = getUpdate.CreatedAt
		company.Creator = getUpdate.Creator
		if num, errUpdate := o.Update(&company); errUpdate != nil {
			errRet = errUpdate
			_ = num
		}
	}
	return errRet
}

//GetCom _
func GetCom(ID int) (mem *Company, errRet error) {
	company := &Company{}
	o := orm.NewOrm()
	o.QueryTable("company").Filter("ID", ID).RelatedSel().One(company)
	return company, errRet
}

//GetComFirst _
func GetComFirst() (mem *Company, errRet error) {
	company := &Company{}
	o := orm.NewOrm()
	o.QueryTable("company").RelatedSel().One(company)
	return company, errRet
}
