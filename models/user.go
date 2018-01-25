package models

import (
	"errors"
	"math/rand"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

//User _
type User struct {
	ID        int
	Lock      bool
	Username  string `orm:"size(50)"`
	Password  string `orm:"size(255)"`
	Name      string `orm:"size(500)"`
	Tel       string `orm:"size(255)"`
	Facebook  string `orm:"size(255)"`
	Line      string `orm:"size(255)"`
	Role      *Role  `orm:"rel(fk)"`
	Active    bool
	Creator    *User       `orm:"rel(fk)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	Editor     *User       `orm:"rel(fk)"`
	EditAt     time.Time   `orm:"auto_now;type(datetime)"`
}

//Permiss เก็บข้อมูลสิทธิ์ใช้งาน
type Permiss struct {
	ID        int
	RoleID    *Role `orm:"rel(fk)"`
	MenuID    *Menu `orm:"rel(fk)"`
	Active    bool
	Creator    *User       `orm:"rel(fk)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	Editor     *User       `orm:"rel(fk)"`
	EditAt     time.Time   `orm:"auto_now;type(datetime)"`
}

//Role _
type Role struct {
	ID        int
	Lock      bool
	Name      string `orm:"size(225)"`
	Creator    *User       `orm:"rel(fk)"`
	CreatedAt  time.Time   `orm:"auto_now_add;type(datetime)"`
	Editor     *User       `orm:"rel(fk)"`
	EditAt     time.Time   `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User), new(Permiss), new(Role)) // Need to register model in init
}

//Login _
func Login(username, password string) (ok bool, errRet string) {
	o := orm.NewOrm()
	user := User{Username: username}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		errRet = "รหัสผู้ใช้/รหัสผ่านไม่ถูกต้อง"
	} else {
		if errCript := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errCript != nil {
			errRet = "รหัสผู้ใช้/รหัสผ่านไม่ถูกต้อง"
		} else {
			ok = true
		}
	}
	return ok, errRet
}

//GetUser _
func GetUser(username string) (ok bool, errRet error) {
	o := orm.NewOrm()
	user := User{Username: username}
	errRet = o.Read(&user, "Username")
	if errRet == orm.ErrNoRows {
		errRet = errors.New("ไม่พบผู้ใช้งานนี้ในระบบ")
	} else {
		ok = true
	}
	return ok, errRet
}

//GetUserByID _
func GetUserByID(ID int) (user *User, errRet error) {
	o := orm.NewOrm()
	userGet := &User{}
	o.QueryTable("users").Filter("ID", ID).RelatedSel().One(userGet)
	if nil != userGet {
		userGet.Password = ""
	} else {
		errRet = errors.New("ไม่พบผู้ใช้งานนี้ในระบบ")
	}
	return userGet, errRet
}

//GetUserByUserName _
func GetUserByUserName(username string) (user *User, errRet error) {
	userGet := &User{}
	o := orm.NewOrm()
	o.QueryTable("users").Filter("Username", username).RelatedSel().One(userGet)
	if nil != userGet {
		userGet.Password = ""
	} else {
		errRet = errors.New("ไม่พบผู้ใช้งานนี้ในระบบ")
	}
	return userGet, errRet
}

//ForgetPass _
func ForgetPass(username, newpass string) (ok bool, errRet error) {
	o := orm.NewOrm()
	user := User{Username: username}
	errRet = o.Read(&user, "Username")
	if errRet == orm.ErrNoRows {
		errRet = errors.New("ไม่พบผู้ใช้งานนี้ในระบบ")
	} else {
		if hash, err := bcrypt.GenerateFromPassword([]byte(newpass), bcrypt.DefaultCost); err != nil {
			errRet = err
		} else {
			user.Password = string(hash)
			if _, errUpdate := o.Update(&user); errUpdate != nil {
				errRet = errUpdate
			} else {
				ok = true
			}
		}
	}
	return ok, errRet
}

//ChangePass _
func ChangePass(username, newpass string) (ok bool, errRet error) {
	o := orm.NewOrm()
	user := User{Username: username}
	errRet = o.Read(&user, "Username")
	if errRet == orm.ErrNoRows {
		errRet = errors.New("ไม่พบผู้ใช้งานนี้ในระบบ")
	} else {
		if hash, err := bcrypt.GenerateFromPassword([]byte(newpass), bcrypt.DefaultCost); err != nil {
			errRet = err
		} else {
			user.Password = string(hash)
			if num, errUpdate := o.Update(&user); errUpdate != nil {
				errRet = errUpdate
				_ = num
			} else {
				ok = true
			}
		}
	}
	return ok, errRet
}

//RandStringRunes _
func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
