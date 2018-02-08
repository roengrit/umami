package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//StockCountPick _
type StockCountPick struct {
	ID                int
	Flag              int
	Active            bool
	FlagTemp          int
	DocType           int
	DocNo             string    `orm:"size(30)"`
	DocDate           time.Time `form:"-" orm:"null"`
	DocTime           string    `orm:"size(6)"`
	DocRefNo          string    `orm:"size(30)"`
	TableNo           string    `orm:"size(300)"`
	Member            *Member   `orm:"rel(fk)"`
	MemberName        string    `orm:"size(300)"`
	DiscountType      int
	DiscountWord      string  `orm:"size(300)"`
	TotalDiscount     float64 `orm:"digits(12);decimals(2)"`
	TotalAmount       float64 `orm:"digits(12);decimals(2)"`
	TotalNetAmount    float64 `orm:"digits(12);decimals(2)"`
	CreditDay         int
	CreditDate        time.Time           `orm:"type(date)"`
	Remark            string              `orm:"size(300)"`
	CancelRemark      string              `orm:"size(300)"`
	Creator           *User               `orm:"rel(fk)"`
	CreatedAt         time.Time           `orm:"auto_now_add;type(datetime)"`
	Editor            *User               `orm:"null;rel(fk)"`
	EditedAt          time.Time           `orm:"null;auto_now;type(datetime)"`
	CancelUser        *User               `orm:"null;rel(fk)"`
	CancelAt          time.Time           `orm:"null;type(datetime)"`
	StockCountPickSub []StockCountPickSub `orm:"-"`
}

