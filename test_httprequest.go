package http_result

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name          string `gorm:"column:name;type:varchar(255);not null;default:''" form:"name" json:"name"`
	Mobile        string `gorm:"column:mobile;type:varchar(255);not null;default:''" form:"mobile" json:"mobile"`
	Password      string `gorm:"column:password;type:varchar(255);not null;default:''" form:"password" json:"password"`
	Area          string `gorm:"column:area;type:varchar(255);not null;default:''" form:"area" json:"area"`
	Contact       string `gorm:"column:contact;type:varchar(255);not null;default:''" form:"contact" json:"contact"`
	ContactMobile string `gorm:"column:contact_mobile;type:varchar(255);not null;default:''" form:"contact_mobile" json:"contact_mobile"`
	Address       string `gorm:"column:address;type:varchar(255);not null;default:''" form:"address" json:"address"`
	IdCard        string `gorm:"column:id_card;type:varchar(255);not null;default:''" form:"id_card" json:"id_card"`
	Pid           uint   `gorm:"column:pid;type:int(11);not null;default:0" form:"pid" json:"pid"`
	Parent        *User  `gorm:"foreignKey:Pid;references:id" form:"parent" json:"parent"`
	Children      []User `gorm:"foreignKey:Pid;references:id" form:"children" json:"children"`
	orm.SoftDeletes
}

func Test_HttpRequest() {

}
