package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity          string             `grom:"column:identity;type:varchar(36);" json:"identity"`   //问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id"`                 //关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`        //标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`            //文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` //最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`         //最大运行内存
}

func (table *Problem) TableName() string {
	return "problem"
}

func GetProblemList(keyword string, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(Problem)).
		Preload("ProblemCategories").Preload("ProblemCategories.Category").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem.id").
			Where("pc.category_id = (SELECT c.id FROM category c WHERE c.identity = ? )", categoryIdentity)
	}
	return tx
}
