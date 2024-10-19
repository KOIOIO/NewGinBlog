package Model

import (
	"NewGinBlog/Utills/ErrMsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID           uint   `gorm:"primary_key;auto_increment" json:"id"`
	CategoryName string `gorm:"type:varchar(20);not null" json:"category_name"`
}

// 查询分类下是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	Db.Select("id").Where("category_name = ?", name).First(&cate)
	if cate.ID > 0 {
		return ErrMsg.ERROR_CATENAME_USED
	}
	return ErrMsg.SUCCESS
}

// 新增分类

func CreateCategory(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, ErrMsg.ERROR
	}
	return cate, ErrMsg.SUCCESS
}

//查分类下的所有文章

// 编辑分类
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["category_name"] = data.CategoryName
	err = Db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 删除用户
func DeleteCategory(id int) int {
	var cate Category
	err = Db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}
