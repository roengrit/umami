package models

import (
	"github.com/astaxie/beego/orm"
)

//StockAdj _
type StockAdj struct {
	ID      int
	Flag    int      //0 not process //1 in process
	Product *Product `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(StockAdj)) // Need to register model in init
}

//ProcessAvgVal _
func ProcessAvgVal() {

}
