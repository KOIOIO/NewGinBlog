package Model

import (
	"NewGinBlog/Utills/ErrMsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(10);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(255);not null" json:"desc"`
	Content string `gorm:"type:text;not null" json:"content"`
	Img     string `gorm:"type:varchar(100);not null" json:"img"`
}

// 新增文章
// todo 使用钩子函数进行加密
func CreateArticle(data *Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

//todo 查询分类下所有文章
//todo 查询单个文章

// 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, ErrMsg.ERROR
	}
	return articleList, ErrMsg.SUCCESS
}

// 查分类下的所有文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArticles []Article
	err := Db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArticles).Error
	if err != nil {
		return nil, ErrMsg.ERROR_CATE_NOT_EXIST
	}
	return cateArticles, ErrMsg.SUCCESS
}

// 查询单个文章
func GetArtInfro(id int) (Article, int) {
	var article Article
	err := Db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, ErrMsg.ERROR_ART_NOT_EXIST
	}
	return article, ErrMsg.SUCCESS
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	var Article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = Db.Model(&Article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var Article Article
	err = Db.Where("id = ? ", id).Delete(&Article).Error
	if err != nil {
		return ErrMsg.ERROR
	}
	return ErrMsg.SUCCESS
}
