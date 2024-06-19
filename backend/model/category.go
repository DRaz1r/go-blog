package model

// GORM 会自动将模型名称转换为复数形式作为表名，除非你显式指定表名
type Category struct {
	ID           uint   `json:"id" gorm:"type:uint;primary_key;"`
	CategoryName string `json:"name" gorm:"type:varchar(50);not null"`
}
