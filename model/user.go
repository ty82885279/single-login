package model

type User struct {
	ID    int64  `json:"id" form:"id" gorm:"column:id" `
	Name  string `json:"name" form:"name" gorm:"column:name" binding:"required"`
	Psw   string `json:"psw" form:"psw" gorm:"column:psw" binding:"required"`
	ResID string `json:"res_id" form:"res_id"`
}

func (User) TableName() string {
	return "user"
}
