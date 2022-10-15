package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);" json:"identity"`                 //唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"` //问题唯一标识
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`       //用户唯一标识
	Path            string `gorm:"column:path;type:varchar(255);" json:"path"`                        //问题存储地址
	Status          int    `gorm:"column:status;type:tinyint(1);" json:"status"`                      //状态，0-已提交待判断；1-答案正确；2-答案错误；3-运行超时；4-运行超内存
}

func (table *Submit) TableName() string {
	return "Submit"
}