//StockCountPickSub _
type StockCountPickSub struct {
	ID          int
	Flag        int
	Active      bool
	DocNo       string    `orm:"size(30)"`
	DocDate     time.Time `form:"-" orm:"null"`
	Product     *Product  `orm:"rel(fk)"`
	Unit        *Unit     `orm:"rel(fk)"`
	BalanceQty  float64   `orm:"digits(12);decimals(2)"`
	Qty         float64   `orm:"digits(12);decimals(2)"`
	DiffQty     float64   `orm:"digits(12);decimals(2)"`
	RemainQty   float64   `orm:"digits(12);decimals(2)"`
	AverageCost float64   `orm:"digits(12);decimals(2)"`
	Price       float64   `orm:"digits(12);decimals(2)"`
	TotalPrice  float64   `orm:"digits(12);decimals(2)"`
	Remark      string    `orm:"size(300)"`
	Creator     *User     `orm:"rel(fk)"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(StockCountPick), new(StockCountPickSub)) // Need to register model in init
}

//CreateStockCountPick _
func CreateStockCountPick(StockCountPick StockCountPick, user User) (retID int64, errRet error) {
	StockCountPick.DocNo = GetMaxDoc("stock_count_pick", "SPK")
	StockCountPick.Creator = &user
	StockCountPick.CreatedAt = time.Now()
	StockCountPick.CreditDay = 0
	StockCountPick.CreditDate = time.Now()
	StockCountPick.Active = true
	var fullDataSub []StockCountPickSub
	for _, val := range StockCountPick.StockCountPickSub {
		if val.Product.ID != 0 {
			Product, _ := GetProduct(val.Product.ID)
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = StockCountPick.DocNo
			val.Flag = StockCountPick.Flag
			val.BalanceQty = Product.BalanceQty
			val.DiffQty = val.Qty - Product.BalanceQty
			if StockCountPick.FlagTemp == 0 {
				val.Active = true
				val.Remark = ""
			} else {
				val.Active = false
				val.Remark = "รอการปรับปรุง"
			}
			val.DocDate = StockCountPick.DocDate
			val.AverageCost = val.Price
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(&StockCountPick)
	id, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	if err == nil {
		retID = id
		o.Commit()
	} else {
		o.Rollback()
	}

	errRet = err
	return retID, errRet
}

//UpdateStockCountPick _
func UpdateStockCountPick(StockCountPick StockCountPick, user User) (retID int64, errRet error) {
	docCheck, _ := GetStockCountPick(StockCountPick.ID)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	StockCountPick.Creator = docCheck.Creator
	StockCountPick.CreatedAt = docCheck.CreatedAt
	StockCountPick.CreditDay = docCheck.CreditDay
	StockCountPick.CreditDate = docCheck.CreditDate
	StockCountPick.EditedAt = time.Now()
	StockCountPick.Editor = &user
	StockCountPick.Active = docCheck.Active
	var fullDataSub []StockCountPickSub
	for _, val := range StockCountPick.StockCountPickSub {
		if val.Product.ID != 0 {
			Product, _ := GetProduct(val.Product.ID)
			val.CreatedAt = time.Now()
			val.Creator = &user
			val.DocNo = StockCountPick.DocNo
			val.Flag = StockCountPick.Flag
			val.BalanceQty = Product.BalanceQty
			val.DiffQty = val.Qty - Product.BalanceQty
			if StockCountPick.FlagTemp == 0 {
				val.Active = true
				val.Remark = ""
			} else {
				val.Active = false
				val.Remark = "รอการปรับปรุง"
			}
			val.AverageCost = val.Price
			val.DocDate = StockCountPick.DocDate
			fullDataSub = append(fullDataSub, val)
		}
	}
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Update(&StockCountPick)
	if err == nil {
		_, err = o.QueryTable("stock_count_pick_sub").Filter("doc_no", StockCountPick.DocNo).Delete()
	}
	if err == nil {
		_, err = o.InsertMulti(len(fullDataSub), fullDataSub)
	}
	if err == nil {
		o.Commit()
		retID = id
	} else {
		o.Rollback()
	}
	errRet = err
	return retID, errRet
}

//GetStockCountPick _
func GetStockCountPick(ID int) (doc *StockCountPick, errRet error) {
	StockCountPick := &StockCountPick{}
	o := orm.NewOrm()
	o.QueryTable("stock_count_pick").Filter("ID", ID).RelatedSel().One(StockCountPick)
	o.QueryTable("stock_count_pick_sub").Filter("doc_no", StockCountPick.DocNo).RelatedSel().All(&StockCountPick.StockCountPickSub)
	doc = StockCountPick
	return doc, errRet
}

//GetStockCountPickList _
func GetStockCountPickList(term string, limit int, dateBegin, dateEnd string) (sup *[]StockCountPick, rowCount int, errRet error) {
	StockCountPick := &[]StockCountPick{}
	o := orm.NewOrm()
	qs := o.QueryTable("stock_count_pick")
	condSub1 := orm.NewCondition()
	condSub2 := orm.NewCondition()
	cond1 := condSub1.And("doc_date__gte", dateBegin).And("doc_date__lte", dateEnd)
	qs = qs.SetCond(cond1)
	if dateBegin != "" && dateEnd != "" {
		cond2 := condSub2.Or("Member__Name__icontains", term).Or("DocNo__icontains", term).Or("Remark__icontains", term)
		cond1 = cond1.AndCond(cond2)
		qs = qs.SetCond(cond1)
	}
	qs.RelatedSel().Limit(limit).All(StockCountPick)
	return StockCountPick, len(*StockCountPick), errRet
}

//UpdateCancelStockCountPick _
func UpdateCancelStockCountPick(ID int, remark string, user User) (retID int64, errRet error) {
	docCheck := &StockCount{}
	o := orm.NewOrm()
	o.QueryTable("stock_count_pick").Filter("ID", ID).RelatedSel().One(docCheck)
	if docCheck.ID == 0 {
		errRet = errors.New("ไม่พบข้อมูล")
	}
	docCheck.Active = false
	docCheck.CancelRemark = remark
	docCheck.CancelAt = time.Now()
	docCheck.CancelUser = &user
	o.Begin()
	_, err := o.Update(docCheck)
	if err == nil {
		_, err = o.Raw("update stock_count_picl_sub set active = false where doc_no = ?", docCheck.DocNo).Exec()
	}
	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	errRet = err
	return retID, errRet
}

//UpdateActiveStockCountPick _
func UpdateActiveStockCountPick(ID int, user User) (retID int64, errRet error) {
	o := orm.NewOrm()
	o.Begin()
	orm.Debug = true
	_, err := o.Raw("update stock_count_pick set active = true,flag_temp = 0,editor_id = ?,edited_at = now() where i_d = ?", user.ID, ID).Exec()
	if err != nil {
		o.Rollback()
	} else {
		_, err := o.Raw("update stock_count_pick_sub set active = true where doc_no = (select stock_count_pick.doc_no from stock_count_pick where stock_count_pick.i_d = ? limit 1)", ID).Exec()
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}
	errRet = err
	return retID, errRet
}
