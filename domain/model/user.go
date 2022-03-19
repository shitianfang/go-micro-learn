package model

// 定义一个结构体类型，使用时创建结构体User，直接使用结构体创建出来是对象
type User struct {
	ID       int64  `gorm:"primary_key;not_null;auto_increment"`
	UserName string `gorm:"unique_index;not_null"`
	PassWord string
}
