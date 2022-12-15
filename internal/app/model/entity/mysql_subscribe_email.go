package entity

const TableNameSubscribeEmail = "subscribe_email"

type SubscribeEmail struct {
	ID       int64  `gorm:"column:id;comment:主键" json:"id"`
	Email    string `gorm:"column:email;comment:邮箱" json:"email"`
	Deleted  bool   `gorm:"column:deleted;comment:逻辑删除" json:"deleted"`
	CreateAt int64  `gorm:"column:create_at" json:"create_at"`
	UpdateAt int64  `gorm:"column:update_at" json:"update_at"`
	DeleteAt int64  `gorm:"column:delete_at" json:"delete_at"`
}

func (*SubscribeEmail) TableName() string {
	return TableNameSubscribeEmail
}
