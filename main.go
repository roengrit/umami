package main

import (
	"fmt"
	c "umami/controllers"
	h "umami/helpers"

	_ "umami/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "host=localhost port=5432 user=postgres password=P@ssw0rd dbname=umami sslmode=disable")
}

func main() {

	name := "default"
	force := false                             // Drop table and re-create.
	verbose := true                            // Print log.
	err := orm.RunSyncdb(name, force, verbose) // Error.

	if err != nil {
		fmt.Println(err)
	}

	beego.Router("/", &c.AppController{})
	beego.Router("/service/secure/json/", &c.ServiceController{}, "get:GetXSRF")
	beego.Router("/service/entity/list/json", &c.ServiceController{}, "get:ListEntityJSON")
	beego.Router("/product/list/json", &c.ProductController{}, "get:ListProductJSON")
	beego.Router("/product/json", &c.ProductController{}, "get:GetProductJSON")

	beego.Router("/supplier/?:id", &c.SupplierController{}, "get:CreateSupplier;post:UpdateSupplier;delete:DeleteSupplier")
	beego.Router("/supplier/read/?:id", &c.SupplierController{}, "get:CreateSupplier")
	beego.Router("/supplier/list", &c.SupplierController{}, "get:SupplierList;post:GetSupplierList")
	beego.Router("/receive", &c.ReceiveController{})
	beego.Router("/receive/read", &c.ReceiveController{})
	beego.Router("/receive/cancel", &c.ReceiveController{}, "get:CancelReceive;post:UpdateCancelReceive")
	beego.Router("/receive/list", &c.ReceiveController{}, "get:ReceiveList;post:GetReceiveList")
	beego.AddFuncMap("ThCommaSep", h.ThCommaSep)
	beego.AddFuncMap("TextThCommaSep", h.TextThCommaSep)
	beego.Run()
}
