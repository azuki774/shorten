package model

import "time"

func (URLTable) TableName() string {
	return "URL_table"
}

type URLTable struct {
	ID        int64     `gorm:"id"`
	Source    string    `gorm:"column:source"`
	Target    string    `gorm:"column:target"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
