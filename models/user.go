package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity         string `grom:"column:identity;type:varchar(36);" json:"identity"`                 //用户表的唯一标识
	Name             string `gorm:"column:name;type:varchar(100);" json:"name"`                        //用户名
	Password         string `gorm:"column:password;type:varchar(32);" json:"password"`                 //密码
	Phone            string `gorm:"column:phone;type:varchar(20);" json:"phone"`                       //手机号
	Mail             string `gorm:"column:mail;type:varchar(100);" json:"mail"`                        //邮箱
	FinishProblemNum int    `gorm:"column:finish_problem_num;type:int(11);" json:"finish_problem_num"` //完成问题个数
	SubmitNum        int    `gorm:"column:submit_num;type:int(11);" json:"submit_num"`                 //提交次数
}

func (table *User) TableName() string {
	return "user"
}
