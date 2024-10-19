package Model

import (
	"NewGinBlog/Utills/ErrMsg"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=20"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=8,max=20"`
	Role     int    `gorm:"type:int;DEFAULT:2;not null" json:"role" validate:"required,gte=2"`
}

func CheckUser(name string) (code int) {
	var users User
	Db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return ErrMsg.ERROR_USERNAME_USED
	}
	return ErrMsg.SUCCESS
}

// 新增用户
// todo 使用钩子函数进行加密
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := Db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, ErrMsg.ERROR
	}
	return users, ErrMsg.SUCCESS
}

// 密码加密
func ScryptPw(password string) string {
	const Keylen = 8
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, Keylen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = Db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = Db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 钩子函数 用于密码加密
func (user *User) BeforeSave() {
	user.Password = ScryptPw(user.Password)
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User

	Db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return ErrMsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return ErrMsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return ErrMsg.ERROR_USER_NO_RIGHT
	}
	return ErrMsg.SUCCESS
}
