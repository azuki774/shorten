package model

import "time"

func (URLTable) TableName() string {
	return "URL_table"
}

type URLTable struct {
	ID        int64     `gorm:"id"`
	ShortKey  string    `gorm:"column:short_key"`
	TargetURL string    `gorm:"column:target_url"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
