package model

import (
	"context"
	"short-url/pkg/database/mysql"
	"time"
)

type TinyUrl struct {
	ID          int       `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	TinyUrl     string    `gorm:"column:tiny_url"`              // 短链
	OriginalUrl string    `gorm:"column:original_url;NOT NULL"` // 原始链接
	CreatorIP   int       `gorm:"column:creator_ip;default:0;NOT NULL"`
	InstanceID  int       `gorm:"column:instance_id;default:0;NOT NULL"`
	ExpireTime  time.Time `gorm:"column:expire_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 过期时间
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

func (m *TinyUrl) TableName() string {
	return "tiny_url"
}

// GetOriginUrlByTinyUrl 获取原地址
func (m *TinyUrl) GetOriginUrlByTinyUrl(ctx context.Context, tinyUrl string) error {
	return mysql.MysqlClient.Where("tiny_url = ?", tinyUrl).First(m).Error
}

func (m *TinyUrl) Create(ctx context.Context) error {
	return mysql.MysqlClient.Create(m).Error
}
