package main

import (
	"fmt"
	c "umami/controllers"

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
	force := true                              // Drop table and re-create.
	verbose := true                            // Print log.
	err := orm.RunSyncdb(name, force, verbose) // Error.

	if err != nil {
		fmt.Println(err)
		fmt.Println("regis model")
	}

	beego.Router("/", &c.AppController{})
	beego.Run()
}
